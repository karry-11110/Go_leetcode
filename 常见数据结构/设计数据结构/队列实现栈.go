type MyStack struct {
	queue0, queue1 []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.queue1 = append(this.queue1, x) //首先把要进的元素放进辅助队列
	for len(this.queue0) > 0 {
		this.queue1 = append(this.queue1, this.queue0[0])
		this.queue0 = this.queue0[1:]
	} //再把主要队列里的全部移到辅助队列
	this.queue0, this.queue1 = this.queue1, this.queue0 //此时辅助队列里面的顺序即满足条件来，交换即可
}

func (this *MyStack) Pop() int {
	x := this.queue0[0] //直接把主要队列的第一个元素弄出去就行
	this.queue0 = this.queue0[1:]
	return x
}

func (this *MyStack) Top() int {
	return this.queue0[0] //主要队列的第一个元素就是即将要出栈的
}

func (this *MyStack) Empty() bool { //判断主要的队列是否否为空就行
	return len(this.queue0) == 0
}