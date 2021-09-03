package object

import "fmt"

const (
	ClosureObj = "CLOSURE"
)

type Closure struct {
	Fn   *CompiledFunction
	Free []Object
}

func (c *Closure) Type() ObjectType {
	return ClosureObj
}

func (c *Closure) Inspect() string {
	return fmt.Sprintf("Closure[%p]", c)
}
