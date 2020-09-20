package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/lag13/practice/leetcode"
)

func TestIsMatch(t *testing.T) {
	tt := []struct {
		s    string
		p    string
		want bool
	}{
		{
			s:    "",
			p:    "",
			want: true,
		},
		{
			s:    "a",
			p:    "a",
			want: true,
		},
		{
			s:    "",
			p:    "a",
			want: false,
		},
		{
			s:    "aa",
			p:    "a",
			want: false,
		},
		{
			s:    "aa",
			p:    "aa",
			want: true,
		},
		{
			s:    "aa",
			p:    "a.",
			want: true,
		},
		{
			s:    "aa",
			p:    "a*",
			want: true,
		},
		{
			s:    "aab",
			p:    "c*a*b",
			want: true,
		},
		{
			s:    "ccaab",
			p:    "c*ccaab",
			want: true,
		},
		{
			s:    "abcdefg",
			p:    ".*",
			want: true,
		},
		{
			s:    "",
			p:    ".*a*b*c*",
			want: true,
		},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("IsMatch(%q, %q)", tc.s, tc.p), func(t *testing.T) {
			if got, want := leetcode.IsMatch(tc.s, tc.p), tc.want; got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
