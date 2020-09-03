package main

import (
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Profession string `json:"profession"`
}

func main() {
	// an instance
	person := Person{Name: "Prabhat Ravi", Profession: "Software Engineer"}
	fmt.Println("person: ", person)

	// To print key:value use %+v
	fmt.Printf("person: %+v", person)
}
