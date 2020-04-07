package leetcode

import (
	"testing"
)

func TestCountStr(t *testing.T) {
	s := "aab"
	i := CountStr(s)
	t.Logf("expect: %d, act: %d", 6, i)
}
