package leetcode

func LengthOfLongestSubstring(s string) int {
	longest := 0
	for i := 0; i < len(s); i++ {
		seen := map[byte]bool{}
		tmpLongest := 0
		for j := i; j < len(s); j++ {
			if seen[s[j]] {
				break
			}
			seen[s[j]] = true
			tmpLongest++
		}
		if tmpLongest > longest {
			longest = tmpLongest
		}
	}
	return longest
}
