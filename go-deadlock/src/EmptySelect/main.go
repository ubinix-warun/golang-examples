package main

// fatal error: all goroutines are asleep - deadlock!

func main() {
    select{}
}