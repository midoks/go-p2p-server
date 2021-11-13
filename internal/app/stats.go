package app

import (
	// "fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"

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

	//mem
	memInfo, _ := mem.VirtualMemory()

	c.JSON(http.StatusOK, gin.H{
		"peers":          hub.GetClientNum(),
		"load_avg_per":   info.Load1 / max,
		"cpu_per":        percent[0],
		"mem_per":        memInfo.UsedPercent,
		"goroutine_num":  runtime.NumGoroutine(),
		"server_runtime": time.Now().Format("2006-01-02 15:04:05"),
	})
}
