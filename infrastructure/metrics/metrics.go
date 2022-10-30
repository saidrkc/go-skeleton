//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package metrics

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type MetricsInterface interface {
	AddToResponseTime(method string)
	NewTimer(method string) *prometheus.Timer
	IncrementTotalRequests(method string)
	IncrementResponseStatus(status int)
}

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

func (m Metrics) AddToResponseTime(method string) {
	m.HttpResponseTime.WithLabelValues(method)
}

func (m Metrics) NewTimer(method string) *prometheus.Timer {
	return prometheus.NewTimer(m.HttpResponseTime.WithLabelValues(method))
}

func (m Metrics) IncrementTotalRequests(method string) {
	m.TotalRequest.WithLabelValues(method).Inc()
}

func (m Metrics) IncrementResponseStatus(status int) {
	m.ResponseStatus.WithLabelValues(strconv.Itoa(status)).Inc()
}
