package main

func maxSlidingWindow(nums []int, k int) []int {
	q := NewMyQueue()
	result := []int{}
	//先将前k个数放入单调队列中
	for i := 0; i < k; i++ {
		q.Push(nums[i])
	}
	//记录前k个最大值
	result = append(result, q.Front())
	//处理剩下的
	for i := k; i < len(nums); i++ {
		//移除滑动窗口最前面的元素
		q.Pop(nums[i-k])
		//滑动窗口添加最后面的元素
		q.Push(nums[i])
		result = append(result, q.Front())
	}

	return result
}

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
func (m *MyQueue) Front() int {
	return m.queue[0]
}
func (m *MyQueue) Back() int {
	return m.queue[len(m.queue)-1]
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
