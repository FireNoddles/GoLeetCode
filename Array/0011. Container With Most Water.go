package Array

func MaxArea(height []int) int {
	i, j, maxVal, temp := 0, len(height)-1, 0, 0
	for i < j {
		temp = min(height[i], height[j]) * (j - i)
		maxVal = max(maxVal, temp)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxVal
}

//func min(a int, b int) int {
//	if a > b {
//		return b
//	} else {
//		return a
//	}
//}
//
//func max(a int, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}
