package leetcode

import (
	"math"
)

func Atoi(s string) int {
	for len(s) > 0 && s[0] == ' ' {
		s = s[1:]
	}
	if len(s) == 0 {
		return 0
	}
	negate := false
	if s[0] == '-' {
		negate = true
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	var maxVal uint32 = math.MaxInt32
	if negate {
		maxVal++
	}
	var n uint32 = 0
	for _, ch := range []byte(s) {
		ch = ch - '0'
		if ch > 9 {
			break
		}
		if n > maxVal/10 {
			n = maxVal
			break
		}
		n = n * 10
		n1 := n + uint32(ch)
		if n1 < n || n1 > maxVal {
			n = maxVal
			break
		}
		n = n1
	}
	n1 := int32(n)
	if negate {
		// Remember that the most negative int32 number times
		// -1 is the same. Math with limitations amiright?
		n1 = -n1
	}
	return int(n1)
}
