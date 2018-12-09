package practice_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/lag13/practice"
)

// TestMaxsInSlidingWindow tests that we correctly identify the all
// maximum numbers within a window of some size as it goes across the
// array.
func TestMaxsInSlidingWindow(t *testing.T) {
	tt := []struct {
		arr         []int
		windowSize  int
		wantMaxNums []int
	}{
		{
			[]int{},
			3,
			[]int{},
		},
		{
			[]int{-4, -3, -2, -1, 0, 1, 2, 3, 4},
			3,
			[]int{-2, -1, 0, 1, 2, 3, 4},
		},
		{
			[]int{4, 3, 2, 1, 0, -1, -2, -3, -4},
			3,
			[]int{4, 3, 2, 1, 0, -1, -2},
		},
		{
			[]int{0, -1, 10, 100, 2, 2, 4},
			2,
			[]int{0, 10, 100, 100, 2, 4},
		},
		{
			[]int{0, -1, 10, 100, 2, 2, 4, -1, 5, 6},
			4,
			[]int{100, 100, 100, 100, 4, 5, 6},
		},
	}
	// TODO: One time a test panic'd and it was tough for me to
	// determine which test case caused the panic. I thought go
	// test's would handle panic's nicely and still output which
	// test caused the failure but I guess that is not the case.
	// Figure out if there is a nice way to fix this issue or if
	// we'll need to defer a recover statement.
	for _, tc := range tt {
		t.Run(fmt.Sprintf("MaxsInSlidingWindow(%#v,%d)", tc.arr, tc.windowSize), func(t *testing.T) {
			if got, want := practice.MaxsInSlidingWindow2(tc.arr, tc.windowSize), tc.wantMaxNums; !reflect.DeepEqual(got, want) {
				t.Errorf("got %#v, want %#v", got, want)
			}
		})
	}
}

// TODO: Like the binary work, figure out how to better analyze these
// functions to see where the parts of code which caused more
// slowdown.

// TODO: Figure out how to do more detailed analysis of performance.
// Like as I change the window size how does the algorithm's runtime
// change? Or if the list is increasing, decreasing, or random how
// does that affect performance? I should be able generate graphs and
// answer these types of questions.

// TODO: In the future, add a couple Example tests to have examples of
// using that facility.

// TODO: Also in the future talk about some of the more advanced
// options for documenting go source code.

func BenchmarkMaxsInSlidingWindow1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practice.MaxsInSlidingWindow(bigOrderedList, 10)
	}
}

func BenchmarkMaxsInSlidingWindow2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practice.MaxsInSlidingWindow2(bigOrderedList, 10)
	}
}
