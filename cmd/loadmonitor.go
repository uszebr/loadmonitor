package main

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-gonic/gin"
	"github.com/uszebr/loadmonitor/inner/domain/collector"
	"github.com/uszebr/loadmonitor/inner/domain/jobproducer"
	"github.com/uszebr/loadmonitor/inner/domain/metric"
	"github.com/uszebr/loadmonitor/inner/domain/workerpool"
	"github.com/uszebr/loadmonitor/inner/handler/jobmonitorhandl"
	"github.com/uszebr/loadmonitor/inner/handler/loadmanagerhandl"
	"github.com/uszebr/loadmonitor/inner/handler/runtimehandl"
)

const (
	VERSION = "0.2" //TODO: move to config or env var
)

func main() {
	fmt.Println("!!!!!!!!!!!!!!!Load Monitor started..")

	//PROMETHEUS
	promReg := prometheus.NewRegistry()
	//Default metrics(memory usage cpu etc, loading)
	//promReg.MustRegister(collectors.NewGoCollector())
	m := metric.NewMetrics(promReg)
	m.Info.With(prometheus.Labels{"version": VERSION}).Set(1)
	promCustomHandler := promhttp.HandlerFor(promReg, promhttp.HandlerOpts{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // TODO add to graceful shutdown

	jProducer := jobproducer.New(jobproducer.WithJobComplexity(50), jobproducer.WithMemoryLoad(10))
	jobQueue := jProducer.Start(ctx)

	wPool, proccessedJobx := workerpool.NewWorkerPool(ctx, 3, jobQueue)

	collector := collector.NewCollector(20, m)
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

	//prometheus router
	enginePrometheus := gin.Default()
	//engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	enginePrometheus.GET("/metrics", gin.WrapH(promCustomHandler))
	runtimeUpdater := metric.NewRuntimeUpdater(m, time.Second*2) //TODO: move to config.. or env.. time every
	runtimeUpdater.StartUpdating(ctx)

	// Start Load Monitor
	go func() {
		if err := engine.Run(":8085"); err != nil {
			fmt.Printf("Monitor server failed: %v\n", err)
		}
	}()

	// Start Prometheus metrics server using Gin on a separate port
	go func() {
		if err := enginePrometheus.Run(":8081"); err != nil {
			fmt.Printf("Metrics server failed: %v\n", err)
		}
	}()

	select {} //blocking forever

	//TODO: MAIN
	// DONE // refactor to JobI interface	// with all public methods??
	// config with default/start options
	// config path in env for prod/dev
	// logger with prod/dev
	// Multiplicator for job complexity store in config
	// DONE// htmx separate update for done jobs quantity. Commulative complexity..
}
