package main

//滑动窗口解法：起始位置和结束位置如何移动，窗口内是什么。
func minSubArrayLen(target int, nums []int) int {
	result := len(nums) + 1
	for left, right, sum := 0, 0, 0; right < len(nums); right++ {
		sum += nums[right]
		for left <= right && sum >= target {
			length := right - left + 1
			if length < result {
				result = length
			}
			sum -= nums[left]
			left++
		}
	}
	if result == len(nums)+1 {
		return 0
	}
	return result
}
