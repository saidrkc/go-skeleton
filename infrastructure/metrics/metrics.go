package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	HttpResponseTime *prometheus.HistogramVec
	TotalRequest     *prometheus.CounterVec
	ResponseStatus   *prometheus.CounterVec
}

func NewMetrics(HttpResponseTime *prometheus.HistogramVec, TotalRequest *prometheus.CounterVec, ResponseStatus *prometheus.CounterVec) Metrics {
	return Metrics{
		HttpResponseTime,
		TotalRequest,
		ResponseStatus,
	}
}
