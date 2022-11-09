package main
//将指定数据结构创建后构建堆
h := &IHeap{}
heap.Init(h)
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
