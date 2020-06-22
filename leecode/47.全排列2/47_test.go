package leetcode

import (
	"testing"
)

func TestA(t *testing.T) {
	a := []int{1,1,2}
	b := permuteUnique(a)
	t.Log(b)
}