package practice_test

import (
	"fmt"
	"testing"

	"github.com/lag13/practice"
)

// TestLowHighIndex does some suweet stuff.
func TestLowHighIndex(t *testing.T) {
	tt := []struct {
		arr      []int
		key      int
		wantLow  int
		wantHigh int
	}{
		{
			arr:      []int{10},
			key:      1,
			wantLow:  -1,
			wantHigh: -1,
		},
		{
			arr:      []int{1, 2, 5, 5, 5, 5, 5, 5, 5, 5, 20},
			key:      1,
			wantLow:  0,
			wantHigh: 0,
		},
		{
			arr:      []int{1, 2, 5, 5, 5, 5, 5, 5, 5, 5, 20},
			key:      2,
			wantLow:  1,
			wantHigh: 1,
		},
		{
			arr:      []int{1, 2, 5, 5, 5, 5, 5, 5, 5, 5, 20},
			key:      5,
			wantLow:  2,
			wantHigh: 9,
		},
		{
			arr:      []int{1, 2, 5, 5, 5, 5, 5, 5, 5, 5, 20},
			key:      20,
			wantLow:  10,
			wantHigh: 10,
		},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("LowHighIndex(%v, %d)", tc.arr, tc.key), func(t *testing.T) {
			low, high := practice.LowHighIndex(tc.arr, tc.key)
			if got, want := low, tc.wantLow; got != want {
				t.Errorf("got low %d, want %d", got, want)
			}
			if got, want := high, tc.wantHigh; got != want {
				t.Errorf("got high %d, want %d", got, want)
			}
		})
	}
}
