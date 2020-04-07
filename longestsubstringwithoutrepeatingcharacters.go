package leetcode

/*
  题目：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
  难度：中等
*/

// CountStr CountStr
func CountStr(s string) int {
	length := len(s)
	var maxLen, currLen int
	usedChar := make(map[string]int)
	i := 0
	for i < length {
		si := string(s[i])
		if _, ok := usedChar[si]; ok {
			maxLen = max(maxLen, currLen)
			currLen = 0
			i = usedChar[si] + 1
			usedChar = make(map[string]int)
			continue
		}
		usedChar[si] = i
		currLen = currLen + 1
		i = i + 1
	}

	return max(maxLen, currLen)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
