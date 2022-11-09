package main

import "fmt"

func main() {
	fmt.Println(singleNumbers([]int{1, 2, 1, 2, 7, 9}))
}
func singleNumbers(nums []int) []int {
	x, y, n, m := 0, 0, 0, 1
	for _, num := range nums {
		n ^= num
	}
	for n&m == 0 {
		m <<= 1
	}
	for _, num := range nums {
		if num&m == 0 {
			x ^= num
		} else {
			y ^= num
		}
	}
	return []int{x, y}
}
