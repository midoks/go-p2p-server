package handler

import (
	"github.com/midoks/go-p2p-server/internal/client"
	"github.com/midoks/go-p2p-server/internal/hub"
	"github.com/midoks/go-p2p-server/internal/queue"
)

type TranxHandler struct {
	Msg *SignalMsg
	Cli *client.Client
}

func (s *TranxHandler) Handle() {

	toPeerId := s.Msg.ToPeerId

	if c, ok := hub.GetClient(toPeerId); ok {
		lat := s.Cli.Latitude
		long := s.Cli.Longitude

		to_lat := c.Latitude
		to_long := c.Longitude

		queue.PushConnection("connection", lat, long, to_lat, to_long)
	}

}
