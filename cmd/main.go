package main

import (
	"fmt"

	"github.com/jogotests/jogolib"
)

func main() {
	g := jogolib.NewGreeter("Hello")
	m := g.Greet("joshuagame")
	fmt.Println(m)
}
