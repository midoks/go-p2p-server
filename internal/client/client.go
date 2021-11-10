package client

import (
	"encoding/json"
	"fmt"
	// "log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	MAX_NOT_FOUND_PEERS_LIMIT = 5
)

type SignalVerResponse struct {
	Action string `json:"action"`
	Ver    int    `json:"ver"`
}

type Client struct {
	Ws            *websocket.Conn
	PeerId        string //唯一标识
	LocalNode     bool   // 是否本地节点
	Timestamp     int64
	NotFoundPeers []string // 记录没有找到的peer的队列
}

func New(peerId string, ws *websocket.Conn, localNode bool) *Client {
	return &Client{
		Ws:        ws,
		PeerId:    peerId,
		LocalNode: localNode,
		Timestamp: time.Now().Unix(),
	}
}

func (c *Client) SendMsgVersion(version int) error {
	resp := SignalVerResponse{
		Action: "ver",
		Ver:    version,
	}
	b, err := json.Marshal(resp)
	if err != nil {
		// log.Error("json.Marshal", err)
		fmt.Println("json.Marshal", err)
		return err
	}
	err = c.SendMessage(b)
	return err
}

func (c *Client) UpdateTs() {
	//log.Warnf("%s UpdateTs", c.PeerId)
	c.Timestamp = time.Now().Unix()
}

func (c *Client) SendMessage(msg []byte) error {
	return c.Ws.WriteMessage(1, msg)
}

func (c *Client) Close() error {
	return c.Ws.Close()
}

func (c *Client) IsExpired(now, limit int64) bool {
	return now-c.Timestamp > limit
}

func (c *Client) EnqueueNotFoundPeer(id string) {
	c.NotFoundPeers = append(c.NotFoundPeers, id)
	if len(c.NotFoundPeers) > MAX_NOT_FOUND_PEERS_LIMIT {
		c.NotFoundPeers = c.NotFoundPeers[1:len(c.NotFoundPeers)]
	}
}

func (c *Client) HasNotFoundPeer(id string) bool {
	for _, v := range c.NotFoundPeers {
		if id == v {
			return true
		}
	}
	return false
}
