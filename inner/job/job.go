package job

import (
	"context"
	"fmt"
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

var (
	NotFinished = fmt.Errorf("job is not finished")
)

type Job struct {
	Id         uuid.UUID
	Complexity complexity // represent how hard/complex job to execute.
	Start      time.Time
	End        time.Time
	Finished   bool
	Result     int //for storing job result, 0 if canceld.(to show on frontend dynamic results calculated)
}

// NewJob creating pointer to the new Job with given complexity
func NewJob(complexity int64) *Job {
	return &Job{Id: uuid.New(), Complexity: newComplexity(complexity), Finished: false}
}

// Start starting particular job
func (j *Job) StartJob(ctx context.Context) {
	j.Start = time.Now()
	j.Result = execute(ctx, j)
	j.Finished = true
	j.End = time.Now()
}

func (j *Job) JobDuration() (time.Duration, error) {
	if !j.Finished {
		return 0, NotFinished
	}
	return j.End.Sub(j.Start), nil
}

// execute executing load for particular job depending on complexity. Returns int result, 0 if canceled
func execute(ctx context.Context, j *Job) int {
	iterations := int(j.Complexity) * COMPLEXITY_INTERACTIONS_MULTIPLIER
	result := 1
	for i := 0; i < iterations; i++ {
		select {
		case <-ctx.Done():
			return 0
		default:
			result += rand.Intn(COMPLEXITY_PARTICULAR_VALUE_MAX) * rand.Intn(COMPLEXITY_PARTICULAR_VALUE_MAX)
		}
	}
	if result == 0 {
		result = 1
	}
	return result
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
