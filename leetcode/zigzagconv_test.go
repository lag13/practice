package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/lag13/practice/leetcode"
)

func TestZigZagConvert(t *testing.T) {
	tt := []struct {
		s       string
		numRows int
		want    string
	}{
		{
			s:       "PAYPALISHIRING",
			numRows: 0,
			want:    "",
		},
		{
			s:       "PAYPALISHIRING",
			numRows: 1,
			want:    "PAYPALISHIRING",
		},
		{
			s:       "PAYPALISHIRING",
			numRows: 2,
			want:    "PYAIHRNAPLSIIG",
		},
		{
			s:       "PAYPALISHIRING",
			numRows: 3,
			want:    "PAHNAPLSIIGYIR",
		},
		{
			s:       "PAYPALISHIRING",
			numRows: 4,
			want:    "PINALSIGYAHRPI",
		},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("Convert(%q,%d)", tc.s, tc.numRows), func(t *testing.T) {
			if got, want := leetcode.Convert(tc.s, tc.numRows), tc.want; got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}
