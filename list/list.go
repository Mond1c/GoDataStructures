package list

import "fmt"

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type List[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func (this *List[T]) IsEmpty() bool {
	return this.size == 0
}

func (this *List[T]) Add(value T) {
	ptr := Node[T]{value: value, next: nil}
	if this.IsEmpty() {
		this.head = &ptr
		this.tail = &ptr
		this.size++
		return
	}
	this.tail.next = &ptr
	this.tail = &ptr
	this.size++
}

func (this *List[T]) RemoveFirst() {
	if this.IsEmpty() {
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

func (this *List[T]) RemoveLast() {
	if this.IsEmpty() {
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

func (this *List[T]) RemoveByIndex(index int) {
	if index < 0 || index >= this.size {
		panic("Invalid index")
	}
	if this.IsEmpty() {
		panic("List is empty")
	}
	if index == 0 {
		this.RemoveFirst()
		return
	}
	if index == this.size-1 {
		this.RemoveLast()
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

func (this *List[T]) RemoveByValue(value T) {
	if this.IsEmpty() {
		panic("List is empty")
	}
	for this.head.value == value {
		this.RemoveFirst()
	}
	if this.IsEmpty() || this.size == 1 {
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

func (this *List[T]) First() T {
	return this.head.value
}

func (this *List[T]) Last() T {
	return this.tail.value
}

func (this *List[T]) Get(index int) T {
	if index < 0 || index >= this.size {
		panic("Invalid index")
	}
	ptr := this.head
	for i := 0; i < index; i++ {
		ptr = ptr.next
	}
	return ptr.value
}

func (this *List[T]) Print() {
	if this.IsEmpty() {
		return
	}
	ptr := this.head
	for ptr != nil {
		fmt.Printf("%v ", ptr.value)
		ptr = ptr.next
	}
	fmt.Println()
}
