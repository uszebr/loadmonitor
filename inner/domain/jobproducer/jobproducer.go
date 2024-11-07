package jobproducer

import (
	"context"
	"fmt"
	"sync"

	"github.com/uszebr/loadmonitor/inner/domain/job"
)

type JobProducerOption func(*JobProducer)

func WithJobComplexity(jobComplexity int64) JobProducerOption {
	return func(jp *JobProducer) {
		jp.jobComplexity = jobComplexity
	}
}

func WithMemoryLoad(jobMemoryLoad int64) JobProducerOption {
	return func(jp *JobProducer) {
		jp.jobMemoryLoad = jobMemoryLoad
	}
}

type JobProducer struct {
	jobComplexity int64
	jobMemoryLoad int64
	mu            sync.RWMutex
}

func New(options ...JobProducerOption) *JobProducer {
	res := &JobProducer{mu: sync.RWMutex{}}
	for _, option := range options {
		option(res)
	}
	return res
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
