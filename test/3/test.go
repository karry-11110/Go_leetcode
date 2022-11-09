package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	resultLeft, resultRight := 0, 0
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			resultLeft = mid
			right = mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			resultRight = mid
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	fmt.Println(resultRight, resultLeft)
	return resultRight - resultLeft + 1
}
func main() {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
}
