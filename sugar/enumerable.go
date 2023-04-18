package sugar

import "golang.org/x/exp/constraints"

// Any checks arr whether any of the elements satistifed the fn criteria AND then return the first found element.
// If no fn criteria given, return default value as the function assume that it can't check the criteria.
func Any[T any](arr []T, fn func(elem T) bool) T {
	var defaultValue T
	if fn == nil {
		return defaultValue
	}

	for i := range arr {
		if fn(arr[i]) {
			return arr[i]
		}
	}
	return defaultValue
}

// Count returns arr length
func Count[T any](arr []T) int {
	return len(arr)
}

// IsEmpty checks whether arr is empty
func IsEmpty[T any](arr []T) bool {
	return len(arr) == 0
}

// IsAll checks arr whether all of the elements satisfied the fn criteria
// If no fn criteria given, return false as the function assume that it can't check the criteria.
func IsAll[T any](arr []T, fn func(elem T) bool) bool {
	if fn == nil {
		return false
	}

	for i := range arr {
		if !fn(arr[i]) {
			return false
		}
	}
	return true
}

// IsAny checks arr whether any of the elements satisfied the fn criteria
// True if at least 1 element satisfied
// If no fn criteria given, return false as the function assume that it can't check the criteria.
func IsAny[T any](arr []T, fn func(elem T) bool) bool {
	if fn == nil {
		return false
	}

	for i := range arr {
		if fn(arr[i]) {
			return true
		}
	}
	return false
}

// IsNone checks arr whether none of the elements satisfied the fn criteria.
// True if none satisfied.
// If no fn criteria given, return default value as the function assume that it can't check the criteria.
func IsNone[T any](arr []T, fn func(elem T) bool) bool {
	if fn == nil {
		return false
	}

	for i := range arr {
		if fn(arr[i]) {
			return false
		}
	}
	return true
}

// Map maps new value(s) from arr, you can provide your own map function via fn.
// E.g: Map(int{1, 2, 3}, func(i int) string { strconv.Itoa(i) }) to produce new array of string
func Map[T any, V any](arr []T, fn func(elem T) V) []V {
	res := make([]V, len(arr))
	for i := range arr {
		res[i] = fn(arr[i])
	}
	return res
}

// Reverse reverse the array
func Reverse[T any](arr []T) []T {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// Select filters elements that satisfied the fn criteria
// If no fn criteria given, return the original arr
func Select[T any](arr []T, fn func(elem T) bool) []T {
	if fn == nil {
		return arr
	}

	res := make([]T, 0)
	for i := range arr {
		if fn(arr[i]) {
			res = append(res, arr[i])
		}
	}
	return res
}

// Sum sums value from evaluation fn function
// you can provide initialValue of the sum
// if no fn criteria given, return default value of V
func Sum[T any, V constraints.Integer | constraints.Float](initialValue V, arr []T, fn func(elem T) V) V {
	var sum V
	if fn == nil {
		return sum
	}

	for i := range arr {
		sum += fn(arr[i])
	}
	return sum
}
