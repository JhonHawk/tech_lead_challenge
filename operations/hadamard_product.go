package operations

func HadamardProduct(mat1 [][]int, mat2 [][]int) [][]int {
	var product [][]int
	for i, row := range mat1 {
		var rowProduct []int
		for j, col := range row {
			rowProduct = append(rowProduct, col*mat2[i][j])
		}
		product = append(product, rowProduct)
	}
	return product
}
