package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/midoks/go-p2p-server/internal/announce"
	"github.com/midoks/go-p2p-server/internal/tools/uniqid"
)

func initAnnounce() {
	announce.Init()
}

//接收announce信息
func p2pChannel(c *gin.Context) {
	postJson := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&postJson)

	uniqid_id := uniqid.New(uniqid.Params{"", false})
	gPeers := []string{}
	if channel, ok := postJson["channel"]; ok {
		key := channel.(string)
		if peer, ok := announce.Get(key); ok {
			for _, p := range peer {
				fmt.Println(p)
				gPeers = append(gPeers, p)
			}
			announce.Set(key, uniqid_id)
		} else {
			announce.Set(key, uniqid_id)
		}

		fmt.Println("hubAnn count:", announce.KeyCount(key))
	}

	c.JSON(http.StatusOK, gin.H{
		"ret": 0,
		"data": gin.H{
			"id":              uniqid_id,
			"peers":           gPeers,
			"report_interval": 3,
			"v":               uniqid_id,
		},
	})
}

//接收announce信息
func p2pChannelPeers(c *gin.Context) {
	postJson := make(map[string]interface{}) //注意该结构接受的内容
	c.ShouldBind(&postJson)

	channel_id := c.Param("channel_id")
	peers := c.Param("peer")

	gPeers := []string{}
	key := channel_id
	if peer, ok := announce.Get(key); ok {
		for _, p := range peer {
			gPeers = append(gPeers, p)
		}
	}

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
	fmt.Println(json)

	channel_id := c.Param("channel_id")
	peers := c.Param("peer")

	gPeers := []string{}
	key := channel_id
	if peer, ok := announce.Get(key); ok {
		for _, p := range peer {
			gPeers = append(gPeers, p)
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
