// Code generated by goadlc v3 - DO NOT EDIT.
package test31

import (
	"fmt"
	"github.com/adl-lang/goadl_rt/v3/sys/types"
)

type A1 struct {
	Branch A1Branch
}

type A1Branch interface {
	isA1Branch()
}

func (*A1) MakeNewBranch(key string) (any, error) {
	switch key {
	case "a":
		return &_A1_A{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _A1_A struct {
	V A2 `branch:"a"`
}

func (_A1_A) isA1Branch() {}

func Make_A1_a(v A2) A1 {
	return A1{
		_A1_A{v},
	}
}

func (un A1) Cast_a() (A2, bool) {
	br, ok := un.Branch.(_A1_A)
	return br.V, ok
}

func Handle_A1[T any](
	_in A1,
	a func(a A2) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _A1_A:
		if a != nil {
			return a(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : A1")
}

func HandleWithErr_A1[T any](
	_in A1,
	a func(a A2) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _A1_A:
		if a != nil {
			return a(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : A1")
}

type A2 struct {
	Branch A2Branch
}

type A2Branch interface {
	isA2Branch()
}

func (*A2) MakeNewBranch(key string) (any, error) {
	switch key {
	case "b":
		return &_A2_B{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _A2_B struct {
	V A3 `branch:"b"`
}

func (_A2_B) isA2Branch() {}

func Make_A2_b(v A3) A2 {
	return A2{
		_A2_B{v},
	}
}

func (un A2) Cast_b() (A3, bool) {
	br, ok := un.Branch.(_A2_B)
	return br.V, ok
}

func Handle_A2[T any](
	_in A2,
	b func(b A3) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _A2_B:
		if b != nil {
			return b(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : A2")
}

func HandleWithErr_A2[T any](
	_in A2,
	b func(b A3) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _A2_B:
		if b != nil {
			return b(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : A2")
}

type A3 struct {
	Branch A3Branch
}

type A3Branch interface {
	isA3Branch()
}

func (*A3) MakeNewBranch(key string) (any, error) {
	switch key {
	case "c":
		return &_A3_C{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _A3_C struct {
	V A4 `branch:"c"`
}

func (_A3_C) isA3Branch() {}

func Make_A3_c(v A4) A3 {
	return A3{
		_A3_C{v},
	}
}

func (un A3) Cast_c() (A4, bool) {
	br, ok := un.Branch.(_A3_C)
	return br.V, ok
}

func Handle_A3[T any](
	_in A3,
	c func(c A4) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _A3_C:
		if c != nil {
			return c(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : A3")
}

func HandleWithErr_A3[T any](
	_in A3,
	c func(c A4) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _A3_C:
		if c != nil {
			return c(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : A3")
}

type A4 struct {
	Branch A4Branch
}

type A4Branch interface {
	isA4Branch()
}

func (*A4) MakeNewBranch(key string) (any, error) {
	switch key {
	case "d":
		return &_A4_D{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _A4_D struct {
	V string `branch:"d"`
}

func (_A4_D) isA4Branch() {}

func Make_A4_d(v string) A4 {
	return A4{
		_A4_D{v},
	}
}

func (un A4) Cast_d() (string, bool) {
	br, ok := un.Branch.(_A4_D)
	return br.V, ok
}

func Handle_A4[T any](
	_in A4,
	d func(d string) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _A4_D:
		if d != nil {
			return d(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : A4")
}

func HandleWithErr_A4[T any](
	_in A4,
	d func(d string) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _A4_D:
		if d != nil {
			return d(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : A4")
}

type B1 struct {
	Branch B1Branch
}

type B1Branch interface {
	isB1Branch()
}

func (*B1) MakeNewBranch(key string) (any, error) {
	switch key {
	case "f1":
		return &_B1_F1{}, nil
	case "f2":
		return &_B1_F2{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _B1_F1 struct {
	V MyString `branch:"f1"`
}
type _B1_F2 struct {
	V []MyString `branch:"f2"`
}

func (_B1_F1) isB1Branch() {}
func (_B1_F2) isB1Branch() {}

func Make_B1_f1(v MyString) B1 {
	return B1{
		_B1_F1{v},
	}
}

func Make_B1_f2(v []MyString) B1 {
	return B1{
		_B1_F2{v},
	}
}

func (un B1) Cast_f1() (MyString, bool) {
	br, ok := un.Branch.(_B1_F1)
	return br.V, ok
}

func (un B1) Cast_f2() ([]MyString, bool) {
	br, ok := un.Branch.(_B1_F2)
	return br.V, ok
}

func Handle_B1[T any](
	_in B1,
	f1 func(f1 MyString) T,
	f2 func(f2 []MyString) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _B1_F1:
		if f1 != nil {
			return f1(_b.V)
		}
	case _B1_F2:
		if f2 != nil {
			return f2(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : B1")
}

func HandleWithErr_B1[T any](
	_in B1,
	f1 func(f1 MyString) (T, error),
	f2 func(f2 []MyString) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _B1_F1:
		if f1 != nil {
			return f1(_b.V)
		}
	case _B1_F2:
		if f2 != nil {
			return f2(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : B1")
}

type B2 struct {
	Branch B2Branch
}

type B2Branch interface {
	isB2Branch()
}

func (*B2) MakeNewBranch(key string) (any, error) {
	switch key {
	case "f3":
		return &_B2_F3{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _B2_F3 struct {
	V MyVecString `branch:"f3"`
}

func (_B2_F3) isB2Branch() {}

func Make_B2_f3(v MyVecString) B2 {
	return B2{
		_B2_F3{v},
	}
}

func (un B2) Cast_f3() (MyVecString, bool) {
	br, ok := un.Branch.(_B2_F3)
	return br.V, ok
}

func Handle_B2[T any](
	_in B2,
	f3 func(f3 MyVecString) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _B2_F3:
		if f3 != nil {
			return f3(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : B2")
}

func HandleWithErr_B2[T any](
	_in B2,
	f3 func(f3 MyVecString) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _B2_F3:
		if f3 != nil {
			return f3(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : B2")
}

type B3 struct {
	Branch B3Branch
}

type B3Branch interface {
	isB3Branch()
}

func (*B3) MakeNewBranch(key string) (any, error) {
	switch key {
	case "f3":
		return &_B3_F3{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _B3_F3 struct {
	V MyVecMyString `branch:"f3"`
}

func (_B3_F3) isB3Branch() {}

func Make_B3_f3(v MyVecMyString) B3 {
	return B3{
		_B3_F3{v},
	}
}

func (un B3) Cast_f3() (MyVecMyString, bool) {
	br, ok := un.Branch.(_B3_F3)
	return br.V, ok
}

func Handle_B3[T any](
	_in B3,
	f3 func(f3 MyVecMyString) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _B3_F3:
		if f3 != nil {
			return f3(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : B3")
}

func HandleWithErr_B3[T any](
	_in B3,
	f3 func(f3 MyVecMyString) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _B3_F3:
		if f3 != nil {
			return f3(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : B3")
}

type B4 struct {
	Branch B4Branch
}

type B4Branch interface {
	isB4Branch()
}

func (*B4) MakeNewBranch(key string) (any, error) {
	switch key {
	case "b":
		return &_B4_B{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _B4_B struct {
	V GenU[string] `branch:"b"`
}

func (_B4_B) isB4Branch() {}

func Make_B4_b(v GenU[string]) B4 {
	return B4{
		_B4_B{v},
	}
}

func (un B4) Cast_b() (GenU[string], bool) {
	br, ok := un.Branch.(_B4_B)
	return br.V, ok
}

func Handle_B4[T any](
	_in B4,
	b func(b GenU[string]) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _B4_B:
		if b != nil {
			return b(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : B4")
}

func HandleWithErr_B4[T any](
	_in B4,
	b func(b GenU[string]) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _B4_B:
		if b != nil {
			return b(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : B4")
}

type B5 struct {
	Branch B5Branch
}

type B5Branch interface {
	isB5Branch()
}

func (*B5) MakeNewBranch(key string) (any, error) {
	switch key {
	case "b":
		return &_B5_B{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _B5_B struct {
	V MyGenUString `branch:"b"`
}

func (_B5_B) isB5Branch() {}

func Make_B5_b(v MyGenUString) B5 {
	return B5{
		_B5_B{v},
	}
}

func (un B5) Cast_b() (MyGenUString, bool) {
	br, ok := un.Branch.(_B5_B)
	return br.V, ok
}

func Handle_B5[T any](
	_in B5,
	b func(b MyGenUString) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _B5_B:
		if b != nil {
			return b(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : B5")
}

func HandleWithErr_B5[T any](
	_in B5,
	b func(b MyGenUString) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _B5_B:
		if b != nil {
			return b(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : B5")
}

type GenU[T any] struct {
	Branch GenUBranch[T]
}

type GenUBranch[T any] interface {
	isGenUBranch()
}

func (*GenU[T]) MakeNewBranch(key string) (any, error) {
	switch key {
	case "a":
		return &_GenU_A[T]{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _GenU_A[T any] struct {
	V T `branch:"a"`
}

func (_GenU_A[T]) isGenUBranch() {}

func Make_GenU_a[T any](v T) GenU[T] {
	return GenU[T]{
		_GenU_A[T]{v},
	}
}

func (un GenU[T]) Cast_a() (T, bool) {
	br, ok := un.Branch.(_GenU_A[T])
	return br.V, ok
}

func Handle_GenU[T any, T2 any](
	_in GenU[T],
	a func(a T) T2,
	_default func() T2,
) T2 {
	switch _b := _in.Branch.(type) {
	case _GenU_A[T]:
		if a != nil {
			return a(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : GenU")
}

func HandleWithErr_GenU[T any, T2 any](
	_in GenU[T],
	a func(a T) (T2, error),
	_default func() (T2, error),
) (T2, error) {
	switch _b := _in.Branch.(type) {
	case _GenU_A[T]:
		if a != nil {
			return a(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : GenU")
}

type Measure struct {
	Branch MeasureBranch
}

type MeasureBranch interface {
	isMeasureBranch()
}

func (*Measure) MakeNewBranch(key string) (any, error) {
	switch key {
	case "count":
		return &_Measure_Count{}, nil
	case "multiple":
		return &_Measure_Multiple{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _Measure_Count struct {
	V uint64 `branch:"count"`
}
type _Measure_Multiple struct {
	V types.Pair[uint64, float64] `branch:"multiple"`
}

func (_Measure_Count) isMeasureBranch()    {}
func (_Measure_Multiple) isMeasureBranch() {}

func Make_Measure_count(v uint64) Measure {
	return Measure{
		_Measure_Count{v},
	}
}

func Make_Measure_multiple(v types.Pair[uint64, float64]) Measure {
	return Measure{
		_Measure_Multiple{v},
	}
}

func (un Measure) Cast_count() (uint64, bool) {
	br, ok := un.Branch.(_Measure_Count)
	return br.V, ok
}

func (un Measure) Cast_multiple() (types.Pair[uint64, float64], bool) {
	br, ok := un.Branch.(_Measure_Multiple)
	return br.V, ok
}

func Handle_Measure[T any](
	_in Measure,
	count func(count uint64) T,
	multiple func(multiple types.Pair[uint64, float64]) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _Measure_Count:
		if count != nil {
			return count(_b.V)
		}
	case _Measure_Multiple:
		if multiple != nil {
			return multiple(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : Measure")
}

func HandleWithErr_Measure[T any](
	_in Measure,
	count func(count uint64) (T, error),
	multiple func(multiple types.Pair[uint64, float64]) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _Measure_Count:
		if count != nil {
			return count(_b.V)
		}
	case _Measure_Multiple:
		if multiple != nil {
			return multiple(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : Measure")
}

type MyGenUString = GenU[string]

type MyString = string

type MyStructV1 struct {
	_MyStructV1
}

type _MyStructV1 struct {
	Quantity uint64 `json:"quantity"`
}

func MakeAll_MyStructV1(
	quantity uint64,
) MyStructV1 {
	return MyStructV1{
		_MyStructV1{
			Quantity: quantity,
		},
	}
}

func Make_MyStructV1(
	quantity uint64,
) MyStructV1 {
	ret := MyStructV1{
		_MyStructV1{
			Quantity: quantity,
		},
	}
	return ret
}

type MyStructV2 struct {
	_MyStructV2
}

type _MyStructV2 struct {
	Quantity Measure `json:"quantity"`
}

func MakeAll_MyStructV2(
	quantity Measure,
) MyStructV2 {
	return MyStructV2{
		_MyStructV2{
			Quantity: quantity,
		},
	}
}

func Make_MyStructV2(
	quantity Measure,
) MyStructV2 {
	ret := MyStructV2{
		_MyStructV2{
			Quantity: quantity,
		},
	}
	return ret
}

type MyVecMyString = []MyString

type MyVecString = []string

type NullableTest struct {
	Branch NullableTestBranch
}

type NullableTestBranch interface {
	isNullableTestBranch()
}

func (*NullableTest) MakeNewBranch(key string) (any, error) {
	switch key {
	case "a":
		return &_NullableTest_A{}, nil
	case "b":
		return &_NullableTest_B{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _NullableTest_A struct {
	V *string `branch:"a"`
}
type _NullableTest_B struct {
	V struct{} `branch:"b"`
}

func (_NullableTest_A) isNullableTestBranch() {}
func (_NullableTest_B) isNullableTestBranch() {}

func Make_NullableTest_a(v *string) NullableTest {
	return NullableTest{
		_NullableTest_A{v},
	}
}

func Make_NullableTest_b() NullableTest {
	return NullableTest{
		_NullableTest_B{struct{}{}},
	}
}

func (un NullableTest) Cast_a() (*string, bool) {
	br, ok := un.Branch.(_NullableTest_A)
	return br.V, ok
}

func (un NullableTest) Cast_b() (struct{}, bool) {
	br, ok := un.Branch.(_NullableTest_B)
	return br.V, ok
}

func Handle_NullableTest[T any](
	_in NullableTest,
	a func(a *string) T,
	b func(b struct{}) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _NullableTest_A:
		if a != nil {
			return a(_b.V)
		}
	case _NullableTest_B:
		if b != nil {
			return b(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : NullableTest")
}

func HandleWithErr_NullableTest[T any](
	_in NullableTest,
	a func(a *string) (T, error),
	b func(b struct{}) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _NullableTest_A:
		if a != nil {
			return a(_b.V)
		}
	case _NullableTest_B:
		if b != nil {
			return b(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : NullableTest")
}

type S1 struct {
	_S1
}

type _S1 struct {
	Quant uint64  `json:"quant"`
	Value float64 `json:"value"`
}

func MakeAll_S1(
	quant uint64,
	value float64,
) S1 {
	return S1{
		_S1{
			Quant: quant,
			Value: value,
		},
	}
}

func Make_S1(
	quant uint64,
	value float64,
) S1 {
	ret := S1{
		_S1{
			Quant: quant,
			Value: value,
		},
	}
	return ret
}

type Struct02Test struct {
	Branch Struct02TestBranch
}

type Struct02TestBranch interface {
	isStruct02TestBranch()
}

func (*Struct02Test) MakeNewBranch(key string) (any, error) {
	switch key {
	case "u1":
		return &_Struct02Test_U1{}, nil
	case "s1":
		return &_Struct02Test_S1{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _Struct02Test_U1 struct {
	V U1 `branch:"u1"`
}
type _Struct02Test_S1 struct {
	V S1 `branch:"s1"`
}

func (_Struct02Test_U1) isStruct02TestBranch() {}
func (_Struct02Test_S1) isStruct02TestBranch() {}

func Make_Struct02Test_u1(v U1) Struct02Test {
	return Struct02Test{
		_Struct02Test_U1{v},
	}
}

func Make_Struct02Test_s1(v S1) Struct02Test {
	return Struct02Test{
		_Struct02Test_S1{v},
	}
}

func (un Struct02Test) Cast_u1() (U1, bool) {
	br, ok := un.Branch.(_Struct02Test_U1)
	return br.V, ok
}

func (un Struct02Test) Cast_s1() (S1, bool) {
	br, ok := un.Branch.(_Struct02Test_S1)
	return br.V, ok
}

func Handle_Struct02Test[T any](
	_in Struct02Test,
	u1 func(u1 U1) T,
	s1 func(s1 S1) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _Struct02Test_U1:
		if u1 != nil {
			return u1(_b.V)
		}
	case _Struct02Test_S1:
		if s1 != nil {
			return s1(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : Struct02Test")
}

func HandleWithErr_Struct02Test[T any](
	_in Struct02Test,
	u1 func(u1 U1) (T, error),
	s1 func(s1 S1) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _Struct02Test_U1:
		if u1 != nil {
			return u1(_b.V)
		}
	case _Struct02Test_S1:
		if s1 != nil {
			return s1(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : Struct02Test")
}

type StructTest struct {
	Branch StructTestBranch
}

type StructTestBranch interface {
	isStructTestBranch()
}

func (*StructTest) MakeNewBranch(key string) (any, error) {
	switch key {
	case "abc":
		return &_StructTest_Abc{}, nil
	case "def":
		return &_StructTest_Def{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _StructTest_Abc struct {
	V S1 `branch:"abc"`
}
type _StructTest_Def struct {
	V S1 `branch:"def"`
}

func (_StructTest_Abc) isStructTestBranch() {}
func (_StructTest_Def) isStructTestBranch() {}

func Make_StructTest_abc(v S1) StructTest {
	return StructTest{
		_StructTest_Abc{v},
	}
}

func Make_StructTest_def(v S1) StructTest {
	return StructTest{
		_StructTest_Def{v},
	}
}

func (un StructTest) Cast_abc() (S1, bool) {
	br, ok := un.Branch.(_StructTest_Abc)
	return br.V, ok
}

func (un StructTest) Cast_def() (S1, bool) {
	br, ok := un.Branch.(_StructTest_Def)
	return br.V, ok
}

func Handle_StructTest[T any](
	_in StructTest,
	abc func(abc S1) T,
	def func(def S1) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _StructTest_Abc:
		if abc != nil {
			return abc(_b.V)
		}
	case _StructTest_Def:
		if def != nil {
			return def(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : StructTest")
}

func HandleWithErr_StructTest[T any](
	_in StructTest,
	abc func(abc S1) (T, error),
	def func(def S1) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _StructTest_Abc:
		if abc != nil {
			return abc(_b.V)
		}
	case _StructTest_Def:
		if def != nil {
			return def(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : StructTest")
}

type U1 struct {
	Branch U1Branch
}

type U1Branch interface {
	isU1Branch()
}

func (*U1) MakeNewBranch(key string) (any, error) {
	switch key {
	case "a":
		return &_U1_A{}, nil
	case "b":
		return &_U1_B{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _U1_A struct {
	V string `branch:"a"`
}
type _U1_B struct {
	V int64 `branch:"b"`
}

func (_U1_A) isU1Branch() {}
func (_U1_B) isU1Branch() {}

func Make_U1_a(v string) U1 {
	return U1{
		_U1_A{v},
	}
}

func Make_U1_b(v int64) U1 {
	return U1{
		_U1_B{v},
	}
}

func (un U1) Cast_a() (string, bool) {
	br, ok := un.Branch.(_U1_A)
	return br.V, ok
}

func (un U1) Cast_b() (int64, bool) {
	br, ok := un.Branch.(_U1_B)
	return br.V, ok
}

func Handle_U1[T any](
	_in U1,
	a func(a string) T,
	b func(b int64) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _U1_A:
		if a != nil {
			return a(_b.V)
		}
	case _U1_B:
		if b != nil {
			return b(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : U1")
}

func HandleWithErr_U1[T any](
	_in U1,
	a func(a string) (T, error),
	b func(b int64) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _U1_A:
		if a != nil {
			return a(_b.V)
		}
	case _U1_B:
		if b != nil {
			return b(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : U1")
}

type UTypeDiscriminationofMaybe struct {
	Branch UTypeDiscriminationofMaybeBranch
}

type UTypeDiscriminationofMaybeBranch interface {
	isUTypeDiscriminationofMaybeBranch()
}

func (*UTypeDiscriminationofMaybe) MakeNewBranch(key string) (any, error) {
	switch key {
	case "a":
		return &_UTypeDiscriminationofMaybe_A{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _UTypeDiscriminationofMaybe_A struct {
	V types.Maybe[UnionUnion] `branch:"a"`
}

func (_UTypeDiscriminationofMaybe_A) isUTypeDiscriminationofMaybeBranch() {}

func Make_UTypeDiscriminationofMaybe_a(v types.Maybe[UnionUnion]) UTypeDiscriminationofMaybe {
	return UTypeDiscriminationofMaybe{
		_UTypeDiscriminationofMaybe_A{v},
	}
}

func (un UTypeDiscriminationofMaybe) Cast_a() (types.Maybe[UnionUnion], bool) {
	br, ok := un.Branch.(_UTypeDiscriminationofMaybe_A)
	return br.V, ok
}

func Handle_UTypeDiscriminationofMaybe[T any](
	_in UTypeDiscriminationofMaybe,
	a func(a types.Maybe[UnionUnion]) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _UTypeDiscriminationofMaybe_A:
		if a != nil {
			return a(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : UTypeDiscriminationofMaybe")
}

func HandleWithErr_UTypeDiscriminationofMaybe[T any](
	_in UTypeDiscriminationofMaybe,
	a func(a types.Maybe[UnionUnion]) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _UTypeDiscriminationofMaybe_A:
		if a != nil {
			return a(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : UTypeDiscriminationofMaybe")
}

type Ua struct {
	Branch UaBranch
}

type UaBranch interface {
	isUaBranch()
}

func (*Ua) MakeNewBranch(key string) (any, error) {
	switch key {
	case "u_b":
		return &_Ua_U_b{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _Ua_U_b struct {
	V []Ub `branch:"u_b"`
}

func (_Ua_U_b) isUaBranch() {}

func Make_Ua_u_b(v []Ub) Ua {
	return Ua{
		_Ua_U_b{v},
	}
}

func (un Ua) Cast_u_b() ([]Ub, bool) {
	br, ok := un.Branch.(_Ua_U_b)
	return br.V, ok
}

func Handle_Ua[T any](
	_in Ua,
	u_b func(u_b []Ub) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _Ua_U_b:
		if u_b != nil {
			return u_b(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : Ua")
}

func HandleWithErr_Ua[T any](
	_in Ua,
	u_b func(u_b []Ub) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _Ua_U_b:
		if u_b != nil {
			return u_b(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : Ua")
}

type Ub struct {
	Branch UbBranch
}

type UbBranch interface {
	isUbBranch()
}

func (*Ub) MakeNewBranch(key string) (any, error) {
	switch key {
	case "u_c":
		return &_Ub_U_c{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _Ub_U_c struct {
	V []Uc `branch:"u_c"`
}

func (_Ub_U_c) isUbBranch() {}

func Make_Ub_u_c(v []Uc) Ub {
	return Ub{
		_Ub_U_c{v},
	}
}

func (un Ub) Cast_u_c() ([]Uc, bool) {
	br, ok := un.Branch.(_Ub_U_c)
	return br.V, ok
}

func Handle_Ub[T any](
	_in Ub,
	u_c func(u_c []Uc) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _Ub_U_c:
		if u_c != nil {
			return u_c(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : Ub")
}

func HandleWithErr_Ub[T any](
	_in Ub,
	u_c func(u_c []Uc) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _Ub_U_c:
		if u_c != nil {
			return u_c(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : Ub")
}

type Uc struct {
	Branch UcBranch
}

type UcBranch interface {
	isUcBranch()
}

func (*Uc) MakeNewBranch(key string) (any, error) {
	switch key {
	case "a":
		return &_Uc_A{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _Uc_A struct {
	V []string `branch:"a"`
}

func (_Uc_A) isUcBranch() {}

func Make_Uc_a(v []string) Uc {
	return Uc{
		_Uc_A{v},
	}
}

func (un Uc) Cast_a() ([]string, bool) {
	br, ok := un.Branch.(_Uc_A)
	return br.V, ok
}

func Handle_Uc[T any](
	_in Uc,
	a func(a []string) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _Uc_A:
		if a != nil {
			return a(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : Uc")
}

func HandleWithErr_Uc[T any](
	_in Uc,
	a func(a []string) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _Uc_A:
		if a != nil {
			return a(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : Uc")
}

type UnionOfLiftedUnion struct {
	Branch UnionOfLiftedUnionBranch
}

type UnionOfLiftedUnionBranch interface {
	isUnionOfLiftedUnionBranch()
}

func (*UnionOfLiftedUnion) MakeNewBranch(key string) (any, error) {
	switch key {
	case "abc":
		return &_UnionOfLiftedUnion_Abc{}, nil
	case "def":
		return &_UnionOfLiftedUnion_Def{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _UnionOfLiftedUnion_Abc struct {
	V UnionOfLiftedUnion `branch:"abc"`
}
type _UnionOfLiftedUnion_Def struct {
	V S1 `branch:"def"`
}

func (_UnionOfLiftedUnion_Abc) isUnionOfLiftedUnionBranch() {}
func (_UnionOfLiftedUnion_Def) isUnionOfLiftedUnionBranch() {}

func Make_UnionOfLiftedUnion_abc(v UnionOfLiftedUnion) UnionOfLiftedUnion {
	return UnionOfLiftedUnion{
		_UnionOfLiftedUnion_Abc{v},
	}
}

func Make_UnionOfLiftedUnion_def(v S1) UnionOfLiftedUnion {
	return UnionOfLiftedUnion{
		_UnionOfLiftedUnion_Def{v},
	}
}

func (un UnionOfLiftedUnion) Cast_abc() (UnionOfLiftedUnion, bool) {
	br, ok := un.Branch.(_UnionOfLiftedUnion_Abc)
	return br.V, ok
}

func (un UnionOfLiftedUnion) Cast_def() (S1, bool) {
	br, ok := un.Branch.(_UnionOfLiftedUnion_Def)
	return br.V, ok
}

func Handle_UnionOfLiftedUnion[T any](
	_in UnionOfLiftedUnion,
	abc func(abc UnionOfLiftedUnion) T,
	def func(def S1) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _UnionOfLiftedUnion_Abc:
		if abc != nil {
			return abc(_b.V)
		}
	case _UnionOfLiftedUnion_Def:
		if def != nil {
			return def(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : UnionOfLiftedUnion")
}

func HandleWithErr_UnionOfLiftedUnion[T any](
	_in UnionOfLiftedUnion,
	abc func(abc UnionOfLiftedUnion) (T, error),
	def func(def S1) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _UnionOfLiftedUnion_Abc:
		if abc != nil {
			return abc(_b.V)
		}
	case _UnionOfLiftedUnion_Def:
		if def != nil {
			return def(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : UnionOfLiftedUnion")
}

type UnionOfVectorOfTypeDiscriminationUnionTest struct {
	Branch UnionOfVectorOfTypeDiscriminationUnionTestBranch
}

type UnionOfVectorOfTypeDiscriminationUnionTestBranch interface {
	isUnionOfVectorOfTypeDiscriminationUnionTestBranch()
}

func (*UnionOfVectorOfTypeDiscriminationUnionTest) MakeNewBranch(key string) (any, error) {
	switch key {
	case "b":
		return &_UnionOfVectorOfTypeDiscriminationUnionTest_B{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _UnionOfVectorOfTypeDiscriminationUnionTest_B struct {
	V []Struct02Test `branch:"b"`
}

func (_UnionOfVectorOfTypeDiscriminationUnionTest_B) isUnionOfVectorOfTypeDiscriminationUnionTestBranch() {
}

func Make_UnionOfVectorOfTypeDiscriminationUnionTest_b(v []Struct02Test) UnionOfVectorOfTypeDiscriminationUnionTest {
	return UnionOfVectorOfTypeDiscriminationUnionTest{
		_UnionOfVectorOfTypeDiscriminationUnionTest_B{v},
	}
}

func (un UnionOfVectorOfTypeDiscriminationUnionTest) Cast_b() ([]Struct02Test, bool) {
	br, ok := un.Branch.(_UnionOfVectorOfTypeDiscriminationUnionTest_B)
	return br.V, ok
}

func Handle_UnionOfVectorOfTypeDiscriminationUnionTest[T any](
	_in UnionOfVectorOfTypeDiscriminationUnionTest,
	b func(b []Struct02Test) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _UnionOfVectorOfTypeDiscriminationUnionTest_B:
		if b != nil {
			return b(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : UnionOfVectorOfTypeDiscriminationUnionTest")
}

func HandleWithErr_UnionOfVectorOfTypeDiscriminationUnionTest[T any](
	_in UnionOfVectorOfTypeDiscriminationUnionTest,
	b func(b []Struct02Test) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _UnionOfVectorOfTypeDiscriminationUnionTest_B:
		if b != nil {
			return b(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : UnionOfVectorOfTypeDiscriminationUnionTest")
}

type UnionUnion struct {
	Branch UnionUnionBranch
}

type UnionUnionBranch interface {
	isUnionUnionBranch()
}

func (*UnionUnion) MakeNewBranch(key string) (any, error) {
	switch key {
	case "abc":
		return &_UnionUnion_Abc{}, nil
	case "def":
		return &_UnionUnion_Def{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _UnionUnion_Abc struct {
	V U1 `branch:"abc"`
}
type _UnionUnion_Def struct {
	V S1 `branch:"def"`
}

func (_UnionUnion_Abc) isUnionUnionBranch() {}
func (_UnionUnion_Def) isUnionUnionBranch() {}

func Make_UnionUnion_abc(v U1) UnionUnion {
	return UnionUnion{
		_UnionUnion_Abc{v},
	}
}

func Make_UnionUnion_def(v S1) UnionUnion {
	return UnionUnion{
		_UnionUnion_Def{v},
	}
}

func (un UnionUnion) Cast_abc() (U1, bool) {
	br, ok := un.Branch.(_UnionUnion_Abc)
	return br.V, ok
}

func (un UnionUnion) Cast_def() (S1, bool) {
	br, ok := un.Branch.(_UnionUnion_Def)
	return br.V, ok
}

func Handle_UnionUnion[T any](
	_in UnionUnion,
	abc func(abc U1) T,
	def func(def S1) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _UnionUnion_Abc:
		if abc != nil {
			return abc(_b.V)
		}
	case _UnionUnion_Def:
		if def != nil {
			return def(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : UnionUnion")
}

func HandleWithErr_UnionUnion[T any](
	_in UnionUnion,
	abc func(abc U1) (T, error),
	def func(def S1) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _UnionUnion_Abc:
		if abc != nil {
			return abc(_b.V)
		}
	case _UnionUnion_Def:
		if def != nil {
			return def(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : UnionUnion")
}

type UofMaybe struct {
	Branch UofMaybeBranch
}

type UofMaybeBranch interface {
	isUofMaybeBranch()
}

func (*UofMaybe) MakeNewBranch(key string) (any, error) {
	switch key {
	case "a":
		return &_UofMaybe_A{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _UofMaybe_A struct {
	V types.Maybe[UnionUnion] `branch:"a"`
}

func (_UofMaybe_A) isUofMaybeBranch() {}

func Make_UofMaybe_a(v types.Maybe[UnionUnion]) UofMaybe {
	return UofMaybe{
		_UofMaybe_A{v},
	}
}

func (un UofMaybe) Cast_a() (types.Maybe[UnionUnion], bool) {
	br, ok := un.Branch.(_UofMaybe_A)
	return br.V, ok
}

func Handle_UofMaybe[T any](
	_in UofMaybe,
	a func(a types.Maybe[UnionUnion]) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _UofMaybe_A:
		if a != nil {
			return a(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : UofMaybe")
}

func HandleWithErr_UofMaybe[T any](
	_in UofMaybe,
	a func(a types.Maybe[UnionUnion]) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _UofMaybe_A:
		if a != nil {
			return a(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : UofMaybe")
}

type VectorErrorTest struct {
	Branch VectorErrorTestBranch
}

type VectorErrorTestBranch interface {
	isVectorErrorTestBranch()
}

func (*VectorErrorTest) MakeNewBranch(key string) (any, error) {
	switch key {
	case "a":
		return &_VectorErrorTest_A{}, nil
	case "b":
		return &_VectorErrorTest_B{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _VectorErrorTest_A struct {
	V []string `branch:"a"`
}
type _VectorErrorTest_B struct {
	V []Struct02Test `branch:"b"`
}

func (_VectorErrorTest_A) isVectorErrorTestBranch() {}
func (_VectorErrorTest_B) isVectorErrorTestBranch() {}

func Make_VectorErrorTest_a(v []string) VectorErrorTest {
	return VectorErrorTest{
		_VectorErrorTest_A{v},
	}
}

func Make_VectorErrorTest_b(v []Struct02Test) VectorErrorTest {
	return VectorErrorTest{
		_VectorErrorTest_B{v},
	}
}

func (un VectorErrorTest) Cast_a() ([]string, bool) {
	br, ok := un.Branch.(_VectorErrorTest_A)
	return br.V, ok
}

func (un VectorErrorTest) Cast_b() ([]Struct02Test, bool) {
	br, ok := un.Branch.(_VectorErrorTest_B)
	return br.V, ok
}

func Handle_VectorErrorTest[T any](
	_in VectorErrorTest,
	a func(a []string) T,
	b func(b []Struct02Test) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _VectorErrorTest_A:
		if a != nil {
			return a(_b.V)
		}
	case _VectorErrorTest_B:
		if b != nil {
			return b(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : VectorErrorTest")
}

func HandleWithErr_VectorErrorTest[T any](
	_in VectorErrorTest,
	a func(a []string) (T, error),
	b func(b []Struct02Test) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _VectorErrorTest_A:
		if a != nil {
			return a(_b.V)
		}
	case _VectorErrorTest_B:
		if b != nil {
			return b(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : VectorErrorTest")
}

type VectorOfTypeDiscriminationUnionTest struct {
	Branch VectorOfTypeDiscriminationUnionTestBranch
}

type VectorOfTypeDiscriminationUnionTestBranch interface {
	isVectorOfTypeDiscriminationUnionTestBranch()
}

func (*VectorOfTypeDiscriminationUnionTest) MakeNewBranch(key string) (any, error) {
	switch key {
	case "b":
		return &_VectorOfTypeDiscriminationUnionTest_B{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _VectorOfTypeDiscriminationUnionTest_B struct {
	V []Struct02Test `branch:"b"`
}

func (_VectorOfTypeDiscriminationUnionTest_B) isVectorOfTypeDiscriminationUnionTestBranch() {}

func Make_VectorOfTypeDiscriminationUnionTest_b(v []Struct02Test) VectorOfTypeDiscriminationUnionTest {
	return VectorOfTypeDiscriminationUnionTest{
		_VectorOfTypeDiscriminationUnionTest_B{v},
	}
}

func (un VectorOfTypeDiscriminationUnionTest) Cast_b() ([]Struct02Test, bool) {
	br, ok := un.Branch.(_VectorOfTypeDiscriminationUnionTest_B)
	return br.V, ok
}

func Handle_VectorOfTypeDiscriminationUnionTest[T any](
	_in VectorOfTypeDiscriminationUnionTest,
	b func(b []Struct02Test) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _VectorOfTypeDiscriminationUnionTest_B:
		if b != nil {
			return b(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : VectorOfTypeDiscriminationUnionTest")
}

func HandleWithErr_VectorOfTypeDiscriminationUnionTest[T any](
	_in VectorOfTypeDiscriminationUnionTest,
	b func(b []Struct02Test) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _VectorOfTypeDiscriminationUnionTest_B:
		if b != nil {
			return b(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : VectorOfTypeDiscriminationUnionTest")
}

type VectorTest struct {
	Branch VectorTestBranch
}

type VectorTestBranch interface {
	isVectorTestBranch()
}

func (*VectorTest) MakeNewBranch(key string) (any, error) {
	switch key {
	case "a":
		return &_VectorTest_A{}, nil
	case "b":
		return &_VectorTest_B{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _VectorTest_A struct {
	V []string `branch:"a"`
}
type _VectorTest_B struct {
	V []Struct02Test `branch:"b"`
}

func (_VectorTest_A) isVectorTestBranch() {}
func (_VectorTest_B) isVectorTestBranch() {}

func Make_VectorTest_a(v []string) VectorTest {
	return VectorTest{
		_VectorTest_A{v},
	}
}

func Make_VectorTest_b(v []Struct02Test) VectorTest {
	return VectorTest{
		_VectorTest_B{v},
	}
}

func (un VectorTest) Cast_a() ([]string, bool) {
	br, ok := un.Branch.(_VectorTest_A)
	return br.V, ok
}

func (un VectorTest) Cast_b() ([]Struct02Test, bool) {
	br, ok := un.Branch.(_VectorTest_B)
	return br.V, ok
}

func Handle_VectorTest[T any](
	_in VectorTest,
	a func(a []string) T,
	b func(b []Struct02Test) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _VectorTest_A:
		if a != nil {
			return a(_b.V)
		}
	case _VectorTest_B:
		if b != nil {
			return b(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : VectorTest")
}

func HandleWithErr_VectorTest[T any](
	_in VectorTest,
	a func(a []string) (T, error),
	b func(b []Struct02Test) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _VectorTest_A:
		if a != nil {
			return a(_b.V)
		}
	case _VectorTest_B:
		if b != nil {
			return b(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : VectorTest")
}

type VoidTest struct {
	Branch VoidTestBranch
}

type VoidTestBranch interface {
	isVoidTestBranch()
}

func (*VoidTest) MakeNewBranch(key string) (any, error) {
	switch key {
	case "abc":
		return &_VoidTest_Abc{}, nil
	case "def":
		return &_VoidTest_Def{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _VoidTest_Abc struct {
	V struct{} `branch:"abc"`
}
type _VoidTest_Def struct {
	V S1 `branch:"def"`
}

func (_VoidTest_Abc) isVoidTestBranch() {}
func (_VoidTest_Def) isVoidTestBranch() {}

func Make_VoidTest_abc() VoidTest {
	return VoidTest{
		_VoidTest_Abc{struct{}{}},
	}
}

func Make_VoidTest_def(v S1) VoidTest {
	return VoidTest{
		_VoidTest_Def{v},
	}
}

func (un VoidTest) Cast_abc() (struct{}, bool) {
	br, ok := un.Branch.(_VoidTest_Abc)
	return br.V, ok
}

func (un VoidTest) Cast_def() (S1, bool) {
	br, ok := un.Branch.(_VoidTest_Def)
	return br.V, ok
}

func Handle_VoidTest[T any](
	_in VoidTest,
	abc func(abc struct{}) T,
	def func(def S1) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _VoidTest_Abc:
		if abc != nil {
			return abc(_b.V)
		}
	case _VoidTest_Def:
		if def != nil {
			return def(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : VoidTest")
}

func HandleWithErr_VoidTest[T any](
	_in VoidTest,
	abc func(abc struct{}) (T, error),
	def func(def S1) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _VoidTest_Abc:
		if abc != nil {
			return abc(_b.V)
		}
	case _VoidTest_Def:
		if def != nil {
			return def(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : VoidTest")
}