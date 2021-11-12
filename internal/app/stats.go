package app

import (
	// "fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"

	"github.com/midoks/go-p2p-server/internal/hub"
)

func p2pGetStats(c *gin.Context) {

	//cpu
	percent, _ := cpu.Percent(time.Second, false)

	//load
	info, _ := load.Avg()
	var max float64
	if info.Load1 > info.Load5 {
		max = info.Load1
	} else {
		max = info.Load5
	}

	if max < info.Load15 {
		max = info.Load15
	}

	c.JSON(http.StatusOK, gin.H{
		"peers":          hub.GetClientNum(),
		"load_avg_per":   info.Load1 / max,
		"cpu":            percent[0],
		"server_runtime": time.Now().Format("2006-01-02 15:04:05"),
	})
}
