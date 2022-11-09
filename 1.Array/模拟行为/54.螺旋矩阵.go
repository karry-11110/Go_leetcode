package 模拟行为

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	var (
		order                    = make([]int, len(matrix)*len(matrix[0]))
		index                    = 0
		left, right, top, bottom = 0, len(matrix[0]) - 1, 0, len(matrix) - 1
	)
	for left <= right && top <= bottom {
		for col := left; col <= right; col++ {
			order[index] = matrix[top][col]
			index++
		}
		for row := top + 1; row <= bottom; row++ {
			order[index] = matrix[row][right]
			index++
		}
		if left < right && top < bottom { //当最里层不是一个环的时候就没必要了
			for col := right - 1; col > left; col-- {
				order[index] = matrix[bottom][col]
				index++
			}
			for row := bottom; row > top; row-- {
				order[index] = matrix[row][left]
				index++
			}
		}

		left++
		right--
		top++
		bottom--
	}
	return order
}
