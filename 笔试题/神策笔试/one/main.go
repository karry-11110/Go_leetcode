package main

import "fmt"
func main() {
	fmt.Println(solve([][]int{{10,11,15,16},{12,14,17,22},{13,18,21,23},{19,20,24,25}}))
}
func solve(grid [][]int) []int {
	if len(grid) == 0 {
		return nil
	}
	if len(grid) == 1 {
		return []int{grid[0][0]}
	}
	res := []int{grid[0][0]}
	flag := 1
	for sum := 1; sum <= len(grid)-1; sum++ {
		if flag > 0 {
			for row := 0; row <= sum; row++ {
				res = append(res, grid[row][sum-row])
			}
			flag = -flag
		} else {
			for row := sum; row >= 0; row-- {
				res = append(res, grid[row][sum-row])
			}
			flag = -flag
		}
	}
	for sum := len(grid); sum <= 2*len(grid)-2 ; sum++ {
		if flag > 0 {
			for row := 0; row <= len(grid)-1; row++ {
				if sum-row <= len(grid)-1 {
					res = append(res, grid[row][sum-row])
				}
			}
			flag = -flag
		} else {
			for row := len(grid)-1; row >= 0; row-- {
				if sum-row <= len(grid)-1 {
					res = append(res, grid[row][sum-row])
				}
			}
			flag = -flag
		}
	}
	return res
}


