package leetcode

func rotate(matrix [][]int)  {
	l := len(matrix)
	for i := 0; i < l - 1; i++ {
		a := matrix[i][l]
		matrix[i][0] = a
		b := matrix[l][l]
		matrix[i][l] = b
	}
}


func main() {
	matrix := [][]int{}
	matrix = append(matrix, []int{5, 1, 9, 11})
	matrix = append(matrix, []int{2, 4, 8,10})
	matrix = append(matrix, []int{13, 3, 6, 7})
	matrix = append(matrix, []int{15,14,12,16})
	rotate(matrix)
}