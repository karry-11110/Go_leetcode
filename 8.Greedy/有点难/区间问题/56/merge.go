package main

import "sort"

func merge(intervals [][]int) [][]int {
	//从小到大排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//针对重复区间做操作
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= intervals[i-1][1] {
			intervals[i-1][1] = max(intervals[i-1][1], intervals[i][1])
			intervals = append(intervals[:i], intervals[i+1:]...)
			i-- //非常重要，有一个元素相当于没有了
		}
	}
	return intervals
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
