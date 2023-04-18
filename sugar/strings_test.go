package sugar_test

import (
	"testing"

	"github.com/raymondwongso/gogox/sugar"
	"github.com/stretchr/testify/assert"
)

func Test_Concat(t *testing.T) {
	t.Run("empty array, get empty string", func(t *testing.T) {
		res := sugar.Concat([]int{}, ",")
		assert.Equal(t, "", res)
	})

	t.Run("delimiter by comma", func(t *testing.T) {
		res := sugar.Concat(arr, ",")
		assert.Equal(t, "1,2,3,4", res)
	})

	t.Run("delimiter by pipe, using float", func(t *testing.T) {
		res := sugar.Concat([]float64{1.2, 2.3, 3.4}, "|")
		assert.Equal(t, "1.2|2.3|3.4", res)
	})
}
