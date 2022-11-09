package main

//处理下一个元素更大的问题
//首先需要明白单调栈里存放的是元素下标（当前元素右边比较大的下标），结果数组里存放的是右边第一个比自己大的元素在自己右边的第几个
//单调栈始终存放元素下标，使用直接nums[i]，只不过在存放之前可能会淘汰一些下标
func nextGreaterElement(nums []int) []int {
	result, stack := make([]int, len(nums)), []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) != 0 && nums[i] >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 { //淘汰一批后发现没有比当前元素大的了
			result[i] = 0
		} else {
			result[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i) //单调栈始终存放元素下标，使用直接nums[i]
	}
	return result
}
