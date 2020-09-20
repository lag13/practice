package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/lag13/practice/leetcode"
)

func TestReverseInt(t *testing.T) {
	tt := []struct {
		x    int
		want int
	}{
		{
			x:    123,
			want: 321,
		},
		{
			x:    0,
			want: 0,
		},
		{
			x:    2,
			want: 2,
		},
		{
			x:    -123,
			want: -321,
		},
		{
			x:    120,
			want: 21,
		},
		{
			x:    1534236469,
			want: 0,
		},
		{
			x:    -2147483648,
			want: 0,
		},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("Reverse(%d)", tc.x), func(t *testing.T) {
			if got, want := leetcode.Reverse(tc.x), tc.want; got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
