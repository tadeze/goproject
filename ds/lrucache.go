package main

import "fmt"

type DoubleLinkedList struct {
	prev *DoubleLinkedList
	next *DoubleLinkedList
	key  int
	val  int
}

type LRU struct {
	capacity int
	head     *DoubleLinkedList
	tail     *DoubleLinkedList
	cache    map[int]*DoubleLinkedList
}

func NewLRU(capacity int) LRU {
	tail := new(DoubleLinkedList)
	head := new(DoubleLinkedList) //(nil, nil, 0,0))
	tail.prev = head
	tail.key, tail.val = 0, 0
	head.next = tail
	head.key, head.val = 0, 0

	lru := LRU{capacity: capacity, head: head, tail: tail,
		cache: make(map[int]*DoubleLinkedList)}
	/*lru.capacity = capacity
	lru.head, lru.tail = head, tail
	lru.cache = make(map[int]*DoubleLinkedList)*/
	return lru
}

func (lru *LRU) AddNode(node *DoubleLinkedList) {
	prev := lru.tail.prev
	prev.next = node
	node.next = lru.tail
	node.prev = prev
	lru.tail.prev = node
}
func (lru *LRU) RemoveNode(node *DoubleLinkedList) {
	prev := node.prev
	prev.next = node.next
	node.next.prev = prev

}

func (lru *LRU) Get(key int) int {
	if node, key := lru.cache[key]; key {
		lru.RemoveNode(node)
		lru.AddNode(node)
		return node.val
	}
	return -1
}
func (lru *LRU) Set(key, val int) {
	if node, key := lru.cache[key]; key {
		lru.RemoveNode(node)
	}
	node := new(DoubleLinkedList)
	node.key, node.val = key, val
	go lru.AddNode(node)
	lru.cache[key] = node
	if len(lru.cache) > lru.capacity {
		// delete the last used node
		headNode := lru.head.next
		lru.RemoveNode(headNode)
		delete(lru.cache, headNode.key)
	}
}

func main() {

	lru := NewLRU(2)
	// fmt.Println(lru.head.val)
	lru.Set(1, 1)
	lru.Set(2, 2)
	fmt.Println(lru.Get(1))
	lru.Set(3, 3)
	fmt.Println(lru.Get(2))
	fmt.Println(lru.Get(1))

}
