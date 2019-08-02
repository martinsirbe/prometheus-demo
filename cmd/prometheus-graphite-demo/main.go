package main

import (
	"context"
	"net/http"
	"time"

	"github.com/martinsirbe/prometheus-graphite-demo/internal/db/postgres"
	"github.com/martinsirbe/prometheus-graphite-demo/internal/service"
	"github.com/prometheus/client_golang/prometheus/graphite"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {
	storageClient := postgres.NewInstrumentedClientPrometheus()
	http.Handle("/metrics", promhttp.Handler())

	// Example for other clients, such as Graphite, DataDog where the client post metrics
	//storageClient := postgres.InstrumentedClientOther()

	ll := service.NewLogicLayer(storageClient)
	go func() {
		if err := ll.Run(); err != nil {
			log.WithError(err).Fatal("failed to run")
		}
	}()

	bridge, err := graphite.NewBridge(&graphite.Config{
		URL:           "localhost:2003",
		Prefix:        "poc",
		Timeout:       5 * time.Second,
		ErrorHandling: graphite.AbortOnError,
		Logger:        log.StandardLogger(),
	})

	if err != nil {
		logrus.WithError(err).Fatal("failed to create Prometheus Graphite bridge")
	}

	go bridge.Run(context.Background())

	if err := http.ListenAndServe(":1337", nil); err != nil {
		log.WithError(err).Fatal("failed to listen and serve metrics")
	}
}
