package main

import "fmt"

type person struct {
	first string
	last string
	age int
}
func (p person)	speak() {
   fmt.Println("I am", p.first, p.last, "and I am", p.age, "years old.")
}

func main() {
	p := person {
		first: "Prabhat",
		last: "Ravi",
		age: 29,
	}
	p.speak()
}
