// Code generated by goadlc v3 - DO NOT EDIT.
package generics

import (
	goadl "github.com/adl-lang/goadl_rt/v3"
	"github.com/adl-lang/goadl_rt/v3/customtypes"
	"github.com/adl-lang/goadl_rt/v3/sys/adlast"
	"github.com/adl-lang/goadl_rt/v3/sys/types"
)

func Texpr_Abc[A comparable, B any](a adlast.ATypeExpr[A], b adlast.ATypeExpr[B]) adlast.ATypeExpr[Abc[A, B]] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("exer03.generics", "Abc"),
		),
		[]adlast.TypeExpr{a.Value, b.Value},
	)
	return adlast.Make_ATypeExpr[Abc[A, B]](te)
}

func AST_Abc() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"Abc",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{
					"A",
					"B",
				},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"a",
						"a",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_typeParam(
								"A",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"b",
						"b",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Vector",
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_typeParam(
										"B",
									),
									[]adlast.TypeExpr{},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"kids",
						"kids",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Vector",
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"exer03.generics",
											"Abc",
										),
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"A",
											),
											[]adlast.TypeExpr{},
										),
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"B",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
							},
						),
						types.Make_Maybe_just[any](
							[]interface{}{},
						),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"grandkids_named",
						"grandkids_named",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"StringMap",
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"exer03.generics",
											"Abc",
										),
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"Int64",
											),
											[]adlast.TypeExpr{},
										),
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"String",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
							},
						),
						types.Make_Maybe_just[any](
							map[string]interface{}{"fav": map[string]interface{}{"a": 4321, "b": []interface{}{"aaa"}, "c": 99, "d": map[string]interface{}{"d": map[string]interface{}{"a": 22}}, "e": map[string]interface{}{"a": 1234}, "f": interface{}(nil), "grandkids_named": map[string]interface{}{}}},
						),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"c",
						"c",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Int64",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_just[any](
							66,
						),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"d",
						"d",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"exer03.generics",
									"Def",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_typeParam(
										"A",
									),
									[]adlast.TypeExpr{},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_typeParam(
										"B",
									),
									[]adlast.TypeExpr{},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"e",
						"e",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"exer03.generics",
									"Def",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"Int64",
									),
									[]adlast.TypeExpr{},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"String",
									),
									[]adlast.TypeExpr{},
								),
							},
						),
						types.Make_Maybe_just[any](
							map[string]interface{}{"a": 1234},
						),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"f",
						"f",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Nullable",
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"exer03.generics",
											"Abc",
										),
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"Int64",
											),
											[]adlast.TypeExpr{},
										),
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"String",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
							},
						),
						types.Make_Maybe_just[any](
							map[string]interface{}{"a": 4321, "b": []interface{}{"aaa"}, "c": 99, "d": map[string]interface{}{"d": map[string]interface{}{"a": 22}}, "e": map[string]interface{}{"a": 1234}, "f": interface{}(nil)},
						),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"comparable", "any"}},
	)
	return adlast.Make_ScopedDecl("exer03.generics", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("exer03.generics", "Abc"),
		AST_Abc(),
	)
}

func Texpr_Def[C comparable, D any](c adlast.ATypeExpr[C], d adlast.ATypeExpr[D]) adlast.ATypeExpr[Def[C, D]] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("exer03.generics", "Def"),
		),
		[]adlast.TypeExpr{c.Value, d.Value},
	)
	return adlast.Make_ATypeExpr[Def[C, D]](te)
}

