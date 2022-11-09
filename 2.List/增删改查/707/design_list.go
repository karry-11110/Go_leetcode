package main

type MyLinkedList struct {
	dummy *Node
}
type Node struct {
	Val  int
	next *Node
	pre  *Node
}

func Constructor() MyLinkedList {
	node := &Node{
		Val:  -1,
		next: nil,
		pre:  nil,
	}
	node.next = node
	node.pre = node
	return MyLinkedList{node}
}

func (this *MyLinkedList) Get(index int) int { //链表中节点是0-index的。
	result := -1
	node := this.dummy.next
	for index > 0 && node != this.dummy {
		index--
		node = node.next
	}
	if index != 0 {
		return -1
	}
	result = node.Val
	return result
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{
		Val:  val,
		next: this.dummy.next,
		pre:  this.dummy,
	}
	this.dummy.next.pre = node
	this.dummy.next = node
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{
		Val:  val,
		next: this.dummy,
		pre:  this.dummy.pre,
	}
	this.dummy.pre.next = node
	this.dummy.pre = node
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	insertNode := &Node{
		Val:  val,
		next: nil,
		pre:  nil,
	}
	node := this.dummy.next
	for index > 0 && node != this.dummy {
		index--
		node = node.next
	}
	if index <= 0 {
		insertNode.next = node
		insertNode.pre = node.pre
		node.pre = insertNode
		insertNode.pre.next = insertNode

	}

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	node := this.dummy.next
	for index > 0 && node != this.dummy {
		index--
		node = node.next
	}
	if node == this.dummy {
		return
	}
	if index == 0 {
		node.next.pre = node.pre
		node.pre.next = node.next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
