package nop

import "github.com/raymondwongso/gogox/stats"

// Nop implements Stats but with no operation.
type Nop struct{}

// New creates new Nop stats
func New() *Nop {
	return &Nop{}
}

func (n *Nop) Increment(metric string, opt stats.Option) error {
	return nil
}

func (n *Nop) Histogram(metric string, value float64, opt stats.Option) error {
	return nil
}

func (n *Nop) Gauge(mmetric string, value float64, opt stats.Option) error {
	return nil
}

func (n *Nop) Add(metric string, value float64, opt stats.Option) error {
	return nil
}
