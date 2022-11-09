package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 3}
	if len(nums) == 1 {
		fmt.Println(nums[0])
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if left+1 == right {
			fmt.Println(max(nums[left], nums[right]))
			return
		}
		if mid >= 1 {
			if nums[mid] > nums[mid-1] {
				left = mid
			} else {
				right = mid
			}
		}
	}
	return
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
