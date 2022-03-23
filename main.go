package main

import (
	"fmt"
	. "structures/stack"
)

func main() {
	stack := Stack[int]{}
	for i := 1; i <= 10; i++ {
		stack.Add(i)
	}
	fmt.Printf("stack: %v\n", stack)
	fmt.Printf("stack.Top(): %v\n", stack.Top())
	stack.Pop()
	fmt.Printf("stack.Top(): %v\n", stack.Top())
}
