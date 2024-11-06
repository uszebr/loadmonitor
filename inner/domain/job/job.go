package job

import (
	"cmp"
	"context"
	"sync"

	"time"

	"math/rand"

	"github.com/google/uuid"
)

const (
	// used to muliply complexity
	COMPLEXITY_INTERACTIONS_MULTIPLIER = 1_000_000
	// used for comlexity in each iteration
	COMPLEXITY_PARTICULAR_VALUE_MAX = 10
)

type Job struct {
	id         uuid.UUID
	complexity complexity // represent how hard/complex job to execute.
	start      time.Time
	end        time.Time
	status     JobStatus
	result     int64 //for storing job result, 0 if canceld.(to show on frontend dynamic results calculated)
	mu         sync.RWMutex
}

// NewJob creating pointer to the new Job with given complexity
func NewJob(complexity int64) *Job {
	return &Job{id: uuid.New(), complexity: newComplexity(complexity), status: NEW}
}
func (j *Job) Id() uuid.UUID {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.id
}

func (j *Job) Complexity() complexity {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.complexity
}

func (j *Job) Start() time.Time {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.start
}

func (j *Job) SetStart(t time.Time) {
	j.mu.Lock()
	defer j.mu.Unlock()
	j.start = t
}

func (j *Job) End() time.Time {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.end
}

func (j *Job) SetEnd(t time.Time) {
	j.mu.Lock()
	defer j.mu.Unlock()
	j.end = t
}

func (j *Job) Status() JobStatus {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.status
}

func (j *Job) SetStatus(js JobStatus) {
	j.mu.Lock()
	defer j.mu.Unlock()
	j.status = js
}

func (j *Job) Result() int64 {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.result
}

func (j *Job) SetResult(res int64) {
	j.mu.Lock()
	defer j.mu.Unlock()
	j.result = res
}

// Start starting particular job
func (j *Job) Do(ctx context.Context) {
	defer func() { j.SetEnd(time.Now()) }()
	j.SetStart(time.Now())
	if j.Status() != NEW {
		panic("job should be NEW to be started")
	}
	j.SetStatus(STARTED)
	iterations := int(j.complexity) * COMPLEXITY_INTERACTIONS_MULTIPLIER
	result := int64(1)
	for i := 0; i < iterations; i++ {
		select {
		case <-ctx.Done():
			j.result = 0
			j.status = CANCELED
			return
		default:
			result += int64(rand.Intn(COMPLEXITY_PARTICULAR_VALUE_MAX) * rand.Intn(COMPLEXITY_PARTICULAR_VALUE_MAX))
		}
	}
	j.SetResult(cmp.Or(result, 1))
	j.SetStatus(FINISHED)
}

func (j *Job) JobDuration() time.Duration {
	if j.Status() == NEW || j.Status() == STARTED {
		return 0
	}
	return j.End().Sub(j.Start())
}

// abstracting complexity
type complexity int64

// newComplexity creating with validation logic
func newComplexity(initial int64) complexity {
	if initial < 0 {
		return complexity(0)
	}
	return complexity(initial)
}

type JobStatus string

const (
	NEW      JobStatus = "New"
	STARTED  JobStatus = "Started"
	CANCELED JobStatus = "Canceled"
	FINISHED JobStatus = "Finished"
)
