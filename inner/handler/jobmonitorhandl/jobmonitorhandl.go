package jobmonitorhandl

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/uszebr/loadmonitor/inner/domain/collector"
	"github.com/uszebr/loadmonitor/inner/logger"
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
	err := ginutil.Render(c, 200, jobmonitorview.JobMonitorPage())
	if err != nil {
		slog.Error("Error Rendering", logger.Err(err))
	}
}

func (h JobMonitorHandler) HandlePost(c *gin.Context) {
	err := ginutil.Render(c, 200, jobmonitorview.JobMonitorPost(h.jobCollector))
	if err != nil {
		slog.Error("Error Rendering", logger.Err(err))
	}
}
