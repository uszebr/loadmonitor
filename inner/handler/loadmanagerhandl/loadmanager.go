package loadmanagerhandl

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uszebr/loadmonitor/inner/domain/jobproducer"
	"github.com/uszebr/loadmonitor/inner/domain/workerpool"
	"github.com/uszebr/loadmonitor/inner/util/ginutil"
	"github.com/uszebr/loadmonitor/inner/view/loadmanagerview"
)

type LoadManagerHandler struct {
	jobProducer *jobproducer.JobProducer
	workerPool  *workerpool.WorkerPool
}

// TODO refactor to options when all parameters are known
func New(jp *jobproducer.JobProducer, wp *workerpool.WorkerPool) LoadManagerHandler {
	return LoadManagerHandler{jobProducer: jp, workerPool: wp}
}

func (h *LoadManagerHandler) HandlePage(c *gin.Context) {
	jobProducerFormData := loadmanagerview.JobProducerFormData{JobProducer: h.jobProducer, Success: true}
	_ = ginutil.Render(c, 200, loadmanagerview.LoadManagerPage(jobProducerFormData, h.workerPool))
	// TODO: log err here
}

func (h *LoadManagerHandler) HandleProducer(c *gin.Context) {

	// TODO: DELETE
	time.Sleep(3 * time.Second)
	// TODO: DELETE^^^
	jobProducerFormData := loadmanagerview.JobProducerFormData{JobProducer: h.jobProducer, Success: true, ErrorComplexity: "asdfasdf", ErrorMemoryLoad: "qrwerqew"}
	_ = ginutil.Render(c, 200, loadmanagerview.ProducerForm(jobProducerFormData))
	// TODO: log err here
}
