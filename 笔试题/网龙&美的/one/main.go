package main

import (
	"fmt"
	"bufio"
	"os"
)
func main() {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	str := reader.Text()
	res := 0
	m := make(map[byte]int)
	dp := make([]int, len(str))
	dp[0] = 1
	m[str[0]] = 0
	for i := 1; i < len(str); i++ {
		preidx, ok := m[str[i]]
		if !ok {
			dp[i] = dp[i-1] +i + 1
		} else {
			dp[i] = dp[i-1] + (i-preidx)
		}
		m[str[i]] = i
	}
	for _, v := range dp{
		res += v
	}
	fmt.Println(res)
}
//package main
//
//import "fmt"
//func main() {
//	reader := bufio.NewReader(os.Stdin)
//	line, _, _ := reader.ReadLine()
//	str := string(line)
//	sum, end := 0, [25]int{}
//	for i := range end {
//		end[i] = -1
//	}
//	res := 0
//	for index , value := range str {
//		value -= 'a'
//		sum += index - end[value]
//		res += sum
//		end[value] = index
//	}
//	fmt.Println(res)
//}


package main

import (
"fmt"
"bufio"
"os"
"strings"
"strconv"
"sort"
)
func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := strings.Split(input.Text(), ",")
	m := make(map[string]string)
	for i := 0; i < len(str); i++ {
		v:= strings.Split(str[i], "|")
		m[v[1]] = v[0]
	}

	input.Scan()
	str2 := strings.Split(input.Text(), ",")
	res := make(map[string]int)
	for i := 0; i < len(str2); i++ {
		v:= strings.Split(str2[i], "#")
		tmp, _ := strconv.Atoi(v[0])
		s:= strings.Split(v[1], "|")
		for j := 0; j < len(s); j++ {
			res[m[s[j]]] += tmp
		}
	}
	arr := make([]string,0)

	for i:= range res {
		arr = append(arr, i)
	}
	sort.Strings(arr)
	for _, v := range arr {
		bufio.
		fmt.Print(v)
		fmt.Print(" ")
		fmt.Print(res[v])
		fmt.Println()
	}
}