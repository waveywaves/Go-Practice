package main;

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
/*
	var s string; //Declaring variable s as type string
	fmt.Println("Enter text with spaces in between to be taken in as arguments :");
	fmt.Scanln(&s);
	fmt.Println(s); // how many number of arguments we pass we return the string
*/
	// This is similar to the Java Buffered Reader 
	reader := bufio.NewReader(os.Stdin); // input from where ?

	fmt.Println("Enter text to be taken in as an argument :");
	str, err:= reader.ReadString('\n'); //Passing a single byte
	if err !=nil {
		fmt.Println(err);
	} else {
		fmt.Println("Value of String : ",str);
	}

	// Read and write Numebrs to the console as float
	fmt.Print("Enter a Number: \n");
	str, _ = reader.ReadString('\n');
	f, err2 := strconv.ParseFloat(strings.TrimSpace(str),64);
	if err2 !=nil {
		fmt.Println(err2);
	} else {
		fmt.Println("Value of Number : ",f);
	}

}