package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	str := string(line)
	fmt.Println(solve(str))
}
func solve(str string) int {
	if len(str) < 5 {
		return 0
	}
	res := 0
	for i := 0; i <= len(str)-5; i++ {
		fmt.Println(str[i:i+5])
		if check(str[i:i+5]) {
			res++
		}
	}
	return res
}
func check(str string) bool {
	m := map[byte]int{}
	for i := 0; i < len(str); i++ {
		m[str[i]]++
	}
	if len(m) < 5 {
		return false
	}
	set := map[byte]bool {
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	if  !set[str[0]] && set[str[1]] && set[str[2]] && !set[str[3]] && set[str[4]] {
		return true
	}

		//if (str[0] != 'a' && str[0] != 'e' && str[0] != 'i' && str[0] != 'o' && str[0] != 'u') && (str[3] != 'a' && str[3] != 'e' && str[3] != 'i' && str[3] != 'o' && str[3] != 'u') {
		//	if str[1] == 'a' || str[1] == 'e' ||str[1] == 'i' || str[1] == 'o' || str[1] == 'u' {
		//		if str[2] == 'a' || str[2] == 'e' || str[2] == 'i' || str[2] == 'o' || str[2] == 'u' {
		//			if str[4] == 'a' || str[4] == 'e' || str[4] == 'i' || str[4] == 'o' || str[4] == 'u' {
		//				return true
		//			}
		//		}
		//	}
		//}


	return false
}



//package main
//
//import (
//	"fmt"
//	"bufio"
//	"os"
//	"strings"
//	"strconv"
//
//)
//var res int
//func main() {
//	//n, m := 0, 0
//	//fmt.Scan(&n, &m)
//	input := bufio.NewScanner(os.Stdin)
//	input.Scan()
//	number :=strings.Split(input.Text()," ")
//	n, _ := strconv.Atoi(number[0])
//	m, _ := strconv.Atoi(number[1])
//	arr := make([][]byte, n)
//	for i:= range arr {
//		arr[i] = make([]byte, m)
//	}
//	//input = bufio.NewScanner(os.Stdin)
//
//
//	for i := 0; i < n; i++ {
//		input.Scan()
//		str := input.Text()
//		for j := 0; j < m; j++ {
//
//			arr[i][j] = str[j]
//		}
//
//	}
//
//	is_go := make([][]int, n)
//	for i:= range is_go {
//		is_go[i] = make([]int, m)
//	}
//	is_go[0][0] = 1
//	res = n*m
//	dfs(0, 0, 0, n, m, arr, is_go)
//	fmt.Println(res)
//}
//func dfs(x, y, step, n, m int, arr [][]byte, is_go [][]int) {
//	if x == n-1 && y == m-1 {
//		if step < res {
//			res = step
//		}
//		return
//	}
//	if step >= res {
//		return
//	}
//	direc := [4][2]int{{-1,0},{1,0},{0,-1},{0,1}}
//	stop := map[byte]byte{
//		'r':'d',
//		'e':'r',
//		'd':'e',
//	}
//	for i := 0; i < 4; i++ {
//		x_ := x + direc[i][0]
//		y_ := y + direc[i][1]
//		if x_ >= 0 && x_ < n && y_ >= 0 && y_ < m && is_go[x_][y_] != 1 && arr[x_][y_] != stop[arr[x][y]] {
//			is_go[x_][y_] = 1
//			dfs(x_, y_, step+1, n, m, arr, is_go)
//			is_go[x_][y_] = 0
//		}
//	}
//
//}

//package main
//
//import (
//	"fmt"
//	"bufio"
//	"os"
//)
//
//func main() {
//	q :=  0
//	fmt.Scan(&q)
//	input := bufio.NewScanner(os.Stdin)
//	input.Scan()
//	for i := 0 ; i < q; i++ {
//		input.Scan()
//		str := input.Text()
//		fmt.Println(solve(str))
//	}
//}
//func solve(str string) string {
//	n_0, n_1 := 0, 0
//	for i := 0; i< len(str); i++ {
//		if str[i] == '0' {
//			n_0++
//		} else {
//			n_1++
//		}
//	}
//	if n_0 % 2 == 1 && n_1 % 2 == 1{
//		return "False"
//	}else {
//		return "True"
//	}
//}