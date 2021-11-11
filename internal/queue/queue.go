package queue

import (
	// "encoding/json"
	"fmt"
)

type MSlice struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type MSliceMap []MSlice

var (
	ValChan chan MSliceMap
)

func Init() {
	ValChan = make(chan MSliceMap)
}

func PushText(action string, val string) {
	a := make(MSliceMap, 0)
	v := MSlice{Type: action, Data: val}
	b := append(a, v)
	ValChan <- b
}

func Receive() {

	for {
		data := <-ValChan
		fmt.Println("queue", data)
	}

}
