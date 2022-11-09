package main

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, value := range nums {
		m[value]++
	}
	//将指定数据结构创建后构建堆
	h := &IHeap{}
	heap.Init(h)
	//把map中的元素放入堆中,同时保证里面只有k个元素
	for key, value := range m {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	//把结果输出来，注意顺序,这里切片大小一开始就要定义好
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return result

}

//将一种数据结构实现heap接口,小顶堆
type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}
func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}
func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *IHeap) Pop() interface{} {
	n := len(*h) - 1
	x := (*h)[n]
	*h = (*h)[:n]
	return x
}
