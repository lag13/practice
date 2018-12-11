package practice

// SearchRotated will search for a number in a sorted array which has
// been rotated (i.e. the last element becomes the first element) an
// arbitrary number of times. Linear search is not an acceptable
// solution.
func SearchRotated(nums []int, num int) int {
	const notFoundIndex = -1
	if len(nums) == 0 {
		return notFoundIndex
	}
	mid := len(nums) / 2
	if num == nums[mid] {
		return mid
	}
	// lastElem := nums[len(nums)-1]
	if num > nums[mid] {
		// mid+1?
		if nums[mid] < nums[len(nums)-1] { // this portion is increasing
			if num <= nums[len(nums)-1] { // and this is where we might find num
				maybeIndex := SearchRotated(nums[mid+1:], num)
				if maybeIndex == notFoundIndex {
					return notFoundIndex
				}
				return maybeIndex + mid + 1
			}
			// The largest element must be in the lower
			// half of the array
			return SearchRotated(nums[:mid], num)
		}
		// The rotation point happens here and we'll find num
		// here as well
		if nums[mid] > nums[len(nums)-1] {
			maybeIndex := SearchRotated(nums[mid+1:], num)
			if maybeIndex == notFoundIndex {
				return notFoundIndex
			}
			return maybeIndex + mid + 1
		}
		// TODO: What to do when nums[mid] == nums[len(nums)-1]
	}
	// this portion is increasing as normal
	if nums[0] < nums[mid] {
		if nums[0] <= num {
			return SearchRotated(nums[:mid], num)
		}
		// The smallest element will be in the other half
		maybeIndex := SearchRotated(nums[mid+1:], num)
		if maybeIndex == notFoundIndex {
			return notFoundIndex
		}
		return maybeIndex + mid + 1
	}
	// rotation point happens in this array and we'll find num here
	if nums[0] > nums[mid] {
		return SearchRotated(nums[:mid], num)
	}
	// TODO: What to do when nums[0] == nums[mid]?
	return -1
}
