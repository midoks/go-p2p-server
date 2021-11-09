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
	fmt.Println("app init")
	hub.Init()

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
					//log.Warnf("client %s is expired for %d, close it", cli.PeerId, now-cli.Timestamp)
					if ok := hub.DoUnregister(cli.PeerId); ok {
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
	gin.SetMode(gin.ReleaseMode)
	r.Use(httpCors())

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	r.POST("/channel", p2pChannel)
	r.POST("/channel/:channel_id/node/:peer/stats", p2pChannelStats)
	r.POST("/channel/:channel_id/node/:peer/peers", p2pChannelPeers)

	r.GET("/ws", websocketReqMethod)
	r.GET("/ws?id=:id", websocketReqMethod)
	r.GET("count", websocketConnCount)

	r.Run(fmt.Sprintf(":%s", httpPort))
}
