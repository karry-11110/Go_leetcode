////1
//package main
//
////func main() {
////
////	fmt.Println(solve("1001101"))
////}
////func solve(str string) int {
////
////}
////1
////func minOperations(str string) int {
////	if len(str) <= 1 {
////		return 0
////	}
////	chars := []byte(str)
////	count0 := 0
////	for i := 0; i < len(chars); i++ {
////		if chars[i] == '0' {
////			count0++
////		}
////	}
////	count1 := len(chars) - count0
////	key := '0'
////	if count0 < count1 {
////		key = '1'
////	}
////	index, res := 0, 0
////	for index < len(chars) {
////		if chars[index] == byte(key) {
////			index++
////		} else {
////			index += 2
////			res++
////		}
////	}
////	return res
////}
//
////2
////package main
////
////import (
////	"fmt"
////)
////
////func main() {
////
////	fmt.Println(getSubarrayNum([]int{5, 2, 3, 50, 4, 3, 4, 6}, 2))
////}
////
////var res int64
////
////func getSubarrayNum(a []int, x int) int {
////	var key int64 = 1
////	for i := 0; i < x; i++ {
////		key *= 10
////	}
////	for i := 0; i < len(a); i++ {
////		for j := i + 1; j < len(a); j++ {
////			if mutil(a[i:j+1])%key == 0 {
////				res++
////			}
////		}
////	}
////	res = res % 1000000007
////	return int(res)
////}
////func mutil(a []int) int64 {
////	var res int64 = 1
////	for i := 0; i < len(a); i++ {
////		res *= int64(a[i])
////	}
////	return res
////}
//
////var res int64
////
////func getSubarrayNum(a []int, x int) int {
////	var key int64 = 1
////	for i := 0; i < x; i++ {
////		key *= 10
////	}
////	for start := 0; start < len(a); start++ {
////
////		dfs(a, start, int64(a[start]), key)
////
////	}
////	return int(res % 1000000007)
////}
////func dfs(nums []int, index int, sum, key int64) {
////	if sum%key == 0 && sum/key != 0 {
////		res = res + 1
////	}
////	i := index + 1
////	if i >= len(nums) {
////		return
////	}
////	dfs(nums, i, sum*int64(nums[i]), key)
////	return
////}
//
////3
//func numsOfGoodMatrix(n int, m int, x int) int {
//	var first, second, third, res, mo float64 = x*x*x*x/2, x*x/2,x / 2
//	first, second :=
//	third :=
//	res := 1
//	mo := 1e9 + 7
//	for i := 0; i < n-1; i++ {
//		for j := 0; j < m-1; j++ {
//			if i == 0 && j == 0 {
//				res = (res*first) % mo
//			} else if i == 0 || j == 0 {
//				res = (res * second) % mo
//			} else {
//				res = (res * third) % mo
//			}
//		}
//	}
//	return int(res)
//}

//3
//package main
//
//import "fmt"
//
//func main() {
//	n, x := 0, 0
//	fmt.Scan(&n, &x)
//	fmt.Print(solve(n, x))
//}
//func solve(n, x int) int64 {
//	if n == 0 {
//		return int64(1)
//	} else if n == 1 {
//		return int64(n * 2)
//	} else {
//		res := 2*int64(x)*solve(n-1, x) - 2*int64(n-1)*solve(n-2, x)
//		return res
//	}
//}

//2
package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	if n == 1 {
		nums := [4][5]string{
			{" ", " ", "*", " ", " "},
			{" ", "*", " ", "*", " "},
			{"*", " ", "*", " ", "*"},
			{" ", " ", "*", " ", " "},
		}
		for i := 0; i < len(nums); i++ {
			for j := 0; j < len(nums[0]); j++ {
				fmt.Print(nums[i][j])
			}
			fmt.Println()
		}
	}
	if n == 2 {
		nums := [8][11]string{
			{" ", " ", " ", " ", " ", "*", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", "*", " ", "*", " ", " ", " ", " "},
			{" ", " ", " ", "*", " ", "*", " ", "*", " ", " ", " "},
			{" ", " ", "*", " ", " ", " ", " ", " ", "*", " ", " "},
			{" ", "*", " ", "*", " ", " ", " ", "*", " ", "*", " "},
			{"*", " ", "*", " ", "*", " ", "*", " ", "*", " ", "*"},
			{" ", " ", " ", " ", " ", "*", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", "*", " ", " ", " ", " ", " "},
		}
		for i := 0; i < len(nums); i++ {
			for j := 0; j < len(nums[0]); j++ {
				fmt.Print(nums[i][j])
			}
			fmt.Println()
		}
	}
	if n == 2 {
		nums := [15][23]string{
			{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "*", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "*", " ", "*", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", " ", " ", " ", " ", "*", " ", "*", " ", "*", " ", " ", " ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", " ", " ", " ", "*", " ", " ", " ", " ", " ", "*", " ", " ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", " ", " ", "*", " ", "*", " ", " ", " ", "*", " ", "*", " ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", " ", "*", " ", "*", " ", "*", " ", "*", " ", "*", " ", "*", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", "*", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "*", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", "*", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "*", " ", " ", " ", " "},
			{" ", " ", " ", "*", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "*", " ", " ", " "},
			{" ", " ", "*", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "*", " ", " "},
			{" ", "*", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "*", " "},
			{"*", " ", "*", " ", "*", " ", "*", " ", "*", " ", "*", " ", "*", " ", "*", " ", "*", " ", "*", " ", "*", " ", "*"},
			{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "*", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "*", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "*", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		}
		for i := 0; i < len(nums); i++ {
			for j := 0; j < len(nums[0]); j++ {
				fmt.Print(nums[i][j])
			}
			fmt.Println()
		}
	}

}

//1
//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"sort"
//	"strings"
//)
//
//func main() {
//	input := bufio.NewScanner(os.Stdin)
//	input.Scan()
//	str := strings.Split(input.Text(), "#")
//	nums := make([][4]string, len(str))
//	for i := range nums {
//		s := strings.Split(str[i], ".")
//		nums[i] = [4]string{s[0], s[1], s[2], s[3]}
//	}
//	sort.Slice(nums, func(i, j int) bool {
//		if nums[i][0] != nums[j][0] {
//			return nums[i][0] < nums[j][0]
//		} else if nums[i][1] != nums[j][1] {
//			return nums[i][1] < nums[j][1]
//		} else if nums[i][2] != nums[j][2] {
//			return nums[i][2] < nums[j][2]
//		} else if nums[i][3] != nums[j][3] {
//			return nums[i][3] < nums[j][3]
//		} else {
//			return true
//		}
//	})
//	for i := 0; i < len(nums); i++ {
//		tmp := nums[i][0] + "." + nums[i][1] + "." + nums[i][2] + "." + nums[i][3]
//		if i != len(nums)-1 {
//			fmt.Println(tmp)
//		} else {
//			fmt.Print(tmp)
//		}
//	}
//}
