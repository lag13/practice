package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/lag13/practice/leetcode"
)

func TestMedianSorted(t *testing.T) {
	tt := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2,
		},
		{
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
		{
			nums1: []int{0, 0},
			nums2: []int{0, 0},
			want:  0,
		},
		{
			nums1: []int{},
			nums2: []int{1},
			want:  1,
		},
		{
			nums1: []int{2},
			nums2: []int{},
			want:  2,
		},
		{
			nums1: []int{1, 1},
			nums2: []int{2},
			want:  1,
		},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("FindMedianSortedArrays(%v,%v)", tc.nums1, tc.nums2), func(t *testing.T) {
			if got, want := leetcode.FindMedianSortedArrays(tc.nums1, tc.nums2), tc.want; got != want {
				t.Errorf("got %f, want %f", got, want)
			}
		})
	}
}
