package app

import (
	"fmt"
	"net/http"

	go_logger "github.com/phachon/go-logger"

	"github.com/gin-gonic/gin"
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

func Run() {
	httpPort := "3030"
	r := gin.Default()
	r.Use(httpCors())

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	r.POST("/channel", p2pChannel)
	r.POST("/channel/:channel_id/node/:peer/stats", p2pChannelStats)
	r.POST("/channel/:channel_id/node/:peer/peers", p2pChannelPeers)

	r.GET("/ws", websocketReqMethod)
	r.GET("/ws?id=:id", websocketReqMethod)

	r.Run(fmt.Sprintf(":%s", httpPort))
}
