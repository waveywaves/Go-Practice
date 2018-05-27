# Go-Practice

This repository will focus on learing Go.

## Progress

### 1> Hello World !

> Go packages and directory strucure for compilation
```sh
> /nameofgocomppath
> /nameofgocomppath/src
> /nameofgocomppath/src/helloworld/main.go
> /nameofgocomppath/bin
> /nameofgocomppath/pkg
```
```sh
$ export GOPATH="/name/of/go/comp/path"
$ go install package1
$ go install package2
```
> Go packages and directory strucure after compilation
```sh
> /nameofgocomppath
> /nameofgocomppath/src
> /nameofgocomppath/src/helloworld/main.go
> /nameofgocomppath/bin
> /nameofgocomppath/bin/helloworld
> /nameofgocomppath/pkg
```

### 2> Strings

Different ways of formatting and printing stuff.


### 3> Console Input

Taking in User Input from the console with bufio and fmt's Scanln.
Apart from this we also see how to handle the input which we return from the console.


### 4> Variables
Types for variables can be set explicitly or implicitly.
"=" operator is for explicit and ":=" is for implicit.
As Go is a statically typed language these values cannot be changed during runtime.

Constants in Go use "=" for both implicit and explicit declaration.
"time" package allows to use different time operations.

#### Arithmatic Operations
+ , - . * , / , %
"&"  Bitwise AND
"|"  Bitwise OR
"^"  Bitwise XOR
"&^" Bit Clear
"<<" LeftShift
">>" Rightshift

Mathematics requires the same types for the operators to function properly

#### Types
bool
string 
uint8-64 (number of bits)
int8-64 (number of bits)
byte 
int
uintptr (large enough to hold the bit pattern of any pointer)
float32-64
complex64-128 (Complex Numbers)
Arrays
Slices
Maps
Structs
Functions
Interfaces
Channels
Pointers

### 5> Pointers

These are the same as they would be in C, C++.
Similarly "&" would be to get the address of the value and "*" would be to point to the value.

### 6> Arrays

#### Arrays

Arrays are declared in the following way
```sh
var arr [1]string;
arr[0] = "value"
```

#### Slices

Slices are no different from Arrays with the only difference being that their size is not fixed, they are declared differently and they make up arrays.
```sh
var sli1 []string;
var arr [5]int = [3]int{1,2,3,4,5};
sli2 := append(arr[1:3]);
```

### 7> Memory Allocation

- The runtime operates in its own background threads. Memory is managed completely in the background and no allocation of deallocation of memory needs to be handled by the programmer. 
- make(), new() are used to create instances of Maps, Slices, Channels
```sh
new()  // Allocates but does not initalize memory, returns address
make() // Allocates and initalizes memory, returns address
```
- The Go Garbage Collector deallocated memory for objects which are out of scope or set to nil.

#### Maps

So maps can be declared in the following way
```sh
mp := make(map[int]string) 
```
In the above example we are creating maps which would have the keys as integers and the value for the key as a string. 
Maps are similar to dictionaries in Python with the difference being that the Keys and Values can be of only one type.

#### Structs

Structs are the same as they are in C. They are declared as follows.
```sh
type Dog struct {
    Breed string
    Weight int
}
```
They are meant to be used like classes, but Go doesn't have the advanced OOP concepts such as inheritance and such. The only OOP Concepts it has are Encapslation and Type Memeber functions.
Go thorugh this [link](https://www.ardanlabs.com/blog/2013/07/object-oriented-programming-in-go.html) for more information about Object Oriented Programming in Go.


### 8> Functions

Basic functions in Go are little different than the normal way we end up writing functions.
```sh
func function(value int) int {
    integervalue := 100
    return integervalue
}

func function(value int, value2 string) (int, string){
    return 10, "word"
}

func function(value int, value string) (number int, sentence string){
    number = value+10
    sentence = value+"adddedForReturn"
    return
}
```

