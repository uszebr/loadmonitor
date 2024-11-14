package metric

import (
	"context"
	"runtime"
	"time"
)

type RuntimeUpdater struct {
	metric *Metrics
	every  time.Duration
}

func NewRuntimeUpdater(metric *Metrics, every time.Duration) RuntimeUpdater {
	return RuntimeUpdater{metric: metric, every: every}
}

func (rtu RuntimeUpdater) StartUpdating(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(rtu.every)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				break
			case <-ticker.C:
				rtu.metric.GoRoutines.Set(float64(runtime.NumGoroutine()))
			}
		}
	}()
}
