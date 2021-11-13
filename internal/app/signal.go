package app

import (
	"fmt"
	// "log"
	"bytes"
	// "encoding/json"
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

	uniqidId := c.Query("id")
	fmt.Println("websocket id:[", uniqidId, "]")

	//use webSocket pro
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	clientId := client.New(uniqidId, ws, true)
	clientId.SendMsgVersion(tools.GetVersionNum(conf.App.Version))
	hub.DoRegister(clientId)

	go func() {
		ipAddr := c.ClientIP()
		if ipAddr == "127.0.0.1" {
			ipAddr = tools.GetNetworkIp()
		}

		lat, lang := geoip.GetLatLongByIpAddr(ipAddr)
		clientId.SetLatLong(lat, lang)

		go func() {
			queue.PushText("join", uniqidId, lat, lang, -121.9829, 37.567)
		}()
	}()

	go func() {
		for {
			mt, message, err := ws.ReadMessage()
			if err != nil {
				logger.Errorf("path[ws][%s] error: %v", uniqidId, err)
				break
			}
			clientId.SetMT(mt)

			// fmt.Println("go func:", string(message))
			data := bytes.TrimSpace(bytes.Replace(message, []byte{'\n'}, []byte{' '}, -1))
			hdr, err := handler.NewHandler(data, clientId)
			if err != nil {
				clientId.UpdateTs()
				logger.Errorf("path[ws][%s] hander error: %v", uniqidId, err)
			} else {
				hdr.Handle()
			}
		}
	}()
}
