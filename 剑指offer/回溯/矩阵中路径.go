func exist(board [][]byte, word string) bool {

	str := []byte(word)
	var backtracking func(row, col, start int) bool
	backtracking = func(row, col, start int) bool {
		if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || board[row][col] != str[start] {
			return false
		}
		if start == len(str) {
			return true
		}
		board[row][col] = ' '
		result := backtracking(row-1, col, start+1) || backtracking(row+1, col, start+1) || backtracking(row, col-1, start+1) || backtracking(row, col+1, start+1)
		board[row][col] = str[start]
		return result
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if backtracking(i, j, 0) {
				return true
			}
		}
	}
	return false
}