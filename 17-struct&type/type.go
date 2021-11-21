package main

import "fmt"

type T1 int
type T2 T1
type T3 string

type M map[int]string
type S []string

type (
	T4 int
	T5 string
)

type T = S

type TT = string

func main() {
	var n1 T1
	var n2 T2
	n1 = T1(n2)

	var s T3 = "hello"
	//n1 = T1(s)

	fmt.Println(n1, n2, s)

	var ss string = "hello"
	var t TT = ss
	fmt.Printf("%T\n", t)
}
