package practice_test

import (
	"fmt"
	"testing"

	"github.com/lag13/practice"
)

// TestIsHappyNumber checks that the function correctly identifies
// happy numbers.
func TestIsHappyNumber(t *testing.T) {
	tt := []struct {
		input  int
		output bool
	}{
		{
			input:  1,
			output: true,
		},
		{
			input:  2,
			output: false,
		},
		{
			input:  4,
			output: false,
		},
		{
			input:  7,
			output: true,
		},
	}
	// The naming of these varaibles "tt" (table tests) and "tc"
	// (test case) comes from this lovely video:
	// https://www.youtube.com/watch?v=hVFEV-ieeew&feature=youtu.be&t=868
	for _, tc := range tt {
		t.Run(fmt.Sprintf("IsHappyNumber(%d)", tc.input), func(t *testing.T) {
			if got, want := practice.IsHappyNumber(tc.input), tc.output; got != want {
				t.Errorf("got %t, wanted %t", got, want)
			}
		})
	}
}
