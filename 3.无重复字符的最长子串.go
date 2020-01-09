package main

import "fmt"

//双指针
func lengthOfLongestSubstring(s string) int {
	i, j := 0, 0

	result := 0
	for i < len(s) {
		tmp := make(map[byte]int)
		for j = i; j < len(s); {

			needBreak := false
			tmpStr := s[j]
			if num, ok := tmp[tmpStr]; !ok {
				//不存在
				tmp[tmpStr] = j
				j++
				if j == len(s) {
					i = j
				}
			} else {
	
				if i < len(s) {
					//已经存在
					i = num + 1
				}

				needBreak = true
			}

			if result < len(tmp) {
				result = len(tmp)
			}

			if (needBreak) {
				break
			}
		}
	}

	return result
}

func main() {
	s := " "

	result := lengthOfLongestSubstring(s)
	fmt.Println(result)

}
