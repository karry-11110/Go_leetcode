package main
import (
	"fmt"
)

func main() {
	t := 0
	fmt.Scan(&t)
	for t > 0 {
		n, m := 0, 0
		fmt.Scan(&n, &m)
		nums := make([][]int, n)
		for i := range nums {
			nums[i] = make([]int, m)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				fmt.Scan(&nums[i][j])
			}
		}
		res := solve(nums)
		fmt.Println(res)
		t--
	}
}
func solve(nums [][]int) int {
	n, m := len(nums), len(nums[0])
	maxSize := max(n, m)
	for i := 1; i <= maxSize; i++ {
		for x := 0; x <= n-i; x++{
			for y := 0; y <= m-i; y++{
				value := check(nums,x, y, i)
				if value != -1{
					maxSize = max(value, maxSize)
				}
			}
		}
	}
	return maxSize
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func check(nums [][]int, x, y,size int)int{
	sum := 0;
	flag := false
Outerloop :
	for i := x; i < x+size; i++{
		for j := y; j < y+size; j++ {
			if nums[i][j] < 0 {
				flag = true
				break Outerloop
			}
			sum += nums[i][j]
		}
	}
	if flag{
		return -1
	}
	return sum
}