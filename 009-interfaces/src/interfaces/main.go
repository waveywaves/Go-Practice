package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}
type Cat struct{}
type Cow struct{}

func (d Dog) Speak() string {
	return "Woof !!"
}
func (d Cat) Speak() string {
	return "Meow !!"
}
func (d Cow) Speak() string {
	return "Moo !!"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Cow{}}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
