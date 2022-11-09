package main

func subsets(nums []int) [][]int {
	var (
		result [][]int
		path   []int
	)
	var backtracking func(nums []int, startIndex int)
	backtracking = func(nums []int, startIndex int) {
		result = append(result, append([]int{}, path...)) //存放结果 收集全部节点
		if startIndex > len(nums) {                       //注意截至条件
			return
		}
		for i := startIndex; i < len(nums); i++ {
			path = append(path, nums[i])
			backtracking(nums, i+1)
			path = path[:len(path)-1]
		}
	}
	backtracking(nums, 0)
	return result
}
