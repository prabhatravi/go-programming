package main

import "fmt"

func main() {
	ii := []int{1,2,3,4,5,6,7,8,9}
	s := foo(ii...)
	fmt.Println("sum of the mentioned slice: ", s)
	ii2 := []int{1,2,3,4,5,6,7,8,9}
	t := bar(ii2)
	fmt.Println("sum of the mentioned slice: ", t)
}

func foo(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
