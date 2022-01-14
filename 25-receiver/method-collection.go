package main

import (
	"fmt"
	"reflect"
)

type Interface interface {
	M1()
	M2()
}

func dumpMethodSet(i interface{}) {
	dynTyp := reflect.TypeOf(i)

	if dynTyp == nil {
		fmt.Printf("there is no dynamic type\n")
		return
	}

	n := dynTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", dynTyp)
		fmt.Printf("\n")
		return
	}

	fmt.Printf("%s's method set:\n", dynTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", dynTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}

type T1 struct {}

func (t T1) M1() {}
func (t *T1) M2() {}

func (t *T1) M3() {}
func (t *T1) M4() {}

func main() {
	//var t T1
	//var pt *T1
	//var i Interface
	//
	//i = pt
	//i = t
	//print(i)

	var n int
	dumpMethodSet(n)
	dumpMethodSet(&n)

	var t T1
	dumpMethodSet(t)
	dumpMethodSet(&t)

	type S T1
	var s S
	dumpMethodSet(s)
	dumpMethodSet(&s)
}