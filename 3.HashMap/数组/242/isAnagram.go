package main

import "sort"

//map作为哈希表
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := map[rune]int{}
	for _, value := range s {
		m[value]++
	}
	for _, value := range t {
		m[value]--
		if m[value] < 0 {
			return false
		}
	}
	return true
}

//数组作为哈希表
// func isAnagram(s string, t string) bool {
// 	a1, a2 := [26]int{}, [26]int{}
// 	for _, value := range s {
// 		a1[value-'a']++
// 	}
// 	for _, value := range s {
// 		a2[value-'a']++
// 	}
// 	return a1 == a2
// }

// //将自负床排序后来比较，注意要想将字符串转化为byte数组再来排序
// func isAnagram(s string, t string) bool {
// 	c1, c2 := []byte(s), []byte(t)
// 	sort.Slice(c1, func(i, j int) bool { return c1[i] < c1[j] })
// 	sort.Slice(c2, func(i, j int) bool { return c2[i] < c2[j] })
// 	return string(c1) == string(c2)
// }
