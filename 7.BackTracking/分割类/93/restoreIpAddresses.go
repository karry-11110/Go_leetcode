package main //第一种写法
import "strconv"

func restoreIpAddresses(s string) []string {
	var result []string
	var path []string
	var backtracking func(s string, startIndex int)
	backtracking = func(s string, startIndex int) {
		if startIndex == len(s) && len(path) == 4 {
			tmp := path[0] + "." + path[1] + "." + path[2] + "." + path[3]
			result = append(result, tmp)

			return
		}
		for i := startIndex; i < len(s); i++ {
			//处理
			path = append(path, s[startIndex:i+1])
			if i-startIndex+1 <= 3 && len(path) <= 4 && isNormalIp(s, startIndex, i) {

				//递归
				backtracking(s, i+1)
			}
			//回溯
			path = path[:len(path)-1]
		}
	}
	backtracking(s, 0)
	return result
}
func isNormalIp(s string, startIndex, end int) bool {
	checkInt, _ := strconv.Atoi(s[startIndex : end+1])
	if end-startIndex+1 > 1 && s[startIndex] == '0' {
		//对于前导 0的IP（特别注意s[startIndex]=='0'的判断，不应该写成s[startIndex]==0，因为s截取出来不是数字）
		return false
	}
	if checkInt > 255 {
		return false
	}
	return true
}

//第二种写法

// func restoreIpAddresses(s string) []string {
// 	var result []string
// 	var path []string
// 	var backtracking func(s string, startIndex int)
// 	backtracking = func(s string, startIndex int) {
// 		if startIndex == len(s) && len(path) == 4 {
// 			tmp := path[0] + "." + path[1] + "." + path[2] + "." + path[3]
// 			result = append(result, tmp)
// 			return
// 		}
// 		for i := startIndex; i < len(s); i++ {

// 			if i-startIndex+1 <= 3 && len(path) <= 3 && isNormalIp(s, startIndex, i) {
// 				//处理
// 				path = append(path, s[startIndex:i+1])
// 			} else { //如果首尾超过了3个，或路径多余4个，或前导为0，或大于255，直接回退
// 				continue
// 			}
// 			//递归
// 			backtracking(s, i+1)
// 			//回溯
// 			path = path[:len(path)-1]
// 		}
// 	}
// 	backtracking(s, 0)
// 	return result
// }

// func isNormalIp(s string, startIndex, end int) bool {
// 	checkInt, _ := strconv.Atoi(s[startIndex : end+1])
// 	if end-startIndex+1 > 1 && s[startIndex] == '0' {
// 		//对于前导 0的IP（特别注意s[startIndex]=='0'的判断，不应该写成s[startIndex]==0，因为s截取出来不是数字）
// 		return false
// 	}
// 	if checkInt > 255 {
// 		return false
// 	}
// 	return true
// }
