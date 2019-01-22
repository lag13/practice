package practice

// RotateArray rotates an array (i.e. the last element becomes the
// first one).
func RotateArray(arr []int, n int) []int {
	if len(arr) == 0 {
		return arr
	}
	rotated := []int{}
	// https://github.com/golang/go/issues/448
	n = n % len(arr)
	if n < 0 {
		n += len(arr)
	}
	for i := len(arr) - n; i < len(arr); i++ {
		rotated = append(rotated, arr[i])
	}
	for i := 0; i < len(arr)-n; i++ {
		rotated = append(rotated, arr[i])
	}
	return rotated
}
