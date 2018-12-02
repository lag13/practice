/*

Given a sorted array of integers, return the index of the given key.
Return -1 if not found.

*/
package practice

// BinarySearch returns the index of the given number within a sorted
// list of numbers or -1 if the number is not in the list.
func BinarySearch(nums []int, num int) int {
	lowerBound := 0
	upperBoundExclusive := len(nums)
	for lowerBound < upperBoundExclusive {
		middle := (lowerBound + upperBoundExclusive) / 2
		if nums[middle] == num {
			return middle
		}
		if num < nums[middle] {
			upperBoundExclusive = middle
			continue
		}
		lowerBound = middle + 1
	}
	return -1
}

// BinarySearch2 is the same algorithm as above but done recursively
// instead of iteratively.
func BinarySearch2(nums []int, num int) int {
	return binarySearch2(nums, num, 0, len(nums))
}

func binarySearch2(nums []int, num int, lowerBound int, upperBoundExclusive int) int {
	if upperBoundExclusive <= lowerBound {
		return -1
	}
	middle := (lowerBound + upperBoundExclusive) / 2
	if nums[middle] == num {
		return middle
	}
	if num < nums[middle] {
		return binarySearch2(nums, num, lowerBound, middle)
	}
	return binarySearch2(nums, num, middle+1, upperBoundExclusive)
}

// BinarySearch3 is a recursive version of binary search which does
// not keep track of lower and upper bounds as explicitly which I
// like. My only complaint is that when we've determined that the
// element does not exist we need to bubble up that -1 value and
// prevent it from being added to.
func BinarySearch3(nums []int, num int) int {
	const notFoundIndex = -1
	if len(nums) == 0 {
		return notFoundIndex
	}
	middle := len(nums) / 2
	if nums[middle] == num {
		return middle
	}
	if num < nums[middle] {
		return BinarySearch3(nums[:middle], num)
	}
	newLowerBound := middle + 1
	possibleIndex := BinarySearch3(nums[newLowerBound:], num)
	if possibleIndex == notFoundIndex {
		return possibleIndex
	}
	return possibleIndex + newLowerBound
}

// BinarySearch4 works like BinarySearch3 but does not have to bubble
// up the -1 value because we panic when we've determined that the
// element does not exist hehe. I would not recommend using this
// solution in practice, seems pretty hacky. But it is pretty fun.
func BinarySearch4(nums []int, num int) (result int) {
	defer func() {
		if err := recover(); err != nil {
			result = -1
		}
	}()
	return binarySearch4(nums, num)
}

func binarySearch4(nums []int, num int) int {
	if len(nums) == 0 {
		panic("returning -1 in a hacky way woooooo!")
	}
	middle := len(nums) / 2
	if nums[middle] == num {
		return middle
	}
	if num < nums[middle] {
		return binarySearch4(nums[:middle], num)
	}
	newLowerBound := middle + 1
	return newLowerBound + binarySearch4(nums[newLowerBound:], num)
}
