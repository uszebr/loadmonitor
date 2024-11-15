package metric

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	GoRoutines           prometheus.Gauge
	Info                 *prometheus.GaugeVec
	CumulativeComplexity prometheus.Counter
	DoneJobsCounter      prometheus.Counter
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	metr := &Metrics{
		GoRoutines: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "loadmonitor",       //matches name of app
			Name:      "gorutine_quantity", //naming convetion with underscore
			Help:      "Number of gorutines running particular moment",
		}),
		Info: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: "loadmonitor",
				Name:      "info",
				Help:      "Information about Load Monitor env",
			}, []string{"version"}),
		CumulativeComplexity: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "loadmonitor",
			Name:      "cumulative_complexity",
			Help:      "Total accumulated complexity of all jobs processed",
		}),
		DoneJobsCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "loadmonitor",
			Name:      "done_jobs_total",
			Help:      "Total number of jobs completed",
		}),
	}
	reg.MustRegister(metr.GoRoutines, metr.Info, metr.CumulativeComplexity, metr.DoneJobsCounter)
	return metr
}
