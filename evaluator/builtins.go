package evaluator

import (
	"ig/object"
)

var builtins = map[string]*object.Builtin{
	"len":   object.GetBuildinByName("len"),
	"first": object.GetBuildinByName("first"),
	"last":  object.GetBuildinByName("last"),
	"rest":  object.GetBuildinByName("rest"),
	"push":  object.GetBuildinByName("push"),
	"puts":  object.GetBuildinByName("puts"),
}
