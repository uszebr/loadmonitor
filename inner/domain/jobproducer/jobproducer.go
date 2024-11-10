package jobproducer

import (
	"context"
	"fmt"
	"sync"

	"github.com/uszebr/loadmonitor/inner/domain/job"
)

type JobProducer struct {
	jobComplexity int64
	jobMemoryLoad int64
	mu            sync.RWMutex
}

func New(jobComplexity int64, jobMemoryLoad int64) *JobProducer {
	return &JobProducer{jobComplexity: jobComplexity, jobMemoryLoad: jobMemoryLoad, mu: sync.RWMutex{}}
}

func (jp *JobProducer) Start(ctx context.Context) <-chan job.JobI {
	res := make(chan job.JobI)
	go func() {
		defer close(res)
		for {
			select {
			case <-ctx.Done():
				return
			case res <- job.NewJob(jp.JobComplexity(), jp.JobMemoryLoad()):
				// Job successfully sent to the channel
			}
		}
	}()
	return res
}

func (jp *JobProducer) JobComplexity() int64 {
	jp.mu.RLock()
	defer jp.mu.RUnlock()
	return jp.jobComplexity
}

func (jp *JobProducer) JobMemoryLoad() int64 {
	jp.mu.RLock()
	defer jp.mu.RUnlock()
	return jp.jobMemoryLoad
}

func (jp *JobProducer) JobComplexitySt() string {
	return fmt.Sprintf("%d", jp.JobComplexity())
}

func (jp *JobProducer) JobMemoryLoadSt() string {
	return fmt.Sprintf("%d", jp.JobMemoryLoad())
}

func (jp *JobProducer) SetComplexity(newComplexity int64) {
	jp.mu.Lock()
	defer jp.mu.Unlock()
	jp.jobComplexity = newComplexity
}

func (jp *JobProducer) SetMemoryLoad(newMemoryLoad int64) {
	jp.mu.Lock()
	defer jp.mu.Unlock()
	jp.jobMemoryLoad = newMemoryLoad
}
