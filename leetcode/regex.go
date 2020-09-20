package leetcode

func IsMatch(s string, p string) bool {
	return isMatchRec(s, p, false, '0')
}

func isMatchRec(s string, p string, star bool, starCh byte) bool {
	if s == "" && p == "" {
		return true
	}
	si := len(s) - 1
	if star && s != "" && charMatches(s[si], starCh) {
		return isMatchRec(s[:si], p, star, starCh)
	}
	if p != "" {
		pi := len(p) - 1
		if p[pi] == '*' {
			// I'm assuming that the * character preceded by
			// nothing is illegal in this problem
			if len(p) == 1 {
				return false
			}
			pi--
			// I'm also going to assume that * preceded by another
			// * is illegal for these regexes.
			if p[pi] == '*' {
				return false
			}
			return isMatchRec(s, p[:pi], true, p[pi])
		}
		if s != "" && charMatches(s[si], p[pi]) {
			return isMatchRec(s[:si], p[:pi], false, '0')
		}
	}
	return false
}

func charMatches(strCh byte, patCh byte) bool {
	return strCh == patCh || patCh == '.'
}
