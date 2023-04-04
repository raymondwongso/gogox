package log

import (
	"net/http"
)

// Decider defines rules for surpressing log
type Decider func(req *http.Request, err error) bool

type options struct {
	ShouldLog Decider
}

// DefaultOptions creates default options
func DefaultOptions() options {
	return options{
		ShouldLog: DefaultDecider(),
	}
}

// DefaultDecider returns always true decider
func DefaultDecider() Decider {
	return func(req *http.Request, err error) bool {
		return true
	}
}

// SkipPrometheusDecider returns common skip prometheus /metric
func SkipPrometheusDecider() Decider {
	return func(req *http.Request, err error) bool {
		if err == nil && req.URL.Path == "/metrics" {
			return false
		}
		return true
	}
}
