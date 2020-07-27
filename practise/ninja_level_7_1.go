package main

import "fmt"

type person struct {
	first string
	last string
}
func changeMe(p *person) {
	(*p).last = "Kumar"
}
func main() {
	s := 2
	fmt.Println(&s)
	p1 := person{
		first: "Prabhat",
		last: "Ravi",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
