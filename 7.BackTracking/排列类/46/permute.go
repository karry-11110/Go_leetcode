package main

// func permuteUnique(nums []int) [][]int {
// 	var result [][]int
// 	var path []int
// 	var backtracking func(nums []int)
// 	backtracking = func(nums []int) {
// 		if len(nums) == 0 {
// 			result = append(result, append([]int{}, path...))
// 			return
// 		}
// 		for i := 0; i < len(nums); i++ {
// 			curr := nums[i]
// 			path = append(path, curr)
// 			nums = append(nums[:i], nums[i+1:]...) //直接使用切片
// 			backtracking(nums)
// 			nums = append(nums[:i], append([]int{curr}, nums[i:]...)...) //回溯的时候切片也要复原，元素位置不能变
// 			path = path[:len(path)-1]

// 		}
// 	}
// 	backtracking(nums)
// 	return result
// }
func permute(nums []int) [][]int {
	var result [][]int
	var path []int
	used := make([]bool, len(nums))
	var backtracking func(nums []int)
	backtracking = func(nums []int) {
		if len(nums) == len(path) {
			result = append(result, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] == true {
				continue
			}
			used[i] = true
			path = append(path, nums[i])

			backtracking(nums)

			path = path[:len(path)-1]
			used[i] = false

		}
	}
	backtracking(nums)
	return result
}
