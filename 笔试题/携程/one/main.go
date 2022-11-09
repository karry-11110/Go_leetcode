package main

import (
	"fmt"
	"sort"
)

func main() {
	t := 0
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		n, k, x := 0, 0, 0
		fmt.Scan(&n, &k, &x)
		nums := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&nums[j])
		}
		fmt.Println(solve(nums, k, x))
	}
}
func solve(nums []int, k, x int) int {
	Max1, Max2 := 0, 0
	for i := 0; i <= len(nums)-k-1; i++ {
		r := check(i, k, nums)
		x = x - r[0] //剩余次数
		if x < 0 {
			return -1
		}
		Max1 = max(Max1, r[1])
	}
	for i := len(nums) - k; i <= len(nums)-2; i++ {
		Max2 = max(Max2, nums[i])
	}

	return max(Max1+(x/k), Max2+x)
}
func check(index, k int, nums []int) []int {
	count := 0
	tmp := []int{}
	for i := index; i < len(nums); i += k {
		tmp = append(tmp, nums[i])
	}
	sort.Ints(tmp)
	for i := 0; i < len(tmp); i++ {
		count += abs(tmp[len(tmp)-1], tmp[i])
	}
	for i := index; i < len(nums); i += k {
		nums[i] = tmp[len(tmp)-1]
	}

	return []int{count, tmp[len(tmp)-1]}
}
func abs(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//func main() {
//	nums := []int{1, 2, 1, 1}
//	solve(nums, 3, 1)
//}
//func solve(nums []int, k, x int) {
//
//}

//func main() {
//
//	x := "2330000466"
//	res := solve(x)
//	fmt.Println(res)
//
//}
//func solve(x string) int {
//	if x < "233" {
//		return -1
//	}
//	tmp, _ := strconv.Atoi(x)
//	start := float64(tmp)
//	key := 233 * math.Pow10(len(x)-3)
//	res := 0
//	if start < key {
//		key = key / 10
//	}
//	if start == key {
//		return 1
//	}
//	for start-key > 0 {
//		start = start - key
//		res++
//		strStart := strconv.Itoa(int(start))
//		if strStart < "233" {
//			return -1
//		} else if strStart == "233" {
//			return res + 1
//		}
//		key = 233 * math.Pow10(len(strStart)-3)
//		if start < key {
//			key = key / 10
//		}
//	}
//	if start == key {
//		return res
//	}
//	return -1
//}
