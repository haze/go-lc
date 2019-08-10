package main

import "fmt"

type ListNode struct {
	Previous, Next *ListNode
	Key, Value     int
	Sentinel       bool
}

func NormalNode(key, value int) *ListNode {
	return &ListNode{
		Sentinel: false,
		Key:      key,
		Value:    value,
		Previous: nil,
		Next:     nil,
	}
}

func SentinelNode() *ListNode {
	return &ListNode{
		Sentinel: true,
		Key:      -1,
		Value:    -1,
		Previous: nil,
		Next:     nil,
	}
}

type LRUCache struct {
	MaxCapacity int
	Head, Tail  *ListNode
	Map         map[int]*ListNode
}

func (c LRUCache) size() int {
	return len(c.Map)
}

func Constructor(capacity int) LRUCache {
	head := SentinelNode()
	tail := SentinelNode()
	tail.Previous = head
	head.Next = tail
	return LRUCache{
		MaxCapacity: capacity,
		Head:        head,
		Tail:        tail,
		Map:         make(map[int]*ListNode, 0),
	}
}

func (c *LRUCache) popOldest() *ListNode {
	node := c.Tail.Previous
	// remove from map
	delete(c.Map, node.Key)
	return node
}

// moves the node to the head of the list (start = recent)
func (c *LRUCache) moveToStart(node *ListNode) {
	fmt.Printf("================ moveToStart(%+v)\n", node)
	if node.Previous != nil {
		fmt.Printf("node.Previous = %+v\n", node.Previous)
		fmt.Printf("node.Next = %+v\n", node.Next)
		node.Previous.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Previous = node.Previous
	}
	node.Previous = c.Head
	// fmt.Printf("c.Head=%+v\n", c.Head)
	// fmt.Printf("node=%+v, c.Head.Next=%+v\n", node.Next, c.Head.Next)
	node.Next = c.Head.Next
	c.Head.Next.Previous = node
	c.Head.Next = node
	fmt.Println("================ moveToStart end")
}

func (c *LRUCache) put(key, item int) {
	var existingNode *ListNode
	existingNode, exists := c.Map[key]
	// if it doesnt exist, lets get it
	if !exists {
		if c.size()+1 > c.MaxCapacity {
			fmt.Printf("c.Tail=%+v, c.Tail.Previous=%+v\n", c.Tail, c.Tail.Previous)
			existingNode = c.popOldest()
			fmt.Printf("popped oldest node: %+v\n", existingNode)
		} else {
			existingNode = NormalNode(key, item)
		}
		c.Map[key] = existingNode
	}
	// set everything
	existingNode.Key = key
	existingNode.Value = item
	// it was used last
	c.moveToStart(existingNode)
}

func (c *LRUCache) Put(key int, value int) {
	fmt.Printf("\nPut(%d, %d)\n", key, value)
	c.put(key, value)
}

func (c *LRUCache) Get(key int) int {
	fmt.Printf("\nGet(%d)\n", key)
	value, _ := c.get(key)
	return value
}

func (c *LRUCache) get(key int) (int, bool) {
	val, exists := c.Map[key]
	if exists {
		c.moveToStart(val)
		return val.Value, true
	}
	return -1, false
}

func (c LRUCache) list() {
	ptr := c.Head.Next
	counter := 0
	fmt.Print("inspect() [")
	for ptr != nil && !ptr.Sentinel && counter < c.MaxCapacity*2 {
		// fmt.Printf("%p %+v", ptr, ptr)
		if ptr.Next != nil && !ptr.Next.Sentinel {
			fmt.Printf("(%p| %d: %d [%p, %p]), ", ptr, ptr.Key, ptr.Value, ptr.Previous, ptr.Next)
		} else {
			fmt.Printf("(%p| %d: %d [%p, %p])", ptr, ptr.Key, ptr.Value, ptr.Previous, ptr.Next)
		}
		ptr = ptr.Next
		counter += 1
	}
	fmt.Print("]\n")
}

func (c LRUCache) inspect() {
	c.list()
	if c.Head != nil {
		fmt.Printf("head (%p): %+v (%d), ", c.Head, c.Head, c.Head.Value)
	}
	if c.Tail != nil {
		fmt.Printf("tail (%p): %+v (%d)", c.Tail, c.Tail, c.Tail.Value)
	}
	if c.Head == nil && c.Tail == nil {
		fmt.Print("both head and tail are nil")
	}
	fmt.Printf(" %+v\n", c.Map)
}

func main() {
	c := Constructor(2)
	c.Put(1, 1)
	c.inspect()
	c.Put(2, 2)
	c.inspect()
	c.Get(1)
	c.inspect()
	c.Put(3, 3)
	c.Get(2)
	c.inspect()
	c.Put(4, 4)
	c.inspect()
	c.Get(1)
	c.inspect()
	c.Get(3)
	c.inspect()
	c.Get(4)
	c.inspect()
}
