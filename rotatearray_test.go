package practice_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/lag13/practice"
)

// TestRotateArray tests that rotating an array works as expected.
func TestRotateArray(t *testing.T) {
	tt := []struct {
		arr  []int
		n    int
		want []int
	}{
		{
			arr:  []int{},
			n:    10,
			want: []int{},
		},
		{
			arr:  []int{1, 2, 3, 4},
			n:    2,
			want: []int{3, 4, 1, 2},
		},
		{
			arr:  []int{1, 2, 3, 4},
			n:    3,
			want: []int{2, 3, 4, 1},
		},
		{
			arr:  []int{1, 2, 3, 4},
			n:    4,
			want: []int{1, 2, 3, 4},
		},
		{
			arr:  []int{1, 2, 3, 4},
			n:    -1,
			want: []int{2, 3, 4, 1},
		},
		{
			arr:  []int{1, 2, 3, 4, 5, 6, 7},
			n:    -3,
			want: []int{4, 5, 6, 7, 1, 2, 3},
		},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("RotateArray(%+v, %d)", tc.arr, tc.n), func(t *testing.T) {
			if got, want := practice.RotateArray(tc.arr, tc.n), tc.want; !reflect.DeepEqual(got, want) {
				t.Errorf("got %+v, want %+v", got, want)
			}
		})
	}
}
