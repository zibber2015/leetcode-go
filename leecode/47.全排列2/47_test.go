package leetcode

import (
	"testing"
)

func TestA(t *testing.T) {
	a := []int{1,2,3,4}
	b := permuteUnique(a)
	t.Log(b)
}