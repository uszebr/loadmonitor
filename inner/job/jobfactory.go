package job

import (
	"context"
	"sync"
	"sync/atomic"
)

// creating jobs, workers
// start,stop jobs creating
// singleton?
type JobFactory struct {
	initialWorkersQuantity int           //initial maximum of sumultanius gorutines executing jobs
	workerSemaphore        chan struct{} //channel to limit workers quantity
	jobsChan               chan *Job
	jobsCounter            atomic.Int64
	//TODO: counter finihed, unfinished??

}

var (
	instance *JobFactory
	once     sync.Once
)

// GetInstance returns the singleton instance of JobFactory.
func GetInstance(initialWorkersQuantity int) *JobFactory {
	once.Do(func() {
		instance = &JobFactory{
			initialWorkersQuantity: initialWorkersQuantity,
		}
	})
	return instance
}

func (jf *JobFactory) StartJobsProducing(ctx context.Context) {
	panic("TODO: implement")
}
