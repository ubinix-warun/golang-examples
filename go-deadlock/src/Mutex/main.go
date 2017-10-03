package main

// fatal error: all goroutines are asleep - deadlock!

import "sync"

func main() {
    var m sync.Mutex
	m.Lock()
    m.Lock()
}