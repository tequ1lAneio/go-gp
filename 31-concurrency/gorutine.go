package main

import (
	"errors"
	"fmt"
	"time"
)

func createGoRoutine() {
	go fmt.Println("I am a goroutine")

	var c = make(chan int)
	go func(a, b int) {
		c <- a + b
	}(3, 4)

	fmt.Println(<-c)
	println("yo")
}

func spawn(f func() error) <-chan error {
	c := make(chan error)

	go func() {
		c <- f()
	}()

	return c
}

func main() {
	//createGoRoutine()

	c := spawn(func() error {
		time.Sleep(2 * time.Second)
		return errors.New("timeout")
	})
	fmt.Println(<-c)
}
