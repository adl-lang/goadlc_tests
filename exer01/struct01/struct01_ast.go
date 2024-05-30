// Code generated by goadlc v3 - DO NOT EDIT.
package struct01

import (
	goadl "github.com/adl-lang/goadl_rt/v3"
	"github.com/adl-lang/goadl_rt/v3/customtypes"
	"github.com/adl-lang/goadl_rt/v3/sys/adlast"
	"github.com/adl-lang/goadl_rt/v3/sys/types"
)

func Texpr_B() goadl.ATypeExpr[B] {
	return goadl.ATypeExpr[B]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "exer01.struct01",
						Name:       "B",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_B() adlast.ScopedDecl {
	decl := adlast.Decl{
		Name: "B",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: adlast.DeclType{
			Branch: adlast.DeclType_Struct_{
				V: adlast.Struct{
					TypeParams: []adlast.Ident{},
					Fields: []adlast.Field{
						adlast.Field{
							Name:           "a",
							SerializedName: "a",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Primitive{
										V: "String"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: customtypes.MapMap[adlast.ScopedName, any]{},
						},
					},
				}},
		},
		Annotations: customtypes.MapMap[adlast.ScopedName, any]{},
	}
	return adlast.ScopedDecl{
		ModuleName: "exer01.struct01",
		Decl:       decl,
	}
}

func init() {
	goadl.RESOLVER.Register(
		adlast.ScopedName{ModuleName: "exer01.struct01", Name: "B"},
		AST_B(),
	)
}

func Texpr_Struct01() goadl.ATypeExpr[Struct01] {
	return goadl.ATypeExpr[Struct01]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "exer01.struct01",
						Name:       "Struct01",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_Struct01() adlast.ScopedDecl {
	decl := adlast.Decl{
		Name: "Struct01",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: adlast.DeclType{
			Branch: adlast.DeclType_Struct_{
				V: adlast.Struct{
					TypeParams: []adlast.Ident{},
					Fields: []adlast.Field{
						adlast.Field{
							Name:           "A",
							SerializedName: "A",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Primitive{
										V: "Void"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: customtypes.MapMap[adlast.ScopedName, any]{},
						},
						adlast.Field{
							Name:           "B",
							SerializedName: "B",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Primitive{
										V: "Int64"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: customtypes.MapMap[adlast.ScopedName, any]{},
						},
						adlast.Field{
							Name:           "C",
							SerializedName: "C",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Primitive{
										V: "String"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: customtypes.MapMap[adlast.ScopedName, any]{},
						},
						adlast.Field{
							Name:           "d",
							SerializedName: "d",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Primitive{
										V: "Json"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: customtypes.MapMap[adlast.ScopedName, any]{},
						},
						adlast.Field{
							Name:           "e",
							SerializedName: "e",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Primitive{
										V: "Vector"},
								},
								Parameters: []adlast.TypeExpr{
									adlast.TypeExpr{
										TypeRef: adlast.TypeRef{
											Branch: adlast.TypeRef_Primitive{
												V: "String"},
										},
										Parameters: []adlast.TypeExpr{},
									},
								},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: customtypes.MapMap[adlast.ScopedName, any]{},
						},
						adlast.Field{
							Name:           "f",
							SerializedName: "f",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Primitive{
										V: "StringMap"},
								},
								Parameters: []adlast.TypeExpr{
									adlast.TypeExpr{
										TypeRef: adlast.TypeRef{
											Branch: adlast.TypeRef_Primitive{
												V: "Vector"},
										},
										Parameters: []adlast.TypeExpr{
											adlast.TypeExpr{
												TypeRef: adlast.TypeRef{
													Branch: adlast.TypeRef_Primitive{
														V: "String"},
												},
												Parameters: []adlast.TypeExpr{},
											},
										},
									},
								},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: customtypes.MapMap[adlast.ScopedName, any]{},
						},
						adlast.Field{
							Name:           "g",
							SerializedName: "g",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Primitive{
										V: "StringMap"},
								},
								Parameters: []adlast.TypeExpr{
									adlast.TypeExpr{
										TypeRef: adlast.TypeRef{
											Branch: adlast.TypeRef_Primitive{
												V: "StringMap"},
										},
										Parameters: []adlast.TypeExpr{
											adlast.TypeExpr{
												TypeRef: adlast.TypeRef{
													Branch: adlast.TypeRef_Primitive{
														V: "String"},
												},
												Parameters: []adlast.TypeExpr{},
											},
										},
									},
								},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: customtypes.MapMap[adlast.ScopedName, any]{},
						},
						adlast.Field{
							Name:           "h",
							SerializedName: "h",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Primitive{
										V: "StringMap"},
								},
								Parameters: []adlast.TypeExpr{
									adlast.TypeExpr{
										TypeRef: adlast.TypeRef{
											Branch: adlast.TypeRef_Primitive{
												V: "StringMap"},
										},
										Parameters: []adlast.TypeExpr{
											adlast.TypeExpr{
												TypeRef: adlast.TypeRef{
													Branch: adlast.TypeRef_Primitive{
														V: "Nullable"},
												},
												Parameters: []adlast.TypeExpr{
													adlast.TypeExpr{
														TypeRef: adlast.TypeRef{
															Branch: adlast.TypeRef_Primitive{
																V: "String"},
														},
														Parameters: []adlast.TypeExpr{},
													},
												},
											},
										},
									},
								},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: customtypes.MapMap[adlast.ScopedName, any]{},
						},
						adlast.Field{
							Name:           "i",
							SerializedName: "i",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Primitive{
										V: "Nullable"},
								},
								Parameters: []adlast.TypeExpr{
									adlast.TypeExpr{
										TypeRef: adlast.TypeRef{
											Branch: adlast.TypeRef_Primitive{
												V: "String"},
										},
										Parameters: []adlast.TypeExpr{},
									},
								},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: customtypes.MapMap[adlast.ScopedName, any]{},
						},
						adlast.Field{
							Name:           "j",
							SerializedName: "j",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Reference{
										V: adlast.ScopedName{
											ModuleName: "exer01.struct01",
											Name:       "B",
										}},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: customtypes.MapMap[adlast.ScopedName, any]{},
						},
					},
				}},
		},
		Annotations: customtypes.MapMap[adlast.ScopedName, any]{},
	}
	return adlast.ScopedDecl{
		ModuleName: "exer01.struct01",
		Decl:       decl,
	}
}

func init() {
	goadl.RESOLVER.Register(
		adlast.ScopedName{ModuleName: "exer01.struct01", Name: "Struct01"},
		AST_Struct01(),
	)
}
