package handler

import (
	// "fmt"

	"github.com/midoks/go-p2p-server/internal/client"
	"github.com/midoks/go-p2p-server/internal/mem"
	"github.com/midoks/go-p2p-server/internal/queue"
)

type TranxHandler struct {
	Msg *SignalMsg
	Cli *client.Client
}

func (s *TranxHandler) Handle() {

	toPeerId := s.Msg.ToPeerId

	flat, flng, err := mem.PosGeo(s.Cli.PeerId)
	if err != nil {
		return
	}
	tlat, tlng, err := mem.PosGeo(toPeerId)
	if err != nil {
		return
	}
	queue.PushConnection("connection", flat, flng, tlat, tlng)

}
