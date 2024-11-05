package job

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewJob(t *testing.T) {
	// Test job creation with positive complexity
	job := NewJob(5)

	assert.NotNil(t, job)
	assert.Equal(t, complexity(5), job.Complexity)
	//fmt.Println(job.Id)
	assert.Equal(t, 36, len(job.Id.String())) // UUID length should be 36 characters
}

func TestJobExecution(t *testing.T) {
	// Test execution of the job
	job := NewJob(1) // Simple job
	ctx := context.Background()
	job.StartJob(ctx)
	fmt.Println(job.Result)
	// Verify that the job was marked as finished
	assert.True(t, job.Finished)
	assert.Greater(t, job.Result, 0) // Result should be greater than 0 if not canceled
}

func TestJobCancellation(t *testing.T) {
	// Test job cancellation
	job := NewJob(10) // More complex job
	ctx, cancel := context.WithCancel(context.Background())

	// Start the job in a separate goroutine
	go func() {
		time.Sleep(50 * time.Millisecond) // Ensure we wait a bit before canceling
		cancel()
	}()

	job.StartJob(ctx)

	// Verify that the job was marked as finished
	assert.True(t, job.Finished)
	assert.Equal(t, 0, job.Result) // Result should be 0 if canceled
}

func TestJobNoWorkDone(t *testing.T) {
	// Test job with complexity that results in no work done
	job := NewJob(0) // Complexity of 0
	ctx := context.Background()
	job.StartJob(ctx)

	// Verify that the job was marked as finished
	assert.True(t, job.Finished)
	assert.Equal(t, 1, job.Result) // Result should be 1 if no work was done
}
