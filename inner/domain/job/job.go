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
	COMPLEXITY_INTERACTIONS_MULTIPLIER_DEFAULT = 1_000_000
	COMPLEXITY_PARTICULAR_VALUE_MAX_DEFAULT    = 10
)

// used to muliply complexity
// default value(may reset in config)
var complexityMultiplier ComplexityMultiplier = COMPLEXITY_INTERACTIONS_MULTIPLIER_DEFAULT

type ComplexityMultiplier int

func SetComplexityMultiplier(multiplier int) {
	if multiplier < 0 {
		panic("Multiplier can not be less then 0")
	}
	complexityMultiplier = ComplexityMultiplier(multiplier)
}

// used for comlexity in each iteration
// default value(may reset in config)
var multiplyValue MultiplyValue = COMPLEXITY_PARTICULAR_VALUE_MAX_DEFAULT

type MultiplyValue int

func SetMultiplyValue(value int) {
	if value < 0 {
		panic("Multiply Value can not be less then 0")
	}
	multiplyValue = MultiplyValue(value)
}

// stored here because a lot of consumers of this interface
type JobI interface {
	Do(ctx context.Context)
	Id() uuid.UUID
	ComplexityInt() int64
	MemoryLoadInt() int64
	Start() time.Time
	End() time.Time
	JobDuration() time.Duration
	Status() JobStatus
	Result() int64
}

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

// Id getter
func (j *Job) Id() uuid.UUID {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.id
}

// Complexity getter
func (j *Job) Complexity() complexity {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.complexity
}

// Complexity getter
func (j *Job) ComplexityInt() int64 {
	return int64(j.Complexity())
}

// Start getter for start time
func (j *Job) Start() time.Time {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.start
}

// setStart private setter
func (j *Job) setStart(t time.Time) {
	j.mu.Lock()
	defer j.mu.Unlock()
	j.start = t
}

// End getter for end time
func (j *Job) End() time.Time {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.end
}

// setEnd private setter
func (j *Job) setEnd(t time.Time) {
	j.mu.Lock()
	defer j.mu.Unlock()
	j.end = t
}

// Status getter
func (j *Job) Status() JobStatus {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.status
}

// setStatus private setter
func (j *Job) setStatus(js JobStatus) {
	j.mu.Lock()
	defer j.mu.Unlock()
	j.status = js
}

// Result getter
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

func (j *Job) MemoryLoadInt() int64 {
	return int64(j.MemoryLoad())
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
	iterations := int(j.complexity) * int(complexityMultiplier)
	result := int64(1)
	for i := 0; i < iterations; i++ {
		select {
		case <-ctx.Done():
			j.setResult(0)
			j.setStatus(CANCELED)
			return
		default:
			result += int64(rand.Intn(int(multiplyValue)) * rand.Intn(int(multiplyValue)))
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
	j.memoryLoader = []byte(nil)
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
