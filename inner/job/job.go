package job

import (
	"time"

	"github.com/google/uuid"
)

// one gorutine entity, parallel execution add to the worker level
type Job struct {
	id         uuid.UUID
	complexity complexity // represent how hard/complex job to execute.
	start      time.Time
	end        time.Time
	done       bool
	//TODO: add array/slice to work on with complexity load
}

// NewJob creating pointer to the new Job with given complexity
func NewJob(complexity int) *Job {
	return &Job{id: uuid.New(), complexity: newComplexity(complexity)}
}

// Start starting particular job
func (j *Job) Start() {
	j.start = time.Now()
	execute(j)
	j.done = true
	j.end = time.Now()
}

// execute executing load for particular job depending on complexity
func execute(j *Job) {
	panic("TODO: Implement job execution based on complexity")
}

// abstracting complexity
type complexity int

// newComplexity creating with validation logic
func newComplexity(initial int) complexity {
	if initial < 0 {
		return complexity(0)
	}
	return complexity(initial)
}
