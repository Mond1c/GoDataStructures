package main

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
