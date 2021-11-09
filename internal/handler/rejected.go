package handler

import (
	"cbsignal/hub"
	"github.com/midoks/go-p2p-server/internal/client"
)

type RejectHandler struct {
	Msg *SignalMsg
	Cli *client.Client
}

func (s *RejectHandler) Handle() {
	//h := hub.GetInstance()
	//判断节点是否还在线
	if target, ok := hub.GetClient(s.Msg.ToPeerId); ok {
		resp := SignalResp{
			Action:     "reject",
			FromPeerId: s.Cli.PeerId,
			Reason:     s.Msg.Reason,
		}
		hub.SendJsonToClient(target, resp)
	}
}
