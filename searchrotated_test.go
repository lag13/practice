package practice_test

import (
	"fmt"
	"testing"

	"github.com/lag13/practice"
)

// TestSearchRotated tests that searching for a number in an array
// that has been rotated works as expected. This was a much more
// complicated problem than I was expecting and it needed a lot of
// test cases to cover all possible scenarios (or perhaps my solution
// is overly complicated). In total it probably took me 10 hours
// spread across 4 days to complete.
//
// This problem was a perfect example of how, in order to write good
// tests (who's cases are not redundant and they exercise all paths in
// the solution), you must have a solid understanding of the
// algorithm. It is okay to write parts of the algorithm first and
// then add a test covering the logic you just added. Solving this
// also served as a good reminder to bust out the pad and paper if you
// are really struggling with a problem.
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
			num:  8,
			want: -1,
		},
		{
			nums: []int{15, 18, -1, 2, 4, 7, 10},
			num:  10,
			want: 6,
		},
		{
			nums: []int{15, 18, -1, 2, 4, 7, 10},
			num:  15,
			want: 0,
		},
		{
			nums: []int{4, 6, 7, 8, 10, 0, 1},
			num:  9,
			want: -1,
		},
		{
			nums: []int{4, 6, 7, 8, 10, 0, 1},
			num:  10,
			want: 4,
		},
		{
			nums: []int{3, 3, 3, 3, 3, 4, 1, 2, 3},
			num:  4,
			want: 5,
		},
		{
			nums: []int{3, 4, 1, 2, 3, 3, 3, 3, 3},
			num:  4,
			want: 1,
		},
		{
			nums: []int{3, 5, 7, 9, -1, 0, 1},
			num:  5,
			want: 1,
		},
		{
			nums: []int{3, 5, 7, 9, -1, 0, 1},
			num:  -2,
			want: -1,
		},
		{
			nums: []int{3, 5, 7, 9, -1, 0, 1},
			num:  -1,
			want: 4,
		},
		{
			nums: []int{7, -1, 0, 1, 3, 5, 6},
			num:  -1,
			want: 1,
		},
		{
			nums: []int{3, 0, 1, 2, 3, 3, 3, 3, 3},
			num:  0,
			want: 1,
		},
		{
			nums: []int{3, 3, 3, 3, 3, 0, 1, 2, 3},
			num:  0,
			want: 5,
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
