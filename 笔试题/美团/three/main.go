package main
import(
	"fmt"
	"sort"

)
func main() {
	n, m := 0, 0
	fmt.Scan(&n, &m)
	A, B := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&B[i])
	}
	res := 0.00
	for i := 0; i< n; i++ {
		res += 0.01*float64(A[i])*float64(B[i])
	}
	tmp := make([]float64, 0)
	for i := 0; i< n; i++ {
		v := (1- 0.01*float64(A[i]))*float64(B[i])
		tmp = append(tmp, v)
	}
	sort.Slice(tmp, func(i, j int)bool {
		return tmp[i] > tmp[j]
	})
	for i := 0; i < m; i++{
		res += tmp[i]
	}
	fmt.Printf("%.2f", res)
}
	//input := bufio.NewScanner(os.Stdin)
	//input.Scan()
	//str := input.Text()
	//A := make([]int, len(str))
	//strb := strings.Split(str, " ")
	//for i := 0; i < n; i++ {
	//	A[i], _ = strconv.Atoi(strb[i])
	//}
	//fmt.Println(A)


