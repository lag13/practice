package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/lag13/practice/leetcode"
)

func TestLongestPalindrome(t *testing.T) {
	tt := []struct {
		s    string
		want string
	}{
		{
			s:    "",
			want: "",
		},
		{
			s:    "a",
			want: "a",
		},
		{
			s:    "babad",
			want: "bab",
		},
		{
			s:    "abccb",
			want: "bccb",
		},
		{
			s:    "abcdefgaba",
			want: "aba",
		},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("LongestPalindrome(%q)", tc.s), func(t *testing.T) {
			if got, want := leetcode.LongestPalindrome(tc.s), tc.want; got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}
