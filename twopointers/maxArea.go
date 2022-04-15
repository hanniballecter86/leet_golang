package twopointers

func MaxArea(height []int) int {
	l, r := 0, len(height)-1
	maxarea := 0
	for l < r {
		maxarea = max(maxarea, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxarea
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
