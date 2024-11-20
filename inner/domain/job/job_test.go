package job

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewJob(t *testing.T) {
	// Define table of test cases with different complexity and memory load values
	tests := []struct {
		name               string
		complexity         int64
		memoryLoad         int64
		expectedComplexity complexity
		expectedMemoryLoad complexity
	}{
		{
			name:               "Moderate complexity and memory load",
			complexity:         5,
			memoryLoad:         1024,
			expectedComplexity: complexity(5),
			expectedMemoryLoad: complexity(1024),
		},
		{
			name:               "Minimum complexity and memory load",
			complexity:         0,
			memoryLoad:         0,
			expectedComplexity: complexity(0),
			expectedMemoryLoad: complexity(0),
		},
		{
			name:               "High complexity and memory load",
			complexity:         100,
			memoryLoad:         2048,
			expectedComplexity: complexity(100),
			expectedMemoryLoad: complexity(2048),
		},
		{
			name:               "Large complexity and memory load",
			complexity:         12345,
			memoryLoad:         987654,
			expectedComplexity: complexity(12345),
			expectedMemoryLoad: complexity(987654),
		},
		{
			name:               "Negative complexity (should be 0)",
			complexity:         -5,
			memoryLoad:         1024,
			expectedComplexity: complexity(0),
			expectedMemoryLoad: complexity(1024),
		},
		{
			name:               "Negative memory load (should be 0)",
			complexity:         5,
			memoryLoad:         -1024,
			expectedComplexity: complexity(5),
			expectedMemoryLoad: complexity(0),
		},
	}

	// Iterate over test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new job with the given complexity and memory load
			job := NewJob(tt.complexity, tt.memoryLoad)

			// Assert that the job is not nil
			assert.NotNil(t, job)

			// Assert the job's status is NEW
			assert.Equal(t, NEW, job.Status())

			// Assert that the job's complexity matches the expected value
			assert.Equal(t, tt.expectedComplexity, job.Complexity())

			// Assert that the job's memory load matches the expected value
			assert.Equal(t, tt.expectedMemoryLoad, job.MemoryLoad())

			// Assert that the memory loader is empty initially
			assert.Empty(t, job.memoryLoader)

			// Assert that the UUID length is correct (36 characters)
			assert.Equal(t, 36, len(job.Id().String()))
		})
	}
}

func TestJobDo_CompletedSuccessfully(t *testing.T) {
	job := NewJob(1, 1024)
	ctx := context.Background()

	job.Do(ctx)
	assert.Equal(t, FINISHED, job.Status())
	assert.Greater(t, job.Result(), int64(1), "Result should be greater than 1 since it was calculated.")
	assert.False(t, job.Start().IsZero(), "Start time should be set")
	assert.False(t, job.End().IsZero(), "End time should be set")
	assert.Empty(t, job.memoryLoader, "memoryLoader should be empty after job completion") // memoryLoader should be cleared
}

func TestJobDo_Canceled(t *testing.T) {
	job := NewJob(1, 1024)
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
	assert.Empty(t, job.memoryLoader, "memoryLoader should be empty after job cancellation") // memoryLoader should be cleared
}

func TestJobDuration(t *testing.T) {
	job := NewJob(1, 1024)
	ctx := context.Background()
	job.Do(ctx)

	duration := job.JobDuration()
	assert.Greater(t, duration, 0*time.Nanosecond, "Duration should be greater than 0 after job is complete")
}

func TestJobDuration_NotFinished(t *testing.T) {
	job := NewJob(1, 1024)

	duration := job.JobDuration()
	assert.Equal(t, 0*time.Nanosecond, duration, "Duration should be 0 for unfinished job")
}

func TestJob_Do_WithInvalidStatus(t *testing.T) {
	job := NewJob(1, 1024)
	job.setStatus(STARTED) // Simulate incorrect initial state

	assert.Panics(t, func() {
		job.Do(context.Background())
	}, "Starting a job that is not in NEW status should panic")
}

func TestMemoryLoadingDuringJob(t *testing.T) {
	job := NewJob(1, 1024)
	ctx := context.Background()

	// Check memory loader is empty initially
	assert.Empty(t, job.memoryLoader, "memoryLoader should be empty before job starts")

	job.Do(ctx)

	// After job is complete, memory loader should be empty (unloaded)
	assert.Empty(t, job.memoryLoader, "memoryLoader should be empty after job execution")
}
