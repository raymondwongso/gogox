package stats

import "golang.org/x/exp/maps"

// Stats defines commonly used stats usecases
type Stats interface {
	// Increment increase metric by 1
	Increment(metric string, opt Option) error
	// Histogram samples observations and counts them in configurable buckets
	Histogram(metric string, value float64, opt Option) error
	// Gauge represents a single numerical value
	Gauge(metric string, value float64, opt Option) error
	// Add increase metric by value
	Add(metric string, value float64, opt Option) error
}

type Tags map[string]string

type Option struct {
	Rate float64
	Tags Tags
}

func MergeTags(t1, t2 Tags) Tags {
	res := Tags{}
	maps.Copy(res, t1)
	maps.Copy(res, t2)
	return res
}
