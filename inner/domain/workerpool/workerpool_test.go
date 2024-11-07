package workerpool

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/uszebr/loadmonitor/inner/domain/job"
)

type MockJob struct {
	mock.Mock
}

func (m *MockJob) Do(ctx context.Context) {
	m.Called(ctx)
}

func TestNewWorkerPool(t *testing.T) {
	ctx := context.Background()
	jobChan := make(chan *job.Job, 10)
	workerCount := 3

	// Create worker pool
	pool, jp := NewWorkerPool(ctx, workerCount, jobChan)

	// Verify WorkerPool initialization
	assert.NotNil(t, pool)
	assert.Equal(t, workerCount, pool.Workers())
	assert.NotNil(t, jp)
	assert.Len(t, pool.quit, 0) // Ensure quit channel is empty
}

func TestSetWorkerCount(t *testing.T) {
	ctx := context.Background()
	jobChan := make(chan *job.Job, 10)
	workerCount := 3

	// Create worker pool
	pool, _ := NewWorkerPool(ctx, workerCount, jobChan)

	// Ensure the initial worker count is correct
	assert.Equal(t, workerCount, pool.Workers())

	// Increase worker count
	pool.SetWorkerCount(5)
	assert.Equal(t, 5, pool.Workers())

	// Decrease worker count
	pool.SetWorkerCount(2)
	assert.Equal(t, 2, pool.Workers())
}
