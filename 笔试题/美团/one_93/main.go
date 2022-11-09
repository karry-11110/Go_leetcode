//package main
//
//import (
//	"fmt"
//
//)
//func main() {
//	a, b := 0, 0
//	fmt.Scan(&a, &b)
//	if b < 10 {
//		fmt.Println(11-a)
//	} else {
//		fmt.Println(b+2-a)
//	}
//}

//package main
//
//import (
//	"fmt"
//)
//func main() {
//	n := 0
//	fmt.Scan(&n)
//	nums := make([]int, n)
//	maxNumber := 0
//	for i := 0; i < n; i++ {
//		fmt.Scan(&nums[i])
//		maxNumber = max(nums[i], maxNumber)
//	}
//	solve(nums, maxNumber)
//
//}
//func solve(nums []int, maxNumber int) {
//	help := make([]int, maxNumber+1)
//	for i := 0; i < len(nums); i++ {
//		help[nums[i]] = 1
//	}
//	for _, num := range nums {
//		help[num]--
//		for i := 0; i < maxNumber+1; i++{
//			if help[i] == 0 {
//				if i != maxNumber {
//					fmt.Print(i)
//					fmt.Print(" ")
//				} else {
//					fmt.Print(i)
//				}
//				break
//			}
//		}
//		help[num]++
//	}
//}
//func max(i, j int) int {
//	if i > j {
//		return i
//	}
//	return j
//}


package main

import (
	"fmt"
	"bufio"
	"os"
)
func main() {
	n := 0
	fmt.Scan(&n)
	nums := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Scan(&nums[i])
	}
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	str := string(line)
	graph := make([][]int,n + 1)
	for i := 1; i < n + 1; i++ {
		for j := 0; j < len(nums); j++ {
			if i == nums[j] {
				graph[i] = append(graph[i], j+2)
			}
		}
	}
	path := make([][]int, n+1)
	copy(path, graph)
	dfs(graph, 1,  path)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = sum(path[i+1], str)
		fmt.Print(res[i])
		if i == n-1 {
			break
		} else {
			fmt.Print(" ")
		}
	}
}
func sum(path []int, str string) int {

	s := []byte(str)
	m := map[byte]int{}
	for i := 0; i < len(path); i++ {
		m[s[path[i]-1]] = 1
	}
	res := 0
	for  range m {
		res++
	}
	return res
}
func dfs(graph [][]int, index int,  path [][]int) {
	path[index] = append(path[index], index)
	if graph[index] == nil {
		return
	}
	for i := 0; i < len(graph[index]); i++ {
		//length := len(graph[graph[index][i]])
		//path[index] = append(path[index], graph[graph[index][i]]...)
		dfs(graph, graph[index][i],path)
		//
		//path[index] = path[index][:len(path[index])-length]
	}
	for i := 0; i < len(graph[index]); i++ {
		//length := len(graph[graph[index][i]])
		path[index] = append(path[index], path[graph[index][i]]...)

		//path[index] = appen 56d(path[index], graph[index][i])
		//
		//path[index] = path[index][:len(path[index])-length]
	}

	return
}