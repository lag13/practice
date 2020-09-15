package leetcode

func LongestPalindrome(s string) string {
	longest := ""
	for i := 0; i < len(s); i++ {
		tmpLongest := findLongest(s[i:])
		if len(tmpLongest) > len(longest) {
			longest = tmpLongest
		}
	}
	return longest
}

func findLongest(s string) string {
	longest := ""
	for i := 0; i < len(s); i++ {
		p := s[0 : i+1]
		if isPalindrome(p) && len(p) > len(longest) {
			longest = p
		}
	}
	return longest
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
