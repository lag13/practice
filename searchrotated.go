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
	if num > nums[mid] {
		// TODO: At this point I feel like we do not need to
		// care about mid anymore so it should be mid+1. But I
		// don't think we can do that because the array might
		// not be large enough to accomodate it. Maybe we
		// don't even need this check at all? There is lots of
		// repetition on which type of SearchRotated() call to
		// perform so maybe we can restructure things
		// differently to consolidate that logic a bit?
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
		maybeIndex1 := SearchRotated(nums[:mid], num)
		if maybeIndex1 == notFoundIndex {
			maybeIndex2 := SearchRotated(nums[mid+1:], num)
			if maybeIndex2 == notFoundIndex {
				return notFoundIndex
			}
			return maybeIndex2 + mid + 1
		}
		return maybeIndex1
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
	maybeIndex1 := SearchRotated(nums[:mid], num)
	if maybeIndex1 == notFoundIndex {
		maybeIndex2 := SearchRotated(nums[mid+1:], num)
		if maybeIndex2 == notFoundIndex {
			return notFoundIndex
		}
		return maybeIndex2 + mid + 1
	}
	return maybeIndex1
}
