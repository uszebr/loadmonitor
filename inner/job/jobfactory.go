package job

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

// creating jobs, workers
// start,stop jobs creating
// singleton
// TODO: mutex create when dynamic parameters change
type JobFactory struct {
	initialWorkersQuantity int64         //initial maximum of sumultanius gorutines executing jobs
	currentWorkesQuantity  int64         // we can change it dynamicly
	workerSemaphore        chan struct{} //channel to limit workers quantity
	jobsChan               chan *Job
	jobsBuffer             int64
	jobsCounter            atomic.Int64
	//TODO: counter finihed, unfinished??
	initialJobComplexity int64
	currentComplexity    int64

	finishedJobsChan chan *Job

	mu sync.Mutex
}

var (
	instance *JobFactory
	once     sync.Once
)

// GetInstance returns the singleton instance of JobFactory.
// TODO: refactor to use options func
func GetInstance(initialWorkersQuantity int64, jobsBuffer int64, initialcomplexity int64) *JobFactory {
	curWorkerQuant := atomic.Int64{}
	curWorkerQuant.Store(initialWorkersQuantity)
	once.Do(func() {
		instance = &JobFactory{
			initialWorkersQuantity: initialWorkersQuantity,
			currentWorkesQuantity:  initialWorkersQuantity,
			workerSemaphore:        make(chan struct{}, initialWorkersQuantity),
			jobsBuffer:             jobsBuffer,
			jobsChan:               make(chan *Job, jobsBuffer),
			initialJobComplexity:   initialcomplexity,
			currentComplexity:      initialcomplexity, //we can change it dynamically
			// TODO: think how many do I need..
			// TODO: create entity/method that collects all and store (only last 50?? 100 to the structure)
			finishedJobsChan: make(chan *Job, 1000),
		}
	})
	return instance
}
func (jf *JobFactory) StartFactory(ctx context.Context) {
	fmt.Println("Starting Factory..")
	jf.startJobsProducing(ctx)
	//starting collector before producer
	jf.startCollectingFinishedJobs()
	jf.startJobsExecuting(ctx)
}

// StartJobsProducing starting jobs producer into jobs channel
func (jf *JobFactory) startJobsProducing(ctx context.Context) {
	fmt.Println("+Jobs Producing started..")
	go func() {
		defer close(jf.jobsChan)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("+Jobs Producing stopped")
				return
			default:
				jf.mu.Lock()
				curComplexity := jf.currentComplexity
				jf.mu.Unlock()
				jf.jobsChan <- NewJob(curComplexity)
			}
		}
	}()
}

func (jf *JobFactory) startJobsExecuting(ctx context.Context) {
	fmt.Println("=Jobs Executing started..")

	go func() {
		wgWorkers := &sync.WaitGroup{}
		defer fmt.Println("=Jobs Executing ended..")
		defer func() {
			wgWorkers.Wait()
			close(jf.finishedJobsChan)
		}()

		for {
			select {
			case job, ok := <-jf.jobsChan:
				if !ok {
					return
				}
				jf.workerSemaphore <- struct{}{}
				jf.jobsCounter.Add(1)
				wgWorkers.Add(1)
				go func() { //worker
					defer func() {
						<-jf.workerSemaphore       // realising semaphore
						jf.finishedJobsChan <- job // collecting finished jobs
					}()
					defer wgWorkers.Done()
					job.StartJob(ctx)
				}()
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (jf *JobFactory) startCollectingFinishedJobs() {
	fmt.Println("--collecting finished jobs started")
	go func() {
		defer fmt.Println("--collecting finished jobs ended")
		for job := range jf.finishedJobsChan {
			duration, err := job.JobDuration()
			if err != nil {
				fmt.Printf("-----error-duration---JOB\t[%v] err: [%v]\n", job.Id, err)
			}
			fmt.Printf("-------------JOB\t[%v] duration: [%v]\n", job.Id, duration)
		}
	}()
}

// SetComplexity changing complexity for all newly created jobs
func (jf *JobFactory) SetComplexity(newComplexity int64) {
	jf.mu.Lock()
	defer jf.mu.Unlock()
	jf.currentComplexity = newComplexity
}

// TODO :THINK if it is safe to change channel??
// SetWorkersQuantity changing quantity of workers gorutines, changing semaphore channel to the new one with updated size
func (jf *JobFactory) SetWorkersQuantity(newWorkersQuantity int64) {
	jf.mu.Lock()
	defer jf.mu.Unlock()
	if newWorkersQuantity < 1 {
		jf.currentWorkesQuantity = 1
	} else {
		jf.currentWorkesQuantity = newWorkersQuantity
	}
	// TODO: close old semaphore channel??? graceful shot down for all inner jobs??
	jf.workerSemaphore = make(chan struct{}, jf.currentWorkesQuantity)
}
