package main;

import (
	"fmt"
)

type Dog struct {
	Breed string
	Weight int
}

func main(){
	poodle := Dog{"poodle",12}
	fmt.Println(poodle.Breed, poodle.Weight);
}
