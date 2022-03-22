package main

import "fmt"

func main() {
	vec := Vector[int]{}
	for i := 1; i <= 10; i++ {
		vec.add(i)
	}
	fmt.Printf("vec: %v\n", vec)
	vec.remove(2)
	fmt.Printf("vec: %v\n", vec)
}
