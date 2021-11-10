package hub

import (
	"encoding/json"
	"fmt"

	"github.com/midoks/go-p2p-server/internal/client"
	"github.com/midoks/go-p2p-server/internal/hub/cmap"
)

var h *Hub

type Hub struct {
	Clients cmap.ConMap
}

func Init() {
	h = &Hub{
		Clients: cmap.New(),
	}
}

func GetInstance() *Hub {
	return h
}

func DoRegister(client *client.Client) {
	fmt.Println("hub DoRegister:", client.PeerId)
	if client.PeerId != "" {
		h.Clients.Set(client.PeerId, client)
	} else {
		panic("DoRegister")
	}
}

func GetClient(id string) (*client.Client, bool) {
	return h.Clients.Get(id)
}

func GetClientNumPerMap() []int {
	return h.Clients.CountPerMapNoLock()
}

func GetClientNum() int {
	return h.Clients.CountNoLock()
}

func RemoveClient(id string) {
	h.Clients.Remove(id)
}

func DoUnregister(peerId string) bool {
	fmt.Println("hub DoUnregister :", peerId)
	if peerId == "" {
		return false
	}
	if h.Clients.Has(peerId) {
		h.Clients.Remove(peerId)
		return true
	}
	return false
}

// send json object to a client with peerId
func SendJsonToClient(target *client.Client, value interface{}) error {
	b, err := json.Marshal(value)
	if err != nil {
		// log.Error("json.Marshal", err)
		return err
	}
	//if target == nil {
	//	//log.Printf("sendJsonToClient error")
	//	return fmt.Errorf("peer %s not found", target.PeerId)
	//}
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			// log.Warnf(err.(string)) // 这里的err其实就是panic传入的内容
		}
	}()

	return target.SendMessage(b)
}
