package main

// fatal error: all goroutines are asleep - deadlock!

func main() {
	c := make(chan struct{})
    <-c
}