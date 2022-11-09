package main

func candy(ratings []int) int {
	candy := make([]int, len(ratings))
	result := 0
	for i := 0; i < len(ratings); i++ {
		candy[i] = 1
	}
	//从前向后
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candy[i] = candy[i-1] + 1
		}
	}
	//这里要从后往前遍历这个维度好点
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candy[i] = max(candy[i], (candy[i+1] + 1)) //非常重要
		}
	}
	for i := 0; i < len(ratings); i++ {
		result += candy[i]
	}
	return result
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
