package object

import "fmt"

// ObjectType ...
type ObjectType string

// ObjectType
const (
	INTEGEROBJ     = "INTEGER"      // Integer
	BOOLEANOBJ     = "BOOLEAN"      // Boolean
	NULLOBJ        = "NULL"         // Null
	RETURNVALUEOBJ = "RETURN_VALUE" // Return
	ERROROBJ       = "ERROR"        // Error
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
