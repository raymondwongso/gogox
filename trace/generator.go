package trace

import (
	"github.com/segmentio/ksuid"
)

// New creates new random string using ksuid
func New() string {
	return ksuid.New().String()
}
