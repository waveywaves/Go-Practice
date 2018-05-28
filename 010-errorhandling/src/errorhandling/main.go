package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("filename.ext")

	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println(err.Error())
	}

	myError := errors.New("Noooooo errrooorr !")
	fmt.Println(myError.Error())

	attendance := map[string]bool{
		"Vibhav":  false,
		"Khushal": true,
		"Anshul":  true,
		"Maddy":   true,
	}

	for k, v := range attendance {
		if v {
			fmt.Printf("%v has attended \n", k)
		} else {
			fmt.Printf("%v is absent \n", k)
		}
	}
}
