package trace_test

import (
	"context"
	"testing"

	"github.com/raymondwongso/gogox/trace"
	"github.com/stretchr/testify/assert"
)

func Test_NewContext(t *testing.T) {
	t.Run("should not panic when parentCtx is nil", func(t *testing.T) {
		assert.NotNil(t, trace.NewContext(nil, ""))
	})
}

func Test_TraceFromContext(t *testing.T) {
	t.Run("should not panic when ctx is nil", func(t *testing.T) {
		assert.NotNil(t, trace.TraceFromContext(nil))
	})

	t.Run("should not panic when ctx does not contain metadata", func(t *testing.T) {
		assert.NotNil(t, trace.TraceFromContext(context.Background()))
	})

	t.Run("should return trace if exists", func(t *testing.T) {
		ctx := trace.NewContext(context.Background(), "mytrace")
		assert.Equal(t, "mytrace", trace.TraceFromContext(ctx))
	})
}
