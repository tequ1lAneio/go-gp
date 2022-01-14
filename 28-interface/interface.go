package main

import (
	"fmt"
	"reflect"
)

type T string

func (t T) Test() {}
func (t T) notImplemented() {}

type myInterface interface {
	Test()
}

func typeAssertion() {
	var i myInterface = T("test")
	//i.notImplemented() 		// There will be an error
	v, ok := i.(T)
	v.notImplemented()
	fmt.Println(reflect.TypeOf(v), ok)
}

func assertions() {
	var a int64 = 13
	var i interface{} = a
	v1, ok := i.(int64)
	fmt.Printf("v1=%d, the type of v1 is %T, ok=%t\n", v1, v1, ok)
}

type MyInterface interface {
	M1()
}

type TT int

func (TT) M1() {
	println("T's M1")
}

func test() {
	var t TT
	var i interface{} = t
	v1, ok := i.(MyInterface)
	if !ok {
		panic("the value of i is not MyInterface")
	}
	v1.M1()
	fmt.Printf("this type of v1 is %T\n", v1)

	i = int64(13)
	v2, ok := i.(MyInterface)
	fmt.Println(v2, ok)
	fmt.Printf("the type of v2 is %T\n", v2)
}

func main() {
	//typeAssertion()
	test()
}
