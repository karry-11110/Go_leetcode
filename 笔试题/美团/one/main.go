package main

import (
	"fmt"
)
func main() {
    A, B := "", ""
	for i := 0; i < 100; i++ {
		A += "a"
		B += "b"
	}
	res := " "
	for i := 0; i < 100; i++{
		res += string(A[i])
		res += string(B[i])
	}
	fmt.Println(res)
}
//func main() {
//	input := bufio.NewScanner(os.Stdin)
//	input.Scan()
//	n, _ := strconv.Atoi(input.Text())
//	input.Scan()
//	A := input.Text()
//	input.Scan()
//	B := input.Text()

//	res := " "
//	for i := 0; i < n; i++{
//		res += string(A[i])
//		res += string(B[i])
//	}
//	fmt.Println(res)
//}
	//var n, m, k int
	//fmt.Scanln(&n, &m, &k)
	//input := bufio.NewScanner(os.Stdin)
	//input.Scan()
	//orders := input.Text()
	//fmt.Println(orders)
	//
	//var n, m, k int
	//fmt.Scan(&n, &m, &k)
	//input := bufio.NewScanner(os.Stdin)
	//input.Scan()
	//orders := input.Text()
	//fmt.Println(orders)

	//package main
	//
	//import "fmt"
	//
	//func main() {
	//	n, m := 0, 0
	//	fmt.Scanln(&n, &m)
	//	a := make([]int, n+1)
	//	b := make([]int, n+1)
	//	for i := 1; i <= n; i++ {
	//		fmt.Scan(&a[i])
	//		b[i] = a[i] - a[i-1]
	//	}
	//	fmt.Println(n,m)
	//	fmt.Print(a, b)
	//	l, r, c := 0, 0, 0
	//	for m > 0 {
	//		fmt.Scanln(&l, &r, &c)
	//		fmt.Println(l,r,c)
	//		m--
	//	}
	//}

