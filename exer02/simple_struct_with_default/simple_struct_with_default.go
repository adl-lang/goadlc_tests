// Code generated by goadlc v3 - DO NOT EDIT.
package simple_struct_with_default

import ()

type StructOfPrimitivesWithDefault struct {
	A int32   `json:"A"`
	B int64   `json:"B"`
	C bool    `json:"c"`
	D float64 `json:"d"`
	E string  `json:"e"`
}

func New_StructOfPrimitivesWithDefault(
	a int32,
	b int64,
	c bool,
	d float64,
	e string,
) StructOfPrimitivesWithDefault {
	return StructOfPrimitivesWithDefault{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
	}
}

func Make_StructOfPrimitivesWithDefault() StructOfPrimitivesWithDefault {
	ret := StructOfPrimitivesWithDefault{
		A: ((*StructOfPrimitivesWithDefault)(nil)).Default_A(),
		B: ((*StructOfPrimitivesWithDefault)(nil)).Default_B(),
		C: ((*StructOfPrimitivesWithDefault)(nil)).Default_c(),
		D: ((*StructOfPrimitivesWithDefault)(nil)).Default_d(),
		E: ((*StructOfPrimitivesWithDefault)(nil)).Default_e(),
	}
	return ret
}

func (*StructOfPrimitivesWithDefault) Default_A() int32 {
	return 1
}
func (*StructOfPrimitivesWithDefault) Default_B() int64 {
	return 2
}
func (*StructOfPrimitivesWithDefault) Default_c() bool {
	return true
}
func (*StructOfPrimitivesWithDefault) Default_d() float64 {
	return 1.1
}
func (*StructOfPrimitivesWithDefault) Default_e() string {
	return "byebye"
}