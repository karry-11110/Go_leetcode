package main

import "container/list"

//在 container/list 中，这个双向链表的每个结点的类型是 Element。Element 中存了 4 个值，前驱和后继结点，双向链表的头结点，value 值。这里的 value 是 interface 类型。笔者在这个 value 里面存了 node 这个结构体。这就解释了 list 里面存的是什么数据。
type LFUCache struct {
	nodes    map[int]*list.Element
	lists    map[int]*list.List
	capacity int
	min      int
}

type node struct {
	key       int
	value     int
	frequency int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		nodes:    make(map[int]*list.Element),
		lists:    make(map[int]*list.List),
		capacity: capacity,
		min:      0,
	}
}

func (this *LFUCache) Get(key int) int {
	value, ok := this.nodes[key] //在nodes map中通过key获取节点信息
	if !ok {
		return -1
	}
	currentNode := value.Value.(*node)                   //获取list代表节点value中的真正节点
	this.lists[currentNode.frequency].Remove(value)      //在lists删除当前frequencey节点
	currentNode.frequency++                              //更新节点的frequ
	if _, ok := this.lists[currentNode.frequency]; !ok { //如果frequecy节点代表的链表不存在新建一个
		this.lists[currentNode.frequency] = list.New()
	}
	//把当前节点加到表首
	newList := this.lists[currentNode.frequency]
	newNode := newList.PushBack(currentNode)
	this.nodes[key] = newNode                                                                  //更新双向链表节点作为value的map。
	if currentNode.frequency-1 == this.min && this.lists[currentNode.frequency-1].Len() == 0 { //最后更新min值
		this.min++
	}
	return currentNode.value
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	// 如果存在，更新访问次数
	if currentValue, ok := this.nodes[key]; ok {
		currentNode := currentValue.Value.(*node)
		currentNode.value = value
		this.Get(key)
		return
	}
	// 如果不存在且缓存满了，需要删除
	if this.capacity == len(this.nodes) {
		currentList := this.lists[this.min]
		frontNode := currentList.Front()
		delete(this.nodes, frontNode.Value.(*node).key)
		currentList.Remove(frontNode)
	}
	// 新建结点，插入到 2 个 map 中
	this.min = 1
	currentNode := &node{
		key:       key,
		value:     value,
		frequency: 1,
	}
	if _, ok := this.lists[1]; !ok {
		this.lists[1] = list.New()
	}
	newList := this.lists[1]
	newNode := newList.PushBack(currentNode)
	this.nodes[key] = newNode
}
