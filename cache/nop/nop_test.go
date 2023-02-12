package nop_test

import (
	"context"
	"testing"

	"github.com/raymondwongso/gogox/cache/nop"
	"github.com/raymondwongso/gogox/errorx"
	"github.com/stretchr/testify/assert"
)

func Test_Nop(t *testing.T) {
	nopCache := nop.New()

	type testStruct struct{}

	t.Run("Get always return not found error", func(t *testing.T) {
		err := nopCache.Get(context.Background(), "some-key", &testStruct{})
		gogoxErr, ok := errorx.Parse(err)

		assert.True(t, ok)
		assert.Equal(t, errorx.CodeNotFound, gogoxErr.Code)
	})

	t.Run("Set always return nil error", func(t *testing.T) {
		err := nopCache.Set(context.Background(), "some-key", &testStruct{}, 0)
		assert.NoError(t, err)
	})

	t.Run("Del always return nil error", func(t *testing.T) {
		err := nopCache.Del(context.Background(), "some-key")
		assert.NoError(t, err)
	})
}
