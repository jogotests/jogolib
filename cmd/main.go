package main

import (
	"fmt"

	"github.com/jogotests/jogolib"
)

func main() {
	// test
	g := jogolib.NewGreeter("Hello")
	m := g.Greet("joshuagame")
	fmt.Println(m)
}
