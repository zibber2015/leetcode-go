package leetcode

import (
	"fmt"
	"strconv"
)

//去重不是最优解
func permuteUnique(nums []int) [][]int {
	var list []int
	var res [][]int
	use := make([]bool, len(nums))
	distinct := make(map[string]bool)
	backtrack(nums, len(nums), 0, list, use, distinct, &res)

	return res
}

func backtrack(nums []int, length int, depth int, list []int, use []bool, distinct map[string]bool, res *[][]int) {
	//到底部返回当前路线
	if depth == length {
		s := join(list)
		if check, ok := distinct[s]; ok && check {
			return
		}
		*res = append(*res, list)
		distinct[s] = true
		fmt.Println(distinct)
		return
	}

	for i := 0; i < length; i++ {
		if use[i] {
			continue
		}
		list = append(list, nums[i])
		newList := make([]int, len(list))
		copy(newList, list)
		use[i] = true
		backtrack(nums, length, depth+1, newList, use, distinct, res)
		//向上回溯
		list = list[:(len(list) - 1)]
		use[i] = false
	}
}

func join(arr []int) string {
	s := ""
	for _,i:=range arr {
		s = s + strconv.Itoa(i)
	}
	return s
}
