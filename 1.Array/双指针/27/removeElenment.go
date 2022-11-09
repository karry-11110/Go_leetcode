package main

func removeElement(nums []int, val int) int {

	slow := 0
	var n = len(nums)
	for fast := 0; fast < n; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
