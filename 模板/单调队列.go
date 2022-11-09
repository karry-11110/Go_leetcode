package main

func queue(nums []int) []int {
	result := make([]int, len(nums))
	st := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(st) != 0 && st[len(st)-1] <= nums[i] {
			st = st[:len(st)-1]
		}
		if len(st) != 0 {
			result[i] = st[len(st)-1]
		} else {
			result[i] = -1
		}
		st = append(st, nums[i])
	}
	return result
}
