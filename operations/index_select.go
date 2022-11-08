package operations

type Matrix interface {
	[]int | [][]int
}

func IndexSelectWithDimensionAndIndex(arr [][]int, dimension int, index []int) [][]int {
	var result [][]int
	if dimension == 0 {
		for _, i := range index {
			result = append(result, arr[i])
		}
	} else if dimension == 1 {
		for _, i := range index {
			var temp []int
			for _, j := range arr {
				temp = append(temp, j[i])
			}
			result = append(result, temp)
		}
	}
	return result
}

//func IndexSelect[T Matrix](mat T, dimension int, index []int) T {
//	matrixType := fmt.Sprintf("%s", reflect.TypeOf(mat))
//	var selected T
//	if matrixType == "[][]int" {
//		if dimension == 0 {
//			for _, i := range index {
//				selected = append(selected, mat[i].(T))
//			}
//		} else if dimension == 1 {
//			for _, i := range index {
//				var temp []T
//				for _, j := range mat {
//					temp = append(temp, j[i])
//				}
//				selected = append(selected, temp...)
//			}
//		}
//	}
//	return selected
//}

// example of duplex type

//func sampler[T int | string](arg ...T) T {
//	var result T = arg[0]
//	return result
//}
