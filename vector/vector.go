package vector

type Vector[T any] struct {
	data []T
	size int
}

func (this *Vector[T]) Add(value T) {
	this.data = append(this.data, value)
	this.size++
}

func (this *Vector[T]) Get(index int) T {
	return this.data[index]
}

func (this *Vector[T]) Remove(index int) {
	if index < 0 || index >= this.size {
		panic("Invalid argument")
	}
	for i := index + 1; i < this.size; i++ {
		this.data[i-1] = this.data[i]
	}
	this.size--
}
