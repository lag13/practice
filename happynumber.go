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
package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

func isHappyNumber(n int) bool {
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

func boolToEnglish(b bool) string {
	if b {
		return "is"
	}
	return "is not"
}

func happyNumberSequence(n int) []int {
	sequence := []int{}
	for n != 1 && !contains(sequence, n) {
		sequence = append(sequence, n)
		n = numToSumOfSquaresOfDigits(n)
	}
	return append(sequence, n)
}

func sequenceToIsHappy(seq []int) bool {
	return seq[len(seq)-1] == 1
}

// Although I don't use this function I came across it when
// researching if there was a builtin way to do the equivalent of
// strings.Join() but for integers. I kept it because (again) this is
// all just practice:
// https://stackoverflow.com/questions/37532255/one-liner-to-transform-int-into-string/37533144
func intJoin(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func intsToStrings(xs []int) []string {
	strs := []string{}
	for _, x := range xs {
		strs = append(strs, strconv.Itoa(x))
	}
	return strs
}

func main() {
	// On the one hand I felt that displaying the sequence was not
	// necessary but on the other hand it is cool to see the
	// sequence so I decided to do both since I'm just practicing
	// anyway.
	for i := 1; i <= 100; i++ {
		fmt.Println(i, boolToEnglish(isHappyNumber(i)), "a happy number")
	}
	for i := 1; i <= 100; i++ {
		seq := happyNumberSequence(i)
		fmt.Println(i, boolToEnglish(sequenceToIsHappy(seq)), "a happy number, the sequence that was generated to determine that was:", strings.Join(intsToStrings(seq), ","))
	}
}
