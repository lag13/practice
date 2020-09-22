package leetcode

func MaxArea(heights []int) int {
	largest := 0
	for i := 0; i < len(heights); i++ {
		largest = maxAreaRec(heights[i:], largest)
	}
	return largest
}

func maxAreaRec(heights []int, largest int) int {
	if len(heights) == 0 || len(heights) == 1 {
		return largest
	}
	hl := len(heights) - 1
	minHeight := min(heights[0], heights[hl])
	largest = max(largest, minHeight*hl)
	if heights[0] > minHeight && heights[0]*(hl-1) > largest {
		largest = maxAreaRec(heights[:hl], largest)
	}
	return largest
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
