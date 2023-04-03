package log

import (
	"path"

	"google.golang.org/grpc"
)

// Decider defines rules for surpressing log
type Decider func(info *grpc.UnaryServerInfo, err error) bool

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
	return func(info *grpc.UnaryServerInfo, err error) bool {
		return true
	}
}

// SkipHealthCheckDecider returns common health check skip log
func SkipHealthCheckDecider() Decider {
	return func(info *grpc.UnaryServerInfo, err error) bool {
		service := path.Dir(info.FullMethod)[1:]
		method := path.Base(info.FullMethod)

		if err == nil && method == "Check" && service == "grpc.health.v1.Health" {
			return false
		}
		return true
	}
}
