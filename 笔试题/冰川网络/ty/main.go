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

package main
import (
	"fmt"
	"os"
	"bufio"
	"strings")
func main() {
	m ,k := 0, 0
	fmt.Scan(&m, &k)
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := strings.Split(input.Text(), " ")
	fmt.Print(str)
}