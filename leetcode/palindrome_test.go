package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/lag13/practice/leetcode"
)

func TestIsPalindrome(t *testing.T) {
	tt := []struct {
		x    int
		want bool
	}{
		{
			x:    0,
			want: true,
		},
		{
			x:    123,
			want: false,
		},
		{
			x:    12321,
			want: true,
		},
		{
			x:    123321,
			want: true,
		},
		{
			x:    -123321,
			want: false,
		},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("IsPalindrome(%d)", tc.x), func(t *testing.T) {
			if got, want := leetcode.IsPalindrome(tc.x), tc.want; got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
