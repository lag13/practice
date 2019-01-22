package practice

// Given an integer array, move all elements that are equal to 0 to
// the left while maintaining the order of other elements in the
// array.
func MoveZerosLeft(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != 0 {
			continue
		}
		for j := i - 1; j >= 0; j-- {
			if nums[j] == 0 {
				continue
			}
			nums[i] = nums[j]
			nums[j] = 0
			break
		}
	}
}
