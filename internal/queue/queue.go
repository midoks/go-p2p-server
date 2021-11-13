package queue

import (
	// "encoding/json"
	"fmt"
)

type LatLang struct {
	From [][]float64 `json:"from"`
	To   []float64   `json:"to"`
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
	ValChan = make(chan MSliceMap, 100)
}

func PushText(action string, peer string, lat_from float64, lang_from float64, lat_to float64, lang_to float64) {
	msg := make(MSliceMap, 0)

	from_ll := make([]float64, 0)
	from_ll = append(from_ll, lang_from)
	from_ll = append(from_ll, lat_from)
	from_ll_arr := make([][]float64, 0)
	from_ll_arr = append(from_ll_arr, from_ll)

	// from_ll := [1][2]float64{{lang_from, lat_from}}

	to_ll := make([]float64, 0)
	to_ll = append(to_ll, lang_to)
	to_ll = append(to_ll, lat_to)

	ll := LatLang{From: from_ll_arr, To: to_ll}

	ll_data := make([]LatLang, 0)
	ll_data = append(ll_data, ll)

	v := MSlice{Type: action, PeerId: peer, Data: ll_data}

	msg = append(msg, v)
	Push(msg)

}

func PushConnection(action string, lat_from float64, lang_from float64, lat_to float64, long_to float64) {
	msg := make(MSliceMap, 0)

	from_ll := make([]float64, 0)
	from_ll = append(from_ll, lang_from)
	from_ll = append(from_ll, lat_from)
	from_ll_arr := make([][]float64, 0)
	from_ll_arr = append(from_ll_arr, from_ll)

	// from_ll := [1][2]float64{{lang_from, lat_from}}

	to_ll := make([]float64, 0)
	to_ll = append(to_ll, long_to)
	to_ll = append(to_ll, lat_to)

	ll := LatLang{From: from_ll_arr, To: to_ll}

	ll_data := make([]LatLang, 0)
	ll_data = append(ll_data, ll)

	v := MSlice{Type: action, Data: ll_data}
	msg = append(msg, v)

	Push(msg)
}

func PushTextLeave(val string) {
	msg := make(MSliceMap, 0)
	v := MSlice{Type: "leave", Data: val}
	msg = append(msg, v)

	Push(msg)
}

func PushTextLatLang(action string, val string) {
	msg := make(MSliceMap, 0)
	v := MSlice{Type: action, Data: val}
	msg = append(msg, v)

	Push(msg)
}

func Push(msg MSliceMap) {
	// fmt.Println(msg)
	// ValChan <- msg
}

func Receive() {

	for {
		data := <-ValChan
		fmt.Println("queue", data)
	}

}
