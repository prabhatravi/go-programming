package main

import (
	"fmt"
)
type person struct {
	first string
	last string
}
type secretAgent struct {
	person
	ltk bool
}
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}
// interface
type human interface {
	speak() // method
}
func bar(h human){
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrr", h.(person).first) //asserting
	case secretAgent:
		fmt.Println("I was passed into barrr", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

type hotdog int
func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Prabhat",
		last: "Ravi",
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
	fmt.Println(p1)
    // polymorphism
	bar(sa1)
	bar(sa2)
	bar(p1)
	//assertion
	var x hotdog = 32
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
