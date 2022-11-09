package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	var (
		result [][]int
		path   []int
	)
	var backtracking func(nums []int, startIndex int)
	backtracking = func(nums []int, startIndex int) {
		result = append(result, append([]int{}, path...)) //存放结果
		if startIndex > len(nums) {                       //注意截至条件
			return
		}
		for i := startIndex; i < len(nums); i++ {
			if i > startIndex && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backtracking(nums, i+1)
			path = path[:len(path)-1]
		}
	}
	sort.Ints(nums)
	backtracking(nums, 0)
	return result
}

func subsetsWithDup(nums []int) [][]int {
	var (
		result [][]int
		path   []int
	)
	used := make([]bool, len(nums))
	var backtracking func(nums []int, startIndex int)
	backtracking = func(nums []int, startIndex int) {
		result = append(result, append([]int{}, path...)) //存放结果
		if startIndex > len(nums) {                       //注意截至条件
			return
		}
		for i := startIndex; i < len(nums); i++ {
			if i > 0 && used[i-1] == false && nums[i] == nums[i-1] { //注意这个条件
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtracking(nums, i+1)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	sort.Ints(nums)
	backtracking(nums, 0)
	return result
}

func subsetsWithDup(nums []int) [][]int {
	var (
		result [][]int
		path   []int
	)

	var backtracking func(nums []int, startIndex int)
	backtracking = func(nums []int, startIndex int) {
		result = append(result, append([]int{}, path...)) //存放结果
		if startIndex > len(nums) {                       //注意截至条件
			return
		}
		used := make([]bool, len(nums))
		for i := startIndex; i < len(nums); i++ {
			if i > 0 && used[i-1] == true && nums[i] == nums[i-1] { //注意这个条件
				used[i] = true
				continue
			}
			path = append(path, nums[i])
			backtracking(nums, i+1)
			path = path[:len(path)-1]
			used[i] = true
		}
	}
	sort.Ints(nums)
	backtracking(nums, 0)
	return result
}
