package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	vibhav := &Person{
		Name: "Vibhav",
		Age:  22,
	}

	data, err := proto.Marshal(vibhav)
	if err != nil {
		log.Fatal("marshalling error: ", err)
	}

	fmt.Println(data)
	newVibhav := &Person{}
	err = proto.Unmarshal(data, newVibhav)

	if err != nil {
		log.Fatal("unmarshalling error :", err)
	}

	fmt.Println(newVibhav.GetAge())
	fmt.Println(newVibhav.GetName())
}
