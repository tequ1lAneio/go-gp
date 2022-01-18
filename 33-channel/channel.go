package main

import (
	"fmt"
	"sync"
	"time"
)

func declare() {
	var ch chan int
	fmt.Println(ch)

	ch1 := make(chan int)
	ch2 := make(chan int, 5)
	fmt.Println(ch1, ch2)
}

func channelWithNoBuffer() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 13
	}()
	n := <-ch1
	println(n)
}

func channelWithBuffer() {
	//ch := make(chan int, 1)
	//n := <-ch
	//println(n)

	ch1 := make(chan int, 1)
	ch1 <- 17
	ch1 <- 27
}


func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i + 1
		time.Sleep(time.Second)
	}
	close(ch)
}

func consume(ch <-chan int) {
	for n := range ch {
		println(n)
	}
}

func singleDirectChannel() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		produce(ch)
		wg.Done()
	}()

	go func() {
		consume(ch)
		wg.Done()
	}()

	wg.Wait()
}

func whetherOk() {
	//n := <-ch
	//m, ok := <-ch
	//for v := range ch {
	//	... ...
	//}
}

func main() {
	//declare()
	//channelWithNoBuffer()
	//channelWithBuffer()
	singleDirectChannel()
}
