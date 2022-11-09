package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
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
			// 若当前树层有使用过相同的元素，则跳过
			if i > startIndex && candidates[i] == candidates[i-1] { //前提是要将数组先排序
				continue
			}
			path = append(path, candidates[i])         //处理
			sum = sum + candidates[i]                  //处理
			backtracking(candidates, target, sum, i+1) ////
			sum = sum - candidates[i]                  // 回溯
			path = path[:len(path)-1]                  // 回溯
		}
	}
	sort.Ints(candidates)
	backtracking(candidates, target, 0, 0)
	return result
}

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	var path []int
	used := make([]bool, len(candidates))
	//至少一个数字的被选数量不同，则两种组合是不同的。 则需要startIndex
	var backtracking func(candidates []int, target int, startIndex int)
	backtracking = func(candidates []int, target int, startIndex int) {
		if target < 0 {
			return
		}
		if target == 0 {
			result = append(result, append([]int{}, path...))
			return
		}
		for i := startIndex; i < len(candidates); i++ {
			if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
				continue
			}
			path = append(path, candidates[i])
			target = target - candidates[i]
			used[i] = true
			backtracking(candidates, target, i+1)
			used[i] = false
			target = target + candidates[i]
			path = path[:len(path)-1]
		}

	}
	sort.Ints(candidates)
	backtracking(candidates, target, 0)
	return result
}
