package runtimehandl

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uszebr/loadmonitor/inner/util/ginutil"
	"github.com/uszebr/loadmonitor/inner/view/rtimemonitor"
)

type RunTimeHandl struct{}

func New() RunTimeHandl {
	return RunTimeHandl{}
}

func (h RunTimeHandl) HandlePage(c *gin.Context) {

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	gcSinceLast := float64(0)
	if memStats.LastGC != 0 {
		gcSinceLast = float64(time.Now().UnixNano()-int64(memStats.LastGC)) / 1e9
	}
	rtdata := rtimemonitor.RunTimeData{
		NumCpu:       runtime.NumCPU(),
		NumGorutines: runtime.NumGoroutine(),
		GoOs:         runtime.GOOS,
		GoArch:       runtime.GOARCH,
		Version:      runtime.Version(),
		//Mem Data:
		MemAlloc:      memStats.Alloc,
		MemTotalAlloc: memStats.TotalAlloc,
		MemSys:        memStats.Sys,

		//GC:
		GcCycles:    memStats.NumGC,
		GcSys:       memStats.GCSys,
		GcNext:      memStats.NextGC,
		GcSinceLast: gcSinceLast,
	}
	err := ginutil.Render(c, 200, rtimemonitor.RuntimeMonitorPage(rtdata))
	if err != nil {
		fmt.Printf("Error Rendering: %v\n", err.Error())
	}
}
