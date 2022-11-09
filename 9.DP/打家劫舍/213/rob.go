package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	result1 := robRange(nums, 0)
	result2 := robRange(nums, 1)
	return max(result1, result2)
}
func robRange(nums []int, start int) int {
	dp := make([]int, len(nums)-1)
	dp[0] = nums[start]
	dp[1] = max(nums[start], nums[start+1])
	for i := 2; i < len(nums)-1; i++ {
		dp[i] = max(dp[i-2]+nums[i+start], dp[i-1])
	}
	return dp[len(nums)-2]
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
