package leetcode

import (
	"strconv"
	"strings"
)

// 题目：https://leetcode-cn.com/problems/decode-ways
// 难度：中等

// Decode 动态规划
func Decode(s string) int {
	length := len(s)
	arr := make([]int, length+1)
	arr[0] = 1
	arr[1] = between(string(s[0]))

	i := 2
	for i <= length {
		arr[i] = arr[i-1]*between(s[i-1:i]) + arr[i-2]*between(s[i-2:i])
		i++
	}
	return arr[length]
}

// Decode1 简化为用两个变量存值
func Decode1(s string) int {
	length := len(s)

	prev := 1
	curr := between(string(s[0]))

	i := 2
	for i <= length {
		temp := curr
		curr = curr*between(s[i-1:i]) + prev*between(s[i-2:i])
		prev = temp
		i++
	}
	return curr
}

func between(s string) int {
	if strings.HasPrefix(s, "0") {
		return 0
	}
	an, _ := strconv.Atoi(s)

	if an > 0 && an <= 26 {
		return 1
	}
	return 0
}
