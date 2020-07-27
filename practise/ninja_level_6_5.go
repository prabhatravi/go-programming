package main

import "fmt"
var g func() = func(){
	fmt.Println("I am outside of main")
}
var h func()
func main() {
	f := func() {
		fmt.Println("I am anonymous function")
	}
	f()
	g()
	h = f
	h()
	fmt.Printf("this is h %T\n", h)
	i := foo()
	fmt.Println(i())
	fmt.Println(foo()())

}
func foo() func() int {
	return func() int {
		return 42
	}
}