// Code generated by goadlc v3 - DO NOT EDIT.
package a

import (
	b2 "adl_testing/exer02/another/b"
	"adl_testing/exer02/b"
)

type A struct {
	A b.B  `json:"a"`
	B b2.B `json:"b"`
}

func New_A(
	a b.B,
	b b2.B,
) A {
	return A{
		A: a,
		B: b,
	}
}

func Make_A(
	a b.B,
	b b2.B,
) A {
	ret := A{
		A: a,
		B: b,
	}
	return ret
}
