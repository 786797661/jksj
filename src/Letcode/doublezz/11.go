package doublezz

func maxArea(height []int) int {
	if height == nil {
		return 0
	}
	i := 0
	j := len(height) - 1
	maxarea := 0
	for i < j {
		area := 0
		if height[i] > height[j] {
			area = height[j] * (j - i)
			j--
		} else {
			area = height[i] * (j - i)
			i++
		}
		if area > maxarea {
			maxarea = area
		}

	}
	return maxarea
}
