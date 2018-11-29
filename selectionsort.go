/*

Selection sort repeatedly finds the smallest element in a list and
puts it at the front of the list:
https://en.wikipedia.org/wiki/Selection_sort

*/
package practice

// SlectionSort sorts the given list in place in ascending order.
func SelectionSort(lst []int) {
	for i := 0; i < len(lst)-1; i++ {
		min := lst[i]
		minPos := i
		for j := i + 1; j < len(lst); j++ {
			if lst[j] < min {
				min = lst[j]
				minPos = j
			}
		}
		// https://stackoverflow.com/questions/41314959/assigning-values-to-multiple-variables-in-go/41315277
		lst[i], lst[minPos] = min, lst[i]
	}
}
