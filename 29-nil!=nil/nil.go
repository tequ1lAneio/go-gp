package main

import (
	"errors"
	"fmt"
)

func typeOfInterface() {
	 var err error
	 err = errors.New("error1")
	 fmt.Printf("%T\n", errors.New("error1"))
	 fmt.Printf("%T\n", err)
}

type MyError struct {
	error
}

var ErrBad = MyError{
	error: errors.New("bad things happened"),
}

func bad() bool {
	return false
}

func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}
	return p
}

func nilEqualsNil() {
	err := returnsError()
	if err != nil {
		fmt.Printf("error occur: %+v\n", err)
		return
	}
	fmt.Println("ok")
}

func main() {
	//typeOfInterface()
	nilEqualsNil()
}
