package main

import (
	"fmt"
	"math"
)
var res int
func main() {
	nums := [][]int{{-1, 0, -1,67}, {0, -1, 0,-1}, {50, 100, 70,145},{80,200,50,98}}
	if len(nums)==0{
		fmt.Println(0)
	}
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums[i]))
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[0]); j++ {
			dp[i][j] = math.MinInt64
		}
		if nums[0][i] >= 0 {
			dp[0][i] = nums[0][i]
		}
	}

	if len(nums) == 1{
		fmt.Println(maxsums(dp[0]))
	}

	for row := 1; row < len(nums); row++{
		for col := 0; col < len(nums[0]); col++ {
			if check(nums, row, col){
				if check(nums, row-1, col-1) {
					if nums[row][col] >= 0 {
						if nums[row-1][col-1] == -1{
							dp[row][col] = max(dp[row-1][col-1]+nums[row][col], dp[row][col])
						}
					} else {
						if nums[row-1][col-1] == -1{
							dp[row][col] = max(dp[row-1][col-1], dp[row][col])
						}
					}
				}
				if check(nums, row-1, col+1) {
					if nums[row-1][col+1] == -1{
						if nums[row][col] >= 0 {
							if nums[row-1][col+1] == -1{
								dp[row][col] = max(dp[row-1][col+1]+nums[row][col], dp[row][col])
							}
						} else {
							if nums[row-1][col+1] == -1{
								dp[row][col] = max(dp[row-1][col+1], dp[row][col])
							}
						}
					}
				}
				if nums[row-1][col] >= 0 && dp[row-1][col] != math.MinInt64{
					if nums[row][col] >= 0 {
						dp[row][col] = max(dp[row-1][col]+nums[row][col], dp[row][col])
					} else {
						dp[row][col] = max(dp[row-1][col], dp[row][col])
					}
				}
			}
		}

	}
	fmt.Println(maxsums(dp[len(nums)-1]))
}
func maxsums(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res = max(res, nums[i])
	}
	return res
}
func check (nums [][]int, row, col int) bool {
	if row  < 0 || row >= len(nums) {
		return false
	}
	if col < 0 || col >= len(nums[0]) {
		return false
	}
	return true
}
func max(i, j int ) int {
	if i > j {
		return i
	}
	return j
}
