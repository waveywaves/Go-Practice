package main;

import (
	"fmt"
)

func main() {
	states := make(map[int]string);
	fmt.Println(states)

	states[1] = "One";
	states[2] = "Two";
	states[3] = "Three";

	fmt.Println(states)
	for k,v := range states {
		fmt.Printf("%v:%v \n",k,v);
	}

}