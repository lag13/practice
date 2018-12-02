package practice_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/lag13/practice"
)

// TestSelectionSort tests that the selection sort correctly sorts the
// list.
func TestSelectionSort(t *testing.T) {
	tt := []struct {
		input  []int
		output []int
	}{
		{
			input:  nil,
			output: nil,
		},
		{
			input:  []int{},
			output: []int{},
		},
		{
			input:  []int{1},
			output: []int{1},
		},
		{
			input:  []int{5, 4, 3, 2, 1},
			output: []int{1, 2, 3, 4, 5},
		},
		{
			input:  []int{5, 4, 3, 2, 1},
			output: []int{1, 2, 3, 4, 5},
		},
		{
			input:  []int{-10, 2, 0, 1, 1},
			output: []int{-10, 0, 1, 1, 2},
		},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("running selection sort on: %#v", tc.input), func(t *testing.T) {
			practice.SelectionSort(tc.input)
			if !reflect.DeepEqual(tc.input, tc.output) {
				t.Errorf("got %#v wanted %#v", tc.input, tc.output)
			}
		})
	}
}
