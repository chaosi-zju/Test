package main

import (
	"fmt"
	"github.com/chaosi-zju/Test/a"
	"github.com/chaosi-zju/Test/b"
	"github.com/chaosi-zju/Test/c"
)

func main() {
	fmt.Println("main")
	SayAll()
}

func SayAll() {
	a.Say()
	b.Say()
	c.Say()
}
