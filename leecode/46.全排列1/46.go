package leetcode


//var l sync.RWMutex

func permute(nums []int) [][]int {
	var list []int
	var res [][]int
	use := make([]bool, len(nums))
	backtrack(nums, len(nums), 0, list, use, &res)

	return res
}

func backtrack(nums []int, length int, depth int, list []int, use []bool, res *[][]int) {
	//到底部返回当前路线
	if depth == length {
		*res = append(*res, list)
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
		backtrack(nums, length, depth+1, newList, use, res)
		//向上回溯
		list = list[:(len(list) - 1)]
		use[i] = false
	}
}
