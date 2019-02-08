package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Run() {
	fmt.Println("Listening on port 4000 ...")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	req, err := http.NewRequest(http.MethodGet, "http://localhost:4000", nil)
	if err != nil {
		log.Fatal(err)
	}
	req = req.WithContext(ctx)

	log.Printf("Handler Started")
	defer log.Printf("Handler Ended")
	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():
		log.Print(ctx.Err())
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	}

}
