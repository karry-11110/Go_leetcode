package main

import "fmt"
var res int
func main() {
	nums := [][]int{{-1, 0, -1}, {0, -1, 0}, {50, 100, 70},{80,200,50}}
 	row :=  0
	 res = 0
	 value :=  0
	for col := 0; col < len(nums[0]); col++ {
		dfs(nums, row, col, value)
	}
	fmt.Println(res)
}
func dfs(nums [][]int, row, col, value int) {
	if row == len(nums)-1 {
		if nums[row][col] > 0 {
			res = max(res, value+nums[row][col])
		} else {
			res = max(res, value)
		}
	    return 
	}
	if nums[row][col] >= 0 {
		dfs(nums, row+1, col, value+nums[row][col])
	} else {
		if col + 1 <= len(nums[0])-1 {
			dfs(nums, row+1, col+1, value)
		}
		if col -1 >= 0 {
			dfs(nums, row+1, col-1, value)
		}
	}
}
func max(i, j int ) int {
	if i > j {
		return i
	}
	return j
}
