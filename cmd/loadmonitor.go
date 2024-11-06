package main

import (
	"context"
	"fmt"
	"time"

	"github.com/yourusername/loadmonitor/inner/job"
)

func main() {
	fmt.Println("!!!!!!!!!!!!!!!Load Monitor started..")
	factory := job.GetInstance(5, 0, 100)
	ctx, cancel := context.WithCancel(context.Background())
	factory.StartFactory(ctx)

	fmt.Println("!!!!!!!!!!!!!!!FactoryStarted")
	time.Sleep(10 * time.Second)

	cancel()
	fmt.Println("!!!!!!!!!!!!!!!VeryEnd")
}
