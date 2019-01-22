package practice

func LowHighIndex(nums []int, key int) (int, int) {
	low, idx, high := lowHighIndexBinarySearch(nums, key, 0, len(nums)-1)
	if idx == -1 {
		return -1, -1
	}
	return searchLower(nums, key, low, idx), searchUpper(nums, key, idx, high)
}

func lowHighIndexBinarySearch(nums []int, key int, low int, high int) (int, int, int) {
	if high < low {
		return -1, -1, -1
	}
	mid := (low + high) / 2
	if key == nums[mid] {
		return low, mid, high
	}
	if key < nums[mid] {
		return lowHighIndexBinarySearch(nums, key, low, mid-1)
	}
	return lowHighIndexBinarySearch(nums, key, mid+1, high)
}

func searchLower(nums []int, key int, low int, high int) int {
	if nums[low] == key {
		return low
	}
	mid := (low + high) / 2
	if nums[mid] < key {
		return searchLower(nums, key, mid+1, high)
	}
	return searchLower(nums, key, low+1, mid)
}

func searchUpper(nums []int, key int, low int, high int) int {
	if nums[high] == key {
		return high
	}
	mid := (low + high) / 2
	if key < nums[mid] {
		return searchUpper(nums, key, low, mid-1)
	}
	return searchUpper(nums, key, mid, high-1)
}
