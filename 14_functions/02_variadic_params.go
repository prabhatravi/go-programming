package main

import (
	"fmt"
)

func main() {
	x := sum(2,3,4,5,6,7,8,9,10)
	fmt.Println("The total is: ", x)
}
func sum(y ...int) int {
	fmt.Println(y)
	fmt.Println("%T\n", y)
	sum := 0
	for i, v := range y {
		sum += v
		fmt.Println("for the item in index position ", i, "we are now adding ", v, "to the total which is now: ", sum)
	}
	return sum
}