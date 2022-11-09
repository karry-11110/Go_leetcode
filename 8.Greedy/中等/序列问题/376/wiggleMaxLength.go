package main

func wiggleMaxLength(nums []int) int {
	//解决思路就是不统计单调路上的中间节点
	if len(nums) <= 1 {
		return len(nums)
	}
	pre := 0
	result := 1
	for i := 1; i < len(nums); i++ {
		curr := nums[i] - nums[i-1]
		if (curr > 0 && pre <= 0) || (curr < 0 && pre >= 0) {
			result++
			pre = curr
		}
	}
	return result
}
