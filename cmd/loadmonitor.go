package main

import (
	// "context"

	"context"
	"fmt"

	// "time"
	"github.com/gin-gonic/gin"
	"github.com/uszebr/loadmonitor/inner/domain/collector"
	"github.com/uszebr/loadmonitor/inner/domain/jobproducer"
	"github.com/uszebr/loadmonitor/inner/domain/workerpool"
	"github.com/uszebr/loadmonitor/inner/handler/jobmonitorhandl"
	"github.com/uszebr/loadmonitor/inner/handler/loadmanagerhandl"
	"github.com/uszebr/loadmonitor/inner/handler/runtimehandl"
	// "github.com/uszebr/loadmonitor/inner/domain/workerpool"
)

func main() {
	fmt.Println("!!!!!!!!!!!!!!!Load Monitor started..")

	ctx, cancel := context.WithCancel(context.Background())

	jProducer := jobproducer.New(jobproducer.WithJobComplexity(50), jobproducer.WithMemoryLoad(10))
	jobQueue := jProducer.Start(ctx)

	wPool, proccessedJobx := workerpool.NewWorkerPool(ctx, 3, jobQueue)

	// go func() {
	// 	for job := range proccessedJobx {
	// 		fmt.Printf("D[%s] Complex: [%v] MemoryLoad: [%v] Status: [%v] Duration: [%v] \n", job.Id().String(), job.Complexity(), job.MemoryLoad(), job.Status(), job.JobDuration())
	// 	}
	// }()

	collector := collector.NewCollector(20)
	collector.StartCollecting(proccessedJobx)

	engine := gin.Default()
	engine.Static("/assets", "./assets")
	loadManagerHandler := loadmanagerhandl.New(jProducer, wPool)
	//page to change load/workers
	engine.GET("/loadmanager", loadManagerHandler.HandlePage)

	// endpoint to change Load, Memory Load
	engine.POST("/loadmanager-producer", loadManagerHandler.HandleProducer)
	// endpoint to change Workers quantity
	engine.POST("/loadmanager-workers", loadManagerHandler.HandleWorkers)

	// page to monitor collector of last finished jobs, jobs quantity and accumulated complexity for finished jobs
	jmonitor := jobmonitorhandl.New(collector)
	engine.GET("/jobmonitor", jmonitor.HandlePage)
	// endpoint with job collector updates
	engine.POST("/jobmonitor", jmonitor.HandlePost)

	rthandler := runtimehandl.New()
	engine.GET("/runtimedata", rthandler.HandlePage)
	engine.Run(":8085")

	cancel() // TODO add to graceful shutdown

	//TODO: MAIN
	// DONE // refactor to JobI interface	// with all public methods??
	// config with default/start options
	// config path in env for prod/dev
	// logger with prod/dev
	// Multiplicator for job complexity store in config
}
