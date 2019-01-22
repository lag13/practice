package practice

// Given an integer array, move all elements that are equal to 0 to
// the left while maintaining the order of other elements in the
// array.
func MoveZerosLeft(nums []int) {
	writeIdx := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			continue
		}
		nums[writeIdx] = nums[i]
		writeIdx--
	}
	for i := writeIdx; i >= 0; i-- {
		nums[i] = 0
	}
}
