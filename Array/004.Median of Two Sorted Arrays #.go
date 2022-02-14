package Array

//1,4,8
//2,3,4,5,6,7

//8
//2,3,4,5,6,7

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	// nums1:  ……………… nums1[nums1Mid-1] | nums1[nums1Mid] ……………………
	// nums2:  ……………… nums2[nums2Mid-1] | nums2[nums2Mid] ……………………
	midNum := (len(nums1) + len(nums2) + 1) >> 1
	left, right, midOfNum1, midOfNum2 := 0, len(nums1), 0, 0
	for left <= right {
		midOfNum1 = (left + right) >> 1
		midOfNum2 = midNum - midOfNum1
		if midOfNum1 > 0 && nums1[midOfNum1-1] > nums2[midOfNum2] {
			right = midOfNum1 - 1
		} else if midOfNum1 < len(nums1) && nums2[midOfNum2-1] > nums1[midOfNum1] {
			left = midOfNum1 + 1
		} else {
			break
		}
	}
	leftNum, rightNum := 0, 0
	//奇数时 结果为leftNum
	//偶数时 为（leftNum+rightNum）/2

	//先定leftNum
	if midOfNum1 == 0 {
		//8
		//2,3,4,5,6,7
		leftNum = nums2[midOfNum2-1]
	} else if midOfNum2 == 0 {
		//3，4，5
		//6,7，8，9
		leftNum = nums1[midOfNum1-1]
	} else {
		leftNum = max(nums2[midOfNum2-1], nums1[midOfNum1-1])
	}
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(leftNum)
	}
	if midOfNum1 == len(nums1) {
		rightNum = nums2[midOfNum2]
	} else if midOfNum2 == len(nums2) {
		rightNum = nums1[midOfNum1]
	} else {
		rightNum = min(nums2[midOfNum2], nums1[midOfNum1])
	}
	return float64(leftNum+rightNum) / 2

}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
