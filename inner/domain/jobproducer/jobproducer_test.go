package jobproducer

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJobProducer_New(t *testing.T) {
	// Test case for initializing a new JobProducer
	jp := New(10, 1024)

	// Assert that the complexity and memory load are set correctly
	assert.Equal(t, int64(10), jp.JobComplexity(), "Job complexity should be 10")
	assert.Equal(t, int64(1024), jp.JobMemoryLoad(), "Job memory load should be 1024")
}

func TestJobProducer_SetComplexity(t *testing.T) {
	// Test case for setting the job complexity
	jp := New(10, 1024)

	// Set new complexity
	jp.SetComplexity(20)

	// Assert that the complexity is updated correctly
	assert.Equal(t, int64(20), jp.JobComplexity(), "Job complexity should be updated to 20")
}

func TestJobProducer_SetMemoryLoad(t *testing.T) {
	// Test case for setting the job memory load
	jp := New(10, 1024)

	// Set new memory load
	jp.SetMemoryLoad(2048)

	// Assert that the memory load is updated correctly
	assert.Equal(t, int64(2048), jp.JobMemoryLoad(), "Job memory load should be updated to 2048")
}

func TestJobProducer_JobComplexitySt(t *testing.T) {
	// Test case for getting job complexity as a string
	jp := New(10, 1024)

	// Assert the JobComplexitySt returns the correct string representation
	assert.Equal(t, "10", jp.JobComplexitySt(), "Job complexity as string should be '10'")
}

func TestJobProducer_JobMemoryLoadSt(t *testing.T) {
	// Test case for getting job memory load as a string
	jp := New(10, 1024)

	// Assert the JobMemoryLoadSt returns the correct string representation
	assert.Equal(t, "1024", jp.JobMemoryLoadSt(), "Job memory load as string should be '1024'")
}

func TestJobProducer_Start(t *testing.T) {
	// Test case for starting the JobProducer and sending jobs via channel

	// Create a new JobProducer
	jp := New(10, 1024)

	// Create a context with cancellation
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	// Start the JobProducer and get the job channel
	jobChan := jp.Start(ctx)

	// Expect to receive one job from the channel
	select {
	case job := <-jobChan:
		// Assert that the job has the expected complexity and memory load
		assert.Equal(t, int64(10), job.ComplexityInt(), "Job complexity should be 10")
		assert.Equal(t, int64(1024), job.MemoryLoadInt(), "Job memory load should be 1024")
	case <-time.After(time.Millisecond * 150):
		t.Fatal("Expected a job but did not receive one within the timeout")
	}
}

func TestJobProducer_JobProducerStop(t *testing.T) {
	// Test case to ensure the JobProducer stops correctly when context is cancelled

	// Create a new JobProducer
	jp := New(10, 1024)

	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the JobProducer and get the job channel
	jobChan := jp.Start(ctx)

	// Cancel the context after 50ms to stop job production
	time.Sleep(time.Millisecond * 50)
	cancel()

	// Ensure that the job channel is closed and no jobs are received after cancellation
	select {
	case job, ok := <-jobChan:
		if ok {
			t.Fatal("Expected no job, but received a job:", job)
		}
	case <-time.After(time.Millisecond * 100):
		// Successful test if no job is received after context cancellation
	}
}
