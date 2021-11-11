package app

import (
	"fmt"
	// "log"
	"net/http"
	"time"

	go_logger "github.com/phachon/go-logger"

	"github.com/gin-gonic/gin"
	// "github.com/midoks/go-p2p-server/internal/conf"
	"github.com/midoks/go-p2p-server/internal/hub"
	"github.com/midoks/go-p2p-server/internal/queue"
	"github.com/midoks/go-p2p-server/internal/tools"
)

const (
	CHECK_CLIENT_INTERVAL = 10
	EXPIRE_LIMIT          = 100
)

var logger *go_logger.Logger

func httpCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// if origin != "" {
		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		// }
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func init() {
	queue.Init()
	tools.Init()
	hub.Init()
	initAnnounce()

	logger = go_logger.NewLogger()

	// 文件输出配置
	fileConfig := &go_logger.FileConfig{
		Filename: "./logs/test.log", // 日志输出文件名，不自动存在
		// 如果要将单独的日志分离为文件，请配置LealFrimeNem参数。
		LevelFileName: map[int]string{
			logger.LoggerLevel("error"): "./logs/error.log", // Error 级别日志被写入 error .log 文件
			logger.LoggerLevel("info"):  "./logs/info.log",  // Info 级别日志被写入到 info.log 文件中
			logger.LoggerLevel("debug"): "./logs/debug.log", // Debug 级别日志被写入到 debug.log 文件中
		},
		MaxSize:    1024 * 1024, // 文件最大值（KB），默认值0不限
		MaxLine:    100000,      // 文件最大行数，默认 0 不限制
		DateSlice:  "d",         // 文件根据日期切分， 支持 "Y" (年), "m" (月), "d" (日), "H" (时), 默认 "no"， 不切分
		JsonFormat: false,       // 写入文件的数据是否 json 格式化
		Format:     "",          // 如果写入文件的数据不 json 格式化，自定义日志格式
	}
	// 添加 file 为 logger 的一个输出
	logger.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileConfig)

	go func() {
		for {
			time.Sleep(CHECK_CLIENT_INTERVAL * time.Second)
			now := time.Now().Unix()
			fmt.Println("start check client alive...")
			count := 0
			for item := range hub.GetInstance().Clients.IterBuffered() {
				cli := item.Val
				if cli.LocalNode && cli.IsExpired(now, EXPIRE_LIMIT) {
					// 节点过期
					info := fmt.Sprintf("client %s is expired for %d, close it", cli.PeerId, now-cli.Timestamp)
					fmt.Println(info)
					if ok := hub.DoUnregister(cli.PeerId); ok {
						queue.PushTextLeave(cli.PeerId)
						cli.Close()
						count++
					}
				}
			}
			fmt.Println("check client finished, closed  clients:", count)
		}
	}()
}

func websocketConnCount(c *gin.Context) {
	// num := hub.GetClientNum()
	c.String(http.StatusOK, fmt.Sprintf("%d", hub.GetClientNum()))
}

func Run() {
	httpPort := "3030"
	r := gin.Default()
	r.Use(httpCors())

	r.Static("/public", "./public")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	r.POST("/channel", p2pChannel)
	r.POST("/channel/:channel_id/node/:peer/stats", p2pChannelStats)
	r.POST("/channel/:channel_id/node/:peer/peers", p2pChannelPeers)

	r.GET("/ws", wsReqMethod)
	r.GET("/ws?id=:id", wsReqMethod)
	r.GET("/trace", wsTrace)
	r.GET("/trace?id=:id", wsTrace)
	r.GET("/count", websocketConnCount)

	r.Run(fmt.Sprintf(":%s", httpPort))
}
