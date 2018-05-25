package main

import "fmt"
import "strings"

func main(){
	str1 := "This is a string";
	str2 := "This is a string too";
	str3 := "This is a string three";

	aNumber := 69;
	isTrue := true;

	// returning error and string length from the println func
	stringLength, err := fmt.Println(str1,str2,str3);
	fmt.Println(strings.ToUpper(str1));
	
	if err == nil {
		fmt.Println("String length:",stringLength );
	}

	//here I do not want to address the returning error value
	strLen, _ := fmt.Println(str1);
	fmt.Println("String length:",strLen);

	//working with formatting and data types
	fmt.Printf("Value of aNumber %v \n", aNumber);
	fmt.Printf("Value of isTrue %v \n", isTrue);
	fmt.Printf("Value of aNumber as a float %.2f \n", float64(aNumber));
	
	//printing out data types
	fmt.Printf("Data types: %T, %T, %T \n", str1, aNumber, isTrue);
	
	//Retruning the final string to print with Sprintf
	myString := fmt.Sprintf("Data types: %T, %T, %T \n", str1, aNumber, isTrue);
	fmt.Print(myString);


}