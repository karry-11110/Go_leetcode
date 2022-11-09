package main

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	//按绝对值大小从大到小排序
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})
	//从前向后遍历，遇到负数将其变为正数，同时K--
	for i := 0; i < len(nums); i++ {
		if k > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}
	//如果K还大于0，那么反复转变数值最小的元素，将K用完
	if k%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}
	//求和
	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result

}
