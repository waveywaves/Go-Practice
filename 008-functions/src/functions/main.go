package main

import "fmt";

func main() {
	fmt.Println(add100to(10));
	fmt.Println(return10word());
	fmt.Println(returnnumstring(10,"string"));

	poodle := Dog{"Poodle",37,"woof !"}
	poodle.Speak()
}

type Dog struct {
	Breed string
	Weight int
	Sound string
}

func add100to(value int) int {
    integervalue := value+100
    return integervalue
}

func return10word() (int, string){
    return 10, "word"
}

func returnnumstring(value int, str string) (number int, sentence string){
    number = value+10
    sentence = str+"adddedForReturn"
    return
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}