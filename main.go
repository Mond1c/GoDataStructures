package main

import (
	"fmt"
	. "structures/queue"
)

func main() {
	queue := Queue[int]{}
	for i := 1; i <= 10; i++ {
		queue.Push(i)
	}
	for !queue.IsEmpty() {
		fmt.Println(queue.Pop())
	}
}
