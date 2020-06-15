package leetcode

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	mid := 0
	even := false
	if totalLength%2 != 0 {
		mid = (totalLength + 1) / 2
	} else {
		mid = totalLength / 2
		even = true
	}

	i := 0
	j := 0
	result := 0
	num := 0
	var ni int
	var nj int
	for {

		if i >= len(nums1) - 1 {
			ni = 0
		} else {
			ni = nums1[i]
		}

		if j >= len(nums2) - 1 {
			nj = 0
		} else {
			nj = nums2[j]
		}

		if ni > nj {
			i++
			num++

			if num == mid {
				result += nums1[i]
				if !even {
					break
				}
			} else if num == (mid + 1) {
				result += nums1[i]
				break
			}
		} else {
			j++
			num++

			if num == mid {
				result += nums1[j]
				if !even {
					break
				}
			} else if num == (mid + 1) {
				result += nums1[j]
				break
			}
		}
	}

	return float64(result)
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	result := findMedianSortedArrays(nums1, nums2)
	fmt.Println(result)
}