package jobproducer

import (
	"context"
	"sync"

	"github.com/uszebr/loadmonitor/inner/domain/job"
)

type JobProducer struct {
	jobComplexity int64
	mu            sync.Mutex
}

func New(jobComplexity int64) *JobProducer {
	return &JobProducer{jobComplexity: jobComplexity, mu: sync.Mutex{}}
}

func (jp *JobProducer) Start(ctx context.Context) <-chan *job.Job {
	res := make(chan *job.Job)
	go func(jpInner *JobProducer) {
		defer close(res)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				jp.mu.Lock()
				res <- job.NewJob(jpInner.jobComplexity)
				jp.mu.Unlock()
			}
		}
	}(jp)
	return res
}

func (jp *JobProducer) SetComplexity(newComplexity int64) {
	jp.mu.Lock()
	defer jp.mu.Unlock()
	jp.jobComplexity = newComplexity
}
