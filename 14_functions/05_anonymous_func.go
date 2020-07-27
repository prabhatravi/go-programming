package main

import (
	"fmt"
)

func main() {
	foo()

	// anonymous func
	func() {
		fmt.Println("Anonymous func ran")
	}()
	func(x int) {
		fmt.Println("Anonymous func ran with argument: ", x)
	}(42)
	// func expression
	f := func(x int){
		fmt.Println("my first func expression with : ", x)
	}
	f(1988)
	// func return
	s1 := bar()
	fmt.Println(s1)

	m := func() int {
		return 2021
	}
	fmt.Printf("%T\n", m)
	i := m()
	fmt.Println(i)
	fmt.Println(m())

	n := foobar()
	fmt.Printf("%T\n", n)
	j := n()
	fmt.Println(j)
	fmt.Println(foobar()())
}

func foo() {
	fmt.Println("foo ran")
}

func bar() string {
	s := "Hello Dear"
	return s
}

func foobar() func() int {
	return func() int {
		return 2020
	}
}
