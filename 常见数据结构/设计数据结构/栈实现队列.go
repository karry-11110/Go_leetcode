type MyQueue struct {
	inStack, outStack []int //定义一个只管出，只管进的栈
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x) //直接进即可
}

func (this *MyQueue) in2out() {
	for len(this.inStack) > 0 {
		this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
		this.inStack = this.inStack[:len(this.inStack)-1]
	}
}

func (this *MyQueue) Pop() int { //在切片中移除最后一个元素，并返回最后一个元素
	if len(this.outStack) == 0 { //如果出的栈没元素一定要把入的栈里面的元素全移到出栈里
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