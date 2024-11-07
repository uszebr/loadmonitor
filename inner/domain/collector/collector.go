package collector

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/uszebr/loadmonitor/inner/domain/job"
)

// Collector struct to collect last N *Job pointers from a channel safely with concurrency
type Collector struct {
	jobs            []*job.Job
	quantity        int
	count           atomic.Int64 // atomic to avoid data races
	mu              sync.RWMutex
	sumOfComplexity atomic.Int64 // atomic to avoid mutex on getter (faster)
}

// NewCollector initializes a Collector with a specified quantity
func NewCollector(quantity int) *Collector {
	return &Collector{
		jobs:     make([]*job.Job, 0, quantity),
		quantity: quantity,
	}
}

// StartCollecting begins reading jobs from the provided channel
func (c *Collector) StartCollecting(jobChan <-chan *job.Job) {
	go c.collect(jobChan)
}

// collect continuously reads from the job channel and adds job pointers safely
func (c *Collector) collect(jobChan <-chan *job.Job) {
	for job := range jobChan {
		c.mu.Lock()
		if len(c.jobs) < c.quantity {
			// Append until we reach the desired quantity
			c.jobs = append(c.jobs, job)
		} else {
			// Use circular indexing to overwrite the oldest job
			c.jobs[c.count.Load()%int64(c.quantity)] = job
		}
		fmt.Printf("D[%s] Complex: [%v] MemoryLoad: [%v] Status: [%v] Duration: [%v] \n", job.Id().String(), job.Complexity(), job.MemoryLoad(), job.Status(), job.JobDuration())
		c.mu.Unlock()
		c.count.Add(1)
		c.sumOfComplexity.Add(int64(job.Complexity()))
	}
}

// GetLastJobs returns the last N *Job pointers in order from oldest to newest safely
func (c *Collector) GetLastJobs() []*job.Job {
	c.mu.RLock()
	defer c.mu.RUnlock()

	jobCount := int(c.count.Load())
	if jobCount <= c.quantity {
		// Return jobs in their natural order if fewer than quantity items
		return append([]*job.Job(nil), c.jobs...) // Return a copy to avoid external modification
	}

	// Reorder based on circular buffer's current state
	start := jobCount % c.quantity
	return append([]*job.Job(nil), append(c.jobs[start:], c.jobs[:start]...)...)
}

// SumOfComplexity returns the sum of all completed job complexities
func (c *Collector) SumOfComplexity() int64 {
	return c.sumOfComplexity.Load()
}

// SumOfComplexitySt returns the sum of complexities as a string
func (c *Collector) SumOfComplexitySt() string {
	return fmt.Sprintf("%d", c.SumOfComplexity())
}

// Count returns the total count of jobs processed as int64
func (c *Collector) Count() int64 {
	return c.count.Load()
}

// CountSt returns the total count of jobs processed as a string
func (c *Collector) CountSt() string {
	return fmt.Sprintf("%d", c.Count())
}
