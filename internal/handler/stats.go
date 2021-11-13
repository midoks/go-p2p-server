package handler

import (
	// "encoding/json"
	// "fmt"
	// "time"

	// "github.com/midoks/go-p2p-server/internal/logger"
	// "github.com/midoks/go-p2p-server/internal/queue"
	"github.com/midoks/go-p2p-server/internal/client"
)

type StatsHandler struct {
	Msg *SignalMsg
	Cli *client.Client
}

func (s *StatsHandler) Handle() {
	// for {

	// select {
	// case data := <-queue.ValChan:
	// 	b, err := json.Marshal(data)
	// 	if err != nil {
	// 		logger.Errorf("stats handler json error: %v", err)
	// 	} else {
	// 		err := s.Cli.SendMessage(b)
	// 		if err != nil {
	// 			s.Cli.Close()
	// 			break
	// 		}
	// 	}
	// case <-time.After(1 * time.Second):
	// 	fmt.Println("oveer...")
	// 	s.Cli.Close()
	// 	break

	// }

	// fmt.Println("Handle", s.Cli.PeerId)
	// break
	// }
}
