package main

//2.将双链表封装在cache结构中，这里设置头尾节点，方便增减
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

//1.设计一个双链表
type Node struct {
	key, value int
	pre, next  *Node
}

//4.构造双链表节点方法
func initNode(key, value int) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

//3.构造Lrucache结构
func Constructor(capacity int) LRUCache {
	l := &LRUCache{
		capacity: capacity,
		cache:    map[int]*Node{},
		head:     initNode(0, 0),
		tail:     initNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

//5.实现get：里面有这个节点值，要（移动节点到最前面）
func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node) //分为两步，先（删节点），再（添加到前面）
	return node.value
}

//6.实现put:有则改value并（移到最前面），无则（添加节点到最前面），如果超过容量则（添加到最前面）后（删除尾节点）
func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removedNode := this.removeTail()
			delete(this.cache, removedNode.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

//9.添加节点到最前面
func (this *LRUCache) addToHead(node *Node) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}

//8.删除节点
func (this *LRUCache) removedNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

//7.实现移动节点到最前面的算法：分为删除节点，和添加节点到最前面两步
func (this *LRUCache) moveToHead(node *Node) {
	this.removedNode(node)
	this.addToHead(node)
}

//10.删除尾节点
func (this *LRUCache) removeTail() *Node {
	node := this.tail.pre
	this.removedNode(node)
	return node
}
