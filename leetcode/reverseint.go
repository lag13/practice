package leetcode

import (
	"math"
)

func Reverse(x int) int {
	digits := []int{}
	for {
		digit := x % 10
		x = x / 10
		digits = append(digits, digit)
		if x == 0 {
			break
		}
	}
	maxDigitsIn32BitNum := 10
	maxLeadingDigitIn32BitNum := 2
	if len(digits) == maxDigitsIn32BitNum && int(math.Abs(float64(digits[0]))) > maxLeadingDigitIn32BitNum {
		return 0
	}
	mult := int(math.Pow10(len(digits) - 1))
	num := 0
	var ok bool
	for _, d := range digits {
		num, ok = add32(int32(num), int32(d*mult))
		if !ok {
			return 0
		}
		mult = mult / 10
	}
	return num
}

func add32(x, y int32) (int, bool) {
	z := x + y
	if (z > x) == (y > 0) {
		return int(z), true
	}
	return int(z), false
}
