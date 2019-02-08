package main

import (
	"app/client"
	"app/server"
	"context"
	"flag"
	"fmt"
)

func main() {

	var mode string
	flag.StringVar(&mode, "mode", "s", "Pass server or client")

	flag.Parse()

	if mode == "server" {
		server.Run()
	} else if mode == "client" {
		client.Run()
	} else {
		fmt.Println("Please pass server or client")
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
}
