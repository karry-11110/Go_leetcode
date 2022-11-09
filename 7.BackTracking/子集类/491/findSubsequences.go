package main

func findSubsequences(nums []int) [][]int {
	result := [][]int{}
	path := []int{}
	var backtracking func(nums []int, startIndex int)
	backtracking = func(nums []int, startIndex int) {
		if len(path) >= 2 {
			result = append(result, append([]int{}, path...))
		}
		used := [201]int{}
		for i := startIndex; i < len(nums); i++ {
			if (len(path) > 0 && nums[i] < path[len(path)-1]) || used[nums[i]+100] == 1 {
				continue
			}
			used[nums[i]+100] = 1 //记录本层元素是否重复使用，新的一层uset都会重新定义（清空），所以要知道uset只负责本层！
			path = append(path, nums[i])
			backtracking(nums, i+1)
			path = path[:len(path)-1]
		}
	}
	backtracking(nums, 0)
	return result
}

func findSubsequences(nums []int) [][]int {
	result := [][]int{}
	path := []int{}
	var backtracking func(nums []int, startIndex int)
	backtracking = func(nums []int, startIndex int) {
		if len(path) >= 2 {
			result = append(result, append([]int{}, path...))
		}
		used := [201]int{}
		for i := startIndex; i < len(nums); i++ {
			if (len(path) > 0 && nums[i] < path[len(path)-1]) || used[nums[i]+100] == 1 {
				continue
			}
			path = append(path, nums[i])
			backtracking(nums, i+1)
			path = path[:len(path)-1]
			used[nums[i]+100] = 1 //记录本层元素是否重复使用，新的一层uset都会重新定义（清空），所以要知道uset只负责本层！
		}
	}
	backtracking(nums, 0)
	return result
}
