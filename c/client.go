package c

import (
	"fmt"
	"github.com/chaosi-zju/Test/a"
	"github.com/chaosi-zju/Test/b"
)

func Say() {
	fmt.Println("c")
	a.Say()
	b.Say()
}
