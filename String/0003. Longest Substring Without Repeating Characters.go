package String

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	dic := make(map[uint8]bool, len(s))
	maxLen, left, right := 1, 0, 1
	dic[s[left]] = true
	for right < len(s) {
		if _, ok := dic[s[right]]; ok {
			delete(dic, s[left])
			left++
		} else {
			dic[s[right]] = true
			right++
		}
		maxLen = max(right-left, maxLen)
	}
	return maxLen
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
