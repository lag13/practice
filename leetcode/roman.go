package leetcode

// I 1
// V 5
// X 10
// L 50
// C 100
// D 500
// M 1000

// This could be made faster (and maybe more memory efficient?) by
// using strings.Builder but I'll take recursive beauty over that any
// day. I mean, we're converting numbers to roman numerals for gods
// sake, I feel like speed is not a pressing concern here. If I was
// writing this for a phone though, yeah I'd probably try to optimize
// the hell out of it.

func IntToRoman(num int) string {
	if num == 0 {
		return ""
	}
	if num < 4 {
		return "I" + IntToRoman(num-1)
	}
	if num == 4 {
		return "IV"
	}
	if num < 9 {
		return "V" + IntToRoman(num-5)
	}
	if num == 9 {
		return "IX"
	}
	if num < 40 {
		return "X" + IntToRoman(num-10)
	}
	if num < 50 {
		return "XL" + IntToRoman(num-40)
	}
	if num < 90 {
		return "L" + IntToRoman(num-50)
	}
	if num < 100 {
		return "XC" + IntToRoman(num-90)
	}
	if num < 400 {
		return "C" + IntToRoman(num-100)
	}
	if num < 500 {
		return "CD" + IntToRoman(num-400)
	}
	if num < 900 {
		return "D" + IntToRoman(num-500)
	}
	if num < 1000 {
		return "CM" + IntToRoman(num-900)
	}
	return "M" + IntToRoman(num-1000)
}
