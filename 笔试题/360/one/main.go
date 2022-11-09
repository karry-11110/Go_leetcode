
//
//package main
//
//import (
//	"fmt"
//	"sort"
//)
//
//func main() {
//	n, m := 0, 0
//	fmt.Scan(&n, &m)
//	numsN := make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Scan(&numsN[i])
//	}
//	numsM := make([][2]int, m)
//	for i := 0; i < m; i++ {
//		fmt.Scan(&numsM[i][0], &numsM[i][1])
//	}
//	tmp := make([]int, b)
//	for i := 0; i < len(numsM); i++ {
//		a, b := numsM[i][0], numsM[i][1]
//		tmp := make([]int, b)
//		for j := 0; j < b; j++ {
//			tmp[j] = numsN[j]
//		}
//		solve(tmp, a)
//		for j := 0; j < b; j++ {
//			numsN[j] = tmp[j]
//		}
//	}
//	for i, _ := range numsN {
//		fmt.Print(numsN[i])
//		if i != len(numsN)-1 {
//			fmt.Print(" ")
//		}
//	}
//}
//func solve(nums []int, t int) []int {
//	sort.Ints(nums)
//	if t == 0 {
//		return nums
//	} else {
//		reverse(nums)
//		return nums
//	}
//}
//func reverse(nums []int) {
//	left, right := 0, len(nums)-1
//	for left < right {
//		nums[left], nums[right] = nums[right], nums[left]
//		left++
//		right--
//	}
//}