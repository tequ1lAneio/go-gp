package main

import (
	"fmt"
	"io"
	"os"
)

func createFuncionByVar() {
	var myFprintf = func(w io.Writer, format string, a ...interface{}) (int, error) {
		return fmt.Fprintf(w, format, a...)
	}
	fmt.Printf("%T\n", myFprintf)
	myFprintf(os.Stdout, "%s\n", "Hello, Go")
}

func myAppend(sl []int, elems ...int) []int {
	fmt.Printf("%T\n", elems)
	if len(elems) == 0 {
		println("no elems to append")
		return sl
	}
	sl = append(sl, elems...)
	return sl
}

func returnValuesWithNames() (ns int, ss string) {
	return
}

func setup(task string) func() {
	println("do some setup stuff for", task)
	return func() {
		println("do some teardown stuff for", task)
	}
}

func main() {
	sl := []int{1, 2, 3}
	sl = myAppend(sl)
	fmt.Println(sl)
	sl = myAppend(sl, 4, 5, 6)
	fmt.Println(sl)

	ns, ss := returnValuesWithNames()
	fmt.Println(ns, ss)

	teardown := setup("demo")
	defer teardown()
	println("do some bussiness stuff")
}
