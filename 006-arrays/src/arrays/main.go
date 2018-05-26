package main;

import (	
	"fmt"
)

func main() {
	var color [3]string;
	color[0] = "Red";
	color[1] = "Blue";
	color[2] = "Green";

	for i:=0 ; i<3 ; i++ {
		fmt.Println(color[i])
	}

}