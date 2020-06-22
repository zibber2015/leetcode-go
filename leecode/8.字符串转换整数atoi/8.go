package leetcode

import (
	"fmt"
	"math"
)

func MyAtoi(str string) int {
	if str[0] == '-' || str[0] == '+' {
		str = str[1:]
	}

	sb := []byte(str)
	l := len(sb)
	fmt.Println(sb)
	result := 0
	for i := 0; i < l; i++ {
		result += (int(sb[i]) - 49) * int(math.Pow10(l - i - 1))
	}

	return result
}
