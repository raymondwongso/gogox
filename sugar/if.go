package sugar

// If evaluate cond, and return appropriate value.
// It is basically shorthand if.
func If[T any](cond bool, trueValue, falseValue T) T {
	if cond {
		return trueValue
	}
	return falseValue
}
