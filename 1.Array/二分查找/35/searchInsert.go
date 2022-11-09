package main

//普通思路：
func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)
	for low < high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
// 	return low
// }

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

//同时可以理解为找第一个大于等于target的下标
func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)
	result := len(nums)
	for low < high {
		mid := (low + high) >> 1
		if target <= nums[mid] {
			result = mid
			high = mid
		} else {
			low = mid + 1
		}
	}
	return result
}

//找最后一个小于等于target的下标
func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)
	result := len(nums)
	for low < high {
		mid := (low + high) >> 1
		if target >= nums[mid] {
			result = mid + 1
			low = mid + 1
		} else {
			high = mid
		}
	}
	return result
}
