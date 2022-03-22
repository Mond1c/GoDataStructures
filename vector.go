package main

type Vector[T any] struct {
	data     []T
	size     int
	capacity int
}

func (this *Vector[T]) is_full() bool {
	return this.size == this.capacity
}

func (this *Vector[T]) add(value T) {
	this.data = append(this.data, value)
	this.size++
	this.capacity = cap(this.data)
}

func (this *Vector[T]) get(index int) T {
	return this.data[index]
}

func (this *Vector[T]) remove(index int) {
	if index < 0 || index >= this.size {
		panic("Invalid argument")
	}
	for i := index + 1; i < this.size; i++ {
		this.data[i-1] = this.data[i]
	}
	this.size--
}
