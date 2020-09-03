package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

)

// MyStruct is an example structure for this program.
type MyStruct struct {
	StructData string `json:"StructData"`
}

func main() {
	filename := "myFile.json"
	err := checkFile(filename)
    fmt.Println(err)
	file, err := ioutil.ReadFile(filename)
	fmt.Println(err)
	data := []MyStruct{}

	// Here the magic happens!
	json.Unmarshal(file, &data)

	newStruct := &MyStruct{
		StructData: "peanut",
	}

	data = append(data, *newStruct)

	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}
