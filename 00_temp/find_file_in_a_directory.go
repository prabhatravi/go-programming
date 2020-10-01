package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func findFile(targetDir string, pattern string) {
	err := filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if err == nil && (info.Name() == pattern) {
			println(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	if len(os.Args) <= 2 {
		fmt.Printf("USAGE : %s <target_directory> <target_filename or part of filename> \n", os.Args[0])
		os.Exit(0)
	}

	targetDirectory := os.Args[1] // get the target directory
	var pattern string
	pattern = os.Args[2]

	findFile(targetDirectory, pattern)
}
