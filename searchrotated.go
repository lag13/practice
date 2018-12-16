package practice

// SearchRotated will search for a number in a sorted array which has
// been rotated (i.e. the last element becomes the first element) an
// arbitrary number of times. Linear search is not an acceptable
// solution. TODO: See how I could simplify this implementation to
// just 3 basic cases (1) We search in the larger portion (2) we
// search in the smaller portion (3) we search in both portions
// because we cannot tell which one to search in.
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
	// TODO: Was I wrong in assuming there could be duplicates?
	// Because I'm realizing that one algorithm could return a
	// different answer if it called SearchRotated(nums[mid+1...)
	// first as opposed to calling SearchRotated(nums[:mid]...) as
	// I did here. Maybe that's fine but it does feel a bit
	// strange.
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

func SearchRotated2(nums []int, num int) int {
	// should originalMid be (len(nums)-1)/2 ?
	return searchRotated2(nums, num, len(nums)/2, 0, len(nums)-1)
}

// TODO: Finish this implementation. It could end up being similar to
// the first one but maybe I'll come up with other ideas.
func searchRotated2(nums []int, num int, originalMid int, lower int, upper int) int {
	if upper < lower {
		if originalMid == 0 {
			return -1
		}
		if upper < 0 {
			return searchRotated2(nums, num, 0, originalMid, len(nums)-1)
		}
		if lower > len(nums)-1 {
			return searchRotated2(nums, num, 0, 0, originalMid)
		}
		return -1
	}
	mid := (lower + upper) / 2
	if num == nums[mid] {
		return mid
	}
	if num > nums[mid] {
		// TODO: We have to check whether the previous number
		// we saw is smaller than this one. If so we've
		// crossed the rotation point and must try to confine
		// our search in this area
		return searchRotated2(nums, num, originalMid, mid+1, upper)
	}
	return searchRotated2(nums, num, originalMid, lower, mid-1)
}
