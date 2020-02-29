package object

import (
	"ast"
	"bytes"
	"fmt"
	"strings"
)

// ObjectType ...
type ObjectType string

// ObjectType
const (
	INTEGEROBJ     = "INTEGER"      // Integer
	BOOLEANOBJ     = "BOOLEAN"      // Boolean
	NULLOBJ        = "NULL"         // Null
	RETURNVALUEOBJ = "RETURN_VALUE" // Return
	ERROROBJ       = "ERROR"        // Error
	FUNCTIONOBJ    = "FUNCTION"     // Function
	STRINGOBJ      = "STRING"       // String
	BUILTINOBJ     = "BUILTIN"      // BuiltinFunction
)

// Object ...
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer ...
type Integer struct {
	Value int64
}

// Inspect ...
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Type ...
func (i *Integer) Type() ObjectType { return INTEGEROBJ }

// Boolean ...
type Boolean struct {
	Value bool
}

// Type ...
func (b *Boolean) Type() ObjectType { return BOOLEANOBJ }

// Inspect ...
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Null ...
type Null struct{}

// Type ...
func (n *Null) Type() ObjectType { return NULLOBJ }

// Inspect ...
func (n *Null) Inspect() string { return "null" }

// ReturnValue ....
type ReturnValue struct {
	Value Object
}

// Type ...
func (rv *ReturnValue) Type() ObjectType { return RETURNVALUEOBJ }

// Inspect ...
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// Error ...
type Error struct {
	Message string
}

// Type ...
func (e *Error) Type() ObjectType { return ERROROBJ }

// Inspect ...
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

// Function ...
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type ...
func (f *Function) Type() ObjectType { return FUNCTIONOBJ }

// Inspect ...
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

// String ...
type String struct {
	Value string
}

// Type ...
func (s *String) Type() ObjectType { return STRINGOBJ }

// Inspect ...
func (s *String) Inspect() string { return s.Value }

// BuiltinFunction ...
type BuiltinFunction func(args ...Object) Object

// Builtin ...
type Builtin struct {
	Fn BuiltinFunction
}

// Type ...
func (b *Builtin) Type() ObjectType { return BUILTINOBJ }

// Inspect ...
func (b *Builtin) Inspect() string { return "builtin function" }
