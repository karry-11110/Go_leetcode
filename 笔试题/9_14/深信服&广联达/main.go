////package main
////func minDistance( a string ,  b string ) int {
////  m, n := len(a), len(b)
////  dp := make([][]int, m+1)
////  for i := range dp {
////	  dp[i] = make([]int, n+1)
////  }
////  for i := 0 ;i < m+1; i++ {
////	  dp[i][0] = i
////  }
////	for i := 0 ;i < n+1; i++ {
////		dp[0][i] = i
////	}
////	for i := 1; i < m+ 1; i++ {
////		for j := 1; j < n+1; j++ {
////			if a[i-1] == b[j-1] {
////				dp[i][j] =dp[i-1][j-i]
////			} else {
////				dp[i][j] = Min(dp[i][j-1],dp[i-1][j],dp[i-1][j-1]) + 1
////			}
////		}
////	}
////	return dp[m][n]
////}
////func Min(args ...int) int {
////	min := args[0]
////	for _, item := range args {
////		if item < min {
////			min = item
////		}
////	}
////	return min
////}
//
//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	n, x := 0, 0
//	fmt.Scan(&n, &x)
//	nums := make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Scan(&nums[i])
//	}
//
//	res := solve(nums, x)
//	fmt.Println(res)
//
//}
//func solve(nums []int, x int) int {
//	max, min := nums[0], nums[0]
//
//	times := 0
//	for i := 0; i < len(nums); i++ {
//		if max-min > 2*x {
//			max =  nums[i-1]
//			min = max
//			times += 1
//		}
//		max = Max(max, nums[i])
//		min = Min(min, nums[i])
//	}
//	if max-min > 2*x {
//		times += 1
//	}
//	return times
//}
//func Max(i, j int ) int {
//	if i > j {
//		return i
//	}
//	return j
//}
//func Min(i, j int ) int {
//	if i < j {
//		return i
//	}
//	return j
//}
//
//
//
//
//
//
//
//
//
//
package main

import (
	"fmt"
	"bufio"
	"os"

)

func main() {
	t := 0
	//reader := bufio.NewReader(os.Stdin)
	//strt, _, _ := reader.ReadLine()
	//t , _ = strconv.Atoi(string(strt))
	fmt.Scan(&t)
	reader := bufio.NewReader(os.Stdin)
	str, _, _ := reader.ReadLine()
	for i := 0; i < t; i++ {
		str, _, _ = reader.ReadLine()
        res := solve(str)
		if res {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
func solve(str []byte) bool {

	count := 0
	st := []byte{}
	for i := 0; i < len(str); i++ {
		if len(st) == 0 {
			st = append(st, str[i])
		} else {
			if str[i] == st[len(st)-1] {
				st = st[:len(st)-1]
				count++
			} else {
				st = append(st, str[i])
			}
		}
	}
	return count%2 != 0
}


