package main

func searchRange(nums []int, target int) []int {
	var start, end int
	result1, ok1 := findFirst(nums, target)
	if ok1 {
		start = result1
	} else {
		start = -1
	}
	result2, ok2 := findEnd(nums, target)
	if ok2 {
		end = result2 - 1
	} else {
		end = -1
	}

	result := []int{start, end}
	return result
}
func findFirst(nums []int, target int) (int, bool) {
	low, high := 0, len(nums)
	result := len(nums)
	var tag bool
	for low < high {
		mid := (low + high) >> 1
		if target == nums[mid] {
			tag = true
			result = mid
			high = mid
		} else if target < nums[mid] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return result, tag
}
func findEnd(nums []int, target int) (int, bool) {
	low, high := 0, len(nums)
	result := len(nums)
	var tag bool
	for low < high {
		mid := (low + high) >> 1
		if target == nums[mid] {
			tag = true
			result = mid + 1
			low = mid + 1
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return result, tag
}
