package practice

// RotateArray rotates an array (i.e. the last element becomes the
// first one).
func RotateArray(arr []int, n int) []int {
	if len(arr) == 0 {
		return arr
	}
	rotated := []int{}
	// TODO: I believe the proper definition of modulus for
	// negative numbers should still return a positive number but
	// this does not. This calculation I have here seems to work
	// though, learn why. Also modulus is not defined when the
	// second operand is 0 but I feel like it should be defined in
	// that case? Look this up too.
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
