package object

import (
	"bytes"
	"fmt"
	"monkey/ast"
	"monkey/code"
	"strings"
)

const (
	FunctionObj         = "FUNCTION"
	CompiledFunctionObj = "COMPILED_FUNCTION"
)

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType {
	return FunctionObj
}

func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type CompiledFunction struct {
	Instructions code.Instructions
	NumLocals    int
}

func (cf *CompiledFunction) Type() ObjectType {
	return CompiledFunctionObj
}

func (cf *CompiledFunction) Inspect() string {
	return fmt.Sprintf("CompiledFunction[%p]", cf)
}
