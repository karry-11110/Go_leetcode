//回溯
// func combine(n int, k int) [][]int {
// 	var result [][]int
// 	var path []int
// 	var backtracking func(n, k, startIndex int)
// 	backtracking = func(n, k, startIndex int) {
// 		if len(path) == k { //终止条件
// 			// tmp := make([]int, k)
// 			// copy(tmp, path)
// 			result = append(result, append([]int{}, path...))
// 			// result = append(result, path) //存放结果,这种存放方式大错特错，因为你存路径的是引用值
// 			return
// 		}
// 		for i := startIndex; i <= n; i++ {
// 			path = append(path, i)    //处理节点
// 			backtracking(n, k, i+1)   //递归
// 			path = path[:len(path)-1] //回溯，撤销处理的节点
// 		}
// 	}
// 	backtracking(n, k, 1)
// 	return result
// }

//减枝后效果
package main

import "fmt"

func main() {
	combine(4, 2)
}
func combine(n int, k int) [][]int {
	var result [][]int
	var path []int
	var backtracking func(n, k, startIndex int)
	backtracking = func(n, k, startIndex int) {
		if len(path) == k { //终止条件
			// tmp := make([]int, k)
			// copy(tmp, path)
			result = append(result, append([]int{}, path...))
			fmt.Println(result)
			// result = append(result, path) //存放结果,这种存放方式大错特错，因为你存路径的是引用值
			return
		}
		if startIndex > (n-k)+len(path)-1 {
			return
		}
		for i := startIndex; i <= n; i++ {
			path = append(path, i)    //处理节点
			backtracking(n, k, i+1)   //递归
			path = path[:len(path)-1] //回溯，撤销处理的节点
		}
	}
	backtracking(n, k, 1)
	return result
}
