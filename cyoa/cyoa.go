// Code generated by goadlc v3 - DO NOT EDIT.
package cyoa

import (
	"fmt"
)

type ChoiceThree[A any, B any, C any] struct {
	Branch ChoiceThreeBranch[A, B, C]
}

type ChoiceThreeBranch[A any, B any, C any] interface {
	isChoiceThreeBranch()
}

func (*ChoiceThree[A, B, C]) MakeNewBranch(key string) (any, error) {
	switch key {
	case "one":
		return &_ChoiceThree_One[A]{}, nil
	case "two":
		return &_ChoiceThree_Two[B]{}, nil
	case "tri":
		return &_ChoiceThree_Tri[C]{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _ChoiceThree_One[A any] struct {
	V A `branch:"one"`
}
type _ChoiceThree_Two[B any] struct {
	V B `branch:"two"`
}
type _ChoiceThree_Tri[C any] struct {
	V C `branch:"tri"`
}

func (_ChoiceThree_One[A]) isChoiceThreeBranch() {}
func (_ChoiceThree_Two[B]) isChoiceThreeBranch() {}
func (_ChoiceThree_Tri[C]) isChoiceThreeBranch() {}

func Make_ChoiceThree_one[A any, B any, C any](v A) ChoiceThree[A, B, C] {
	return ChoiceThree[A, B, C]{
		_ChoiceThree_One[A]{v},
	}
}

func Make_ChoiceThree_two[A any, B any, C any](v B) ChoiceThree[A, B, C] {
	return ChoiceThree[A, B, C]{
		_ChoiceThree_Two[B]{v},
	}
}

func Make_ChoiceThree_tri[A any, B any, C any](v C) ChoiceThree[A, B, C] {
	return ChoiceThree[A, B, C]{
		_ChoiceThree_Tri[C]{v},
	}
}

func (un ChoiceThree[A, B, C]) Cast_one() (A, bool) {
	br, ok := un.Branch.(_ChoiceThree_One[A])
	return br.V, ok
}

func (un ChoiceThree[A, B, C]) Cast_two() (B, bool) {
	br, ok := un.Branch.(_ChoiceThree_Two[B])
	return br.V, ok
}

func (un ChoiceThree[A, B, C]) Cast_tri() (C, bool) {
	br, ok := un.Branch.(_ChoiceThree_Tri[C])
	return br.V, ok
}

func Handle_ChoiceThree[A any, B any, C any, T any](
	_in ChoiceThree[A, B, C],
	one func(one A) T,
	two func(two B) T,
	tri func(tri C) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _ChoiceThree_One[A]:
		if one != nil {
			return one(_b.V)
		}
	case _ChoiceThree_Two[B]:
		if two != nil {
			return two(_b.V)
		}
	case _ChoiceThree_Tri[C]:
		if tri != nil {
			return tri(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : ChoiceThree")
}

func HandleWithErr_ChoiceThree[A any, B any, C any, T any](
	_in ChoiceThree[A, B, C],
	one func(one A) (T, error),
	two func(two B) (T, error),
	tri func(tri C) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _ChoiceThree_One[A]:
		if one != nil {
			return one(_b.V)
		}
	case _ChoiceThree_Two[B]:
		if two != nil {
			return two(_b.V)
		}
	case _ChoiceThree_Tri[C]:
		if tri != nil {
			return tri(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : ChoiceThree")
}

type ChoiceTwo[A any, B any] struct {
	Branch ChoiceTwoBranch[A, B]
}

type ChoiceTwoBranch[A any, B any] interface {
	isChoiceTwoBranch()
}

func (*ChoiceTwo[A, B]) MakeNewBranch(key string) (any, error) {
	switch key {
	case "one":
		return &_ChoiceTwo_One[A]{}, nil
	case "two":
		return &_ChoiceTwo_Two[B]{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _ChoiceTwo_One[A any] struct {
	V A `branch:"one"`
}
type _ChoiceTwo_Two[B any] struct {
	V B `branch:"two"`
}

func (_ChoiceTwo_One[A]) isChoiceTwoBranch() {}
func (_ChoiceTwo_Two[B]) isChoiceTwoBranch() {}

func Make_ChoiceTwo_one[A any, B any](v A) ChoiceTwo[A, B] {
	return ChoiceTwo[A, B]{
		_ChoiceTwo_One[A]{v},
	}
}

func Make_ChoiceTwo_two[A any, B any](v B) ChoiceTwo[A, B] {
	return ChoiceTwo[A, B]{
		_ChoiceTwo_Two[B]{v},
	}
}

func (un ChoiceTwo[A, B]) Cast_one() (A, bool) {
	br, ok := un.Branch.(_ChoiceTwo_One[A])
	return br.V, ok
}

func (un ChoiceTwo[A, B]) Cast_two() (B, bool) {
	br, ok := un.Branch.(_ChoiceTwo_Two[B])
	return br.V, ok
}

func Handle_ChoiceTwo[A any, B any, T any](
	_in ChoiceTwo[A, B],
	one func(one A) T,
	two func(two B) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _ChoiceTwo_One[A]:
		if one != nil {
			return one(_b.V)
		}
	case _ChoiceTwo_Two[B]:
		if two != nil {
			return two(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : ChoiceTwo")
}

func HandleWithErr_ChoiceTwo[A any, B any, T any](
	_in ChoiceTwo[A, B],
	one func(one A) (T, error),
	two func(two B) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _ChoiceTwo_One[A]:
		if one != nil {
			return one(_b.V)
		}
	case _ChoiceTwo_Two[B]:
		if two != nil {
			return two(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : ChoiceTwo")
}

type Page[A any, B any, C any] struct {
	Branch PageBranch[A, B, C]
}

type PageBranch[A any, B any, C any] interface {
	isPageBranch()
}

func (*Page[A, B, C]) MakeNewBranch(key string) (any, error) {
	switch key {
	case "next_page":
		return &_Page_Next_page{}, nil
	case "two":
		return &_Page_Two[A, B]{}, nil
	case "tri":
		return &_Page_Tri[A, B, C]{}, nil
	}
	return nil, fmt.Errorf("unknown branch is : %s", key)
}

type _Page_Next_page struct {
	V struct{} `branch:"next_page"`
}
type _Page_Two[A any, B any] struct {
	V ChoiceTwo[A, B] `branch:"two"`
}
type _Page_Tri[A any, B any, C any] struct {
	V ChoiceThree[A, B, C] `branch:"tri"`
}

func (_Page_Next_page) isPageBranch()    {}
func (_Page_Two[A, B]) isPageBranch()    {}
func (_Page_Tri[A, B, C]) isPageBranch() {}

func Make_Page_next_page[A any, B any, C any]() Page[A, B, C] {
	return Page[A, B, C]{
		_Page_Next_page{struct{}{}},
	}
}

func Make_Page_two[A any, B any, C any](v ChoiceTwo[A, B]) Page[A, B, C] {
	return Page[A, B, C]{
		_Page_Two[A, B]{v},
	}
}

func Make_Page_tri[A any, B any, C any](v ChoiceThree[A, B, C]) Page[A, B, C] {
	return Page[A, B, C]{
		_Page_Tri[A, B, C]{v},
	}
}

func (un Page[A, B, C]) Cast_next_page() (struct{}, bool) {
	br, ok := un.Branch.(_Page_Next_page)
	return br.V, ok
}

func (un Page[A, B, C]) Cast_two() (ChoiceTwo[A, B], bool) {
	br, ok := un.Branch.(_Page_Two[A, B])
	return br.V, ok
}

func (un Page[A, B, C]) Cast_tri() (ChoiceThree[A, B, C], bool) {
	br, ok := un.Branch.(_Page_Tri[A, B, C])
	return br.V, ok
}

func Handle_Page[A any, B any, C any, T any](
	_in Page[A, B, C],
	next_page func(next_page struct{}) T,
	two func(two ChoiceTwo[A, B]) T,
	tri func(tri ChoiceThree[A, B, C]) T,
	_default func() T,
) T {
	switch _b := _in.Branch.(type) {
	case _Page_Next_page:
		if next_page != nil {
			return next_page(_b.V)
		}
	case _Page_Two[A, B]:
		if two != nil {
			return two(_b.V)
		}
	case _Page_Tri[A, B, C]:
		if tri != nil {
			return tri(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : Page")
}

func HandleWithErr_Page[A any, B any, C any, T any](
	_in Page[A, B, C],
	next_page func(next_page struct{}) (T, error),
	two func(two ChoiceTwo[A, B]) (T, error),
	tri func(tri ChoiceThree[A, B, C]) (T, error),
	_default func() (T, error),
) (T, error) {
	switch _b := _in.Branch.(type) {
	case _Page_Next_page:
		if next_page != nil {
			return next_page(_b.V)
		}
	case _Page_Two[A, B]:
		if two != nil {
			return two(_b.V)
		}
	case _Page_Tri[A, B, C]:
		if tri != nil {
			return tri(_b.V)
		}
	}
	if _default != nil {
		return _default()
	}
	panic("unhandled branch in : Page")
}
