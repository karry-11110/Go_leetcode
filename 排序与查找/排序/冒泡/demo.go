package main

import "fmt"

//冒泡排序
func BubbleSort(list []int) {
	n := len(list)
	// 在第一轮中有没有交换过
	didSwap := false
	// 进行 N-1 轮迭代
	for i := n - 1; i > 0; i-- {
		// 每次从第一位开始比较，比较到第 i 位就不比较了，因为前一轮该位已经有序了
		for j := 0; j < i; j++ {
			// 如果前面的数比后面的大，那么交换
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				didSwap = true
			}
		}
		// 如果在第一轮中没有交换过，那么已经排好序了，直接返回
		if !didSwap {
			return
		}
	}
}
func main() {
	nums := []int{1, 3, 2, 4}
	BubbleSort(nums)
	fmt.Println(nums)
}
