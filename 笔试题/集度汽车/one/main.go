package main

import (
	"fmt"
	"sort"
)
//1
func main() {
	intervals := make([][2]int, 0)

	var a, b int
	for {
		n, _ := fmt.Scan(&a, &b)
		if n == 0 {
			break
		}
		intervals = append(intervals, [2]int{a, b})
	}
	fmt.Println(solve(intervals))
}
func solve(intervals [][2]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= intervals[i-1][1] {
			intervals[i-1][1] = max(intervals[i-1][1], intervals[i][1])
			intervals = append(intervals[:i], intervals[i+1:]...)
			i--
		}
	}
	res := 0
	for i := 0; i < len(intervals); i++ {
		res += (intervals[i][1] - intervals[i][0])
	}
	return res
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//2
//func main() {
//	input := bufio.NewScanner(os.Stdin)
//	bs := make([]byte, 1024*2000)
//	input.Buffer(bs, len(bs))
//	input.Scan()
//	s1 := input.Text()
//	input.Scan()
//	s2 := input.Text()
//	fmt.Println(solve(s1, s2))
//}
//
//func solve(s, t string) int {
//	m, n := len(s), len(t)
//	dp := make([][]int, m+1)
//	for i := range dp {
//		dp[i] = make([]int, n+1)
//	}
//	for i := 0; i < m+1; i++ {
//		dp[i][0] = i
//	}
//	for j := 0; j < n+1; j++ {
//		dp[0][j] = j
//	}
//	for i := 1; i < m+1; i++ {
//		for j := 1; j < n+1; j++ {
//			if s[i-1] == t[j-1] {
//				dp[i][j] = dp[i-1][j-1]
//			} else {
//				dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
//			}
//		}
//	}
//	return dp[m][n]
//}
//func min(i, j int) int {
//	if i < j {
//		return i
//	}
//	return j
//}
