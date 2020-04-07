package leetcode

import (
	"testing"
)

func TestDecode(t *testing.T) {
	s := "12343520342344553"
	i := Decode(s)
	t.Logf("expect: %d, act: %d", 6, i)
}

func TestDecode1(t *testing.T) {
	s := "12343520342344553"
	i := Decode1(s)
	t.Logf("expect: %d, act: %d", 6, i)
}
