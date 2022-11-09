package main

import (
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"os"
)

func main()  {
	t := 0
	fmt.Scanln(&t)
	for t > 0 {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		strs := strings.Split(input.Text(), " ")
		nums := make([]int,len(strs))
		for index, v := range strs {
			nums[index], _ = strconv.Atoi(v)
		}
		HeapSort(nums)
		fmt.Println(nums)
		t--
	}

}
//HeapSort 堆排序
func HeapSort(arr []int) {
	LEN := len(arr)
	for i := LEN/2 - 1; i >= 0; i-- {
		HeapAjust(arr, i, LEN)
	}
	for i := LEN - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		HeapAjust(arr, 0, i)
	}
}

//HeapAjust 堆调整:思想是比较节点 i 和它的孩子节点 left(i),right(i)，选出三者最大(或者最小)者，如果最大（小）值不是节点 i 而是它的一个孩子节点，那边交互节点 i 和该节点，然后再调用调整堆过程，这是一个递归的过程.
func HeapAjust(arr []int, start int, length int) {
	tmp := arr[start]
	for i := 2*start + 1; i < length; i = i * 2 { //从左孩子比起
		if i+1 < length && arr[i] < arr[i+1] { //如果范围内允许，找出左孩子右孩子哪个大点
			i++
		}
		if tmp > arr[i] { //如果待调整节点比孩子们都大，不用调整了
			break
		}
		arr[start] = arr[i]
		start = i
	}
	arr[start] = tmp
}