package main

import "fmt"

type A struct {
	A1 string
	A2 string
}

type B struct {
	B1 string
	B2 string
}

type AB struct {
	A
	B
}

func main() {
	ab := AB{}
	ab.A1 = "A1"
	ab.B2 = "B2"
	fmt.Println("hello world", ab)
}
