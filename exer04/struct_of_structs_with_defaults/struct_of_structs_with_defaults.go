// Code generated by goadlc v3 - DO NOT EDIT.
package struct_of_structs_with_defaults

import ()

type Bar struct {
	_Bar
}

type _Bar struct {
	D   string            `json:"d"`
	Svb map[string][]bool `json:"svb"`
}

func MakeAll_Bar(
	d string,
	svb map[string][]bool,
) Bar {
	return Bar{
		_Bar{
			D:   d,
			Svb: svb,
		},
	}
}

func Make_Bar() Bar {
	ret := Bar{
		_Bar{
			D:   ((*Bar)(nil)).Default_d(),
			Svb: ((*Bar)(nil)).Default_svb(),
		},
	}
	return ret
}

func (*Bar) Default_d() string {
	return "doe a deer"
}
func (*Bar) Default_svb() map[string][]bool {
	return map[string][]bool{
		"a": []bool{
			true,
			false,
		},
		"b": []bool{
			false,
			true,
		},
	}
}

type Fizz struct {
	_Fizz
}

type _Fizz struct {
	A Foo `json:"A"`
	B Bar `json:"B"`
}

func MakeAll_Fizz(
	a Foo,
	b Bar,
) Fizz {
	return Fizz{
		_Fizz{
			A: a,
			B: b,
		},
	}
}

func Make_Fizz(
	a Foo,
	b Bar,
) Fizz {
	ret := Fizz{
		_Fizz{
			A: a,
			B: b,
		},
	}
	return ret
}

type Foo struct {
	_Foo
}

type _Foo struct {
	D int32 `json:"d"`
}

func MakeAll_Foo(
	d int32,
) Foo {
	return Foo{
		_Foo{
			D: d,
		},
	}
}

func Make_Foo() Foo {
	ret := Foo{
		_Foo{
			D: ((*Foo)(nil)).Default_d(),
		},
	}
	return ret
}

func (*Foo) Default_d() int32 {
	return 1
}

type NT StructOfStruct

type StructOfStruct struct {
	_StructOfStruct
}

type _StructOfStruct struct {
	A Foo  `json:"A"`
	B Bar  `json:"B"`
	C Fizz `json:"c"`
}

func MakeAll_StructOfStruct(
	a Foo,
	b Bar,
	c Fizz,
) StructOfStruct {
	return StructOfStruct{
		_StructOfStruct{
			A: a,
			B: b,
			C: c,
		},
	}
}

func Make_StructOfStruct(
	a Foo,
	b Bar,
	c Fizz,
) StructOfStruct {
	ret := StructOfStruct{
		_StructOfStruct{
			A: a,
			B: b,
			C: c,
		},
	}
	return ret
}
