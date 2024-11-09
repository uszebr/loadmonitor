package loadmanagerview

import "github.com/uszebr/loadmonitor/inner/domain/jobproducer"

type JobProducerFormData struct {
	*jobproducer.JobProducer
	Success         bool
	ErrorComplexity string
	ErrorMemoryLoad string
}
