package object

import (
	"ast"
	"bytes"
	"fmt"
	"hash/fnv"
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
	ARRAYOBJ       = "ARRAY"        // array
	HASHOBJ        = "HASH"         // hash
	QUOTEOBJ       = "QUOTE"        // quote
	MACROOBJ       = "MACRO"        // macro
)

// Object ...
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Hashable ...
type Hashable interface {
	HashKey() HashKey
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

// Array ...
type Array struct {
	Elements []Object
}

// Type ...
func (ao *Array) Type() ObjectType { return ARRAYOBJ }

// Inspect ...
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// HashKey ...
type HashKey struct {
	Type  ObjectType
	Value uint64
}

// HashKey ...
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

// HashKey ...
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// HashKey ...
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// HashPair ...
type HashPair struct {
	Key   Object
	Value Object
}

// Hash ...
type Hash struct {
	Pairs map[HashKey]HashPair
}

// Type ...
func (h *Hash) Type() ObjectType { return HASHOBJ }

// Inspect ...
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

// Quote ...
type Quote struct {
	Node ast.Node
}

// Type ...
func (q *Quote) Type() ObjectType { return QUOTEOBJ }

// Inspect ...
func (q *Quote) Inspect() string {
	return "QUOTE(" + q.Node.String() + ")"
}

// Macro ...
type Macro struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type ...
func (m *Macro) Type() ObjectType { return MACROOBJ }

// Inspect ...
func (m *Macro) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range m.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(m.Body.String())
	out.WriteString("\n}")

	return out.String()
}
