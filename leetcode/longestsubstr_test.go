package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/lag13/practice/leetcode"
)

func TestLongestSubstring(t *testing.T) {
	tt := []struct {
		s    string
		want int
	}{
		{
			s:    "abcabcbb",
			want: 3,
		},
		{
			s:    "",
			want: 0,
		},
		{
			s:    "a",
			want: 1,
		},
		{
			s:    "aaaaaa",
			want: 1,
		},
		{
			s:    "aaaabac",
			want: 3,
		},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("LengthOfLongestSubstring(%q)", tc.s), func(t *testing.T) {
			if got, want := leetcode.LengthOfLongestSubstring(tc.s), tc.want; got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
