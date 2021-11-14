package app

import (
	"fmt"
	// "log"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/midoks/go-p2p-server/internal/conf"
	"github.com/midoks/go-p2p-server/internal/geoip"
	"github.com/midoks/go-p2p-server/internal/hub"
	"github.com/midoks/go-p2p-server/internal/logger"
	"github.com/midoks/go-p2p-server/internal/queue"
	"github.com/midoks/go-p2p-server/internal/tools"
)

const (
	CHECK_CLIENT_INTERVAL = 10
	EXPIRE_LIMIT          = 100
)

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
	conf.Init()
	logger.Init()
	geoip.Init()

	//App
	queue.Init()
	hub.Init()
	initAnnounce()

	go func() {
		for {
			time.Sleep(CHECK_CLIENT_INTERVAL * time.Second)
			now := time.Now().Unix()
			// fmt.Println("start check client alive...")
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
			// fmt.Println("check client finished, closed  clients:", count)
		}
	}()
}

func websocketConnCount(c *gin.Context) {
	// num := hub.GetClientNum()
	c.String(http.StatusOK, fmt.Sprintf("%d", hub.GetClientNum()))
}

func Run() {
	r := gin.Default()
	r.Use(httpCors())

	r.Static("/public", "./public")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"version": conf.App.Version})
	})

	r.GET("/t", func(c *gin.Context) {
		ip := "127.0.0.1"
		if find_ip, err := tools.GetOutboundIP(); err == nil {
			fmt.Println("ip:", ip)
			ip = find_ip
		} else {
			fmt.Println("ip:", err)
		}

		lat, lang := geoip.GetLatLongByIpAddr(ip)

		c.JSON(http.StatusOK, gin.H{
			"gnum": runtime.NumGoroutine(),
			"ip":   ip,
			"lat":  lat,
			"lang": lang,
		})
	})

	r.GET("/getStats", p2pGetStats)
	r.GET("/count", websocketConnCount)

	r.POST("/channel", p2pChannel)
	r.POST("/channel/:channel_id/node/:peer/stats", p2pChannelStats)
	r.POST("/channel/:channel_id/node/:peer/peers", p2pChannelPeers)

	r.GET("/ws", wsSignal)
	r.GET("/ws?id=:id", wsSignal)

	r.Run(fmt.Sprintf(":%s", conf.Web.HttpPort))
}
