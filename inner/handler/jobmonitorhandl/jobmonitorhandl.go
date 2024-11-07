package jobmonitorhandl

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/uszebr/loadmonitor/inner/domain/collector"
	"github.com/uszebr/loadmonitor/inner/util/ginutil"
	"github.com/uszebr/loadmonitor/inner/view/jobmonitorview"
)

type JobMonitorHandler struct {
	jobCollector *collector.Collector
}

func New(jc *collector.Collector) JobMonitorHandler {
	return JobMonitorHandler{jobCollector: jc}
}

func (h JobMonitorHandler) HandlePage(c *gin.Context) {
	err := ginutil.Render(c, 200, jobmonitorview.JobMonitorPage(h.jobCollector))
	if err != nil {
		fmt.Printf("Error Rendering: %v\n", err.Error())
	}
}
