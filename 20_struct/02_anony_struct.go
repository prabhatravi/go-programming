package main

import "fmt"

func main() {
	s := struct{
		first string
		friends map[string]int
		favDrinks []string
	}{
		first: "Prabhat",
		friends: map[string]int{
			"Monneypenny": 2017,
			"Q": 77,
			"M": 888,
		},
		favDrinks: []string{
			"Juice",
			"Water",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)
	for k, v := range s.friends{
		fmt.Println(k, v)
	}
	for k, v := range s.favDrinks{
		fmt.Println(k, v)
	}
}
