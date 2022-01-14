package main

import (
	"fmt"
	"unsafe"
)

func sliceToArray() {
	b := []int{11, 12, 13}
	var p = (*[3]int)(unsafe.Pointer(&b[0]))
	p[1] += 10
	fmt.Printf("%v\n", b)
}

func sliceToArrayNew() {
	//b := []int{11, 12, 13}
	//p := (*[3]int)(b)
	//p[1] += 10
	//fmt.Printf("%v\n", b)

	var b = []int{11, 12, 13}
	var h = (*[4]int)(b)
	var i = (*[0]int)(b)
	var j = (*[1]int)(b)
	var k = (*[2]int)(b)
	var l = (*[3]int)(b)
	var m = (*[3]int)(b[:1])

	fmt.Println(b, h, i, j, k, l, m)

	var b1 []int
	p1 := (*[0]int)(b1)
	var b2 = []int{}
	p2 := (*[0]int)(b2)
}

func main() {
	sliceToArray()
	//sliceToArrayNew()
}
