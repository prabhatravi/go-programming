package main

import "fmt"

type person struct {
	first string
	last string
}
func (p *person) speak() {
	fmt.Println("i am in method call", p.first, p.last)
}
type human interface {
	speak()
}
func main() {
	p1 := person{
		first: "Prabhat",
		last:  "Ravi",
	}
	saySomething(&p1)
	//doesnot work
	//saySomething(p1)
}
func saySomething(h human){
	h.speak()
}