package main

func combinationSum3(k int, n int) [][]int {
	var result [][]int
	var path []int
	var backtracking func(k, sum, startIndex int)
	backtracking = func(k, sum, startIndex int) {
		if sum < 0 { //，和已经大了目标数，剪枝，一定是在递归终止的地方剪
			return
		}
		if sum == 0 && len(path) == k {
			result = append(result, append([]int{}, path...))
			return
		}
		for i := startIndex; i <= 9; i++ {
			path = append(path, i) //处理
			sum = sum - i          //处理
			backtracking(k, sum, i+1)
			sum = sum + i             // 回溯
			path = path[:len(path)-1] // 回溯
		}
	}
	backtracking(k, n, 1)
	return result
}
