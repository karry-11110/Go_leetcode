package 模拟行为

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	left, right, top, bottom := 0, n-1, 0, n-1
	startNum := 1
	for left <= right && top <= bottom {
		for col := left; col <= right; col++ {
			matrix[top][col] = startNum
			startNum++
		}
		for row := top + 1; row <= bottom; row++ {
			matrix[row][right] = startNum
			startNum++
		}
		if left < right && top < bottom {
			for col := right - 1; col > left; col-- {
				matrix[bottom][col] = startNum
				startNum++
			}
			for row := bottom; row > top; row-- {
				matrix[row][left] = startNum
				startNum++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return matrix
}
