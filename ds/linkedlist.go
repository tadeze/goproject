package main

import (
	"fmt"
)
// Basic Go datastructure 

type LinkedList struct {
	val  int
	next *LinkedList
}

func NewLinkedList(val int) *LinkedList {
	return &LinkedList{val, nil}
}
func (ll *LinkedList) add(val int) {
	head := ll
	for head.next != nil {
		head = head.next
	}
	head.next = NewLinkedList(val)

}
func (ll *LinkedList) remove(val int) *LinkedList {
	head := ll
	for head != nil {
		if head.val == val {
			head.val = head.next.val
			head.next = head.next.next
		}

	}
	return ll
}
func (ll *LinkedList) display() {
	head := ll
	for head != nil {
		fmt.Sprintf("%d->",head.val)
		head = head.next 
	}
}

func Reverse(s []int) {

	first, second := 0, len(s)-1
	for first < second {
		s[first], s[second] = s[second], s[first]
		first++
		second--
	}

}
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	root := NewLinkedList(1)
	root.add(9)
	root.add(10)
	root.add(3)
	root.display()
	/*intn := []int{2, 90, 1, 4, 9}
	bk [][]float32
	bk := [5][5]float32{}
	fmt.Println(bk)
	//intn := [10]int{3}
	//reflect.ValueOf
	fmt.Println(intn)
	//Reverse(intn)
	fmt.Println(intn)
	fmt.Println(factorial(5))
	*/
}
