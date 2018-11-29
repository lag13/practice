package practice_test

import (
	"fmt"
	"testing"

	"github.com/lag13/practice"
)

// TestIsHappyNumber checks that the function correctly identifies
// happy numbers.
func TestIsHappyNumber(t *testing.T) {
	tests := []struct {
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
	for _, test := range tests {
		t.Run(fmt.Sprintf("IsHappyNumber(%d)", test.input), func(t *testing.T) {
			if got, want := practice.IsHappyNumber(test.input), test.output; got != want {
				t.Errorf("got %t, wanted %t", got, want)
			}
		})
	}
}
