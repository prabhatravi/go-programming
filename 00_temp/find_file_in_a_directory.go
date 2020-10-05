package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func findFile(targetDir, fileName string) string {
	var pathName string
	err := filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if err == nil && (info.Name() == fileName) {
			println(path)
			pathName = path
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return pathName
}

func main() {

	if len(os.Args) <= 2 {
		fmt.Printf("USAGE : %s <target_directory> <target_filename or part of filename> \n", os.Args[0])
		os.Exit(0)
	}

	targetDirectory := os.Args[1] // get the target directory
	var fileName string
	fileName = os.Args[2]

	var pathName string
	pathName = findFile(targetDirectory, fileName)
	fmt.Println(pathName)
}
