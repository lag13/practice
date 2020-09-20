package leetcode

func IsMatch(s string, p string) bool {
	return isMatchRec(s, p, false, '0')
}

func isMatchRec(s string, p string, star bool, starCh byte) bool {
	if s == "" && p == "" {
		return true
	}
	if star && s != "" && charMatches(s[0], starCh) {
		// We consume as much of the character as possible
		// (i.e. greedy) but if it doesn't work out then we
		// have to try consuming one less character (and one
		// less and one less etc...) until it either works or
		// doesn't. We won't know until we try.
		if isMatchRec(s[1:], p, star, starCh) {
			return true
		}
	}
	if p != "" {
		// I'm assuming that a '*' character must always come
		// AFTER a non-* character
		if p[0] == '*' {
			return false
		}
		if len(p) > 1 && p[1] == '*' {
			return isMatchRec(s, p[2:], true, p[0])
		}
		if s != "" && charMatches(s[0], p[0]) {
			return isMatchRec(s[1:], p[1:], false, '0')
		}
	}
	return false
}

func charMatches(strCh byte, patCh byte) bool {
	return strCh == patCh || patCh == '.'
}
