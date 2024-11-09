package main

import (
	// "context"

	"context"
	"fmt"

	// "time"
	"github.com/gin-gonic/gin"
	"github.com/uszebr/loadmonitor/inner/domain/jobproducer"
	"github.com/uszebr/loadmonitor/inner/domain/workerpool"
	"github.com/uszebr/loadmonitor/inner/handler/loadmanagerhandl"
	// "github.com/uszebr/loadmonitor/inner/domain/workerpool"
)

func main() {
	fmt.Println("!!!!!!!!!!!!!!!Load Monitor started..")

	ctx, cancel := context.WithCancel(context.Background())

	jProducer := jobproducer.New(50, 100)
	jobQueue := jProducer.Start(ctx)

	wPool, proccessedJobx := workerpool.NewWorkerPool(ctx, 3, jobQueue)

	go func() {
		for job := range proccessedJobx {
			fmt.Printf("D[%s] Complex: [%v] MemoryLoad: [%v] Status: [%v] Duration: [%v] \n", job.Id().String(), job.Complexity(), job.MemoryLoad(), job.Status(), job.JobDuration())
		}
	}()

	// Adjust workers dynamically
	// time.Sleep(6 * time.Second)
	// fmt.Println("Increasing worker count to 5")
	// pool.SetWorkerCount(ctx, 5)

	// time.Sleep(6 * time.Second)
	// fmt.Println("MemoryLoad to 10_000_000)")
	// jp.SetMemoryLoad(10_000_000)

	// time.Sleep(15 * time.Second)
	// fmt.Println("Complexity to 100")
	// jp.SetComplexity(100)

	// time.Sleep(15 * time.Second)
	// fmt.Println("Decreasing worker count to 2")
	// pool.SetWorkerCount(ctx, 2)
	// time.Sleep(20 * time.Second)
	// cancel()
	// pool.Wait()

	engine := gin.Default()
	engine.Static("/assets", "./assets")
	loadManagerHandler := loadmanagerhandl.New(jProducer, wPool)
	engine.GET("/loadmanager", loadManagerHandler.HandlePage)

	engine.POST("/loadmanager-producer", loadManagerHandler.HandleProducer)
	engine.Run(":8085")

	cancel() // TODO add to graceful shutdown

}
