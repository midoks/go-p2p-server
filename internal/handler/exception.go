package handler

import (
	"github.com/midoks/go-p2p-server/internal/client"
)

type ExceptionHandler struct {
	Msg *SignalMsg
	Cli *client.Client
}

// handle {}
func (s *ExceptionHandler) Handle() {
	s.Cli.UpdateTs()
}
