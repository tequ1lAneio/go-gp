package main

import (
	"fmt"
	"io"
	"reflect"
)

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

type myInt int

func (n *myInt) Add(m int) {
	*n = *n + myInt(m)
}

type t struct {
	a int
	b int
}

type S struct {
	*myInt
	t
	io.Reader
	s string
	n int
}

type I interface {
	M1()
	M2()
}

type T struct {
	I
}

func (T) M3() {}




type T1 struct {}

func (T1) T1M1() { println("T1's M1") }
func (*T1) T2M2() { println("PT1's M2") }

type T2 struct{}

func (T2) T2M1() { println("T2's M1")}
func (*T2) T2M2() { println("T2's M2")}

type TT struct {
	T1
	*T2
}


func main() {
	//m := myInt(17)
	//r := strings.NewReader("hello go")
	//s := S{
	//	myInt: &m,
	//	t: t{
	//		a: 1,
	//		b: 2,
	//	},
	//	Reader: r,
	//	s:      "demo",
	//}

	//var sl = make([]byte, len("hello, go"))
	//s.Reader.Read(sl)
	//fmt.Println(string(sl))
	//s.myInt.Add(5)
	//fmt.Println(*(s.myInt))

	//var t T
	//var p *T
	//dumpMethodSet(t)
	//dumpMethodSet(p)

	t := TT{
		T1: T1{},
		T2: &T2{},
	}

	dumpMethodSet(t)
	dumpMethodSet(&t)
}
