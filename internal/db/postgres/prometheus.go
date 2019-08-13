package postgres

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	TimerHistogram prometheus.Histogram
	TimeSummary    prometheus.Summary
	TimerGauge     prometheus.Gauge
}

// InstrumentedClientPrometheus instrumented postgres client
type InstrumentedClientPrometheus struct {
	counter       *prometheus.CounterVec
	insertMetrics Metrics
	deleteMetrics Metrics
}

// NewInstrumentedClientPrometheus initialises a new fake postgres client
func NewInstrumentedClientPrometheus() *InstrumentedClientPrometheus {
	return &InstrumentedClientPrometheus{
		counter: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "actions_total",
			Help: "Total number of actions per db and table",
		},
			[]string{"type"},
		),
		insertMetrics: Metrics{
			TimerHistogram: promauto.NewHistogram(prometheus.HistogramOpts{
				Name: "db_insert_histogram",
			}),
			TimeSummary: promauto.NewSummary(prometheus.SummaryOpts{
				Name: "db_insert_summary",
			}),
			TimerGauge: promauto.NewGauge(prometheus.GaugeOpts{
				Name: "db_insert_gauge",
			}),
		},
		deleteMetrics: Metrics{
			TimerHistogram: promauto.NewHistogram(prometheus.HistogramOpts{
				Name: "db_delete_histogram",
			}),
			TimeSummary: promauto.NewSummary(prometheus.SummaryOpts{
				Name: "db_delete_summary",
			}),
			TimerGauge: promauto.NewGauge(prometheus.GaugeOpts{
				Name: "db_delete_gauge",
			}),
		},
	}
}

// Insert fake insert
func (c *InstrumentedClientPrometheus) Insert(o string) error {
	hisTimer := prometheus.NewTimer(c.insertMetrics.TimerHistogram)
	sumTimer := prometheus.NewTimer(c.insertMetrics.TimeSummary)
	gaugeTimer := prometheus.NewTimer(prometheus.ObserverFunc(c.insertMetrics.TimerGauge.Set))

	defer hisTimer.ObserveDuration()
	defer sumTimer.ObserveDuration()
	defer gaugeTimer.ObserveDuration()

	sleep(30, 360)

	c.counter.WithLabelValues("insert").Add(1)

	fmt.Printf("inserted - %s\n", o)
	return nil
}

// Delete fake delete
func (c *InstrumentedClientPrometheus) Delete(o string) error {
	hisTimer := prometheus.NewTimer(c.deleteMetrics.TimerHistogram)
	sumTimer := prometheus.NewTimer(c.deleteMetrics.TimeSummary)
	gaugeTimer := prometheus.NewTimer(prometheus.ObserverFunc(c.deleteMetrics.TimerGauge.Set))

	defer hisTimer.ObserveDuration()
	defer sumTimer.ObserveDuration()
	defer gaugeTimer.ObserveDuration()

	sleep(5, 60)

	c.counter.WithLabelValues("delete").Add(1)

	fmt.Printf("deleted - %s\n", o)
	return nil
}

func sleep(minSec, maxSec int) {
	rand.Seed(time.Now().UnixNano())
	sleepSec := rand.Intn(maxSec-minSec) + minSec
	time.Sleep(time.Duration(sleepSec) * time.Second)
}
