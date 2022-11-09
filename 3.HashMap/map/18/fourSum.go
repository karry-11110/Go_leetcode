package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	//首先要判断数组是否有四个元素
	if len(nums) < 4 {
		return nil
	}
	//对数组进行排序
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			n2 := nums[j]
			if j > i+1 && n2 == nums[j-1] {
				continue
			}
			left := j + 1
			right := len(nums) - 1
			for left < right {
				n3, n4 := nums[left], nums[right]
				sum := n1 + n2 + n3 + n4
				if sum < target {
					left++
				} else if sum > target {
					right--
				} else {
					result = append(result, []int{n1, n2, n3, n4})
					//去重，继续寻找
					for left < right && nums[left] == n3 {
						left++
					}
					for left < right && nums[right] == n4 {
						right--
					}

				}
			}
		}
	}
	return result
}
