package sugar

import "fmt"

// Concat creates string containing arr elements with delimiter.
// E.g: Concat([]int{1, 2, 3}, ",") => "1,2,3"
func Concat[T comparable](arr []T, delimiter string) string {
	if IsEmpty(arr) {
		return ""
	}

	res := fmt.Sprintf("%v", arr[0])
	for i := 1; i < Count(arr); i++ {
		res = fmt.Sprintf("%s%s%v", res, delimiter, arr[i])
	}

	return res
}
