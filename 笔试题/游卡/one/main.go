package main

import "fmt"
func main() {
	fmt.Println(solve([]int{2,1,6,2,2,3,2,4}))
}
func solve(nums []int ) int {
	if len(nums) / 4 == 0 {
		return 0
	}
	n := len(nums)/4
	k := 2*n
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for i := 0; i < n; i++ {
		if
	}

	fmt.Println(m)
	if len(m) >= k {
		return k
	}
	return len(m)
}
