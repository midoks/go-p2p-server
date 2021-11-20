package handler

import (
	"encoding/json"
	"github.com/midoks/go-p2p-server/internal/client"
)

type Handler interface {
	Handle()
}

type SignalMsg struct {
	Action   string      `json:"action"`
	ToPeerId string      `json:"to_peer_id,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Reason   string      `json:"reason,omitempty"`
}

type SignalResp struct {
	Action     string      `json:"action"`
	FromPeerId string      `json:"from_peer_id,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Reason     string      `json:"reason,omitempty"`
}

func NewHandler(message []byte, cli *client.Client) (Handler, error) {
	signal := SignalMsg{}
	if err := json.Unmarshal(message, &signal); err != nil {
		return nil, err
	}
	return NewHandlerMsg(signal, cli)
}

func NewHandlerMsg(signal SignalMsg, cli *client.Client) (Handler, error) {
	switch signal.Action {
	case "signal":
		return &SignalHandler{Msg: &signal, Cli: cli}, nil
	case "ping":
		return &HeartbeatHandler{Cli: cli}, nil
	case "tranx":
		return &TranxHandler{Msg: &signal, Cli: cli}, nil
	case "reject":
		return &RejectHandler{Msg: &signal, Cli: cli}, nil
	case "get_stat":
		return &StatsHandler{Msg: &signal, Cli: cli}, nil
	default:
		return &ExceptionHandler{Msg: &signal, Cli: cli}, nil
	}
}
