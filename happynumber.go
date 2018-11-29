/*

A “happy number” is defined by the following process: Starting with
any positive integer, replace the number by the sum of the squares of
its digits, and repeat the process until the number either equals 1
(where it will stay), or it *loops endlessly in a cycle* that does not
include 1. Those numbers for which this process ends in 1 are happy
numbers, while those that do not end in 1 are unhappy numbers (or sad
numbers).

For a given positive integer n determine if n is a happy number.

*/
package practice

func numToDigits(n int) []int {
	ns := []int{}
	for n > 0 {
		ns = append(ns, n%10)
		n = n / 10
	}
	return ns
}

func numToSumOfSquaresOfDigits(n int) int {
	sum := 0
	for _, digit := range numToDigits(n) {
		sum += digit * digit
	}
	return sum
}

func contains(xs []int, y int) bool {
	for _, x := range xs {
		if x == y {
			return true
		}
	}
	return false
}

func IsHappyNumber(n int) bool {
	seen := []int{}
	for {
		if n == 1 {
			return true
		}
		if contains(seen, n) {
			return false
		}
		seen = append(seen, n)
		n = numToSumOfSquaresOfDigits(n)
	}
}
