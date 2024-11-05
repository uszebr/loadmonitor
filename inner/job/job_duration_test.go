package job

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJobDuration_Finished(t *testing.T) {
	job := NewJob(5)

	job.Start = time.Now()
	time.Sleep(10 * time.Millisecond) // Simulate some work
	job.End = time.Now()
	job.Finished = true

	// Get the duration
	duration, err := job.JobDuration()
	assert.NoError(t, err)
	assert.Greater(t, duration, time.Duration(1))
}

func TestJobDuration_NotFinished(t *testing.T) {
	job := NewJob(5)

	job.Start = time.Now()
	job.Finished = false // Job is not finished

	duration, err := job.JobDuration()
	assert.Equal(t, NotFinished, err)           // Assert that the error is NotFinished
	assert.Equal(t, time.Duration(0), duration) // Assert that the duration is 0
}

func TestJobDuration_ImmediateJob(t *testing.T) {
	job := NewJob(5)

	job.Start = time.Now()
	job.End = job.Start
	job.Finished = true

	duration, err := job.JobDuration()
	assert.NoError(t, err)
	assert.Equal(t, time.Duration(0), duration)
}
