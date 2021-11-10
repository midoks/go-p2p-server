package handler

import (
	"github.com/midoks/go-p2p-server/internal/hub"
	// "github.com/lexkong/log"
	"github.com/midoks/go-p2p-server/internal/client"
)

type HeartbeatHandler struct {
	Cli *client.Client
}

func (s *HeartbeatHandler) Handle() {

	// log.Infof("receive heartbeat from %s", s.Cli.PeerId)
	s.Cli.UpdateTs()

	resp := SignalResp{
		Action: "pong",
	}
	hub.SendJsonToClient(s.Cli, resp)
}
