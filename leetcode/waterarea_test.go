package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/lag13/practice/leetcode"
)

func TestMaxArea(t *testing.T) {
	tt := []struct {
		heights []int
		want    int
	}{
		{
			heights: []int{},
			want:    0,
		},
		{
			heights: []int{100},
			want:    0,
		},
		{
			heights: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:    49,
		},
		{
			heights: []int{20, 2, 5, 4, 8, 3, 10},
			want:    60,
		},
		{
			heights: []int{1, 100, 100, 4, 8, 3, 10},
			want:    100,
		},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("MaxArea(%v)", tc.heights), func(t *testing.T) {
			if got, want := leetcode.MaxArea(tc.heights), tc.want; got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
