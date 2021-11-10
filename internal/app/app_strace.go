package app

import (
	"fmt"
	// "log"
	"bytes"
	"net/http"
	// "runtime"
	// "time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/midoks/go-p2p-server/internal/client"
	"github.com/midoks/go-p2p-server/internal/conf"
	"github.com/midoks/go-p2p-server/internal/handler"
	"github.com/midoks/go-p2p-server/internal/hub"
	"github.com/midoks/go-p2p-server/internal/tools"
)

//websocket
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//websocket实现
func websocketReqMethod(c *gin.Context) {
	// c.Request.ParseForm()
	// id := c.Request.Form.Get("id")
	id := c.Query("id")

	fmt.Println("websocket id:[", id, "]")
	//use webSocket pro
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	// defer ws.Close()

	clientId := client.New(id, ws, true)
	clientId.SendMsgVersion(tools.GetVersionNum(conf.App.Version))
	hub.DoRegister(clientId)

	go func() {
		for {

			mt, message, err := ws.ReadMessage()
			if err != nil {
				logger.Errorf("read websocket msg: %v", err)
				fmt.Println("err:", err, "id:", mt)
				break
			}

			fmt.Println("go func:", mt, string(message))
			data := bytes.TrimSpace(bytes.Replace(message, []byte{'\n'}, []byte{' '}, -1))
			hdr, err := handler.NewHandler(data, clientId)
			if err != nil {
				clientId.UpdateTs()
				logger.Errorf("NewHandler %v", err)
			} else {
				hdr.Handle()
			}
		}
	}()

}
