package main

func main() {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1, 2) //链表变为1-> 2-> 3
	linkedList.Get(1)           //返回2
	linkedList.DeleteAtIndex(1) //现在链表是1-> 3
	linkedList.Get(1)

}
