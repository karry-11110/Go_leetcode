package main

import "sort"

func threeSum(nums []int) [][]int {
	//定一双指针移2
	result := [][]int{}
	sort.Ints(nums)
	//这里的循环针对的是定1的那个元素
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		//如果排序后的数组定1的元素为大于0的话，那么就不可能有满足的情况了，及时止损
		if n1 > 0 {
			break
		}
		//怎样排除重复的情况发生呢，就是定1的元素如果有重复的那就跳过，换下一位把
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		//开始我们的正式过程
		left, right := i+1, len(nums)-1
		for left < right {
			n2, n3 := nums[left], nums[right]
			if n1+n2+n3 == 0 {
				result = append(result, []int{n1, n2, n3})
				//继续找看是否还有满足条件的,为什么要用for，防止重复条件出现,达到了去重和同时靠近的效果
				for left < right && nums[left] == n2 {
					left++
				}
				for left < right && nums[right] == n3 {
					right--
				}
			} else if n1+n2+n3 > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return result
}
