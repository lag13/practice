package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/lag13/practice/leetcode"
)

func TestAtoi(t *testing.T) {
	tt := []struct {
		s    string
		want int
	}{
		{
			s:    "",
			want: 0,
		},
		{
			s:    "42",
			want: 42,
		},
		{
			s:    "+123",
			want: 123,
		},
		{
			s:    "-9001",
			want: -9001,
		},
		{
			s:    "    -42",
			want: -42,
		},
		{
			s: "	-42 this string has a tab so should fail because only spaces are ignored",
			want: 0,
		},
		{
			s:    "words then number 123",
			want: 0,
		},
		{
			s:    "123 number then words",
			want: 123,
		},
		{
			s:    "2147483646",
			want: 2147483646,
		},
		{
			s:    "2147483647",
			want: 2147483647,
		},
		{
			s:    "2147483648",
			want: 2147483647,
		},
		{
			s:    "9999999999",
			want: 2147483647,
		},
		{
			s:    "-2147483647",
			want: -2147483647,
		},
		{
			s:    "-2147483648",
			want: -2147483648,
		},
		{
			s:    "-2147483649",
			want: -2147483648,
		},
		{
			s:    "-9999999999",
			want: -2147483648,
		},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("Atoi(%q)", tc.s), func(t *testing.T) {
			if got, want := leetcode.Atoi(tc.s), tc.want; got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
