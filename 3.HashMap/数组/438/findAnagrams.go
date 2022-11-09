package main

import (
	"fmt"
	"sort"
)

// func main() {
// 	findAnagrams("cbaebabacd", "abc")
// }
// func findAnagrams(s string, p string) []int {
// 	c1 := []byte(p)
// 	sort.Slice(c1, func(i, j int) bool {
// 		return c1[i] < c1[j]
// 	})
// 	sByte := []byte(s)
// 	fmt.Println(sByte)
// 	lengthS := len(sByte)
// 	lengthP := len(p)
// 	result := []int{}
// 	for i := 0; i < lengthS; i++ {
// 		j := i + lengthP
// 		if j > lengthS {
// 			break
// 		}
// 		subByte := sByte[i:j]//sByte底层改变了，所以不能用排序
// 		fmt.Println(subByte)
// 		sort.Slice(subByte, func(i, j int) bool {
// 			return subByte[i] < subByte[j]
// 		})
// 		fmt.Println(subByte)
// 		if string(subByte) == string(c1) {
// 			result = append(result, i)
// 		}

// 	}
// 	return result
// }

func findAnagrams(s string, p string) []int {
	var a1 [26]int
	for _, value := range p {
		a1[value-'a']++
	}
	by := []byte(s)
	lenghtBy := len(by)
	lenghtp := len(p)
	result := []int{}
	for i := 0; i < lenghtBy; i++ {
		j := i + lenghtp
		if j > lenghtBy {
			break
		}
		subS := string(by[i:j])
		var a2 [26]int
		for _, value := range subS {
			a2[value-'a']++
		}
		if a1 == a2 {
			result = append(result, i)
		}

	}
	return result
}
