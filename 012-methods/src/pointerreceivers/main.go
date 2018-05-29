package main

import "fmt"

type Man struct {
	walkSpeed uint16
}

func (m *Man) workout() {
	m.walkSpeed++
}

func main() {
	bobby := Man{
		walkSpeed: 100,
	}

	fmt.Println(bobby.walkSpeed)
	bobby.workout()
	fmt.Println(bobby.walkSpeed)
	bobby.workout()
	fmt.Println(bobby.walkSpeed)
	bobby.workout()
	fmt.Println(bobby.walkSpeed)
	bobby.workout()
	fmt.Println(bobby.walkSpeed)
}
