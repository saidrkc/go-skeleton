//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"

	"go-skeleton/infrastructure/metrics"
)

type HttpServer interface {
	BuildHttpServer(metrics metrics.MetricsInterface)
	Gin() *gin.Engine
}

type Config struct {
	DefaultPrometheusMetric string
	AddressPort             int
	AddressIp               string
}

type Server struct {
	GinEngine *gin.Engine
	Metrics   metrics.MetricsInterface
}

func NewHttpServer() *Server {
	return &Server{
		GinEngine: gin.New(),
		Metrics:   metricsRegister(),
	}
}

func (srv *Server) BuildHttpServer(metrics metrics.MetricsInterface) {
	routes := Routes{Gin: srv.GinEngine, Metrics: metrics}
	routes.InitRepository()
	routes.BindRoutes()
}

func (srv *Server) Gin() *gin.Engine {
	return srv.GinEngine
}

func metricsRegister() metrics.Metrics {
	var httpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})
	var totalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of get requests.",
		},
		[]string{"path"},
	)
	var responseStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "response_status",
			Help: "Status of HTTP response",
		},
		[]string{"status"},
	)
	prometheus.MustRegister(httpDuration)
	prometheus.MustRegister(totalRequests)
	prometheus.MustRegister(responseStatus)
	return metrics.NewMetrics(httpDuration, totalRequests, responseStatus)
}
