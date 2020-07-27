package main

import "fmt"

func main() {
	n := foo()
	fmt.Println("foo ran, number is: ", n)
	s, t := bar()
	fmt.Println("bar ran, number: ", s, "and string: ", t)
}

func foo() int {
	return 9
}
func bar() (int, string) {
	return 10, "Big boss is watching"
}