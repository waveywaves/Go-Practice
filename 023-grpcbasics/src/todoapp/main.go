package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"todo"

	"github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

type taskServer struct{}

const (
	dbPath       = "mydb.pb"
	sizeOfLength = 8
)

type length int64

var endianness = binary.LittleEndian

func (s taskServer) List(ctx context.Context, void *todo.Void) (*todo.TaskList, error) {
	b, err := ioutil.ReadFile(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	var tasks todo.TaskList

	for {
		if len(b) == 0 {
			return &tasks, nil
		} else if len(b) < sizeOfLength {
			return nil, fmt.Errorf("remaining odd %d bytes", len(b))
		}

		var l length
		if err = binary.Read(bytes.NewReader(b[:sizeOfLength]), endianness, &l); err != nil {
			return nil, fmt.Errorf("Error while reading")
		}
		b = b[sizeOfLength:]

		var task todo.Task
		if err = proto.Unmarshal(b[:l], &task); err != nil {
			return nil, fmt.Errorf("Error while Unmarshalling")
		}
		b = b[l:]

		tasks.Tasks = append(tasks.Tasks, &task)

		if task.Done {
			fmt.Printf("Task is Complete !")
		} else {
			fmt.Printf("Task is Incomplete \n")
		}
		fmt.Printf("%s \n", task.Todo)
	}

	return nil, err
}

func (s taskServer) Add(ctx context.Context, text *todo.Text) (*todo.Task, error) {
	task := &todo.Task{
		Todo: text.Text,
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

	return task, err
}

func main() {

	ctx := context.Background()

	flag.Parse()
	mode := flag.Arg(0)
	cmd := flag.Arg(1)
	switch mode {
	case "server":
		server()
	case "client":
		ct := client()

		switch cmd {
		case "add":
			add(ctx, ct, strings.Join(flag.Args()[2:], " "))
		case "list":
			list(ctx, ct)
		default:
			fmt.Println("Please provide add or list")
		}
	default:
		fmt.Println("Please provide client or server")
	}
}

func server() {
	var tasks taskServer

	srv := grpc.NewServer()
	todo.RegisterTasksServer(srv, tasks)

	fmt.Println("Listening on Port 4000 ... ")

	l, err := net.Listen("tcp", ":4000")

	if err != nil {
		log.Fatalf("could not listen to :4000")
	}
	log.Fatal(srv.Serve(l))
}

func add(ctx context.Context, client todo.TasksClient, text string) (*todo.Task, error) {
	t, err := client.Add(ctx, &todo.Text{Text: text})
	if err != nil {
		log.Fatal(err)
	}

	return t, err
}

func list(ctx context.Context, client todo.TasksClient) error {

	_, err := client.List(ctx, &todo.Void{})
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func client() todo.TasksClient {
	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := todo.NewTasksClient(conn)

	return client
}
