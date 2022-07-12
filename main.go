package main

import (
	"Test/a"
	"Test/b"
	"fmt"
)

func main() {
	fmt.Println("main")
	SayAll()
}

func SayAll() {
	a.Say()
	b.Say()
}
