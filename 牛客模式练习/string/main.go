package main

import (
	"bufio"
	"os"
	"fmt"
	"sort"

)
func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()
	fmt.Println(permutation(str))
}
func permutation(s string) []string {
	str := []byte(s)
	sort.Slice(str, func(i, j int) bool {
		return str[i] < str[j]
	})
	var result []string
	var path []byte
	used := make([]bool, len(str))
	var backtracking func(str []byte, start int, path []byte)
	backtracking = func(str []byte, start int, path []byte) {
		if len(path) == len(str){
			result = append(result, string(path))
			return
		}
		for i := 0; i < len(str); i++ {
			if i > 0 && str[i] == str[i-1] && used[i-1] == false {  //树层去重
				continue
			}
			if used[i] == true { // 树枝去重
				continue
			}
			used[i] = true
			path = append(path, str[i])
			backtracking(str, start+1, path)
			path = path[:len(path)-1]
			used[i] = false
		}
		return
	}
	backtracking(str, 0, path)
	return result
}

