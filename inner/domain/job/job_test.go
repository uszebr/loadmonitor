package job

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewJob(t *testing.T) {
	// Test job creation with positive complexity
	job := NewJob(5)

	assert.NotNil(t, job)
	assert.Equal(t, NEW, job.Status())
	assert.Equal(t, complexity(5), job.Complexity())
	//fmt.Println(job.Id)
	assert.Equal(t, 36, len(job.id.String())) // UUID length should be 36 characters
}

func TestJobDo_CompletedSuccessfully(t *testing.T) {
	job := NewJob(1)
	ctx := context.Background()

	job.Do(ctx)

	assert.Equal(t, FINISHED, job.Status())
	assert.Greater(t, job.Result(), int64(1)) // Result should be greater than 1 since it was calculated.
	assert.False(t, job.Start().IsZero(), "Start time should be set")
	assert.False(t, job.End().IsZero(), "End time should be set")
}

func TestJobDo_Canceled(t *testing.T) {
	job := NewJob(1)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(10 * time.Millisecond) // Allow some time for the job to start
		cancel()                          // Cancel the job
	}()

	job.Do(ctx)

	assert.Equal(t, CANCELED, job.Status())
	assert.Equal(t, int64(0), job.Result(), "Result should be 0 if canceled")
	assert.False(t, job.Start().IsZero(), "Start time should be set")
	assert.False(t, job.End().IsZero(), "End time should be set")
}

func TestJobDuration(t *testing.T) {
	job := NewJob(1)
	ctx := context.Background()
	job.Do(ctx)

	duration := job.JobDuration()
	assert.Greater(t, duration, 0*time.Nanosecond, "Duration should be greater than 0 after job is complete")
}

func TestJobDuration_NotFinished(t *testing.T) {
	job := NewJob(1)

	duration := job.JobDuration()
	assert.Equal(t, 0*time.Nanosecond, duration, "Duration should be 0 for unfinished job")
}

func TestJob_Do_WithInvalidStatus(t *testing.T) {
	job := NewJob(1)
	job.status = STARTED // Simulate incorrect initial state

	assert.Panics(t, func() {
		job.Do(context.Background())
	}, "Starting a job that is not in NEW status should panic")
}
