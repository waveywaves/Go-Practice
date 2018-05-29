package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GO Yeah !")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is just excellent considering you wanted to know more about it.")
}

func formattedHTMLHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		`
		<h1>This is Beautiful.</h1>
		<p>..and Formatted</p>
		<p>..and Fast</p>
		<p>What else do you need ?</p>
		`,
	)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/formatted", formattedHTMLHandler)
	http.ListenAndServe(":3000", nil)
}
