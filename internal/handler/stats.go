package handler

import (
	"github.com/midoks/go-p2p-server/internal/client"
	"github.com/midoks/go-p2p-server/internal/queue"
)

type StatsHandler struct {
	Msg *SignalMsg
	Cli *client.Client
}

func (s *StatsHandler) Handle() {
	queue.RegisterPeer(s.Cli.PeerId)
}
