package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "12343520342344553"
	i := decode(s)
	fmt.Println(i)
}

func decode(s string) int {
	length := len(s)
	arr := make([]int, length+1)
	arr[0] = 1
	if s[0] == '0' {
		arr[1] = 0
	} else {
		arr[1] = 1
	}

	i := 2
	for i <= length {
		arr[i] = arr[i-1]*between(s[i-1:i]) + arr[i-2]*between(s[i-2:i])
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
