package main

import (
	"fmt"
	"sort"
)

func main() {

	strs := []string{"james", "bond", "moneybenny"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	strs1 := []string{"pk.0", "pk.1", "pk.11", "pk.2"}
	sort.Strings(strs1)
	fmt.Println("Strings:", strs1)

	ints := []int{70, 2, 71}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
