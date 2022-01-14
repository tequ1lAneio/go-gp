package main

import (
	"fmt"
	"time"
)

type T struct {}

func (t T) M(n int) {}

type T1 struct {
	a int
}

func (t T1) Get() int {
	return t.a
}

func (t *T1) Set(a int) int {
	t.a = a
	return t.a
}

type field struct {
	name string
}

func (p *field) print () {
	fmt.Println(p.name)
}

func main() {
	//var t T
	//t.M(1)
	//
	//p := &T{}
	//p.M(2)

	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		//go v.print()
		go (*field).print(v)
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		//go v.print()
		go (*field).print(&v)
	}
	println("main()")
	time.Sleep(3 * time.Second)
}
