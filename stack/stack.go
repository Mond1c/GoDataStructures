package stack

type Stack[T any] struct {
	data []T
	size int
}

func (this *Stack[T]) IsEmpty() bool {
	return this.size == 0
}

func (this *Stack[T]) Add(value T) {
	this.data = append(this.data, value)
	this.size++
}

func (this *Stack[T]) Top() T {
	if this.IsEmpty() {
		panic("List is empty")
	}
	return this.data[this.size-1]
}

func (this *Stack[T]) Pop() T {
	if this.IsEmpty() {
		panic("List is empty")
	}
	this.size--
	return this.data[this.size]
}
