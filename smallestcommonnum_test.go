package practice_test

import (
	"fmt"
	"testing"

	"github.com/lag13/practice"
)

func TestSmallestCommonNumber(t *testing.T) {
	tt := []struct {
		a1     []int
		a2     []int
		a3     []int
		result int
	}{
		{
			a1:     []int{1, 2, 3, 4},
			a2:     []int{4, 5, 6, 7},
			a3:     []int{0, 2, 4, 6},
			result: 4,
		},
		{
			a1:     []int{10},
			a2:     []int{1, 3, 4, 5, 8, 10},
			a3:     []int{10, 12, 14},
			result: 10,
		},
		{
			a1:     []int{0, 1, 2, 5, 6, 9},
			a2:     []int{5},
			a3:     []int{4, 5},
			result: 5,
		},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("case %d: SmallestCommonNumber(%+v,%+v,%+v)", i, tc.a1, tc.a2, tc.a3), func(t *testing.T) {
			if got, want := practice.SmallestCommonNumber(tc.a1, tc.a2, tc.a3), tc.result; got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}
