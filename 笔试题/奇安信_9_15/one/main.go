package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	bs := make([]byte, 2000*1024)
	input.Buffer(bs, len(bs))
	input.Scan()
	A := strings.Split(input.Text(), ",")
	input.Scan()
	B := strings.Split(input.Text(), ",")
	numsA, numsB := make([]int, len(A)), make([]int, len(B))
	for i := 0; i < len(A); i++ {
		numsA[i], _ = strconv.Atoi(A[i])
	}
	for i := 0; i < len(B); i++ {
		numsB[i], _ = strconv.Atoi(B[i])
	}
	r := solve(numsA, numsB)
	fmt.Println(r)
}

func solve(numsA, numsB []int) int {
	res = 10000000
	dfs(0, []int{}, []int{}, len(numsA), numsA, numsB)
	return res
}
func dfs(i int, buy_a, buy_b []int, n int, A, B []int) {
	if i == n {
		sort.Slice(buy_a, func(i , j int)bool {
			return buy_a[i] > buy_b[i]
		})
		sort.Slice(buy_b, func(i , j int)bool {
			return buy_b[i] > buy_b[i]
		})
		tmp := []int{}
		he := 0
		for i := 0; i< len(buy_a); i++ {
			tmp = append(tmp, buy_a[i])
			if len(tmp) == 3{
				he += int(sum(tmp)*6)/10
				tmp = []int{}
			}
		}
		he += sum(tmp)
		tmp = []int{}
		for i := 0; i< len(buy_b); i++ {
			tmp = append(tmp, buy_b[i])
			if len(tmp) == 3{
				he += sum(tmp[:2]))
				tmp = []int{}
			}
		}
		he += sum(tmp)
		if he < res {
			res = he
		}
		return
	}
	buy_a = append(buy_a, A[i])
	dfs(i+1, buy_a, buy_b, n, A, B)
	buy_a = buy_a[:len(buy_a)-1]
	buy_b = append(buy_b, B[i])
	dfs(i+1, buy_a, buy_b, n, A, B)
	buy_b = buy_b[:len(buy_b)-1]

}
func sum (a []int) int{
	count := 0
	for i := 0; i< len(a); i++ {
		count += a[i]
	}
	return count
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

//第2题
//func getMinLen(points [][]int) int {
//	sort.Slice(points, func(i, j int) bool {
//		a := points[i][0]*points[i][0] + points[i][1]*points[i][1]
//		b := points[j][0]*points[j][0] + points[j][1]*points[j][1]
//		return a < b
//	})
//	res := 0
//	res += getDis([]int{0, 0}, points[0])
//	for i := 1; i < len(points); i++ {
//		res += geDis(points[i], points[i-1])
//	}
//	return res
//}
//func getDis(a []int, b []int) int {
//	return abs(a[0]-b[0]) + abs(a[1]-b[1])
//}
//func abs(i int) int {
//	if i < 0 {
//		return -i
//	}
//	return i
//}
var res int

func getMinLen(points [][]int) int {
	n := len(points)
	is_eat := make([]int, n)
	res = 201
	dfs(0, 0, 0, 0, n, points, is_eat)
	return res
}
func dfs(x, y, step, dis, n int, food [][]int, is_eat []int) {
	if step == n {
		if dis < res {
			res = dis
		}
		return
	}
	if dis >= res {
		return
	}
	for i := 0; i < n; i++ {
		if is_eat[i] == 0 {
			d := abs(x-food[i][0]) + abs(y-food[i][1])
			is_eat[i] = 1
			dfs(food[i][0], food[i][1], step+1, dis+d, n, food, is_eat)
			is_eat[i] = 0
		}
	}
}
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}


package main

import (
"bufio"
"fmt"
"os"
"sort"
"strconv"
"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	bs := make([]byte, 2000*1024)
	input.Buffer(bs, len(bs))
	input.Scan()
	A := strings.Split(input.Text(), ",")
	input.Scan()
	B := strings.Split(input.Text(), ",")
	numsA, numsB := make([]int, len(A)), make([]int, len(B))
	for i := 0; i < len(A); i++ {
		numsA[i], _ = strconv.Atoi(A[i])
	}
	for i := 0; i < len(B); i++ {
		numsB[i], _ = strconv.Atoi(B[i])
	}
	r := solve(numsA, numsB)
	fmt.Println(r)
}

func solve(numsA, numsB []int) int {
	sumA := sum(numsA)
	costA := int(sumA*6)/10
	B := make([]int, len(numsB))
	copy(B, numsB)
	sort.Slice(B, func(i, j int) bool {
		return B[i] < B[j]
	})
	costB := 0
	step := 1
	for i := len(B)-1; i>=0; i-- {
		if step %3 != 0{
			costB+=B[i]
		}
		step++
	}
	return min(costA, costB)
}
func sum (a []int) int{
	count := 0
	for i := 0; i< len(a); i++ {
		count += a[i]
	}
	return count
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}