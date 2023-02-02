package log_test

import (
	"testing"

	"github.com/raymondwongso/gogox/log"
	"github.com/stretchr/testify/assert"
)

func Test_MergeMetadata(t *testing.T) {
	t.Run("should not error when merging md1 and nil", func(t *testing.T) {
		expected := log.Metadata{"service": "api"}
		assert.Equal(t,
			expected,
			log.MergeMetadata(log.Metadata{"service": "api"}, nil),
		)
	})

	t.Run("should not error when merging nil and md2", func(t *testing.T) {
		expected := log.Metadata{"service": "api"}
		assert.Equal(t,
			expected,
			log.MergeMetadata(nil, log.Metadata{"service": "api"}),
		)
	})

	t.Run("should not error when merging nil and nil", func(t *testing.T) {
		expected := log.Metadata{}
		assert.Equal(t,
			expected,
			log.MergeMetadata(nil, nil),
		)
	})
}
