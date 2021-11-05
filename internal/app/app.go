package app

import (
	"fmt"
	"net/http"

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

//websocket实现
func websocketReqMethod(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	for {

		mt, message, err := ws.ReadMessage()
		if err != nil {
			logger.Errorf("read websocket msg: %v", err)
			break
		}

		fmt.Println(mt, message, err)

	}
}

func Run() {
	httpPort := "3030"
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	r.GET("/ws", websocketReqMethod)

	r.Run(fmt.Sprintf(":%s", httpPort))
}
