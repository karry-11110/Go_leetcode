package main

//递归
// func countSubstrings(s string) int {
// 	result := 0
// 	dp := make([][]bool, len(s))
// 	for i := 0; i < len(s); i++ {
// 		dp[i] = make([]bool, len(s))
// 	}
// 	for i := len(s) - 1; i >= 0; i-- {
// 		for j := i; j <= len(s)-1; j++ {
// 			if s[i] == s[j] {
// 				if j-i <= 1 {
// 					result++
// 					dp[i][j] = true
// 				} else if dp[i+1][j-1] {
// 					result++
// 					dp[i][j] = true
// 				}
// 			}
// 		}
// 	}
// 	return result
// }

//双指针法
func countSubstrings(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		result += extend(s, i, i, len(s))
		result += extend(s, i, i+1, len(s))
	}
	return result
}
func extend(s string, i, j, length int) int {
	subResult := 0
	for i >= 0 && j < length && s[i] == s[j] {
		i--
		j++
		subResult++
	}
	return subResult
}
