package jogolib

import "fmt"

// Greeter ..
type Greeter interface {
	Greet(name string) string
}

type greeter struct {
	prefix string
}

func NewGreeter(prefix string) Greeter {
	return &greeter{prefix}
}

func (g *greeter) Greet(name string) string {
	return fmt.Sprintf("%s %s", g.prefix, name)
}
