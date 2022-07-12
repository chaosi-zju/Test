package main

import (
	"Test/a"
	"Test/c"
	"fmt"
)

func main() {
	fmt.Println("main")
	SayAll()
}

func SayAll() {
	a.Say()
	c.Say()
}
