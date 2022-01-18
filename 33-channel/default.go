package main

import "time"

var c chan int

func wooorker() {
	select {
	case <-c:
	// do something
	case <-time.After(30 * time.Second):
		return
	}
}

func main() {

}
