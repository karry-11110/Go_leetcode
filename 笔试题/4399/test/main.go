package main

import "fmt"



func main() {

	res := convert("abcdefghijklmn", 4)
	fmt.Println(res)
}
func convert(s string, numCols int) [][]byte {
	n, c := len(s), numCols
	rows := n / c
	if n % c != 0 {
		rows ++
	}
	strs := make([][]byte, rows)
	for i := 0; i < len(strs); i++ {
		strs[i] = make([]byte, c)
	}
	i, j := 0, 0
	t := numCols * 2 - 2
	fmt.Println(strs)
	for index, v := range s {
		strs[i][j] = byte(v)
		if index % t < numCols - 1 {
			j++
		} else {
			i++
			j--
		}
	}
	return strs
}




//func solve(str string) bool {
//	k := 1
//	left, right := 0, len(str)-1
//	for left < right {
//		if str[left] != str[right] {
//			if k > 0 {
//				if str[left] == str[right-1] {
//					k--
//					right--
//				} else if str[left+1] == str[right] {
//					k--
//					left++
//				} else {
//					return false
//				}
//			} else {
//				return false
//			}
//		} else {
//			left++
//			right--
//		}
//	}
//	return true
//}
//func solve(curr, pos, neg int, path []string) {
//	if curr < 0 {
//		return
//	}
//	if curr == 0 {
//		if pos == 0 && neg == 0 {
//			fmt.Println(path)
//			res++
//		}
//		return
//	}
//	if neg >= 1 {
//		path = append(path, "neg")
//		solve(curr-1, pos, neg-1, path)
//		path = path[:len(path)-1]
//	}
//	if pos >= 1 {
//		path = append(path, "pos")
//		solve(curr*2, pos-1, neg, path)
//		path = path[:len(path)-1]
//	}
//	return
//}
