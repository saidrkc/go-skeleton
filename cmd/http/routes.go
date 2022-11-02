package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go-skeleton/infrastructure/metrics"
	"go-skeleton/src/application/ping"
	"go-skeleton/src/application/pong"
	"go-skeleton/src/application/user"
	"go-skeleton/src/infrastructure/bus/command"
	"go-skeleton/src/infrastructure/bus/query"
	"go-skeleton/src/infrastructure/http/get"
	user3 "go-skeleton/src/infrastructure/http/get/user"
	"go-skeleton/src/infrastructure/http/post"
	user2 "go-skeleton/src/infrastructure/http/post/user"
	"go-skeleton/src/infrastructure/memory"
)

const DEFAULT_RELATIVE_RANKING = "/relative"
const DEFAULT_ABSOLUTE_SCORE = "/score"
const DEFAULT_ABSOLUTE_RANKING = "/ranking"

const DEFAULT_PING_URL = "/ping"
const DEFAULT_PONG_URL = "/pong"
const DEFAULT_PROMETHEUS_METRICS = "/metrics"

var singleton *memory.UserRepository

type Routes struct {
	Gin            *gin.Engine
	Metrics        metrics.MetricsInterface
	UserRepository memory.UserRepository
}

func (g *Routes) BindRoutes() {
	g.Gin.GET(DEFAULT_RELATIVE_RANKING, g.buildRelativeRankingHandlersMapping)
	g.Gin.POST(DEFAULT_ABSOLUTE_SCORE, g.buildAbsoluteScoreHandlersMapping)
	g.Gin.GET(DEFAULT_ABSOLUTE_RANKING, g.buildAbsoluteRankingHandlersMapping)

	// Basic testing endpoints
	g.Gin.POST(DEFAULT_PONG_URL, g.buildPongHandlersMapping)
	g.Gin.GET(DEFAULT_PING_URL, g.buildPingHandlersMapping)
	g.Gin.GET(DEFAULT_PROMETHEUS_METRICS, prometheusHandler())
}

func (g *Routes) buildPingHandlersMapping(c *gin.Context) {
	pingQueryHandler := ping.NewPingApplication(c, g.Metrics)
	qbManager := query.NewQueryBus()
	qbManager.RegisterHandler(ping.PingQuery{}, pingQueryHandler)
	pingController := get.NewPingHandler(g.Metrics)
	pingController.Ping(c, qbManager)
}

func (g *Routes) buildPongHandlersMapping(c *gin.Context) {
	pongCommandHandler := pong.NewPongApplication(c, g.Metrics)
	cbManager := command.NewCommandBus()
	cbManager.RegisterHandler(pong.PongCommand{}, pongCommandHandler)
	pongController := post.NewPongHandler(g.Metrics)
	pongController.Pong(c, cbManager)
}

func (g *Routes) buildAbsoluteScoreHandlersMapping(c *gin.Context) {
	absoluteScoreHandler := user.NewAbsoluteScoreApplication(c, g.Metrics, UserRepository())
	cbManager := command.NewCommandBus()
	cbManager.RegisterHandler(user.AbsoluteScoreCommand{}, absoluteScoreHandler)
	absoluteController := user2.NewAbsoluteScoreHandler(g.Metrics)
	absoluteController.AbsoluteScore(c, cbManager)
}

func (g *Routes) buildAbsoluteRankingHandlersMapping(c *gin.Context) {
	absoluteRankingQueryHandler := user.NewAbsoluteRanking(c, g.Metrics, UserRepository())
	qbManager := query.NewQueryBus()
	qbManager.RegisterHandler(user.AbsoluteRankingQuery{}, absoluteRankingQueryHandler)
	absoluteRankingController := user3.NewAbsoluteRankingHandler(g.Metrics)
	absoluteRankingController.AbsoluteRanking(c, qbManager)
}

func (g *Routes) buildRelativeRankingHandlersMapping(c *gin.Context) {
	relativeRankingQueryHandler := user.NewRelativeRanking(c, g.Metrics, UserRepository())
	qbManager := query.NewQueryBus()
	qbManager.RegisterHandler(user.RelativeRankingQuery{}, relativeRankingQueryHandler)
	relativeRankingController := user3.NewRelativeRankingHandler(g.Metrics)
	relativeRankingController.RelativeRanking(c, qbManager)
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func UserRepository() *memory.UserRepository {
	return singleton
}

func (g *Routes) InitRepository() {
	userRepository := memory.NewUserRepository()
	singleton = &userRepository
}
