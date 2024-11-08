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

// idea: TODO: add parameter for complexity of particular value
// idea: TODO: add memory loader parameter... in bytes.. creating array with size, and put each inner result to the array(slice??) but not store this slice(array)into Job.. it shoul load memory only when job is executing
type Job struct {
	id           uuid.UUID
	complexity   complexity // represent how hard/complex job to execute.
	start        time.Time
	end          time.Time
	status       JobStatus
	result       int64      //for storing job result, 0 if canceld.(to show on frontend dynamic results calculated)
	memoryLoad   complexity // represents quantity of bytes job will take during the execution
	memoryLoader []byte
	mu           sync.RWMutex
}

// NewJob creating pointer to the new Job with given complexity
func NewJob(complexity int64, memoryLoadInitial int64) *Job {
	return &Job{id: uuid.New(), complexity: newComplexity(complexity), status: NEW, memoryLoad: newComplexity(memoryLoadInitial)}
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

func (j *Job) setStart(t time.Time) {
	j.mu.Lock()
	defer j.mu.Unlock()
	j.start = t
}

func (j *Job) End() time.Time {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.end
}

func (j *Job) setEnd(t time.Time) {
	j.mu.Lock()
	defer j.mu.Unlock()
	j.end = t
}

func (j *Job) Status() JobStatus {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.status
}

func (j *Job) setStatus(js JobStatus) {
	j.mu.Lock()
	defer j.mu.Unlock()
	j.status = js
}

func (j *Job) Result() int64 {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.result
}

func (j *Job) setResult(res int64) {
	j.mu.Lock()
	defer j.mu.Unlock()
	j.result = res
}

func (j *Job) MemoryLoad() complexity {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.memoryLoad
}

// Start starting particular job
func (j *Job) Do(ctx context.Context) {
	defer func() {
		j.setEnd(time.Now())
		j.unloadMemory() //cleaning memory loader
	}()
	j.setStart(time.Now())
	if j.Status() != NEW {
		panic("job should be NEW to be started")
	}
	j.setStatus(STARTED)
	j.loadMemory()
	iterations := int(j.complexity) * COMPLEXITY_INTERACTIONS_MULTIPLIER
	result := int64(1)
	for i := 0; i < iterations; i++ {
		select {
		case <-ctx.Done():
			j.setResult(0)
			j.setStatus(CANCELED)
			return
		default:
			result += int64(rand.Intn(COMPLEXITY_PARTICULAR_VALUE_MAX) * rand.Intn(COMPLEXITY_PARTICULAR_VALUE_MAX))
		}
	}
	j.setResult(cmp.Or(result, 1))
	j.setStatus(FINISHED)
}

func (j *Job) JobDuration() time.Duration {
	if j.Status() == NEW || j.Status() == STARTED {
		return 0
	}
	return j.End().Sub(j.Start())
}

func (j *Job) loadMemory() {
	j.mu.Lock()
	defer j.mu.Unlock()
	if j.memoryLoad > 0 {
		j.memoryLoader = make([]byte, int(j.memoryLoad))
		j.memoryLoader[rand.Intn(len(j.memoryLoader))] = byte(rand.Uint32())
	}
}

func (j *Job) unloadMemory() {
	j.mu.Lock()
	defer j.mu.Unlock()
	j.memoryLoader = make([]byte, 0)
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
