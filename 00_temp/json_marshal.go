package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Profession Profession `json:"profession"`
}

type Profession struct {
	Hobby     string  `json:"hobby"`
	Age       int  `json:"age"`
	Developer bool `json:"is_developer"`
}

func main() {
	// an instance
	profession := Profession{Hobby: "travelling", Age: 25, Developer: true}
	person := Person{Name: "Prabhat Ravi", Profession: profession}
	fmt.Println("person: ", person)

	// To print key:value use %+v
	fmt.Printf("person: %+v\n", person)

	//byteArray, err := json.Marshal(person)
	byteArray, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
}
