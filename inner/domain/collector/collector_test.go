package collector

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/uszebr/loadmonitor/inner/domain/job"
)

// Smoke
func TestCollector_AddJobs(t *testing.T) {
	// Test case where we collect less than the desired number of jobs
	c := NewCollector(3)
	jobChan := make(chan *job.Job)

	// Start the collector in a separate goroutine
	c.StartCollecting(jobChan)

	// Use NewJob to create jobs
	jobChan <- job.NewJob(1, 1024) // job with complexity 1 and memory load 1024
	jobChan <- job.NewJob(2, 2048) // job with complexity 2 and memory load 2048
	jobChan <- job.NewJob(3, 4096) // job with complexity 3 and memory load 4096

	// Close the channel
	close(jobChan)

	// Wait for the goroutine to finish
	time.Sleep(time.Millisecond * 10)

	// Assert that the collector has collected exactly 3 jobs
	assert.Equal(t, int64(3), c.Count(), "Job count should be 3")
	assert.Equal(t, int64(6), c.SumOfComplexity(), "Sum of complexities should be 6")

	// Retrieve last jobs and check order
	lastJobs := c.GetLastJobs()
	assert.Len(t, lastJobs, 3, "There should be 3 jobs")
	assert.Equal(t, int64(1), int64(lastJobs[0].Complexity()), "First job complexity should be 1")
	assert.Equal(t, int64(2), int64(lastJobs[1].Complexity()), "Second job complexity should be 2")
	assert.Equal(t, int64(3), int64(lastJobs[2].Complexity()), "Third job complexity should be 3")
}

func TestCollector_OverflowJobs(t *testing.T) {
	// Test case where the number of jobs exceeds the collector's capacity
	c := NewCollector(3)
	jobChan := make(chan *job.Job)

	// Start the collector in a separate goroutine
	c.StartCollecting(jobChan)

	// Add 5 jobs to simulate overflow using NewJob
	jobChan <- job.NewJob(1, 1024)
	jobChan <- job.NewJob(2, 2048)
	jobChan <- job.NewJob(3, 4096)
	jobChan <- job.NewJob(4, 82)
	jobChan <- job.NewJob(5, 0)

	// Close the channel
	close(jobChan)

	// Wait for the goroutine to finish
	time.Sleep(time.Millisecond * 10)

	// Assert that the collector only retains the last 3 jobs
	assert.Equal(t, int64(5), c.Count(), "Job count should be 5")
	assert.Equal(t, int64(15), c.SumOfComplexity(), "Sum of complexities should be 15")

	// Retrieve last jobs and check order (should be 3, 4, 5)
	lastJobs := c.GetLastJobs()
	assert.Len(t, lastJobs, 3, "There should be 3 jobs")
	assert.Equal(t, int64(3), int64(lastJobs[0].Complexity()), "First job complexity should be 3")
	assert.Equal(t, int64(4), int64(lastJobs[1].Complexity()), "Second job complexity should be 4")
	assert.Equal(t, int64(5), int64(lastJobs[2].Complexity()), "Third job complexity should be 5")
}

func TestCollector_GetCountAndCountSt(t *testing.T) {
	// Test case for retrieving count in both int64 and string formats
	c := NewCollector(3)
	jobChan := make(chan *job.Job)

	// Start the collector in a separate goroutine
	go c.StartCollecting(jobChan)

	// Add 2 jobs using NewJob
	jobChan <- job.NewJob(1, 1024)
	jobChan <- job.NewJob(2, 2048)

	// Close the channel
	close(jobChan)

	// Wait for the goroutine to finish
	time.Sleep(time.Millisecond * 10)

	// Assert the job count in int64
	assert.Equal(t, int64(2), c.Count(), "Job count should be 2")

	// Assert the job count in string format
	assert.Equal(t, "2", c.CountSt(), "Job count as string should be '2'")
}

func TestCollector_SumOfComplexitySt(t *testing.T) {
	// Test case for retrieving the sum of complexities as a string
	c := NewCollector(3)
	jobChan := make(chan *job.Job)

	// Start the collector in a separate goroutine
	go c.StartCollecting(jobChan)

	// Add 3 jobs using NewJob
	jobChan <- job.NewJob(1, 1024)
	jobChan <- job.NewJob(2, 2048)
	jobChan <- job.NewJob(3, 4096)

	// Close the channel
	close(jobChan)

	// Wait for the goroutine to finish
	time.Sleep(time.Millisecond * 10)

	// Assert the sum of complexities
	assert.Equal(t, int64(6), c.SumOfComplexity(), "Sum of complexities should be 6")

	// Assert the sum of complexities as a string
	assert.Equal(t, "6", c.SumOfComplexitySt(), "Sum of complexities as string should be '6'")
}
