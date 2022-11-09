package main

func travel(matrix [][]int, startCol, endCol, startRow, endRow, target int) bool {
	if endCol < startCol || startRow > endRow {
		return false
	}
	centerCol, centerRow := (startCol+endCol)/2, (startRow+endRow)/2

	if target == matrix[centerCol][centerRow] {
		return true
	} else if target < matrix[centerCol][centerRow] {
		return travel(matrix, centerCol, centerCol, startRow, centerRow-1, target) || travel(matrix, startCol, centerCol-1, startRow, centerRow, target)
	} else {
		return travel(matrix, centerCol+1, endCol, startRow, endRow, target) || travel(matrix, startCol, centerCol, centerRow+1, endRow, target)
	}
	return false
}
func findNumberIn2DArray(matrix [][]int, target int) bool {

	return travel(matrix, 0, len(matrix)-1, 0, len(matrix[0])-1, target)
}

func main() {
	var slice [][]int
	slice = append(slice, []int{1, 4, 7, 11, 15}, []int{2, 5, 8, 12, 19}, []int{3, 6, 9, 16, 22}, []int{10, 13, 14, 17, 24}, []int{18, 21, 23, 26, 30})

	findNumberIn2DArray(slice, 5)
}
