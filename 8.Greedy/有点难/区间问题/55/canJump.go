// func canJump(nums []int) bool {
// 	if len(nums) <= 1 {
// 		return true
// 	}
// 	for i := 0; i < len(nums); {
// 		if nums[i] == 0 && i < len(nums)-1 {
// 			return false
// 		}
// 		start := i + 1
// 		maxValue := nums[start]
// 		end := start
// 		for j := start; j < start+nums[i]; j++ {
// 			if nums[j] > maxValue {
// 				maxValue = nums[j]
// 				end = j
// 			}
// 		}
// 		i = end
// 	}
// 	if i == len(nums) {
// 		return true
// 	}
// }
package main

func canJump(nums []int) bool {
	cover := 0
	if len(nums) <= 1 {
		return true
	}
	for i := 0; i <= cover; i++ { //注意这里是cover
		cover = max(i+nums[i], cover)
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
