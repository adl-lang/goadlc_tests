// Code generated by goadlc v3 - DO NOT EDIT.
package tttest

import (
	"fmt"
	"github.com/adl-lang/goadl_rt/v3/sys/adlast"
)

type S1[A any] struct {
	_S1[A]
}

type _S1[A any] struct {
	A adlast.ATypeExpr[A] `json:"a"`
	B any                 `json:"b"`
}

func MakeAll_S1[A any](
	a adlast.ATypeExpr[A],
	b any,
) S1[A] {
	return S1[A]{
		_S1[A]{
			A: a,
			B: b,
		},
	}
}

// struct S1 contains at least one TypeToken, not generating Make_ funcs

type S2[A any] struct {
	_S2[A]
}

type _S2[A any] struct {
	A S1[string] `json:"a"`
	B S1[S3]     `json:"b"`
	C S1[[]S3]   `json:"c"`
}

func MakeAll_S2[A any](
	a S1[string],
	b S1[S3],
	c S1[[]S3],
) S2[A] {
	return S2[A]{
		_S2[A]{
			A: a,
			B: b,
			C: c,
		},
	}
}

func Make_S2[A any](
	a S1[string],
	b S1[S3],
	c S1[[]S3],
) S2[A] {
	ret := S2[A]{
		_S2[A]{
			A: a,
			B: b,
			C: c,
		},
	}
	return ret
}

type S3 struct {
	_S3
}

type _S3 struct {
	A adlast.ATypeExpr[string] `json:"a"`
}

func MakeAll_S3(
	a adlast.ATypeExpr[string],
) S3 {
	return S3{
		_S3{
			A: a,
		},
	}
}

// struct S3 contains at least one TypeToken, not generating Make_ funcs

type S4[A any] struct {
	_S4[A]
}

type _S4[A any] struct {
	Z S1[S4[A]] `json:"z"`
	A A         `json:"a"`
}

func MakeAll_S4[A any](
	z S1[S4[A]],
	a A,
) S4[A] {
	return S4[A]{
		_S4[A]{
			Z: z,
			A: a,
		},
	}
}

func Make_S4[A any](
	z S1[S4[A]],
	a A,
) S4[A] {
	ret := S4[A]{
		_S4[A]{
			Z: z,
			A: a,
		},
	}
	return ret
}

type S5[A any, B any] struct {
	_S5[A, B]
}

type _S5[A any, B any] struct {
	A adlast.ATypeExpr[A] `json:"a"`
	B adlast.ATypeExpr[B] `json:"b"`
}

func MakeAll_S5[A any, B any](
	a adlast.ATypeExpr[A],
	b adlast.ATypeExpr[B],
) S5[A, B] {
	return S5[A, B]{
		_S5[A, B]{
			A: a,
			B: b,
		},
	}
}

// struct S5 contains at least one TypeToken, not generating Make_ funcs

type T1 = S1[string]

type U1[A any] struct {
	Branch U1Branch[A]
}

type U1Branch[A any] interface {
	isU1Branch()
}

func (*U1[A]) MakeNewBranch(key string) (any, error) {
	switch key {
	case "a":
		return &_U1_A[A]{}, nil
	case "b":
		return &_U1_B{}, nil
	case "c":
		return &_U1_C{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _U1_A[A any] struct {
	V adlast.ATypeExpr[A] `branch:"a"`
}
type _U1_B struct {
	V string `branch:"b"`
}
type _U1_C struct {
	V struct{} `branch:"c"`
}

func (_U1_A[A]) isU1Branch() {}
func (_U1_B) isU1Branch()    {}
func (_U1_C) isU1Branch()    {}

func Make_U1_a[A any](v adlast.ATypeExpr[A]) U1[A] {
	return U1[A]{
		_U1_A[A]{v},
	}
}

func Make_U1_b[A any](v string) U1[A] {
	return U1[A]{
		_U1_B{v},
	}
}

func Make_U1_c[A any]() U1[A] {
	return U1[A]{
		_U1_C{struct{}{}},
	}
}

func (un U1[A]) Cast_a() (adlast.ATypeExpr[A], bool) {
	br, ok := un.Branch.(_U1_A[A])
	return br.V, ok
}

func (un U1[A]) Cast_b() (string, bool) {
	br, ok := un.Branch.(_U1_B)
	return br.V, ok
}

func (un U1[A]) Cast_c() (struct{}, bool) {
	br, ok := un.Branch.(_U1_C)
	return br.V, ok
}

func Handle_U1[A any, T any](
	_in U1[A],
	a func(a adlast.ATypeExpr[A]) T,
	b func(b string) T,
	c func(c struct{}) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _U1_A[A]:
		if a != nil {
			return a(_b.V)
		}
	case _U1_B:
		if b != nil {
			return b(_b.V)
		}
	case _U1_C:
		if c != nil {
			return c(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : U1")
}

func HandleWithErr_U1[A any, T any](
	_in U1[A],
	a func(a adlast.ATypeExpr[A]) (T, error),
	b func(b string) (T, error),
	c func(c struct{}) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _U1_A[A]:
		if a != nil {
			return a(_b.V)
		}
	case _U1_B:
		if b != nil {
			return b(_b.V)
		}
	case _U1_C:
		if c != nil {
			return c(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : U1")
}

type X[A any] S1[A]

type Y struct {
	_Y
}

type _Y struct {
	A S1[S4[string]] `json:"a"`
}

func MakeAll_Y(
	a S1[S4[string]],
) Y {
	return Y{
		_Y{
			A: a,
		},
	}
}

func Make_Y(
	a S1[S4[string]],
) Y {
	ret := Y{
		_Y{
			A: a,
		},
	}
	return ret
}

type Z struct {
	_Z
}

type _Z struct {
	A S1[string]        `json:"a"`
	B X[int64]          `json:"b"`
	C S5[string, int64] `json:"c"`
}

func MakeAll_Z(
	a S1[string],
	b X[int64],
	c S5[string, int64],
) Z {
	return Z{
		_Z{
			A: a,
			B: b,
			C: c,
		},
	}
}

func Make_Z() Z {
	ret := Z{
		_Z{
			A: ((*Z)(nil)).Default_a(),
			B: ((*Z)(nil)).Default_b(),
			C: ((*Z)(nil)).Default_c(),
		},
	}
	return ret
}

func (*Z) Default_a() S1[string] {
	return MakeAll_S1[string](
		adlast.Make_ATypeExpr[string](adlast.MakeAll_TypeExpr(
			adlast.Make_TypeRef_primitive(
				"String",
			),
			[]adlast.TypeExpr{},
		)),
		map[string]interface{}{"a": interface{}(nil)},
	)
}
func (*Z) Default_b() X[int64] {
	return X[int64](
		MakeAll_S1[int64](
			adlast.Make_ATypeExpr[int64](adlast.MakeAll_TypeExpr(
				adlast.Make_TypeRef_primitive(
					"Int64",
				),
				[]adlast.TypeExpr{},
			)),
			map[string]interface{}{"a": interface{}(nil)},
		),
	)
}
func (*Z) Default_c() S5[string, int64] {
	return MakeAll_S5[string, int64](
		adlast.Make_ATypeExpr[string](adlast.MakeAll_TypeExpr(
			adlast.Make_TypeRef_primitive(
				"String",
			),
			[]adlast.TypeExpr{},
		)),
		adlast.Make_ATypeExpr[int64](adlast.MakeAll_TypeExpr(
			adlast.Make_TypeRef_primitive(
				"Int64",
			),
			[]adlast.TypeExpr{},
		)),
	)
}
