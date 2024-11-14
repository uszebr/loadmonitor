package metric

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	GoRoutines prometheus.Gauge
	Info       *prometheus.GaugeVec
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
	}
	reg.MustRegister(metr.GoRoutines, metr.Info)
	return metr
}
