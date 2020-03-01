package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johny"}
	diff2 := &String{Value: "My name is johny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("string with same content have different hash keys")
	}

	if diff2.HashKey() != diff2.HashKey() {
		t.Errorf("string with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("string with differnent content have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	true1 := &Boolean{Value: true}
	true2 := &Boolean{Value: true}
	false1 := &Boolean{Value: false}
	false2 := &Boolean{Value: false}

	if true1.HashKey() != true2.HashKey() {
		t.Errorf("boolean with same content have different hash keys")
	}

	if false1.HashKey() != false2.HashKey() {
		t.Errorf("boolean with same content have different hash keys")
	}

	if true1.HashKey() == false1.HashKey() {
		t.Errorf("boolean with differnent content have same hash keys")
	}
}

func TestIntegerHashKey(t *testing.T) {
	one1 := &Integer{Value: 1}
	one2 := &Integer{Value: 1}
	ten1 := &Integer{Value: 10}
	ten2 := &Integer{Value: 10}

	if one1.HashKey() != one2.HashKey() {
		t.Errorf("integer with same content have different hash keys")
	}

	if ten1.HashKey() != ten2.HashKey() {
		t.Errorf("integer with same content have different hash keys")
	}

	if one1.HashKey() == ten1.HashKey() {
		t.Errorf("integer with differnent content have same hash keys")
	}
}
