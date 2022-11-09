package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4, 5, 4, 6, 3, 56, 2, 6}
	res := solve(input, 56)
	fmt.Println(res)
}

func solve(nums []int, target int) int {
	left, right := 0, len(nums)-1
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if nums[mid] == target {
		return target
	}
	if solve(nums[left:mid], target) != -1 {
		return solve(nums[left:mid+1], target)
	}
	if solve(nums[mid+1:], target) != -1 {
		return solve(nums[mid+1:], target)
	}
	return -1
}
