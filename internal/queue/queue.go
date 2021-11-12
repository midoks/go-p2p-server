package queue

import (
	// "encoding/json"
	"fmt"
)

type LatLang struct {
	From []float64 `json:"from"`
	To   []float64 `json:"to"`
}

type MSlice struct {
	Type   string      `json:"type"`
	Data   interface{} `json:"data"`
	PeerId interface{} `json:"peer_id,omitempty"`
}

type MSliceMap []MSlice

var (
	ValChan chan MSliceMap
)

func Init() {
	ValChan = make(chan MSliceMap)
}

func PushText(action string, peer string, lat_from float64, lang_from float64, lat_to float64, long_to float64) {
	a := make(MSliceMap, 0)

	from_ll := make([]float64, 0)
	from_ll = append(from_ll, lat_from)
	from_ll = append(from_ll, lang_from)

	to_ll := make([]float64, 0)
	to_ll = append(to_ll, lat_to)
	to_ll = append(to_ll, long_to)

	ll := LatLang{From: from_ll, To: to_ll}

	ll_data := make([]LatLang, 0)
	ll_data = append(ll_data, ll)

	v := MSlice{Type: action, PeerId: peer, Data: ll_data}
	b := append(a, v)

	ValChan <- b
}

func PushConnection(action string, lat_from float64, lang_from float64, lat_to float64, long_to float64) {
	a := make(MSliceMap, 0)

	from_ll := make([]float64, 0)
	from_ll = append(from_ll, lat_from)
	from_ll = append(from_ll, lang_from)

	to_ll := make([]float64, 0)
	to_ll = append(to_ll, lat_to)
	to_ll = append(to_ll, long_to)

	ll := LatLang{From: from_ll, To: to_ll}

	ll_data := make([]LatLang, 0)
	ll_data = append(ll_data, ll)

	v := MSlice{Type: action, Data: ll_data}
	b := append(a, v)

	ValChan <- b
}

func PushTextLeave(val string) {
	a := make(MSliceMap, 0)
	v := MSlice{Type: "leave", Data: val}
	b := append(a, v)
	ValChan <- b
}

func PushTextLatLang(action string, val string) {
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
