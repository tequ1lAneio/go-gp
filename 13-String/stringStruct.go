package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type StringHeader struct {
	Data uintptr
	Len  int
}

func dumpBytesArray(arr []byte) {
	fmt.Printf("[ ")
	for _, b := range arr {
		fmt.Printf("%c ", b)
	}
	fmt.Printf("]\n")
}


func main() {
	var s = "hello"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("0x%x\n %d\n", hdr.Data, hdr.Len)
	p := (*[5]byte)(unsafe.Pointer(hdr.Data))
	fmt.Println(p)
	dumpBytesArray((*p)[:])
}
