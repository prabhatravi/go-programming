package main

import (
	"fmt"
	"sort"
)

func main() {

	strs := []string{"james", "bond", "moneybenny"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{70, 2, 71}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
