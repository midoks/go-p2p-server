package app

import (
	// "fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/midoks/go-p2p-server/internal/geoip"
	"github.com/midoks/go-p2p-server/internal/hub"
	"github.com/midoks/go-p2p-server/internal/logger"
	"github.com/midoks/go-p2p-server/internal/mem"
	"github.com/midoks/go-p2p-server/internal/queue"
	"github.com/midoks/go-p2p-server/internal/tools"
	"github.com/midoks/go-p2p-server/internal/tools/uniqid"
)

type AnPeer struct {
	Id string `json:"id"`
}

var mu sync.RWMutex

//接收announce信息
func p2pChannel(c *gin.Context) {

	postJson := make(map[string]interface{})
	c.BindJSON(&postJson)

	uniqidId := uniqid.New(uniqid.Params{"", false})
	uniqidId = uniqidId + tools.RandId()

	gPeers := []AnPeer{}
	if channel, ok := postJson["channel"]; ok {
		key := channel.(string)
		if peer, ok := mem.GetChannel(key); ok {
			for _, p := range peer {
				gPeers = append(gPeers, AnPeer{Id: p})
			}
			mem.SetChannel(key, uniqidId)
		} else {
			mem.SetChannel(key, uniqidId)
		}
	}

	go func() {

		ipAddr := c.ClientIP()
		if ipAddr == "127.0.0.1" {
			ipAddr = tools.GetNetworkIp()
		}

		lat, lng := geoip.GetLatLongByIpAddr(ipAddr)
		//客服端经纬度->保存到redis
		mem.SetPeerLatLang(uniqidId, lat, lng)

		to_lat, to_lng, err := mem.GetServerLatLang()
		if err != nil {
			logger.Errorf("announce.GetServerLatLang error: %v", err)
			to_lat, to_lng = 0, 0
		}

		if !strings.HasPrefix(uniqidId, "p2p") {
			queue.PushText("join", uniqidId, lat, lng, to_lat, to_lng)
		}

	}()

	c.JSON(http.StatusOK, gin.H{
		"ret": 0,
		"data": gin.H{
			"id":              uniqidId,
			"peers":           gPeers,
			"report_interval": 3,
			"v":               uniqidId,
		},
	})

}

//接收announce信息
func p2pChannelPeers(c *gin.Context) {
	postJson := make(map[string]interface{}) //注意该结构接受的内容
	c.ShouldBind(&postJson)

	channel_id := c.Param("channel_id")
	peers := c.Param("peer")

	gPeers := []AnPeer{}
	if peer, ok := mem.GetChannel(channel_id); ok {
		for _, p := range peer {
			gPeers = append(gPeers, AnPeer{Id: p})
		}
	}

	mem.SetPeerHeartbeat(peers, 60*time.Second)
	if c, ok := hub.GetClient(peers); ok {
		c.UpdateTs()
	}

	c.JSON(http.StatusOK, gin.H{
		"ret": 0,
		"data": gin.H{
			"id":              peers,
			"peers":           gPeers,
			"report_interval": 3,
			"v":               peers,
		},
	})
}

//接收announce信息
func p2pChannelStats(c *gin.Context) {
	json := make(map[string]interface{})
	c.ShouldBind(&json)

	channel_id := c.Param("channel_id")
	peers := c.Param("peer")

	//延长缓冲时间
	mem.SetPeerHeartbeat(peers, 60*time.Second)
	if c, ok := hub.GetClient(peers); ok {
		c.UpdateTs()
	}

	//查找缓冲中的peer
	gPeers := []AnPeer{}
	if peer, ok := mem.GetChannel(channel_id); ok {
		for _, p := range peer {
			gPeers = append(gPeers, AnPeer{Id: p})
		}
	}

	// if !strings.HasPrefix(peers, "p2p") {
	// 	ipAddr := c.ClientIP()
	// 	if ipAddr == "127.0.0.1" {
	// 		ipAddr = tools.GetNetworkIp()
	// 	}

	// 	lat, lng := geoip.GetLatLongByIpAddr(ipAddr)
	// 	to_lat, to_lng, _ := mem.GetServerLatLang()

	// 	queue.PushText("join", peers, lat, lng, to_lat, to_lng)
	// }

	c.JSON(http.StatusOK, gin.H{
		"ret":   0,
		"id":    peers,
		"name":  "stats",
		"peers": gPeers,
		"data":  gin.H{},
	})
}
