package log_test

import (
	"context"
	"testing"

	"github.com/raymondwongso/gogox/log"
	"github.com/stretchr/testify/assert"
)

func Test_NewContext(t *testing.T) {
	t.Run("should not panic when parentCtx is nil", func(t *testing.T) {
		assert.NotNil(t, log.NewContext(nil, nil))
	})
}

func Test_MetadataFromContext(t *testing.T) {
	t.Run("should not panic when ctx is nil", func(t *testing.T) {
		assert.NotNil(t, log.MetadataFromContext(nil))
	})

	t.Run("should not panic when ctx does not contain metadata", func(t *testing.T) {
		assert.NotNil(t, log.MetadataFromContext(context.Background()))
	})

	t.Run("should return metadata if exists", func(t *testing.T) {
		md := log.Metadata{"a": "b"}
		ctx := log.NewContext(context.Background(), md)
		assert.Equal(t, md, log.MetadataFromContext(ctx))
	})
}
