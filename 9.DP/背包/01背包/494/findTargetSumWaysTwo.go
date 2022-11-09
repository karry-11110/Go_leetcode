//回溯解法,模仿二叉树的回溯解法
package main

func findTargetSumWays(nums []int, target int) int {
	var backtrack func(int, int)
	var count int
	backtrack = func(startIndex, sum int) {
		if startIndex == len(nums) {
			if sum == target {
				count++
			}
			return
		}
		backtrack(startIndex+1, sum+nums[startIndex])
		backtrack(startIndex+1, sum-nums[startIndex])
	}
	backtrack(0, 0)
	return count
}
