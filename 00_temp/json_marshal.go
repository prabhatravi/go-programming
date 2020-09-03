package main

import (
	"encoding/json"
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
	fmt.Printf("person: %+v\n", person)

	byteArray, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
}
