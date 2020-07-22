package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourwheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourwheel: true,
	}
	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}
	fmt.Println(t1)
	fmt.Println(t1.doors)
	fmt.Println(s1)
	fmt.Println(s1.doors)
}
