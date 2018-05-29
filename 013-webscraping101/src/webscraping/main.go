package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	resp, err := http.Get("http://www.thehindu.com/?service=newssitemap")

	if err == nil {
		bytes, err2 := ioutil.ReadAll(resp.Body) /*
			ioutil.ReadAll()
			takes in io.Reader interface and returns byte and error
		*/
		if err2 == nil {
			defer resp.Body.Close()
			fmt.Println(string(bytes))
		} else {
			fmt.Println(err2)
		}
	} else {
		fmt.Println(err)
	}
}
