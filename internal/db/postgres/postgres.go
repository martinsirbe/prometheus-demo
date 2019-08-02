package postgres

import "github.com/rcrowley/go-metrics"

func init() {
	metrics.UseNilMetrics = true
}
