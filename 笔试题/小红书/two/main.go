package main

import (
	"fmt"
	"math"
)
var res int
func main() {
	nums := []int{-6, 0, 2, -2, 3}
	res = math.MaxInt64
	for i := 0; i < len(nums); i++ {
		dfs(nums, i, 1, 0)
		dfs(nums, i, -1, 0)
	}
	fmt.Println(res)
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func dfs(nums []int, index, chosen, k int)  {
	if check(nums) == 8 {
		res = min(res, k)
		return
	}
	nums[index] += chosen
	k++
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			if check(nums) == 0 || check(nums) == -8{
				break
			} else if check(nums) > 8 ||(check(nums) < 0 && check(nums) > -8) {
				dfs(nums, i, 1, k)
			} else {
				dfs(nums, i, -1, k)
			}
		} else if nums[i] > 0 {
			if check(nums) == 0 || check(nums) == -8{
				break
			} else if check(nums) > 8 ||(check(nums) < 0 && check(nums) > -8) {
				dfs(nums, i, -1, k)
			} else {
				dfs(nums, i, 1, k)
			}
		} else {
			dfs(nums, i, 1, k)
			dfs(nums, i, -1, k)
		}
	}
	return
}
func check(nums []int) int {
	res := 1
	for i := 0; i < len(nums); i++ {
		res *= nums[i]
	}
	return res
}
