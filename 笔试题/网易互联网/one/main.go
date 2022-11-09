package main

import (
	"fmt"
	"strconv"
)
var res []int
func main(){
	n := 10
	k := 6
	t := 3
	if t >= k {
		fmt.Println("-1")
	}
	nums := make([]int, 0)
	res = make([]int, n)
	if dfs(nums, n, k, t) == nil {
		fmt.Println("-1")
	}
	str := ""
	for i := 0; i < n; i++ {
		s := strconv.Itoa(res[i])
		str += s
	}
	fmt.Println(str)
}
func dfs(nums []int, n, k, t int) []int {
	if len(nums) > n {
		return nil
	}
	if len(nums) == n {
		if k == 0 && t == 0 {
			copy(res, nums)
			return res
		}
	}
	nums = append(nums, 1)
	k--
	if len(nums) >= 2 {
		if nums[len(nums)-2] == 1 {
			t--
		}
	}
	if dfs(nums, n, k, t) != nil {
		return dfs(nums, n, k, t)
	}
	nums = nums[:len(nums)-1]
	k++
	if len(nums) >= 1{
		if nums[len(nums)-1] == 1{
			t++
		}
	}
	nums = append(nums, 0)
	if dfs(nums, n, k, t) != nil {
		return dfs(nums, n, k, t)
	}
	nums = nums[:len(nums)-1]
	return nil
}



