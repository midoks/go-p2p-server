package app

import (
	"fmt"
	// "log"
	"bytes"
	"encoding/json"
	"net/http"
	// "runtime"
	// "time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/midoks/go-p2p-server/internal/client"
	"github.com/midoks/go-p2p-server/internal/conf"
	"github.com/midoks/go-p2p-server/internal/geoip"
	"github.com/midoks/go-p2p-server/internal/handler"
	"github.com/midoks/go-p2p-server/internal/hub"
	"github.com/midoks/go-p2p-server/internal/logger"
	"github.com/midoks/go-p2p-server/internal/queue"
	"github.com/midoks/go-p2p-server/internal/tools"
)

//websocket
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//websocket实现
func wsReqMethod(c *gin.Context) {
	// c.Request.ParseForm()
	// id := c.Request.Form.Get("id")
	uniqidId := c.Query("id")

	fmt.Println("websocket id:[", uniqidId, "]")
	//use webSocket pro
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	// defer ws.Close()

	ipAddr := c.ClientIP()
	if ipAddr == "127.0.0.1" {
		ipAddr = tools.GetNetworkIp()
	}

	lat, lang := geoip.GetLatLongByIpAddr(ipAddr)
	queue.PushText("join", uniqidId, lang, lat, -121.9829, 37.567)

	clientId := client.New(uniqidId, ws, true)
	clientId.SendMsgVersion(tools.GetVersionNum(conf.App.Version))
	clientId.SetLatLong(lat, lang)
	hub.DoRegister(clientId)

	go func() {
		for {

			_, message, err := ws.ReadMessage()
			if err != nil {
				logger.Errorf("read websocket msg: %v", err)
				fmt.Println("err:", err, "id:")
				break
			}

			// fmt.Println("go func:", string(message))
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

func wsTrace(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	uniqidId := c.Query("id")
	clientId := client.New(uniqidId, ws, true)
	clientId.SendMsgVersion(tools.GetVersionNum(conf.App.Version))

	go func() {
		for {

			data := <-queue.ValChan
			fmt.Println("queue", data)

			b, err := json.Marshal(data)
			if err != nil {
				// log.Error("json.Marshal", err)
				fmt.Println("trace json.Marshal", err)
			} else {
				ws.WriteMessage(1, b)
			}

			// _, message, err := ws.ReadMessage()
			// if err != nil {
			// 	logger.Errorf("read websocket msg: %v", err)
			// 	fmt.Println("err:", err, "id:")
			// 	break
			// }

		}
	}()
}
