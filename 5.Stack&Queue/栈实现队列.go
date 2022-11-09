package __Stack_Queue

type MyQueue struct {
	inStack, outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) in2out() {
	for len(this.inStack) > 0 {
		this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
		this.inStack = this.inStack[:len(this.inStack)-1]
	}
}

func (this *MyQueue) Pop() int { //在切片中移除最后一个元素，并返回最后一个元素
	if len(this.outStack) == 0 {
		this.in2out()
	}
	x := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return x
}

func (this *MyQueue) Peek() int { //返回末尾元素，切片不改变
	if len(this.outStack) == 0 {
		this.in2out()
	}
	return this.outStack[len(this.outStack)-1]
}

func (this *MyQueue) Empty() bool {
	if len(this.inStack) == 0 && len(this.outStack) == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
