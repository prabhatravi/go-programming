package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("hello from the thing one")
		wg.Done()
	}()

	go func() {
		fmt.Println("hello from the thing two")
		wg.Done()
	}()
	fmt.Println("middle cpu", runtime.NumCPU())
	fmt.Println("middle gs", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("about to exit")
	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())
}
