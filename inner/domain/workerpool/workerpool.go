package workerpool

import (
	"context"
	"sync"

	"github.com/uszebr/loadmonitor/inner/domain/job"
)

type WorkerPool struct {
	jobQueue      <-chan *job.Job
	jobProccessed chan *job.Job
	workers       int
	quit          chan bool
	wg            sync.WaitGroup
}

func NewWorkerPool(ctx context.Context, workerCount int, jobChan <-chan *job.Job) (*WorkerPool, <-chan *job.Job) {
	jp := make(chan *job.Job)
	pool := &WorkerPool{
		jobQueue:      jobChan,
		jobProccessed: jp,
		quit:          make(chan bool),
	}
	pool.SetWorkerCount(ctx, workerCount)
	return pool, jp
}

// Worker that consumes jobs
func (p *WorkerPool) worker(ctx context.Context) {
	for {
		select {
		case job, ok := <-p.jobQueue:
			if !ok {
				return
			}
			job.Do(ctx) // Execute the job
			p.jobProccessed <- job
		case <-p.quit:
			return
		case <-ctx.Done():
			return
		}
	}
}

// SetWorkerCount dynamically adjusts the number of workers
func (p *WorkerPool) SetWorkerCount(ctx context.Context, count int) {
	currentWorkers := p.workers
	if count > currentWorkers {
		for i := 0; i < count-currentWorkers; i++ {
			p.wg.Add(1)
			go func() {
				defer p.wg.Done()
				p.worker(ctx)
			}()
		}
	} else if count < currentWorkers {
		for i := 0; i < currentWorkers-count; i++ {
			p.quit <- true
		}
	}
	p.workers = count
}

// Wait for all workers to finish
// Close proccessed channel
func (p *WorkerPool) Wait() {
	p.wg.Wait()
	close(p.jobProccessed)
}
