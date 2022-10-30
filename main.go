package main

import (
	"go-skeleton/cmd/app"
	"go-skeleton/cmd/http"
)

func main() {
	srv := http.NewHttpServer()
	engine := app.NewEngine(srv, srv.Metrics)
	engine.RunEngine()
}
