package practice

// SmallestCommonNumber takes three arrays in ascending order and
// returns the smallest number common in all three arrays.
func SmallestCommonNumber(arr1 []int, arr2 []int, arr3 []int) int {
	a1i, a2i, a3i := 0, 0, 0
	for {
		for arr1[a1i] < arr2[a2i] || arr1[a1i] < arr3[a3i] {
			a1i++
		}
		for arr2[a2i] < arr1[a1i] || arr2[a2i] < arr3[a3i] {
			a2i++
		}
		for arr3[a3i] < arr1[a1i] || arr3[a3i] < arr2[a2i] {
			a3i++
		}
		if arr1[a1i] == arr2[a2i] && arr2[a2i] == arr3[a3i] {
			return arr1[a1i]
		}
	}
	// Since the problem description never described what should
	// happen if there is no matching number I assume this case
	// will never be reached
	return 42
}
