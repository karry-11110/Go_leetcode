package main

import "fmt"
func main() {
	fmt.Println(solve([]int{1,2,3,1}, 3))
}
func solve(nums []int, k int) bool {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++  {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = i
		} else {
			if abs(i, m[nums[i]]) <= k {
				return true
			}
		}
	}
	return false
}
func abs(i, j int) int {
	if i-j < 0 {
		return j-i
	} else{
		return j-i
	}
}
