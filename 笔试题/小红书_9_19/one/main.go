package main

import (
	"fmt"
	"strconv"
)

//趋势3
func main() {
	fmt.Println(solve("AABBBAAAABBAAABBBBAA", 17))
	//fmt.Println(work("AABBBAAAABBAAABBBB", 17))
	fmt.Println(solve("AABBBABAAA", 4))
}

func solve(coin_series string, k int) []string {
	m := map[string]int{}
	m[coin_series]++
	number2 := 0
	s := coin_series
	for m[work(s, k)] == 0 {
		for i := range m {
			m[i]++
		}
		m[work(s, k)]++
		number2++
		s = work(s, k)
	}
	res3 := 0
	for i := range m {
		res3 = max(res3, m[i])
	}
	res1 := work(s, k)
	res2 := m[work(s, k)]
	res := []string{res1}
	a := strconv.Itoa(res3 - res2)
	b := strconv.Itoa(res2)
	res = append(res, a, b)
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func work(str string, k int) string {
	i, j := k-1, k-1
	for i >= 0 || j <= len(str)-1 {
		flag := false
		if j <= len(str)-1 {
			if str[k-1] == str[j] {
				j++
				flag = true
			}
		}
		if i >= 0 {
			if str[k-1] == str[i] {
				i--
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	if i == -1 {
		return str
	}
	if j == len(str) {
		return str[i+1:] + str[0:i+1]
	}
	return str[i+1:j] + str[0:i+1] + str[j:]
}

//2
//func main() {
//	n, k := 0, 0
//	fmt.Scan(&n, &k)
//	nums := make([]int, n)
//	for i := 0; i < len(nums); i++ {
//		fmt.Scan(&nums[i])
//	}
//	dp := make([]int, n)
//	for i := 0; i < len(dp); i++ {
//		dp[i] = math.MaxInt64
//	}
//	dp[0] = 0
//	for i := 1; i < n; i++ {
//		j := i - 1
//		for j >= 0 && j >= i-k {
//			if nums[j] >= nums[i] {
//				dp[i] = min(dp[i], dp[j])
//			} else {
//				dp[i] = min(dp[i], dp[j]+nums[i]-nums[j])
//			}
//			j--
//		}
//	}
//	fmt.Println(dp[n-1])
//}
//func min(i, j int) int {
//	if i < j {
//		return i
//	}
//	return j
//}

//3
//var res [][2]int
//var res int
//
//func main() {
//	//n, k := 0, 0
//	//fmt.Scan(&n, &k)
//	//nums := make([]int, n)
//	////m := map[int]int{}
//	//for i := 0; i < n; i++ {
//	//	fmt.Scan(&nums[i])
//	//	//m[nums[i]]++
//	//}
//	//for start := 0; start < len(nums); start++ {
//	//
//	//	dfs(nums, start, start+1, k)
//	//}
//	//fmt.Println(res)
//	dp := make([]int, 30)
//	dp[0], dp[1], dp[2] = 1, 1, 2
//	dp[3] = 3
//	for i := 4; i < len(dp); i++ {
//		dp[i] = dp[i-1] + dp[i-3]
//	}
//	fmt.Println(dp[28])
//}

//func dfs(nums []int, start, end, k int) {
//	if end >= len(nums) || end+1 > len(nums) {
//		return
//	}
//	if check(nums[start:end+1], k) {
//		tmp := [2]int{start, end}
//		for i := 0; i < len(res); i++ {
//			if tmp == res[i] {
//				return
//			}
//		}
//		res = append(res, tmp)
//		return
//	}
//	for i := end + 1; i < len(nums); i++ {
//		dfs(nums, start, i, k)
//	}
//	return
//}

//func check(nums []int, k int) bool {
//	m := map[int]int{}
//	for i := 0; i < len(nums); i++ {
//		m[nums[i]]++
//	}
//	//求出了众数数组
//	maxcount, zhongshunums := 0, []int{}
//	for i := range m {
//		if m[i] > maxcount {
//			zhongshunums = zhongshunums[0:0]
//			zhongshunums = append(zhongshunums, i)
//			maxcount = m[i]
//		} else if m[i] == maxcount {
//			zhongshunums = append(zhongshunums, i)
//		}
//	}
//	if m[zhongshunums[0]] < k {
//		return false
//	} else {
//		return true
//	}
//	return true
//}

//1
//func main() {
//	a1, a2, a3, a4, x := 0, 0, 0, 0, 0
//	fmt.Scan(&a1, &a2, &a3, &a4, &x)
//	nums1 := make([]int, a1)
//	nums2 := make([]int, a2)
//	nums3 := make([]int, a3)
//	nums4 := make([]int, a4)
//	countA := 0
//	for i := 0; i < a1; i++ {
//		fmt.Scan(&nums1[i])
//		if nums1[i] > x {
//			countA++
//		}
//	}
//	countB := 0
//	for i := 0; i < a2; i++ {
//		fmt.Scan(&nums2[i])
//		if nums2[i] > x {
//			countB++
//		}
//	}
//	countC := 0
//	for i := 0; i < a3; i++ {
//		fmt.Scan(&nums3[i])
//		if nums3[i] > x {
//			countC++
//		}
//	}
//	countD := 0
//	for i := 0; i < a4; i++ {
//		fmt.Scan(&nums4[i])
//		if nums4[i] > x {
//			countD++
//		}
//	}
//	fmt.Println(min(min(countA, countB), min(countC, countD)))
//}
//func min(i, j int) int {
//	if i < j {
//		return i
//	}
//	return j
//}
