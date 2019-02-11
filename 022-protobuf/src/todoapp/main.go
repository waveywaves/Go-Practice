package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	. "todo"

	"github.com/golang/protobuf/proto"
)

func main() {

	var err error

	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Missing subcommand list or add")
		os.Exit(1)
	} else {

		switch mode := flag.Arg(0); mode {
		case "add":
			err = add(strings.Join(flag.Args()[1:], " "))
		case "list":
			err = list()
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}

type length int64

const (
	sizeOfLength = 8
	dbPath       = "mydb.pb"
)

var endianness = binary.LittleEndian

func add(todo string) error {
	task := &Task{
		Todo: todo,
		Done: false,
	}

	b, err := proto.Marshal(task)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(dbPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	err = binary.Write(f, endianness, length(len(b)))
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(b)
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func list() error {

	b, err := ioutil.ReadFile(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	for {
		if len(b) == 0 {
			return nil
		} else if len(b) < sizeOfLength {
			fmt.Errorf("remaining odd %d bytes, what to do ?")
		}

		var l length
		if err = binary.Read(bytes.NewReader(b[:sizeOfLength]), endianness, &l); err != nil {
			log.Fatal(err)
		}
		b = b[sizeOfLength:]

		task := Task{}
		if err = proto.Unmarshal(b[:l], &task); err != nil {
			log.Fatal(err)
		}
		b = b[l:]

		if task.Done {
			fmt.Printf("Task is Complete !")
		} else {
			fmt.Printf("Task is Incomplete \n")
		}

		fmt.Printf("%s \n", task.Todo)

	}

	return nil
}
