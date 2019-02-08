package client

import (
	"io"
	"log"
	"net/http"
	"os"
)

func Run() {
	res, err := http.Get("http://localhost:4000")
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	io.Copy(os.Stdout, res.Body)
}
