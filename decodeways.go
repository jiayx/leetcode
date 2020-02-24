package main

import (
	"strconv"
	"strings"
)

func main() {
	s := "226"
	decode(s)
}

func decode(s string) int {
	length := len(s)
	arr := make([]int, length+1)
	arr[0] = 0
	if s[0] == '0' {
		arr[1] = 0
	} else {
		arr[1] = 1
	}

	i := 2
	for i <= length {
		var fix int
		if i == 2 {
			fix = 1
		} else {
			fix = arr[i-2]
		}
		arr[i] = arr[i-1]*between(s[i-1:i]) + fix*between(s[i-2:i])
		i = i + 1
	}
	return arr[length]
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
