func movingCount(m int, n int, k int) int {
	board := make([][]bool, m)
	for i := range board {
		board[i] = make([]bool, n)
	}
	result := 0
	board[0][0] = true
	var check func(i int) int
	check = func(i int) int {
		a := i % 10
		b := i / 10
		for b != 0 {
			a += b % 10
			b = b / 10
		}
		return a
	}
	for i := 1; i < m; i++ {
		if check(i) <= k {
			board[i][0] = board[i-1][0]
			if board[i][0] {
				result++
			}
		}
	}
	for i := 1; i < n; i++ {
		if check(i) <= k {
			board[0][i] = board[0][i-1]
			if board[0][i] {
				result++
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if check(i)+check(j) <= k {
				board[i][j] = board[i-1][j] || board[i][j-1]
				if board[i][j] {
					result++
				}
			}
		}
	}
	return result + 1
}