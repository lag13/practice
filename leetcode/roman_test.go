package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/lag13/practice/leetcode"
)

func TestIntToRoman(t *testing.T) {
	tt := []struct {
		num  int
		want string
	}{
		{
			num:  1,
			want: "I",
		},
		{
			num:  3,
			want: "III",
		},
		{
			num:  4,
			want: "IV",
		},
		{
			num:  5,
			want: "V",
		},
		{
			num:  8,
			want: "VIII",
		},
		{
			num:  9,
			want: "IX",
		},
		{
			num:  58,
			want: "LVIII",
		},
		{
			num:  1994,
			want: "MCMXCIV",
		},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("IntToRoman(%d)", tc.num), func(t *testing.T) {
			if got, want := leetcode.IntToRoman(tc.num), tc.want; got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}

func BenchmarkRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		leetcode.IntToRoman(1994)
	}
}
