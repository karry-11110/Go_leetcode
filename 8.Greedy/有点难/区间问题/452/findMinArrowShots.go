package main

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	//对左边界依次排序从小到大
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	//局部最优就是一支箭就可能射更多的气球，射完后就是跳过
	result := 1 //points不为空就至少需要一支
	for i := 1; i < len(points); i++ {
		if points[i][0] > points[i-1][1] { //两相邻气球不挨着，需要另外一支箭
			result++

		} else { //两气球挨着，就需要更新下下一次比较的右边界，这点就非常重要了！！！！！
			points[i][1] = min(points[i-1][1], points[i][1])
		}
	}
	return result
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
