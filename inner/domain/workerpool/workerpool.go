package workerpool

import (
	"context"
	"fmt"
	"sync"

	"github.com/uszebr/loadmonitor/inner/domain/job"
)

type WorkerPool struct {
	jobQueue      <-chan job.JobI
	jobProccessed chan job.JobI
	workers       int
	quit          chan bool
	wg            sync.WaitGroup
	mu            sync.RWMutex
	ctx           context.Context
}

func NewWorkerPool(ctx context.Context, workerCount int, jobChan <-chan job.JobI) (*WorkerPool, <-chan job.JobI) {
	jp := make(chan job.JobI)
	pool := &WorkerPool{
		jobQueue:      jobChan,
		jobProccessed: jp,
		quit:          make(chan bool),
		ctx:           ctx,
	}
	pool.SetWorkerCount(workerCount)
	return pool, jp
}

// worker that consumes jobs
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
func (p *WorkerPool) SetWorkerCount(count int) {
	currentWorkers := p.Workers()
	p.mu.Lock()
	defer p.mu.Unlock()
	if count > currentWorkers {
		for i := 0; i < count-currentWorkers; i++ {
			p.wg.Add(1)
			go func() {
				defer p.wg.Done()
				p.worker(p.ctx)
			}()
		}
	} else if count < currentWorkers {
		for i := 0; i < currentWorkers-count; i++ {
			p.quit <- true
		}
	}
	p.workers = count
}

// Workers getter for workers quantity
func (p *WorkerPool) Workers() int {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.workers
}

// WorkersSt getter for workers quantity in string
func (p *WorkerPool) WorkersSt() string {
	return fmt.Sprintf("%d", p.Workers())
}

// Wait for all workers to finish
func (p *WorkerPool) Wait() {
	p.wg.Wait()
	close(p.jobProccessed)
}
