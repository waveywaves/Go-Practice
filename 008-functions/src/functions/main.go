package main

import "fmt";

func main() {
	fmt.Println(add100to(10));
	fmt.Println(return10word());
	fmt.Println(returnnumstring(10,"string"));
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