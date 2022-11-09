package main

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var path []int
	//至少一个数字的被选数量不同，则两种组合是不同的。 组合问题则需要startIndex
	var backtracking func(candidates []int, target int, sum int, startIndex int)
	backtracking = func(candidates []int, target int, sum int, startIndex int) {
		if sum > target { //，和已经大了目标数，剪枝，一定是在递归终止的地方剪
			return
		}
		if sum == target {
			result = append(result, append([]int{}, path...))
			return
		}
		for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ { //剪枝操作
			path = append(path, candidates[i])       //处理
			sum = sum + candidates[i]                //处理
			backtracking(candidates, target, sum, i) //// 关键点:不用i+1了，表示可以重复读取当前的数
			sum = sum - candidates[i]                // 回溯
			path = path[:len(path)-1]                // 回溯
		}
	}
	backtracking(candidates, target, 0, 0)
	return result
}
