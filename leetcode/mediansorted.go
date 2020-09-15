package leetcode

// This is not O(log(n+m)) complexity as the site demands (I believe
// it's O((n+m)/2) but I wanted to get an initial solution in place. I
// wonder what happens if I submit it. TODO: Revisit this to see if I
// can figure out the O(log(n+m)) solution.
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLenMerged := (len(nums1)+len(nums2))/2 + 1
	merged := buildMerged(totalLenMerged, []int{}, nums1, nums2)
	if (len(nums1)+len(nums2))%2 == 1 {
		return float64(merged[totalLenMerged-1])
	}
	return float64(merged[totalLenMerged-1]+merged[totalLenMerged-2]) / 2
}

func buildMerged(totalLenMerged int, merged []int, nums1 []int, nums2 []int) []int {
	if len(merged) == totalLenMerged {
		return merged
	}
	if len(nums1) == 0 {
		return append(merged, nums2[:totalLenMerged-len(merged)]...)
	}
	if len(nums2) == 0 {
		return append(merged, nums1[:totalLenMerged-len(merged)]...)
	}
	n1 := nums1[0]
	n2 := nums2[0]
	if n1 < n2 {
		return buildMerged(totalLenMerged, append(merged, n1), nums1[1:], nums2)
	}
	return buildMerged(totalLenMerged, append(merged, n2), nums1, nums2[1:])
}
