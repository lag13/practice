/*

Selection sort repeatedly finds the smallest element in a list and
puts it at the front of the list:
https://en.wikipedia.org/wiki/Selection_sort

*/
package main

import "fmt"

// slectionSort sorts the given list in place.
func selectionSort(lst []int) {
	for i := 0; i < len(lst)-1; i++ {
		min := lst[i]
		minPos := i
		for j := i + 1; j < len(lst); j++ {
			if lst[j] < min {
				min = lst[j]
				minPos = j
			}
		}
		swap := lst[i]
		lst[i] = min
		lst[minPos] = swap
	}
}

// selectionSortFn returns a sorted list not modifying the original
// list.
func selectionSortFn(lst []int) []int {
	newLst := make([]int, len(lst))
	copy(newLst, lst)
	selectionSort(newLst)
	return newLst
}

func main() {
	lst := []int{0, 3, -1, 2, 1, 500, 1}
	fmt.Println("Before:", lst)
	selectionSort(lst)
	fmt.Println("After:", lst)
	lst = []int{0, 3, -1, 2, 1, 500, 1, 501}
	fmt.Println(selectionSortFn(lst))
	fmt.Println(lst)
}
