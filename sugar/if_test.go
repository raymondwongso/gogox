package sugar_test

import (
	"testing"

	"github.com/raymondwongso/gogox/sugar"
	"github.com/stretchr/testify/assert"
)

func Test_If(t *testing.T) {
	t.Run("true value", func(t *testing.T) {
		res := sugar.If(true, "true", "false")
		assert.Equal(t, "true", res)
	})

	t.Run("false value", func(t *testing.T) {
		res := sugar.If(false, "true", "false")
		assert.Equal(t, "false", res)
	})
}
