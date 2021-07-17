package main

import (
	"Bluefairy/app"
	"Bluefairy/config"
	"context"
	"flag"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	config.Initconfig()
	//config.SetLogger()
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		_ = http.ListenAndServe(":80", nil)
	}()
	flag.Parse()

	app.Run(context.TODO())
}
