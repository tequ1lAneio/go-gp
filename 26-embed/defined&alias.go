package main

import (
	"fmt"
	"reflect"
)

func dumpMethodSet1(i interface{}) {
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

type TTT struct{}

func (TTT) M1() {}
func (*TTT) M2() {}

type TTT1 TTT
type TTT2 = TTT

func main() {
	var t TTT
	var pt *TTT
	var t1 TTT1
	var pt1 *TTT1
	var t2 TTT2
	var pt2 *TTT2

	dumpMethodSet1(t)
	dumpMethodSet1(pt)
	dumpMethodSet1(t1)
	dumpMethodSet1(pt1)
	dumpMethodSet1(t2)
	dumpMethodSet1(pt2)
}
