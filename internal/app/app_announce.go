package app

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/midoks/go-p2p-server/internal/announce"
	"github.com/midoks/go-p2p-server/internal/hub"
	"github.com/midoks/go-p2p-server/internal/queue"
	"github.com/midoks/go-p2p-server/internal/tools"
	"github.com/midoks/go-p2p-server/internal/tools/uniqid"
)

func initAnnounce() {
	announce.Init()
}

type AnPeer struct {
	Id string `json:"id"`
}

var mu sync.RWMutex

//接收announce信息
func p2pChannel(c *gin.Context) {

	postJson := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&postJson)

	ipAddr := c.ClientIP()
	if ipAddr == "127.0.0.1" {
		ipAddr = tools.GetNetworkIp()
	}

	lat, lang := tools.GetLatLongByIpAddr(ipAddr)
	fmt.Println("lat:", lat)
	fmt.Println("lang:", lang)

	uniqidId := uniqid.New(uniqid.Params{"", false})
	uniqidId = uniqidId + tools.RandId()

	gPeers := []AnPeer{}
	if channel, ok := postJson["channel"]; ok {
		key := channel.(string)
		if peer, ok := announce.Get(key); ok {
			for _, p := range peer {

				gPeers = append(gPeers, AnPeer{Id: p})
			}
			announce.Set(key, uniqidId)
		} else {
			announce.Set(key, uniqidId)
		}
	}

	// for i := 0; i < 99; i++ {
	queue.PushText("join", uniqidId, lang, lat, -121.9829, 37.567)
	// }

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
	key := channel_id
	if peer, ok := announce.Get(key); ok {
		for _, p := range peer {
			gPeers = append(gPeers, AnPeer{Id: p})
		}
	}

	announce.SetPeerHeartbeat(peers, 60*time.Second)

	c.JSON(http.StatusOK, gin.H{
		"ret": 0,
		"data": gin.H{
			"id":              peers,
			"peers":           gPeers,
			"report_interval": 3,
			"v":               "scadasd",
		},
	})
}

//接收announce信息
func p2pChannelStats(c *gin.Context) {
	json := make(map[string]interface{})
	c.ShouldBind(&json)

	channel_id := c.Param("channel_id")
	peers := c.Param("peer")

	announce.SetPeerHeartbeat(peers, 60*time.Second)

	gPeers := []AnPeer{}
	key := channel_id
	if peer, ok := announce.Get(key); ok {
		for _, p := range peer {
			gPeers = append(gPeers, AnPeer{Id: p})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"ret":   0,
		"id":    peers,
		"name":  "stats",
		"peers": gPeers,
		"data":  gin.H{},
	})
}

func p2pGetStats(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"peers":         hub.GetClientNum(),
		"serverRuntime": time.Now().Format("2006-01-02 15:04:05"),
	})
}