func AST_Def() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"Def",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_union_(
			adlast.MakeAll_Union(
				[]adlast.Ident{
					"C",
					"D",
				},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"a",
						"a",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_typeParam(
								"C",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"comparable"}},
					),
					adlast.MakeAll_Field(
						"b",
						"b",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Vector",
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_typeParam(
										"D",
									),
									[]adlast.TypeExpr{},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"any"}},
					),
					adlast.MakeAll_Field(
						"c",
						"c",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Int64",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"d",
						"d",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"exer03.generics",
									"Def",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_typeParam(
										"C",
									),
									[]adlast.TypeExpr{},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_typeParam(
										"D",
									),
									[]adlast.TypeExpr{},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"comparable", "any"}},
					),
					adlast.MakeAll_Field(
						"e",
						"e",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"exer03.generics",
									"Abc",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"sys.types",
											"Either",
										),
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"String",
											),
											[]adlast.TypeExpr{},
										),
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"Int64",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"String",
									),
									[]adlast.TypeExpr{},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"f",
						"f",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Vector",
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"sys.types",
											"Either",
										),
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"String",
											),
											[]adlast.TypeExpr{},
										),
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_reference(
												adlast.MakeAll_ScopedName(
													"sys.types",
													"Either",
												),
											),
											[]adlast.TypeExpr{
												adlast.MakeAll_TypeExpr(
													adlast.Make_TypeRef_primitive(
														"String",
													),
													[]adlast.TypeExpr{},
												),
												adlast.MakeAll_TypeExpr(
													adlast.Make_TypeRef_primitive(
														"Int64",
													),
													[]adlast.TypeExpr{},
												),
											},
										),
									},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"f2",
						"f2",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Vector",
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"sys.types",
											"Either",
										),
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_reference(
												adlast.MakeAll_ScopedName(
													"sys.types",
													"Either",
												),
											),
											[]adlast.TypeExpr{
												adlast.MakeAll_TypeExpr(
													adlast.Make_TypeRef_primitive(
														"String",
													),
													[]adlast.TypeExpr{},
												),
												adlast.MakeAll_TypeExpr(
													adlast.Make_TypeRef_primitive(
														"Int64",
													),
													[]adlast.TypeExpr{},
												),
											},
										),
									},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"comparable"}},
					),
					adlast.MakeAll_Field(
						"g",
						"g",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"sys.types",
									"Either",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"Vector",
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"String",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"sys.types",
											"Either",
										),
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"String",
											),
											[]adlast.TypeExpr{},
										),
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"Int64",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"g1",
						"g1",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"sys.types",
									"Either",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"Vector",
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"sys.types",
											"Either",
										),
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"comparable"}},
					),
					adlast.MakeAll_Field(
						"g2",
						"g2",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"sys.types",
									"Either",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"Nullable",
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"sys.types",
											"Either",
										),
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"comparable"}},
					),
					adlast.MakeAll_Field(
						"g3",
						"g3",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"sys.types",
									"Either",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"StringMap",
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"sys.types",
											"Either",
										),
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"comparable"}},
					),
					adlast.MakeAll_Field(
						"g4",
						"g4",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"sys.types",
									"Either",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"Vector",
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"String",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"sys.types",
											"Either",
										),
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"comparable"}},
					),
					adlast.MakeAll_Field(
						"g5",
						"g5",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"sys.types",
									"Either",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"sys.types",
											"Either",
										),
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"Vector",
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"String",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"comparable"}},
					),
					adlast.MakeAll_Field(
						"g6",
						"g6",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"sys.types",
									"Either",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"sys.types",
											"Either",
										),
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"String",
											),
											[]adlast.TypeExpr{},
										),
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"String",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"Vector",
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"comparable"}},
					),
					adlast.MakeAll_Field(
						"g7",
						"g7",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"sys.types",
									"Either",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"sys.types",
											"Either",
										),
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"D",
											),
											[]adlast.TypeExpr{},
										),
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"String",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"Vector",
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"any", "comparable"}},
					),
					adlast.MakeAll_Field(
						"g8",
						"g8",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"sys.types",
									"Either",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"Vector",
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"Vector",
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"comparable"}},
					),
					adlast.MakeAll_Field(
						"g9",
						"g9",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"sys.types",
									"Either",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"Vector",
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"Vector",
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"D",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"comparable", "any"}},
					),
					adlast.MakeAll_Field(
						"g10",
						"g10",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"sys.types",
									"Either",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"Vector",
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"D",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"Vector",
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"any", "comparable"}},
					),
					adlast.MakeAll_Field(
						"h",
						"h",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Vector",
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_typeParam(
										"C",
									),
									[]adlast.TypeExpr{},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"comparable"}},
					),
					adlast.MakeAll_Field(
						"i",
						"i",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Vector",
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"Vector",
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"C",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"comparable"}},
					),
					adlast.MakeAll_Field(
						"m1",
						"m1",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"sys.types",
									"Map",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"String",
									),
									[]adlast.TypeExpr{},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"Int64",
									),
									[]adlast.TypeExpr{},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"m2",
						"m2",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"sys.types",
									"Map",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_typeParam(
										"C",
									),
									[]adlast.TypeExpr{},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"Int64",
									),
									[]adlast.TypeExpr{},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"comparable"}},
					),
					adlast.MakeAll_Field(
						"m3",
						"m3",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"sys.types",
									"Map",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"String",
									),
									[]adlast.TypeExpr{},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_typeParam(
										"C",
									),
									[]adlast.TypeExpr{},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"comparable"}},
					),
					adlast.MakeAll_Field(
						"m4",
						"m4",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"sys.types",
									"Map",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_typeParam(
										"C",
									),
									[]adlast.TypeExpr{},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_typeParam(
										"D",
									),
									[]adlast.TypeExpr{},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"comparable", "any"}},
					),
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"comparable", "any"}},
	)
	return adlast.Make_ScopedDecl("exer03.generics", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("exer03.generics", "Def"),
		AST_Def(),
	)
}

