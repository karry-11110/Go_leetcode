package main

import (
	"fmt"
	"sort"
)
func main() {
	fmt.Println(solve("fffffailedailedailedailedailed"))
}
func check(str, faield string) bool {
	strIndex, faieldIndex := 0, 0
	for ; strIndex < len(str); strIndex++ {
		if faieldIndex == len(faield) {
			return true
		}
		if str[strIndex] == faield[faieldIndex] {
			faieldIndex++
		}
	}
	if faieldIndex == len(faield) {
		return true
	}
	return false
}

func solve(str string) int {
	if len(str) < 6 {
		return -1
	}
	if len(str) % 6 != 0 {
		return -1
	}
	left, right := []int{}, []int{}
	for i := 0; i < len(str); i++ {
		if str[i] == 'f'{
			left = append(left, i)
		}
		if str[i] == 'd'{
			right = append(right, i)
		}
	}
	nums := [][]int{}
	for i := 0; i < len(left); i++{
		for j := 0; j < len(right); j++{
			if right[j]-left[i]+1 < 6{
				continue
			}
			flag := true
			if check(str[left[i]:right[j]+1], "failed") {
				for _, v := range nums {
					if left[i] == v[0] {
						flag = false
						break
					}
				}
				for _, v := range nums {
					if right[j] == v[1] {
						flag = false
						break
					}
				}
				if flag {
					nums = append(nums,[]int{left[i], right[j]})
				}
			}
		}
	}
	if len(nums) != len(str)/6 {
		return -1
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})
	fmt.Println(nums)
	res := 1
	max := 0
	for i := 1; i < len(nums); {
		if nums[i][0] < nums[i-1][1] {
			res++
			if res > max {
				max = res
			}
			nums = append(nums[0:i], nums[i+1:]...)
		} else {
			res = 0
			i++
		}
	}
	return max
}


