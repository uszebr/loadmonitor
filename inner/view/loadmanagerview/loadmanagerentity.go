package loadmanagerview

import (
	"github.com/uszebr/loadmonitor/inner/domain/jobproducer"
	"github.com/uszebr/loadmonitor/inner/domain/workerpool"
)

type JobProducerFormData struct {
	*jobproducer.JobProducer
	Success         bool
	ErrorComplexity string
	ErrorMemoryLoad string
}

type WorkerPoolFormData struct {
	*workerpool.WorkerPool
	Success             bool
	ErrorWorkerQuantity string
}
