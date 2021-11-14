package app

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/midoks/go-p2p-server/internal/announce"
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
func wsSignal(c *gin.Context) {

	uniqidId := c.Query("id")
	// fmt.Println("websocket id:[", uniqidId, "]")

	//use webSocket pro
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Errorf("ws fail: %v", err)
		ws.Close()
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

		to_lat, to_lang, err := announce.GetServerLatLang()
		if err != nil {
			logger.Errorf("announce.GetServerLatLang error: %v", err)
			to_lat, to_lang = 0, 0
		}

		if !strings.HasPrefix(uniqidId, "p2p") {
			queue.PushText("join", uniqidId, lat, lang, to_lat, to_lang)
		}

	}()

	go func() {
		for {
			mt, message, err := ws.ReadMessage()
			if err != nil {
				hub.DoUnregister(uniqidId)
				queue.PushTextLeave(uniqidId)
				announce.DelPeer(uniqidId)

				// 主动关闭,非异常
				// logger.Debugf("path[ws][%s] %v", uniqidId, err)
				break
			}
			clientId.SetMT(mt)

			data := bytes.TrimSpace(bytes.Replace(message, []byte{'\n'}, []byte{' '}, -1))
			hdr, err := handler.NewHandler(data, clientId)
			if err != nil {
				logger.Errorf("path[ws][%s] hander error: %v", uniqidId, err)
			} else {
				clientId.UpdateTs()
				hdr.Handle()
			}
		}
	}()
}
