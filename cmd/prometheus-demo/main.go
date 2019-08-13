package main

import (
	"net/http"
	"os"

	cli "github.com/jawher/mow.cli"
	"github.com/martinsirbe/prometheus-demo/internal/db/postgres"
	"github.com/martinsirbe/prometheus-demo/internal/service"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

const (
	appName        = "prometheus-demo"
	appDescription = "An example app which exposes Prometheus metrics on /metrics endpoint."
)

func main() {
	app := cli.App(appName, appDescription)

	app.Action = func() {
		storageClient := postgres.NewInstrumentedClientPrometheus()
		http.Handle("/metrics", promhttp.Handler())

		// Example for other clients, such as Graphite, DataDog where the client post metrics.
		// Re Graphite metrics, it's also possible to configure Prometheus Graphite bridge for this.
		// 		bridge, err := graphite.NewBridge(&graphite.Config{
		//			URL:           *carbonURL,
		//			Prefix:        "poc",
		//			Timeout:       5 * time.Second,
		//			ErrorHandling: graphite.AbortOnError,
		//			Logger:        log.StandardLogger(),
		//		})
		//		if err != nil {
		//			logrus.WithError(err).Fatal("failed to create Prometheus Graphite bridge")
		//		}
		//      go bridge.Run(context.Background())
		//storageClient := postgres.InstrumentedClientOther()

		ll := service.NewLogicLayer(storageClient)
		go func() {
			if err := ll.Run(); err != nil {
				log.WithError(err).Fatal("failed to run")
			}
		}()

		if err := http.ListenAndServe(":1337", nil); err != nil {
			log.WithError(err).Fatal("failed to listen and serve metrics")
		}
	}

	if err := app.Run(os.Args); err != nil {
		log.WithError(err).Panicf("app failed to run")
	}
}
