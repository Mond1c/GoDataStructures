package queue

import . "structures/list"

type Queue[T comparable] struct {
	data *List[T]
}

func (this *Queue[T]) IsEmpty() bool {
	return this.data.IsEmpty()
}

func (this *Queue[T]) Push(value T) {
	if this.data == nil {
		this.data = new(List[T])
	}
	this.data.Add(value)
}

func (this *Queue[T]) Front() T {
	return this.data.First()
}

func (this *Queue[T]) Pop() T {
	value := this.data.First()
	this.data.RemoveFirst()
	return value
}
