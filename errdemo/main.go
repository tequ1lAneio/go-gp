package main

import (
	"./sub1"
	"fmt"
)

func main() {
	err := sub1.Diff(1, 2)
	fmt.Println(err)
	fmt.Printf("%+v",err)
}
