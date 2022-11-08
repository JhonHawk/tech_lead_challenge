package main

import (
	"fmt"
	"tech_lead_challenge/operations"
)

func main() {
	// Test Case 1 | Reshape possible
	arr := [][]int{{1, 2}, {3, 4}, {5, 6}}
	reshaped := operations.MatrixReshape(arr, 2, 3)
	fmt.Println("Reshaped Matrix: ", reshaped)

	// Test Case 2 | Hadamard Product
	arrHp1 := [][]int{{1, 2}, {3, 4}}
	arrHp2 := [][]int{{5, 6}, {7, 8}}
	hadamardProduct := operations.HadamardProduct(arrHp1, arrHp2)
	fmt.Println("Hadamard Product: ", hadamardProduct)

	// Test Case 3 | Index Select
	arrIs := [][]int{{1, 2}, {3, 4}, {5, 6}}
	//arrIs := []int{1, 2, 3, 4, 5, 6}
	dimension := 0
	index := []int{0, 1}
	indexSelect := operations.IndexSelectWithDimensionAndIndex(arrIs, dimension, index)
	//indexSelect := operations.IndexSelect(arrIs, dimension, index)
	fmt.Println("Index Select: ", indexSelect)
}
