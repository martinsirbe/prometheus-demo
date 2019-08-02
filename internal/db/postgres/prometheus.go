package postgres

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// InstrumentedClientPrometheus instrumented postgres client
type InstrumentedClientPrometheus struct {
	counter *prometheus.CounterVec
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
	}
}

// Insert fake insert
func (c *InstrumentedClientPrometheus) Insert(o string) error {
	c.counter.WithLabelValues("insert").Add(1)
	fmt.Printf("inserted - %s\n", o)
	return nil
}

// Delete fake delete
func (c *InstrumentedClientPrometheus) Delete(o string) error {
	c.counter.WithLabelValues("delete").Add(1)
	fmt.Printf("deleted - %s\n", o)
	return nil
}
