package main

import "fmt"
func main() {
	row, col := 0, 0
	fmt.Scanln(&row, &col)
	grid := make([][]int, row)
	for index,_ := range grid{
		grid[index] = make([]int, col)
		for i := 0; i < col; i++ {
			m := 0
			fmt.Scan(&m)
			grid[index][i] = m
		}
	}
	fmt.Println(solve(grid))
}
func solve(grid [][]int) []int {
	if grid == nil {
		return nil
	}
	if len(grid) == 0 {
		return []int{0, 0}
	}
	count, max := 0, 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++{
			if grid[row][col] == 1 {
				tmp := make([][]int, 0)
				count++
				m := dfs(row, col, tmp, grid, max)
				max = Max(m, max)
			}
		}
	}
	result := []int{count, max}
	return result
}
func dfs(row, col int, tmp, grid [][]int, max int) int {
	if !inArea(grid, row, col) {
		return max
	}
	if grid[row][col] == 0 {
		return max
	}
	for _, value := range tmp {
		max = Max(max, abs(row-value[0]) + abs(col-value[1]))
	}
	tmp = append(tmp, []int{row, col})
	grid[row][col] = 0
	max = Max(dfs(row-1, col, tmp, grid, max), dfs(row+1, col, tmp, grid, max))
	max = Max(dfs(row, col-1, tmp, grid, max), dfs(row, col+1, tmp, grid, max))
	return max
}
func inArea(grid [][]int, row, col int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0])
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}