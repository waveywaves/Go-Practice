package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type ToDo struct {
	Title  string
	Status string
}

type Tasks int

type EditedToDo struct {
	Title     string
	NewTitle  string
	NewStatus string
}

var todoSlice []ToDo

func (t *Tasks) GetToDo(title string, reply *ToDo) error {
	var err error
	for _, v := range todoSlice {
		if v.Title == title {
			*reply = v
			break
		} else {
			err = errors.New("ToDo not found ..")
		}
	}
	return err
}

func (t *Tasks) CreateToDo(todo ToDo, reply *ToDo) error {
	todoSlice = append(todoSlice, todo)
	*reply = todo
	return nil
}

func (t *Tasks) EditToDo(editedToDo EditedToDo, reply *ToDo) error {
	var edited ToDo

	for i, v := range todoSlice {
		if v.Title == editedToDo.Title {
			todoSlice[i] = ToDo{editedToDo.NewTitle, editedToDo.NewStatus}
			edited = ToDo{editedToDo.NewTitle, editedToDo.NewStatus}
		}
	}
	*reply = edited
	return nil
}

func (t *Tasks) DeleteToDo(toDo ToDo, reply *ToDo) error {
	var deleted ToDo
	for i, v := range todoSlice {
		if v.Title == toDo.Title && v.Status == toDo.Status {
			todoSlice = append(todoSlice[:i], todoSlice[1+i:]...)
			deleted = toDo
			break
		}
	}
	*reply = deleted
	return nil
}

func (t *Tasks) GetToDos(title string, reply *[]ToDo) error {
	*reply = todoSlice
	return nil
}

func server() {
	port := ":4000"
	tasks := new(Tasks)
	var err error
	err = rpc.Register(tasks)
	rpc.HandleHTTP()

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server running on Port ", port)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = http.Serve(listener, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func client() {
	var reply ToDo
	var slice []ToDo
	client, err := rpc.DialHTTP("tcp", "localhost:4000")
	if err == nil {
		fmt.Println(err)
	}

	todo1 := ToDo{"Finish RPC", "Started"}
	todo2 := ToDo{"Read RPC RFC", "Not Started"}
	todo3 := ToDo{"Finish GRPC", "Not Started"}

	client.Call("Tasks.GetTodos", "", &slice)
	log.Println("1. Todos : ", slice, err)

	// Create Todo
	err = client.Call("Tasks.CreateToDo", todo1, &reply)
	err = client.Call("Tasks.CreateToDo", todo2, &reply)
	err = client.Call("Tasks.CreateToDo", todo3, &reply)
	client.Call("Tasks.GetTodos", "", &reply)
	log.Println("2. Todos : ", reply, err)

	// Get Todo
	err = client.Call("Tasks.GetToDo", "Finish RPC", &reply)
	log.Println("Get ToDo : ", reply, err)

	// Edit Todo
	editedTodo := ToDo{"Read RPC RFC", "Started"}
	err = client.Call("Tasks.EditToDo", editedTodo, &reply)
	client.Call("Tasks.GetTodos", "", &reply)
	log.Println("3. Todos : ", reply)

	// Delete Todo
	err = client.Call("Tasks.DeleteToDo", editedTodo, &reply)
	client.Call("Tasks.GetTodos", "", &reply)
	log.Println("4. Todos : ", reply)
}

func main() {

	var mode string
	flag.StringVar(&mode, "mode", "s", "Pass server or client")

	flag.Parse()

	if mode == "server" {
		server()
	} else if mode == "client" {
		client()
	} else {
		fmt.Println("Please pass server or client")
	}

}
