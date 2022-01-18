package main

import (
	workpool "../35-workerPool"
	"time"
)

func main() {
	p := workpool.New(5)

	for i := 0; i < 10; i++ {
		err := p.Schedule(func() {
			time.Sleep(time.Second * 3)
		})

		if err != nil {
			println("task: ", "err", err)
		}
	}

	p.Free()
}
