// Code generated by goadlc v3 - DO NOT EDIT.
package simple_union

import (
	goadl "github.com/adl-lang/goadl_rt/v3"
	"github.com/adl-lang/goadl_rt/v3/customtypes"
	"github.com/adl-lang/goadl_rt/v3/sys/adlast"
	"github.com/adl-lang/goadl_rt/v3/sys/types"
)

func Texpr_UnionOfPrimitives() adlast.ATypeExpr[UnionOfPrimitives] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("exer01.simple_union", "UnionOfPrimitives"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[UnionOfPrimitives](te)
}

func AST_UnionOfPrimitives() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"UnionOfPrimitives",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_union_(
			adlast.MakeAll_Union(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"A",
						"A",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Int32",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"B",
						"B",
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
						"c",
						"c",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Bool",
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
							adlast.Make_TypeRef_primitive(
								"Double",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"e",
						"e",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"String",
							),
							[]adlast.TypeExpr{},
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
						"g",
						"g",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Void",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{},
	)
	return adlast.Make_ScopedDecl("exer01.simple_union", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("exer01.simple_union", "UnionOfPrimitives"),
		AST_UnionOfPrimitives(),
	)
}

func Texpr_UnionOfVoids() adlast.ATypeExpr[UnionOfVoids] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("exer01.simple_union", "UnionOfVoids"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[UnionOfVoids](te)
}

func AST_UnionOfVoids() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"UnionOfVoids",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_union_(
			adlast.MakeAll_Union(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"A",
						"A",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Void",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"B",
						"B",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Void",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"c",
						"c",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Void",
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
							adlast.Make_TypeRef_primitive(
								"Void",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"e",
						"e",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Void",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"f",
						"f",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Void",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"g",
						"g",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Void",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{},
	)
	return adlast.Make_ScopedDecl("exer01.simple_union", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("exer01.simple_union", "UnionOfVoids"),
		AST_UnionOfVoids(),
	)
}
