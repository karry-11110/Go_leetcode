package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var resMap = map[int]int{}

func main() {
	n := 0
	fmt.Scanln(&n)
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()
	edge := make([][]int, n+1)
	for n-1 > 0 {
		n--
		input.Scan()
		s := input.Text()
		node1, _ := strconv.Atoi(strings.Split(s, " ")[0])
		node2, _ := strconv.Atoi(strings.Split(s, " ")[1])
		edge[node1] = append(edge[node1], node2)
		edge[node2] = append(edge[node2], node1)
	}
	res := solve(str, edge)
	fmt.Println(res)
}

func solve(str string, edge [][]int) int {
	resMap[1] = 1
	dfs(str, edge, 1)
	res := 0
	for _, v := range resMap {
		res += v
	}
	return res
}
func dfs(str string, edge [][]int, node int) {
	if len(resMap) == len(str) {
		return
	}
	childs := edge[node]
	for i := 0; i < len(childs); i++ {
		if _, ok := resMap[childs[i]]; !ok {
			if str[childs[i]-1] == str[node-1] {
				resMap[childs[i]] = resMap[node] + 1
			} else {
				resMap[childs[i]] = abs(resMap[node] - 1)
			}
			dfs(str, edge, childs[i])
		}
	}
	return
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
