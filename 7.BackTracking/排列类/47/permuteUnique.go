package main

import "sort"

func permuteUnique(nums []int) [][]int {

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
			// used[i - 1] == true，说明同一树枝nums[i - 1]使用过
			// used[i - 1] == false，说明同一树层nums[i - 1]使用过
			// 如果同一树层nums[i - 1]使用过则直接跳过
			if i > 0 && used[i-1] == false && nums[i] == nums[i-1] { // i>0必须写在最前面
				continue
			}
			if used[i] == false {
				used[i] = true
				path = append(path, nums[i])
				backtracking(nums)
				path = path[:len(path)-1]
				used[i] = false
			}

		}
	}
	sort.Ints(nums)
	backtracking(nums)
	return result

}