func Texpr_Zyx[Z comparable, Y any](z adlast.ATypeExpr[Z], y adlast.ATypeExpr[Y]) adlast.ATypeExpr[Zyx[Z, Y]] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("exer03.generics", "Zyx"),
		),
		[]adlast.TypeExpr{z.Value, y.Value},
	)
	return adlast.Make_ATypeExpr[Zyx[Z, Y]](te)
}

func AST_Zyx() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"Zyx",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{
					"Z",
					"Y",
				},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"a",
						"a",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"exer03.generics",
									"Abc",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_typeParam(
										"Z",
									),
									[]adlast.TypeExpr{},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_typeParam(
										"Y",
									),
									[]adlast.TypeExpr{},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"e",
						"e",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"exer03.generics",
									"Abc",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"sys.types",
											"Either",
										),
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"String",
											),
											[]adlast.TypeExpr{},
										),
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"Int64",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"String",
									),
									[]adlast.TypeExpr{},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"f",
						"f",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Vector",
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"sys.types",
											"Either",
										),
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"String",
											),
											[]adlast.TypeExpr{},
										),
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_reference(
												adlast.MakeAll_ScopedName(
													"sys.types",
													"Either",
												),
											),
											[]adlast.TypeExpr{
												adlast.MakeAll_TypeExpr(
													adlast.Make_TypeRef_primitive(
														"String",
													),
													[]adlast.TypeExpr{},
												),
												adlast.MakeAll_TypeExpr(
													adlast.Make_TypeRef_primitive(
														"Int64",
													),
													[]adlast.TypeExpr{},
												),
											},
										),
									},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"f1",
						"f1",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Vector",
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"sys.types",
											"Either",
										),
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"Z",
											),
											[]adlast.TypeExpr{},
										),
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_reference(
												adlast.MakeAll_ScopedName(
													"sys.types",
													"Either",
												),
											),
											[]adlast.TypeExpr{
												adlast.MakeAll_TypeExpr(
													adlast.Make_TypeRef_typeParam(
														"Z",
													),
													[]adlast.TypeExpr{},
												),
												adlast.MakeAll_TypeExpr(
													adlast.Make_TypeRef_primitive(
														"Int64",
													),
													[]adlast.TypeExpr{},
												),
											},
										),
									},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"f2",
						"f2",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Vector",
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"sys.types",
											"Either",
										),
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_typeParam(
												"Z",
											),
											[]adlast.TypeExpr{},
										),
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_reference(
												adlast.MakeAll_ScopedName(
													"sys.types",
													"Either",
												),
											),
											[]adlast.TypeExpr{
												adlast.MakeAll_TypeExpr(
													adlast.Make_TypeRef_typeParam(
														"Z",
													),
													[]adlast.TypeExpr{},
												),
												adlast.MakeAll_TypeExpr(
													adlast.Make_TypeRef_typeParam(
														"Z",
													),
													[]adlast.TypeExpr{},
												),
											},
										),
									},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"f3",
						"f3",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Vector",
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"sys.types",
											"Either",
										),
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"String",
											),
											[]adlast.TypeExpr{},
										),
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_reference(
												adlast.MakeAll_ScopedName(
													"sys.types",
													"Either",
												),
											),
											[]adlast.TypeExpr{
												adlast.MakeAll_TypeExpr(
													adlast.Make_TypeRef_typeParam(
														"Z",
													),
													[]adlast.TypeExpr{},
												),
												adlast.MakeAll_TypeExpr(
													adlast.Make_TypeRef_typeParam(
														"Y",
													),
													[]adlast.TypeExpr{},
												),
											},
										),
									},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"g",
						"g",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"sys.types",
									"Either",
								),
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"Vector",
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"String",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"sys.types",
											"Either",
										),
									),
									[]adlast.TypeExpr{
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"String",
											),
											[]adlast.TypeExpr{},
										),
										adlast.MakeAll_TypeExpr(
											adlast.Make_TypeRef_primitive(
												"Int64",
											),
											[]adlast.TypeExpr{},
										),
									},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("adlc.config.go_", "TypeParamConstraintList"): []interface{}{"comparable", "any"}},
	)
	return adlast.Make_ScopedDecl("exer03.generics", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("exer03.generics", "Zyx"),
		AST_Zyx(),
	)
}
