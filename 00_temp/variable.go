package main

import "fmt"

var a int = 5
var b string = "Prabhat Ravi"
var c bool = true
type d int
var e d
func main() {
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	e = 9
	a = int(e)
	fmt.Println(a)
}
