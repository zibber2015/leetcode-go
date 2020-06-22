package leetcode

import "fmt"

func permuteUnique(nums []int) [][]int {
	var list []int
	use := make([]bool, len(nums))
	result := backTrack(nums, len(nums),0, list , use)

	return result
}

func backTrack(nums []int, length int, depth int, list []int, use []bool) [][]int {
	var res [][]int

	//到底部返回当前路线
	if depth == length {
		res = append(res, list)
		return res
	}

	for i := 0; i < length; i++ {
		if use[i] {
			continue
		}

		list = append(list, nums[i])
		use[i] = true
		result := backTrack(nums, length, depth + 1, list, use)
		fmt.Println(list)
		fmt.Println(use)
		res = append(res, result...)
		//向上回溯
		list = list[:len(list)-1]
		use[i] = false
	}

	return res
}