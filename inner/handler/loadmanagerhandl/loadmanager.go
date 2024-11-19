package loadmanagerhandl

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/uszebr/loadmonitor/inner/domain/jobproducer"
	"github.com/uszebr/loadmonitor/inner/domain/workerpool"
	"github.com/uszebr/loadmonitor/inner/logger"
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

func (h LoadManagerHandler) HandlePage(c *gin.Context) {
	// success set to false. To prevent showing Fade-Out icon when request is done(not implemented yet)
	jobProducerFormData := loadmanagerview.JobProducerFormData{JobProducer: h.jobProducer, Success: false}
	workerPoolFormData := loadmanagerview.WorkerPoolFormData{WorkerPool: h.workerPool, Success: false}
	err := ginutil.Render(c, 200, loadmanagerview.LoadManagerPage(jobProducerFormData, workerPoolFormData))
	if err != nil {
		slog.Error("Error Rendering", logger.Err(err))
	}
}

func (h LoadManagerHandler) HandleProducer(c *gin.Context) {
	complexityForm := c.PostForm("complexity")
	memoryLoadForm := c.PostForm("memory-load")
	comlexity, memoryLoad, errComplexity, errMemoryLoad := validateProducerFormValues(complexityForm, memoryLoadForm)
	if errComplexity != "" || errMemoryLoad != "" {
		jobProducerFormDataerr := loadmanagerview.JobProducerFormData{JobProducer: h.jobProducer, Success: false, ErrorComplexity: errComplexity, ErrorMemoryLoad: errMemoryLoad}
		_ = ginutil.Render(c, 200, loadmanagerview.ProducerForm(jobProducerFormDataerr))
		return
	}
	h.jobProducer.SetComplexity(int64(comlexity))
	h.jobProducer.SetMemoryLoad(int64(memoryLoad))
	jobProducerFormData := loadmanagerview.JobProducerFormData{JobProducer: h.jobProducer, Success: true, ErrorComplexity: "", ErrorMemoryLoad: ""}
	err := ginutil.Render(c, 200, loadmanagerview.ProducerForm(jobProducerFormData))
	if err != nil {
		slog.Error("Error Rendering", logger.Err(err))
	}
}

// validateProducerFormValues return issues(main or secondary) for all form inputs. Client can fix both
func validateProducerFormValues(complexitySt, memoryLoadSt string) (int, int, string, string) {
	errComplexity := ""
	errMemoryLoad := ""
	comlexity, err := strconv.Atoi(complexitySt)
	if err != nil {
		//first priority issue
		errComplexity = fmt.Sprintf("complexity issue: %v", err.Error())
	} else {
		if comlexity < 0 {
			comlexity = 0
			errComplexity = "complexity should be >= 0"
		}
	}

	memoryLoad, err := strconv.Atoi(memoryLoadSt)
	if err != nil {
		//first priority issue
		errMemoryLoad = fmt.Sprintf("memory load issue: %v", err.Error())
	} else {
		if memoryLoad < 0 {
			memoryLoad = 0
			errMemoryLoad = "memory load should be >= 0"
		}
	}
	return comlexity, memoryLoad, errComplexity, errMemoryLoad
}

func (h LoadManagerHandler) HandleWorkers(c *gin.Context) {
	workersForm := c.PostForm("workers")
	workers, errWorkers := validateWorkersFormValues(workersForm)
	if errWorkers != "" {
		workerPoolFormData := loadmanagerview.WorkerPoolFormData{WorkerPool: h.workerPool, Success: false, ErrorWorkerQuantity: errWorkers}
		err := ginutil.Render(c, 200, loadmanagerview.WorkerForm(workerPoolFormData))
		if err != nil {
			slog.Error("Error Rendering", logger.Err(err))
		}
		return
	}
	h.workerPool.SetWorkerCount(workers)
	workerPoolFormData := loadmanagerview.WorkerPoolFormData{WorkerPool: h.workerPool, Success: true, ErrorWorkerQuantity: ""}
	err := ginutil.Render(c, 200, loadmanagerview.WorkerForm(workerPoolFormData))
	if err != nil {
		slog.Error("Error Rendering", logger.Err(err))
	}
}

// validateWorkersFormValues return string instead of error just to consitancy(follow pattern in previous validator)
func validateWorkersFormValues(workersSt string) (int, string) {
	errWorkers := ""

	workers, err := strconv.Atoi(workersSt)
	if err != nil {
		//first priority issue
		errWorkers = fmt.Sprintf("workers quantity: %v", err.Error())
	} else {
		if workers < 0 {
			workers = 0
			errWorkers = "complexity should be >= 0"
		}
	}
	return workers, errWorkers
}
