package leetcode

// TODO: I'm sure there's a pattern that can be exploited in order to
// achieve a solution that doesn't take up any unnecessary space but
// I'll submit this for now.
func Convert(s string, numRows int) string {
	if numRows == 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}
	rows := [][]byte{}
	for i := 0; i < numRows; i++ {
		rows = append(rows, []byte{})
	}
	zigZag := 0
	inc := true
	for i := 0; i < len(s); i++ {
		rows[zigZag] = append(rows[zigZag], s[i])
		if inc && zigZag == numRows-1 {
			inc = false
			zigZag--
			continue
		}
		if inc {
			zigZag++
			continue
		}
		if zigZag == 0 {
			inc = true
			zigZag++
			continue
		}
		zigZag--
	}
	ret := []byte{}
	for i := 0; i < numRows; i++ {
		ret = append(ret, rows[i]...)
	}
	return string(ret)
}
