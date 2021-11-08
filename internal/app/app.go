package app

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	go_logger "github.com/phachon/go-logger"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var logger *go_logger.Logger

//websocket
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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

//websocket实现
func websocketReqMethod(c *gin.Context) {

	id := c.Query("id")

	fmt.Println("websocketReqMethod id", id)
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	// defer ws.Close()

	for {

		// go func(ws *websocket.Conn) {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			// logger.Errorf("read websocket msg: %v", err)
			fmt.Println("err:", err, "id:", mt)
			// logger.Errorf("read websocket msg: %v", err)
			break
		}

		fmt.Println(mt, string(message), err, runtime.NumGoroutine())

		// }(ws)
		now := time.Now().Format("2006-01-02 15:04:05")
		log.Printf("recv: %s", message)

		err = ws.WriteMessage(mt, []byte(now))
		if err != nil {
			log.Println("write:", err)
			break
		}

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
