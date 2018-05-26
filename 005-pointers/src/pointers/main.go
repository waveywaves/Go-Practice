package main;

import "fmt";

var p *int;

func main(){
	if p != nil {
		fmt.Println(*p);
	} else {
		//Printing initial value of pointer
		fmt.Println(p);
		fmt.Println("p is Nil");
	}

	v := 50;
	p = &v; // Value of p is the address of v

	// Printing pointer values after 
	fmt.Println(&p); // Address of p
	fmt.Println(p);
	fmt.Println(*p); // Value where p is pointing to
}
