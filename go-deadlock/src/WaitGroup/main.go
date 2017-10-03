package main

import "sync"

// fatal error: all goroutines are asleep - deadlock!

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    wg.Wait()
}