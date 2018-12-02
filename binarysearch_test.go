package practice_test

import (
	"fmt"
	"testing"

	"github.com/lag13/practice"
)

// TestBinarySearch tests that the binary search works as expected.
func TestBinarySearch(t *testing.T) {
	tt := []struct {
		nums   []int
		num    int
		output int
	}{
		{
			nums:   []int{},
			num:    2,
			output: -1,
		},
		{
			nums:   []int{3},
			num:    1,
			output: -1,
		},
		{
			nums:   []int{1, 2, 3, 4},
			num:    0,
			output: -1,
		},
		{
			nums:   []int{1, 2, 3, 4},
			num:    5,
			output: -1,
		},
		{
			nums:   []int{1, 2, 4, 5},
			num:    3,
			output: -1,
		},
		{
			nums:   []int{3},
			num:    3,
			output: 0,
		},
		{
			nums:   []int{1, 2, 3, 4},
			num:    1,
			output: 0,
		},
		{
			nums:   []int{2, 3, 4, 5, 10},
			num:    10,
			output: 4,
		},
		{
			nums:   []int{2, 3, 4, 5, 10},
			num:    3,
			output: 1,
		},
		{
			nums:   []int{2, 3, 4, 5, 10},
			num:    5,
			output: 3,
		},
	}
	for _, tc := range tt {
		t.Run(fmt.Sprintf("BinarySearch(%v, %d)", tc.nums, tc.num), func(t *testing.T) {
			if got, want := practice.BinarySearch(tc.nums, tc.num), tc.output; got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		})
	}
}

func genBigOrderedList(size int) []int {
	lst := make([]int, size)
	for i := 0; i < size; i++ {
		lst = append(lst, i)
	}
	return lst
}

var bigOrderedList = genBigOrderedList(10000000)

// Originally I used this function to benchmark each of the binary
// search implementations but when I was pprof'ing
// (https://golang.org/pkg/runtime/pprof/) the various solutions to
// try and see why some were faster than others I found that not using
// this template made pprof's output easier to read. Let this serve as
// a reminder, do not refactor benchmarks too much.

// func benchmarkBinarySearchTemplate(b *testing.B, fn func([]int, int) int) {
// 	for i := 0; i < b.N; i++ {
// 		fn(bigOrderedList, -1)
// 	}
// }

/*

Quick Primer on profiling with pprof
(https://golang.org/pkg/runtime/pprof/ is just as good really, I
wanted the commands inline here):

    go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
    go tool pprof -http=localhost:8080 cpu.prof

One thing to note is that benchmarks in Go cause the code under test
to run as many times as it takes to get consistent timings. This tends
to mean that your fastest code will be executed more which could
translate into a larger overall runtime. So when you see pprof saying
that your fastest algorithm was actually your slowest, don't freak
out. It's just that this algorithm had to be run more to get reliable
timings. This also makes me think that maybe generating the pprof
profile based on a benchmark is not necessarily good for comparing
algorithms like I tried here because each algorithm was run a
different number of times. I wish the libraries and tools around this
were smart enough to pick a representative sample from the benchmarks
and make a pprof profile out of those. Or maybe I'm trying to use the
tool wrong and really the tool is better suited for analyzing where
the slowdowns are in a particular algorithm?

One thought I had on why the 3rd algorithm was faster than the 2nd was
because the second algorithm has to call one more function but then I
thought that seemed a little silly because if simply calling a
function caused that much of a slowdown then the 1st algorithm should
be *much* faster when comparing it to the other recursive ones.

TODO: Figure out how to profile and compare different implementations
of the same algorithm in Go to better understand why one
implementation is faster/slower.

*/

const numToSearch = -1

func BenchmarkBinary1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practice.BinarySearch(bigOrderedList, numToSearch)
	}
}

func BenchmarkBinary2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practice.BinarySearch2(bigOrderedList, numToSearch)
	}
}

func BenchmarkBinary3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practice.BinarySearch3(bigOrderedList, numToSearch)
	}
}

func BenchmarkBinary4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practice.BinarySearch4(bigOrderedList, numToSearch)
	}
}
