package practice_test

import (
	"fmt"
	"testing"

	"github.com/lag13/practice"
)

// TestSearchRotated tests that searching for a number in an array
// that has been rotated works as expected.
func TestSearchRotated(t *testing.T) {
	tt := []struct {
		nums []int
		num  int
		want int
	}{
		{
			nums: []int{},
			num:  1,
			want: -1,
		},
		{
			nums: []int{15, 18, -1, 2, 4, 7, 10},
			num:  2,
			want: 3,
		},
		{
			nums: []int{15, 18, -1, 2, 4, 5, 7, 10},
			num:  7,
			want: 6,
		},
		{
			nums: []int{6, 7, 8, 9, 10, 0, 1, 2, 3, 4, 5},
			num:  -3,
			want: -1,
		},
		// you either first jump into the part that has
		// wrapped around or you jump into the part that has
		// not been wrapped. If you do the former then if you
		// keep searching for smaller numbers you will have to
		// check the other side of the list assuming you do
		// not find it. But if you are searching for a smaller
		// number and the number you're looking at is larger
		// than the last number you've seen in the array, then
		// you know you've searched all you can.

		// Do you always jump to the rotated point? No... that
		// feels wrong. But there does seem to be something to
		// it...

		// You will either jump into the portion of the array
		// that has been wrapped around or the portion that
		// has not. And actually, how can you tell the
		// difference? Given the array could you find the
		// rotation number? Probably if you did a linear
		// search you could.
		{
			nums: []int{0, -10, -9, -8, -7},
			num:  0,
			want: 0,
		},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("SearchRotated(%v,%d)", tc.nums, tc.num), func(t *testing.T) {
			if got, want := practice.SearchRotated(tc.nums, tc.num), tc.want; got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
