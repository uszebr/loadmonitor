package main

import (
	"context"
	"fmt"
	"time"

	"github.com/uszebr/loadmonitor/inner/domain/jobproducer"
	"github.com/uszebr/loadmonitor/inner/domain/workerpool"
)

func main() {
	fmt.Println("!!!!!!!!!!!!!!!Load Monitor started..")

	ctx, cancel := context.WithCancel(context.Background())

	jp := jobproducer.New(50)
	jobQueue := jp.Start(ctx)

	pool, proccessedJobx := workerpool.NewWorkerPool(ctx, 3, jobQueue)

	go func() {
		for job := range proccessedJobx {
			fmt.Printf("D[%s] Complex: [%v] Status: [%v] Duration: [%v]\n", job.Id().String(), job.Complexity(), job.Status(), job.JobDuration())
		}
	}()

	// Adjust workers dynamically
	time.Sleep(6 * time.Second)
	fmt.Println("Increasing worker count to 5")
	pool.SetWorkerCount(ctx, 5)
	time.Sleep(6 * time.Second)
	fmt.Println("Complexity to 100")
	jp.SetComplexity(100)

	time.Sleep(15 * time.Second)
	fmt.Println("Decreasing worker count to 2")
	pool.SetWorkerCount(ctx, 2)
	time.Sleep(20 * time.Second)
	cancel()
	pool.Wait()

}
