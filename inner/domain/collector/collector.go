package collector

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/uszebr/loadmonitor/inner/domain/job"
)

//idea TODO: implement Listener Notify patern for websocket/SSE/Console output
// Collector struct to collect last N *Job pointers from a channel safely with concurrency
type Collector struct {
	jobs            []job.JobI
	quantity        int
	count           atomic.Int64 // atomic to avoid data races
	mu              sync.RWMutex
	sumOfComplexity atomic.Int64 // atomic to avoid mutex on getter (faster)
}

// NewCollector initializes a Collector with a specified quantity
func NewCollector(quantity int) *Collector {
	return &Collector{
		jobs:     make([]job.JobI, 0, quantity),
		quantity: quantity,
	}
}

// StartCollecting begins reading jobs from the provided channel
func (c *Collector) StartCollecting(jobChan <-chan job.JobI) {
	go c.collect(jobChan)
}

// collect continuously reads from the job channel and adds job pointers safely
func (c *Collector) collect(jobChan <-chan job.JobI) {
	for job := range jobChan {
		c.mu.Lock()
		if len(c.jobs) < c.quantity {
			// Append until we reach the desired quantity
			c.jobs = append(c.jobs, job)
		} else {
			// Use circular indexing to overwrite the oldest job
			c.jobs[c.count.Load()%int64(c.quantity)] = job
		}
		fmt.Printf("D[%s] Complex: [%v] MemoryLoad: [%v] Status: [%v] Duration: [%v] \n", job.Id().String(), job.ComplexityInt(), job.MemoryLoadInt(), job.Status(), job.JobDuration())
		c.mu.Unlock()
		c.count.Add(1)
		c.sumOfComplexity.Add(job.ComplexityInt())
	}
}

// GetLastJobs returns the last N *Job pointers in order from oldest to newest safely
func (c *Collector) GetLastJobs() []job.JobI {
	c.mu.RLock()
	defer c.mu.RUnlock()

	jobCount := int(c.count.Load())
	if jobCount <= c.quantity {
		// Return jobs in their natural order if fewer than quantity items
		// creating shallow copy of jobs
		return append([]job.JobI(nil), c.jobs...) // Return a copy to avoid external modification
	}

	// Reorder based on circular buffer's current state
	start := jobCount % c.quantity
	return append([]job.JobI(nil), append(c.jobs[start:], c.jobs[:start]...)...)
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
