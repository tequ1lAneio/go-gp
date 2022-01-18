package main

import (
	"fmt"
	"sync"
)

//type counter struct {
//	sync.Mutex
//	i int
//}
//
//var cter counter
//
//func Increase() int {
//	cter.Lock()
//	defer cter.Unlock()
//	cter.i++
//	return cter.i
//}
//
//func channelAsLock() {
//	var wg sync.WaitGroup
//
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go func(i int) {
//			v := Increase()
//			fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
//			wg.Done()
//		}(i)
//	}
//
//	wg.Wait()
//}

type counter struct {
	c chan int
	i int
}

func NewCounter() *counter {
	cter := &counter{
		c: make(chan int),
	}
	go func() {
		for {
			cter.i++
			cter.c <- cter.i
		}
	}()
	return cter
}

func (cter *counter) Increase() int {
	return <-cter.c
}

func channelAsLock() {
	cter := NewCounter()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := cter.Increase()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {
	channelAsLock()
}
