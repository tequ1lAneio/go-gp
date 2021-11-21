package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	Name  string
	Phone string
	Addr  string
}

type Book struct {
	Title   string
	Pages   int
	Author  Person
	Indexes map[string]int
}

type Picture struct {
	Title string
	Person
}

type Empty struct{}

type TTT struct {
	t *TTT
	st []TTT
	m map[string]TTT
}

//func NewT(field1, field2, ...) *T {
//	... ...
//}

func main() {
	var s Empty
	fmt.Println(unsafe.Sizeof(s))

	var book Book
	book.Author = Person{
		Name: "heiheihei", Phone: "18266666666", Addr: "unknown",
	}
	fmt.Println(book.Author.Name)

	var pic Picture
	pic.Name = "hahaha"
	fmt.Println(pic.Name)
}
