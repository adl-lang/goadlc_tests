// Code generated by goadlc v3 - DO NOT EDIT.
package struct01

import ()

type B struct {
	_B
}

type _B struct {
	A string `json:"a"`
}

func MakeAll_B(
	a string,
) B {
	return B{
		_B{
			A: a,
		},
	}
}

func Make_B(
	a string,
) B {
	ret := B{
		_B{
			A: a,
		},
	}
	return ret
}

type Struct01 struct {
	_Struct01
}

type _Struct01 struct {
	A struct{}                      `json:"A"`
	B int64                         `json:"B"`
	C string                        `json:"C"`
	D any                           `json:"d"`
	E []string                      `json:"e"`
	F map[string][]string           `json:"f"`
	G map[string]map[string]string  `json:"g"`
	H map[string]map[string]*string `json:"h"`
	I *string                       `json:"i"`
	J B                             `json:"j"`
}

func MakeAll_Struct01(
	b int64,
	c string,
	d any,
	e []string,
	f map[string][]string,
	g map[string]map[string]string,
	h map[string]map[string]*string,
	i *string,
	j B,
) Struct01 {
	return Struct01{
		_Struct01{
			A: struct{}{},
			B: b,
			C: c,
			D: d,
			E: e,
			F: f,
			G: g,
			H: h,
			I: i,
			J: j,
		},
	}
}

func Make_Struct01(
	b int64,
	c string,
	d any,
	e []string,
	f map[string][]string,
	g map[string]map[string]string,
	h map[string]map[string]*string,
	i *string,
	j B,
) Struct01 {
	ret := Struct01{
		_Struct01{
			A: struct{}{},
			B: b,
			C: c,
			D: d,
			E: e,
			F: f,
			G: g,
			H: h,
			I: i,
			J: j,
		},
	}
	return ret
}
