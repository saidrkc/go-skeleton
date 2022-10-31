package main

import (
	"go-skeleton/cmd/engine"
	"go-skeleton/cmd/http"
)

func main() {
	srv := http.NewHttpServer()
	eng := engine.NewEngine(srv, srv.Metrics)
	cfg := eng.BuildEngine()
	eng.RunEngine(cfg.AddressPort)
}
