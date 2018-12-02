/*

Given a large array of integers and a window of size 'w', find the
current maximum in the window as the window slides through the entire
array.

*/
package practice

// MaxsInSlidingWindow returns the list of maximum values within the
// array within a window as it slides across the array.
func MaxsInSlidingWindow(arr []int, windowSize int) []int {
	maxNums := []int{}
	for i := 0; i < len(arr)-windowSize+1; i++ {
		max := arr[i]
		for j := i + 1; j < i+windowSize; j++ {
			if arr[j] > max {
				max = arr[j]
			}
		}
		maxNums = append(maxNums, max)
	}
	return maxNums
}

// We can use this function to get more consistency in the overall
// algorithm. It bugs me to no end that the loop to initialize the
// window of largest values and then adding the values successively
// are so similar in construction.
func addElementToWindow(window []int, nums []int, i int, windowSize int) []int {
	// Remove the element that is no longer within the window
	if len(window) > 0 && window[0] <= i-windowSize {
		copy(window, window[1:])
		window = window[:len(window)-1]
	}
	// Remove elements from this window structure if they are <=
	// than the new element we're adding (because if the newest
	// element is larger than the previous elements we no longer
	// need to care about the previous ones anymore since this new
	// element is larger and will be in the window longer than
	// them).
	for len(window) > 0 && nums[i] >= nums[window[len(window)-1]] {
		window = window[:len(window)-1]
	}
	return append(window, i)
}

// MaxsInSlidingWindow2 keeps track of the largest element we've seen
// so far in the window so we do not need to do the more brute force
// approach in the other algorithm.
func MaxsInSlidingWindow2(nums []int, windowSize int) []int {
	maxNums := []int{}
	if len(nums) < windowSize {
		return maxNums
	}
	window := []int{}
	for i := 0; i < windowSize; i++ {
		window = addElementToWindow(window, nums, i, windowSize)
	}
	maxNums = append(maxNums, nums[window[0]])
	for i := windowSize; i < len(nums); i++ {
		window = addElementToWindow(window, nums, i, windowSize)
		maxNums = append(maxNums, nums[window[0]])
	}
	return maxNums
}
