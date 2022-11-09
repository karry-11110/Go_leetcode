package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2,-1,-1}
	res := solve(nums)
	fmt.Println(res)
}
func solve(nums []int) int{
	sort.Slice(nums, func(i, j int) bool {
		if abs(nums[i]) == abs(nums[j]) {
			return nums[i] < nums[j]
		}
		return abs(nums[i]) < abs(nums[j])
	})
	res := 0
	for i := 0; i < len(nums); i++{
		if nums[i] == 0 {
			res++
			nums[i] = 1
		}
	}
	if sums(nums) == 0 {
		return res+1
	}
	return res

}
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func sums(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}
	return res
}
