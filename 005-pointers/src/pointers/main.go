package main

import "fmt"

var p *int

func main() {
	if p != nil {
		fmt.Println(*p)
	} else {
		//Printing initial value of pointer
		fmt.Println(p)
		fmt.Println("p is Nil")
	}

	v := 50
	p = &v // Value of p is the address of v

	// Printing pointer values after
	fmt.Println(&p) // Address of p
	fmt.Println(p)
	fmt.Println(*p) // Value where p is pointing to

	/*
		To understand pointers without actually making pointers we can look at it in the following way
	*/

	// We have two variable x and a

	x := 15
	a := &x

	fmt.Println(a)  // Printing the address of x
	fmt.Println(*a) // accessing the value through the address

}
