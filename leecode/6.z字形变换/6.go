package leetcode

import (
	"fmt"
)

func Convert(s string, numRows int) string {
	if s == "" {
		return s
	}

	if numRows == 1 {
		return s
	}

	z := make(map[int]map[int]string)
	i := 0
	j := 0
	l := 0
	down := true
	var k int
	for ; i < numRows; i++ {
		if !down {
			k = numRows - i - 1
		} else {
			k = i
		}

		if z[k] == nil {
			z[k] = make(map[int]string)
		}

		z[k][j] = string(s[l])
		fmt.Println(i, k, j, z[k][j])
		l++
		if i == numRows-1 {
			down = !down
			i = 0
		}

		if !down {
			j++
		}

		if l == len(s) {
			break
		}
	}

	//遍历
	var result string
	for x := 0; x <= numRows; x++  {
		for y := 0; y <= len(s); y++ {
			result = result + z[x][y]
		}
	}

	return result
}
