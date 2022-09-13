package main

type node struct {
	key   int
	value int

	next *node
	prev *node
}

type LRUCache struct {
	size int
	cap  int
	head *node
	tail *node
	hMap map[int]*node
}

func Constructor(capacity int) LRUCache {
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		size: 0,
		cap:  capacity,
		head: head,
		tail: tail,
		hMap: make(map[int]*node, capacity),
	}
}
func (this *LRUCache) moveToHead(node *node) {
	if node.next != nil && node.prev != nil {
		node.prev.next = node.next
		node.next.prev = node.prev
		node.next = nil
		node.prev = nil
	}
	if node.prev == nil && node.next == nil {
		node.next = this.head.next
		node.next.prev = node
		this.head.next = node
		node.prev = this.head
	}

}
func (this *LRUCache)removeNode(){
	tailNode :=this.tail.prev
	delete(this.hMap,tailNode.key)
	this.tail.prev = tailNode.prev
	tailNode.prev.next = this.tail
	tailNode.prev = nil
	tailNode.next = nil
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.hMap[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1

}

func (this *LRUCache) Put(key int, value int) {
	//存在，直接更新
	if node, ok := this.hMap[key]; ok {
		node.value = value
		this.moveToHead(node)
		return
	}
	if this.size == this.cap {
		this.removeNode()
		this.size--
	}
	this.size++
	node := &node{key: key,value: value}
	this.hMap[key] = node
	this.moveToHead(node)

}
func main()  {
	lru :=Constructor(4)
	lru.Put(2,0)
	lru.Put(1,1)
	lru.Get(2)

}