//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func main() {
//	input := bufio.NewScanner(os.Stdin)
//	input.Scan()
//	s := input.Text()[1 : len(input.Text())-1]
//	str := strings.Split(s, ",")
//	nums := []int{}
//	for _, v := range str {
//		value, _ := strconv.Atoi(v)
//		nums = append(nums, value)
//	}
//	head := create(nums)
//	res := reverse(head)
//	fmt.Print("{")
//	for res != nil {
//		if res.Next != nil {
//			fmt.Printf("%d,", res.Val)
//		} else {
//			fmt.Printf("%d", res.Val)
//		}
//		res = res.Next
//	}
//	fmt.Print("}")
//}
//func create(nums []int) *ListNode {
//	if len(nums) == 0 {
//		return nil
//	}
//	h := &ListNode{}
//	p := h
//	for _, v := range nums {
//		p.Next = &ListNode{
//			Val: v,
//		}
//		p = p.Next
//	}
//	return h.Next
//}
//func reverse(head *ListNode) *ListNode {
//	return help(nil, head)
//}
//func help(pre *ListNode, curr *ListNode) *ListNode {
//	if curr == nil {
//		return pre
//	}
//	next := curr.Next
//	curr.Next = pre
//	return help(curr, next)
//
//}

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//func main() {
//	input := bufio.NewScanner(os.Stdin)
//	input.Scan()
//
//	str := strings.Split(input.Text(), ";")
//
//	s1 := str[0][1 : len(str[0])-1]
//	str1 := strings.Split(s1, ",")
//	nums1 := []int{}
//	for _, v := range str1 {
//		value, _ := strconv.Atoi(v)
//		nums1 = append(nums1, value)
//	}
//
//	s2 := str[1][1 : len(str[1])-1]
//	str2 := strings.Split(s2, ",")
//	nums2 := []int{}
//	for _, v := range str2 {
//		value, _ := strconv.Atoi(v)
//		nums2 = append(nums2, value)
//	}
//
//	res := solve(nums1, nums2)
//	fmt.Print("[")
//	for i := 0; i < len(res); i++ {
//		if i != len(res)-1 {
//			fmt.Printf("%d,", res[i])
//		} else {
//			fmt.Printf("%d", res[i])
//		}
//	}
//
//	fmt.Print("]")
//
//}
//func solve(nums1 []int, nums2 []int) []int {
//	res := []int{}
//	tmp := map[int]struct{}{}
//	for _, val := range nums2 {
//		if _, ok := tmp[val]; !ok {
//			tmp[val] = struct{}{}
//		}
//	}
//	for _, val := range nums1 {
//		if _, ok := tmp[val]; !ok {
//			res = append(res, val)
//		}
//	}
//	tmp = map[int]struct{}{}
//	for _, val := range nums1 {
//		if _, ok := tmp[val]; !ok {
//			tmp[val] = struct{}{}
//		}
//	}
//	for _, val := range nums2 {
//		if _, ok := tmp[val]; !ok {
//			res = append(res, val)
//		}
//	}
//	sort.Slice(res, func(i, j int) bool {
//		return res[i] < res[j]
//	})
//	return res
//}

//1
//package main
//
//import "fmt"
//
//func main() {
//	nums := []int{1, -1, 1, 1, 1, 1}
//
//	dp := make([]int, len(nums)+1)
//	dp[0] = 0
//	for i := 1; i <= len(nums); i++ {
//		dp[i] = nums[i-1] + dp[i-1]
//	}
//	m := map[int]int{}
//	res := 0
//	for i := range dp {
//		if _, ok := m[dp[i]]; !ok {
//			m[dp[i]] = i
//		} else {
//			res = max(res, i-m[dp[i]])
//		}
//	}
//
//	fmt.Println(res)
//}
//func max(i, j int) int {
//	if i > j {
//		return i
//	}
//	return j
//}

//
package main

import "fmt"

var path int

func main() {
	res := solve(6, 12)
	fmt.Println(res)
}
func solve(x, y int) int {
	distance := abs(x, y)
	if distance <= 3 {
		return distance
	}
	dfs(distance-1, 1, 1)
	return path
}
func dfs(dis int, step int, count int) {
	if step >= stiff(dis) {
		return
	}
	if dis <= 0 {
		return
	}
	if dis == 1 {
		if choice(step)[0] == 1 || choice(step)[1] == 1 || choice(step)[2] == 1 {
			path = min(path, count+1)
		} else {
			return
		}
	}
	cho := choice(step)

	dfs(dis-cho[0], cho[0], count+1)
	dfs(dis-cho[1], cho[1], count+1)
	dfs(dis-cho[2], cho[2], count+1)
	return
}
func stiff(dis int) int {
	tar := dis / 2
	for i := 1; i < tar/2; i++ {
		tar = tar - i
		if tar <= 0 {
			return i
		}
	}
	return dis / 2
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func choice(i int) [3]int {
	return [3]int{i, i + 1, i - 1}
}
func abs(x, y int) int {
	if x-y > 0 {
		return x - y
	}
	return y - x
}
