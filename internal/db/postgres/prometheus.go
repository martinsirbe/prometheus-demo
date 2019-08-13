package postgres

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	TimerHistogram           prometheus.Histogram
	TimeSummary              prometheus.Summary
	TimerGauge               prometheus.Gauge
	RangeMinSec, RangeMaxSec int
}

// InstrumentedClientPrometheus instrumented postgres client
type InstrumentedClientPrometheus struct {
	counter       *prometheus.CounterVec
	insertMetrics Metrics
	deleteMetrics Metrics
}

// NewInstrumentedClientPrometheus initialises a new fake postgres client
func NewInstrumentedClientPrometheus(insertRange, deleteRange string) *InstrumentedClientPrometheus {
	iMin, iMax := getRange(insertRange)
	dMin, dMax := getRange(deleteRange)

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
			RangeMinSec: iMin,
			RangeMaxSec: iMax,
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
			RangeMinSec: dMin,
			RangeMaxSec: dMax,
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

	sleep(c.insertMetrics.RangeMinSec, c.insertMetrics.RangeMaxSec)

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

	sleep(c.deleteMetrics.RangeMinSec, c.deleteMetrics.RangeMaxSec)

	c.counter.WithLabelValues("delete").Add(1)

	fmt.Printf("deleted - %s\n", o)
	return nil
}

func sleep(minSec, maxSec int) {
	rand.Seed(time.Now().UnixNano())
	sleepSec := rand.Intn(maxSec-minSec) + minSec
	time.Sleep(time.Duration(sleepSec) * time.Second)
}

func getRange(rangeString string) (int, int) {
	r := strings.Split(rangeString, ":")
	min, err := strconv.Atoi(r[0])
	if err != nil {
		panic(errors.Errorf("failed to obtain min range from range string, range string - %s", rangeString))
	}

	max, err := strconv.Atoi(r[1])
	if err != nil {
		panic(errors.Errorf("failed to obtain max range from range string, range string - %s", rangeString))
	}

	return min, max
}
