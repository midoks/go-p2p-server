package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//接收announce信息
func p2pChannel(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	log.Printf("%v", &json)

	c.JSON(http.StatusOK, gin.H{
		"ret": 0,
		"data": gin.H{

			"id":              "123123123",
			"peers":           []string{},
			"report_interval": 10,
			"v":               "scadasd",
		},
	})
}

//接收announce信息
func p2pChannelPeers(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	log.Printf("%v", &json)

	channel_id := c.Param("channel_id")
	peers := c.Param("peer")

	fmt.Println("peers:", channel_id, peers)

	c.JSON(http.StatusOK, gin.H{
		"ret": 0,
		"data": gin.H{

			"id":              "123123123",
			"peers":           []string{},
			"report_interval": 10,
			"v":               "scadasd",
		},
	})
}

//接收announce信息
func p2pChannelStats(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	log.Printf("%v", &json)

	channel_id := c.Param("channel_id")
	peers := c.Param("peer")

	fmt.Println("stats:", channel_id, peers)

	c.JSON(http.StatusOK, gin.H{
		"ret":  0,
		"name": "stats",
		"data": gin.H{},
	})
}
