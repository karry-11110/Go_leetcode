package main

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	currCover, nextCover := 0, 0
	result := 0
	for i := 0; i < len(nums); i++ {
		nextCover = max(nums[i]+i, nextCover) // 更新下一步覆盖最远距离下标
		if i == currCover {                   // 遇到当前覆盖最远距离下标
			if currCover != len(nums)-1 { // 如果当前覆盖最远距离下标不是终点
				result++                      // 需要走下一步
				currCover = nextCover         // 更新当前覆盖最远距离下标（相当于加油了）
				if nextCover >= len(nums)-1 { // 下一步的覆盖范围已经可以达到终点，结束循环
					break
				}

			} else { // 当前覆盖最远距离下标是集合终点，不用做ans++操作了，直接结束
				break
			}
		}
	}
	return result
}
func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
