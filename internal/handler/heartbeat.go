package handler

import (
	"github.com/midoks/go-p2p-server/internal/client"
	"github.com/midoks/go-p2p-server/internal/hub"
	// log "github.com/phachon/go-logger"
)

type HeartbeatHandler struct {
	Cli *client.Client
}

func (s *HeartbeatHandler) Handle() {

	s.Cli.UpdateTs()

	resp := SignalResp{
		Action: "pong",
	}
	hub.SendJsonToClient(s.Cli, resp)
}
