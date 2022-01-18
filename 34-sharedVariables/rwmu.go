package main

import "sync"

func myRwmu() {
	var rwmu sync.RWMutex
	rwmu.RLock()
	// read something
	rwmu.RUnlock()
	rwmu.Lock()
	// change something
	rwmu.Unlock()
}

func main() {

}
