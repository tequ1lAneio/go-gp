package main

import (
	"fmt"
	"runtime"
	trace "github.com"
)

func Trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("caller not found")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	gId := trace.curGoroutineId()
	fmt.Printf("g[%05d]: enter: [%s]\n", gId, name)
	return func() { fmt.Printf("g[%05d]: exit: [%s]\n", gId, name) }
}

func foo() {
	defer Trace()()
	bar()
}

func bar() {
	defer Trace()()
}

func main() {
	defer Trace()()
	foo()
}
