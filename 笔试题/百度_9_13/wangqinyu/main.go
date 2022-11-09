//package main
//
//import "fmt"
//
//func main() {
//	nums := [2]int{50, 50}
//	fmt.Println(solve(nums))
//}
//func solve(nums [2]int) int {
//	n, m := nums[0], nums[1]
//	dp := make([][]int, n)
//	for i := range dp {
//		dp[i] = make([]int, m)
//	}
//	dp[0][0] = 1
//	for i := 1; i < m; i++ {
//		res := 0
//		k := i
//		res += dp[0][i-1]
//		k = k - 2
//		for k > 0 {
//			res += dp[0][k-1]
//			k = k - 2
//		}
//		dp[0][i] = res
//	}
//	for i := 1; i < n; i++ {
//		res := 0
//		k := i
//		res += dp[i-1][0]
//		k = k - 2
//		for k > 0 {
//			res += dp[k-1][0]
//			k = k - 2
//		}
//		dp[i][0] = res
//	}
//	for i := 1; i < n; i++ {
//		for j := 1; j < m; j++ {
//			res := 0
//			row, col := 0, 0
//			k, v := j, i
//			k = k - 2
//			for k > 0 {
//				row += dp[i][k-1]
//				k = k - 2
//			}
//			v = v - 2
//			for v > 0 {
//				col += dp[v-1][j]
//				v = v - 2
//			}
//			res += row + col + dp[i-1][j] + dp[i][j-1]
//			dp[i][j] = res % (1000000007)
//		}
//	}
//	return dp[len(dp)-1][len(dp[0])-1]
//}

package main

import "fmt"

func main() {
	nums := []int{998, 570, 876, 200, 877}
	s := make([]int, 0)
	res := 0
	for i := 0; i < len(nums); i++ {
		if len(s) == 0 || nums[i] < nums[s[len(s)-1]] {
			s = append(s, i)
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		tmp := make([]int, len(s))
		copy(tmp, s)
		for len(tmp) != 0 {
			index := tmp[len(tmp)-1]
			tmp = tmp[:len(tmp)-1]
			if i == index {
				continue
			}
			res = max(res, nums[i]+nums[index]-(abs(i-index)-1))
		}
	}
	fmt.Println(res)
}
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
