package main

import "fmt"
func main() {
	r := solve("rgbbrg",[][2]int{{1,2},{2,3},{2,4},{4,5},{5,6}})
	fmt.Println(r)
}
func solve(str string, edge [][2]int) int {
	var check func(str string, zone []int) bool
	check = func(str string, zone []int) bool{
		color := [3]int{}
		for i := 0; i < len(zone); i++ {
			if str[zone[i]-1] == 'r' {
				color[0]++
			} else if str[zone[i]-1] == 'g' {
				color[1]++
			} else {
				color[2]++
			}
		}
		for i := 0; i < len(color); i++ {
			if color[i] == 0 {
				return false
			}
		}
		return true
	}
	var dfs func(str string, graph [][]int, index int, zone []int) bool
	dfs = func(str string, graph [][]int, index int, zone []int) bool {
		zone = append(zone, index)
		var k = -1
		for i := range graph[index]{
			if graph[index][i] == 1 {
				k = i
				break
			}
		}
		if k == -1 {
			zone2 := []int{}
			for i := 1; i <= len(str); i++ {
				flag := 1
				for j := 0; j < len(zone); j++ {
					if i == zone[j] {
						flag = 0
						break
					}
				}
				if flag == 1{
					zone2 = append(zone2, i)
				}
			}
			if check(str, zone) && check(str, zone2) {
				return true
			} else {
				return false
			}
		}
		graph[index][k] = 0
		graph[k][index] = 0
		res := dfs(str, graph, k, zone)
		graph[index][k] = 1
		graph[k][index] = 1
		return res
	}
	graph := make([][]int, len(str)+1)
	for i, _ := range graph {
		graph[i] = make([]int, len(str)+1)
	}
	for index := range edge {
		start, end := edge[index][0], edge[index][1]
		graph[start][end] = 1
		graph[end][start] = 1
	}
	fmt.Println(graph)
	res := 0
	for index := range edge {
		start, end := edge[index][0], edge[index][1]
		graph[start][end] = 0
		graph[end][start] = 0
		if dfs(str, graph, 1, []int{}) {
			fmt.Println(index)
			res++
		}
		graph[start][end] = 1
		graph[end][start] = 1
	}
	return res
}
