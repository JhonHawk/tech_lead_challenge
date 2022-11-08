package operations

import "fmt"

func MatrixReshape(mat [][]int, r int, c int) [][]int {
	// Check if the reshape is possible
	if len(mat)*len(mat[0]) != r*c {
		fmt.Println("Cannot reshape matrix")
		return mat
	}
	var simple []int
	for _, row := range mat {
		for _, col := range row {
			simple = append(simple, col)
		}
	}
	//fmt.Println(simple)
	var reshaped [][]int
	for i := 0; i < r; i++ {
		var row []int
		for j := 0; j < c; j++ {
			row = append(row, simple[i*c+j])
		}
		reshaped = append(reshaped, row)
	}

	return reshaped
}
