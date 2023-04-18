package stats

import "golang.org/x/exp/maps"

//go:generate mockgen -destination=mock/stats.go -package=statsmock -source=stats.go

// Stats defines commonly used stats usecases.
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

// Tags defines key value mapping commonly used to give tags for metrices
type Tags map[string]string

// Option defines stats configuration
type Option struct {
	Rate float64
	Tags Tags
}

// MergeTags merge two tags
func MergeTags(t1, t2 Tags) Tags {
	res := Tags{}
	maps.Copy(res, t1)
	maps.Copy(res, t2)
	return res
}
