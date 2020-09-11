package leetcode_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/lag13/practice/leetcode"
)

func TestTwoSum(t *testing.T) {
	tt := []struct {
		nums   []int
		target int
		output []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			output: []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			output: []int{1, 2},
		},
		{
			nums:   []int{10, 0},
			target: 10,
			output: []int{0, 1},
		},
		{
			nums:   []int{},
			target: 1,
			output: nil,
		},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("TwoSum(%+v, %d)", tc.nums, tc.target), func(t *testing.T) {
			if got, want := leetcode.TwoSum(tc.nums, tc.target), tc.output; !reflect.DeepEqual(got, want) {
				t.Errorf("got %+v, want %+v", got, want)
			}
		})
	}
}
