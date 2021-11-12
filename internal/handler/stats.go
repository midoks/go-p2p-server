package handler

import (
	"encoding/json"
	"fmt"
	// "net/http"
	// "runtime"

	"github.com/midoks/go-p2p-server/internal/client"
	"github.com/midoks/go-p2p-server/internal/queue"
)

type StatsHandler struct {
	Msg *SignalMsg
	Cli *client.Client
}

func (s *StatsHandler) Handle() {
	go func() {
		for {

			data := <-queue.ValChan
			fmt.Println("queue", data)

			b, err := json.Marshal(data)
			if err != nil {
				// log.Error("json.Marshal", err)
				fmt.Println("trace json.Marshal", err)
			} else {
				err := s.Cli.SendMessage(b)
				if err != nil {
					break
				}
			}
		}
	}()
}
