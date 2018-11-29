package practice_test

import (
	"fmt"
	"testing"

	"github.com/lag13/practice"
)

// TestBinarySearch tests that the binary search works as expected.
func TestBinarySearch(t *testing.T) {
	tests := []struct {
		nums   []int
		num    int
		output int
	}{
		{
			nums:   []int{},
			num:    2,
			output: -1,
		},
		{
			nums:   []int{3},
			num:    1,
			output: -1,
		},
		{
			nums:   []int{1, 2, 3, 4},
			num:    0,
			output: -1,
		},
		{
			nums:   []int{1, 2, 3, 4},
			num:    5,
			output: -1,
		},
		{
			nums:   []int{1, 2, 4, 5},
			num:    3,
			output: -1,
		},
		{
			nums:   []int{3},
			num:    3,
			output: 0,
		},
		{
			nums:   []int{1, 2, 3, 4},
			num:    1,
			output: 0,
		},
		{
			nums:   []int{2, 3, 4, 5, 10},
			num:    10,
			output: 4,
		},
		{
			nums:   []int{2, 3, 4, 5, 10},
			num:    3,
			output: 1,
		},
		{
			nums:   []int{2, 3, 4, 5, 10},
			num:    5,
			output: 3,
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("BinarySearch(%v, %d)", test.nums, test.num), func(t *testing.T) {
			if got, want := practice.BinarySearch(test.nums, test.num), test.output; got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
