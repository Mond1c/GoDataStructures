package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type List struct {
	head *Node
	tail *Node
	size int
}

func (this *List) is_empty() bool {
	return this.size == 0
}

func (this *List) add(value int) {
	ptr := Node{value: value, next: nil}
	if this.is_empty() {
		this.head = &ptr
		this.tail = &ptr
		this.size++
		return
	}
	this.tail.next = &ptr
	this.tail = &ptr
	this.size++
}

func (this *List) remove_first() {
	if this.is_empty() {
		panic("List is empty")
	}
	if this.size == 1 {
		this.size--
		this.head = nil
		this.tail = nil
		return
	}
	this.head = this.head.next
	this.size--
}

func (this *List) remove_last() {
	if this.is_empty() {
		panic("List is empty")
	}
	if this.size == 1 {
		this.head = nil
		this.tail = nil
		this.size--
	}
	ptr := this.head
	for ptr.next != this.tail {
		ptr = ptr.next
	}
	fmt.Println(ptr.value)
	this.tail = ptr
	this.tail.next = nil
	this.size--
}

func (this *List) remove_by_index(index int) {
	if index < 0 || index >= this.size {
		panic("Invalid index")
	}
	if this.is_empty() {
		panic("List is empty")
	}
	if index == 0 {
		this.remove_first()
		return
	}
	if index == this.size-1 {
		this.remove_last()
		return
	}
	slow_ptr := this.head
	fast_ptr := this.head.next
	for i := 1; i < index; i++ {
		slow_ptr = slow_ptr.next
		fast_ptr = fast_ptr.next
	}
	slow_ptr.next = fast_ptr.next
	fast_ptr.next = nil
}

func (this *List) remove_by_value(value int) {
	if this.is_empty() {
		panic("List is empty")
	}
	for this.head.value == value {
		this.remove_first()
	}
	if this.is_empty() || this.size == 1 {
		return
	}
	slow_ptr := this.head
	fast_ptr := this.head.next
	for fast_ptr != nil {
		if fast_ptr.value == value {
			slow_ptr.next = fast_ptr.next
		} else {
			slow_ptr = slow_ptr.next
		}
		fast_ptr = fast_ptr.next
	}
}

func (this *List) print() {
	if this.is_empty() {
		return
	}
	ptr := this.head
	for ptr != nil {
		fmt.Printf("%d ", ptr.value)
		ptr = ptr.next
	}
	fmt.Println()
}

func main() {
	l := List{head: nil, tail: nil, size: 0}
	for i := 1; i <= 10; i++ {
		l.add(i)
		l.add(3)
	}
	l.print()
	l.remove_by_value(3)
	l.print()
}
