package main

//队列没有必要维护窗口里的所有元素，只需要维护有可能成为窗口里最大值的元素就可以了，同时保证队列里的元素数值是由大到小的。
//单调队列不是单纯的给队列中元素排序，那和优先级队列没有什么区别了
//封装单调队列
type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}
func (m *MyQueue) Empty() bool {
	return len(m.queue) == 0
}
func (m *MyQueue) Back() int {
	return m.queue[len(m.queue)-1]
}
func (m *MyQueue) Front() int {
	return m.queue[0]
}

func (m *MyQueue) Push(val int) {
	for !m.Empty() && val > m.Back() {
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, val)
}
func (m *MyQueue) Pop(val int) {
	if m.Front() == val {
		m.queue = m.queue[1:len(m.queue)]
	}
}
