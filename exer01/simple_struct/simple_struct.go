// Code generated by goadlc v3 - DO NOT EDIT.
package simple_struct

import ()

type StructOfPrimitives struct {
	_StructOfPrimitives
}

type _StructOfPrimitives struct {
	A int32   `json:"A"`
	B int64   `json:"B"`
	C bool    `json:"c"`
	D float64 `json:"d"`
	E string  `json:"e"`
}

func MakeAll_StructOfPrimitives(
	a int32,
	b int64,
	c bool,
	d float64,
	e string,
) StructOfPrimitives {
	return StructOfPrimitives{
		_StructOfPrimitives{
			A: a,
			B: b,
			C: c,
			D: d,
			E: e,
		},
	}
}

func Make_StructOfPrimitives(
	a int32,
	b int64,
	c bool,
	d float64,
	e string,
) StructOfPrimitives {
	ret := StructOfPrimitives{
		_StructOfPrimitives{
			A: a,
			B: b,
			C: c,
			D: d,
			E: e,
		},
	}
	return ret
}
