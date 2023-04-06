package log

import (
	"net/http"
)

// Decider defines rules for surpressing log
type Decider func(req *http.Request, status int) bool

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
	return func(req *http.Request, status int) bool {
		return true
	}
}

// SkipPrometheusDecider returns common skip prometheus /metric
func SkipPrometheusDecider() Decider {
	return func(req *http.Request, status int) bool {
		if status == 200 && req.URL.Path == "/metrics" {
			return false
		}
		return true
	}
}
