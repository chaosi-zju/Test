package main

import (
	"Test/c"
	"fmt"
	"github.com/chaosi-zju/Test/a"
)

func main() {
	fmt.Println("main")
	SayAll()
}

func SayAll() {
	a.Say()
	c.Say()
}
