package sugar_test

import (
	"strconv"
	"testing"

	"github.com/raymondwongso/gogox/sugar"
	"github.com/stretchr/testify/assert"
)

var (
	arr = []int{1, 2, 3, 4}
)

func Test_Any(t *testing.T) {
	t.Run("fn criteria is not provided, get default value of type", func(t *testing.T) {
		res := sugar.Any(arr, nil)
		assert.Equal(t, 0, res)
	})

	t.Run("success return any", func(t *testing.T) {
		res := sugar.Any(arr, func(x int) bool {
			return x > 2
		})
		assert.Equal(t, 3, res)
	})

	t.Run("no criteria found, get default value of type", func(t *testing.T) {
		res := sugar.Any(arr, func(x int) bool {
			return x > 6
		})
		assert.Equal(t, 0, res)
	})
}

func Test_Count(t *testing.T) {
	t.Run("return arr length", func(t *testing.T) {
		res := sugar.Count(arr)
		assert.Equal(t, 4, res)
	})
}

func Test_IsAll(t *testing.T) {
	t.Run("fn criteria is not provided, get false", func(t *testing.T) {
		res := sugar.IsAll(arr, nil)
		assert.Equal(t, false, res)
	})

	t.Run("success get true criteria is fulfilled", func(t *testing.T) {
		res := sugar.IsAll(arr, func(x int) bool {
			return x > 0
		})
		assert.Equal(t, true, res)
	})

	t.Run("some criteria not fulfilled, get false", func(t *testing.T) {
		res := sugar.IsAll(arr, func(x int) bool {
			return x > 2
		})
		assert.Equal(t, false, res)
	})
}

func Test_IsAny(t *testing.T) {
	t.Run("fn criteria is not provided, get false", func(t *testing.T) {
		res := sugar.IsAny(arr, nil)
		assert.Equal(t, false, res)
	})

	t.Run("success get true criteria is fulfilled", func(t *testing.T) {
		res := sugar.IsAny(arr, func(x int) bool {
			return x > 2
		})
		assert.Equal(t, true, res)
	})

	t.Run("no criteria fulfilled, get false", func(t *testing.T) {
		res := sugar.IsAny(arr, func(x int) bool {
			return x > 6
		})
		assert.Equal(t, false, res)
	})
}

func Test_IsNone(t *testing.T) {
	t.Run("fn criteria is not provided, get false", func(t *testing.T) {
		res := sugar.IsNone(arr, nil)
		assert.Equal(t, false, res)
	})

	t.Run("success get true no criteria is fulfilled", func(t *testing.T) {
		res := sugar.IsNone(arr, func(x int) bool {
			return x > 6
		})
		assert.Equal(t, true, res)
	})

	t.Run("some criteria fulfilled, get false", func(t *testing.T) {
		res := sugar.IsNone(arr, func(x int) bool {
			return x > 2
		})
		assert.Equal(t, false, res)
	})
}

func Test_Map(t *testing.T) {
	t.Run("array int to string", func(t *testing.T) {
		res := sugar.Map(arr, func(x int) string {
			return strconv.Itoa(x)
		})

		assert.Equal(t, []string{"1", "2", "3", "4"}, res)
	})
}

func Test_Reserve(t *testing.T) {
	t.Run("reverse empty array, got empty array", func(t *testing.T) {
		res := sugar.Reverse([]int{})
		assert.Equal(t, []int{}, res)
	})

	t.Run("reverse array", func(t *testing.T) {
		res := sugar.Reverse([]float64{1.2, 2.3, 3.4})
		assert.Equal(t, []float64{3.4, 2.3, 1.2}, res)
	})
}

func Test_Select(t *testing.T) {
	t.Run("fn criteria is not provided, get original arr", func(t *testing.T) {
		res := sugar.Select(arr, nil)
		assert.Equal(t, arr, res)
	})

	t.Run("select according to fn criteria", func(t *testing.T) {
		res := sugar.Select(arr, func(x int) bool {
			return x > 2
		})
		assert.Equal(t, []int{3, 4}, res)
	})
}

func Test_Sum(t *testing.T) {
	t.Run("fn criteria is not provided, get zero value", func(t *testing.T) {
		res := sugar.Sum(0, arr, nil)
		assert.Equal(t, 0, res)

		res2 := sugar.Sum(0, []float64{0.1, 0.2, 0.3, 0.4}, nil)
		assert.Equal(t, 0, res2)
	})

	t.Run("sum the int arr", func(t *testing.T) {
		res := sugar.Sum(0, arr, func(x int) int {
			return x
		})
		assert.Equal(t, 10, res)
	})

	t.Run("sum struct attribute", func(t *testing.T) {
		type thing struct {
			price float64
		}

		things := []thing{
			{
				price: 1.23,
			},
			{
				price: 4.56,
			},
			{
				price: 7.89,
			},
		}

		res := sugar.Sum(0, things, func(x thing) float64 {
			return x.price
		})
		assert.Equal(t, 13.68, res)
	})
}
