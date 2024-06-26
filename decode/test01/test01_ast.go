// Code generated by goadlc v3 - DO NOT EDIT.
package test01

import (
	goadl "github.com/adl-lang/goadl_rt/v3"
	"github.com/adl-lang/goadl_rt/v3/customtypes"
	"github.com/adl-lang/goadl_rt/v3/sys/adlast"
	"github.com/adl-lang/goadl_rt/v3/sys/types"
)

func Texpr_A() adlast.ATypeExpr[A] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "A"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[A](te)
}

func AST_A() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"A",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_newtype_(
			adlast.MakeAll_NewType(
				[]adlast.Ident{},
				adlast.MakeAll_TypeExpr(
					adlast.Make_TypeRef_primitive(
						"Int64",
					),
					[]adlast.TypeExpr{},
				),
				types.Make_Maybe_nothing[any](),
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{},
	)
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "A"),
		AST_A(),
	)
}

func Texpr_B() adlast.ATypeExpr[B] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "B"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[B](te)
}

func AST_B() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"B",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"a",
						"a",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Int64",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("decode.test01", "A"): 123},
	)
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "B"),
		AST_B(),
	)
}

func Texpr_Bool() adlast.ATypeExpr[Bool] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "Bool"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[Bool](te)
}

func AST_Bool() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"Bool",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"a",
						"a",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Bool",
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
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "Bool"),
		AST_Bool(),
	)
}

func Texpr_C() adlast.ATypeExpr[C] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "C"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[C](te)
}

func AST_C() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"C",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"b",
						"b",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"decode.test01",
									"B",
								),
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_just[any](
							map[string]interface{}{"a": 1234},
						),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"c",
						"c",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"decode.test01",
									"B",
								),
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
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "C"),
		AST_C(),
	)
}

func Texpr_D() adlast.ATypeExpr[D] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "D"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[D](te)
}

func AST_D() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"D",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_union_(
			adlast.MakeAll_Union(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"a",
						"a",
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
						"b",
						"b",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"decode.test01",
									"B",
								),
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
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "D"),
		AST_D(),
	)
}

func Texpr_E() adlast.ATypeExpr[E] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "E"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[E](te)
}

func AST_E() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"E",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"d",
						"d",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"decode.test01",
									"D",
								),
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
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "E"),
		AST_E(),
	)
}

func Texpr_F() adlast.ATypeExpr[F] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "F"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[F](te)
}

func AST_F() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"F",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"a",
						"a",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Int8",
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
								"Int16",
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
								"Int32",
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
								"Int64",
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
								"Word8",
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
								"Word16",
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
								"Word32",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"h",
						"h",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Word64",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"i",
						"i",
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
						"j",
						"j",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Float",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"k",
						"k",
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
						"l",
						"l",
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
						"n",
						"n",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Json",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"o",
						"o",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Vector",
							),
							[]adlast.TypeExpr{
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
						"p",
						"p",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"StringMap",
							),
							[]adlast.TypeExpr{
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
						"q",
						"q",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Nullable",
							),
							[]adlast.TypeExpr{
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
						"r",
						"r",
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
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "F"),
		AST_F(),
	)
}

func Texpr_G() adlast.ATypeExpr[G] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "G"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[G](te)
}

func AST_G() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"G",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"a",
						"a",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Vector",
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_reference(
										adlast.MakeAll_ScopedName(
											"decode.test01",
											"G",
										),
									),
									[]adlast.TypeExpr{},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{},
	)
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "G"),
		AST_G(),
	)
}

func Texpr_GenericF[AA any](aa adlast.ATypeExpr[AA]) adlast.ATypeExpr[GenericF[AA]] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "GenericF"),
		),
		[]adlast.TypeExpr{aa.Value},
	)
	return adlast.Make_ATypeExpr[GenericF[AA]](te)
}

func AST_GenericF() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"GenericF",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{
					"AA",
				},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"a",
						"a",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_typeParam(
								"AA",
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
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "GenericF"),
		AST_GenericF(),
	)
}

func Texpr_HasDefault() adlast.ATypeExpr[HasDefault] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "HasDefault"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[HasDefault](te)
}

func AST_HasDefault() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"HasDefault",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"a",
						"a",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"String",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_just[any](
							"1234567890",
						),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{},
	)
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "HasDefault"),
		AST_HasDefault(),
	)
}

func Texpr_Int() adlast.ATypeExpr[Int] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "Int"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[Int](te)
}

func AST_Int() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"Int",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"a",
						"a",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Int8",
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
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "Int"),
		AST_Int(),
	)
}

func Texpr_Json() adlast.ATypeExpr[Json] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "Json"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[Json](te)
}

func AST_Json() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"Json",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"a",
						"a",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Json",
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
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "Json"),
		AST_Json(),
	)
}

func Texpr_MapTest() adlast.ATypeExpr[MapTest] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "MapTest"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[MapTest](te)
}

func AST_MapTest() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"MapTest",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"my_set",
						"my_set",
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
						types.Make_Maybe_just[any](
							[]interface{}{map[string]interface{}{"k": "a", "v": 1}},
						),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{},
	)
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "MapTest"),
		AST_MapTest(),
	)
}

func Texpr_MyV[A any](a adlast.ATypeExpr[A]) adlast.ATypeExpr[MyV[A]] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "MyV"),
		),
		[]adlast.TypeExpr{a.Value},
	)
	return adlast.Make_ATypeExpr[MyV[A]](te)
}

func AST_MyV() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"MyV",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_newtype_(
			adlast.MakeAll_NewType(
				[]adlast.Ident{
					"A",
				},
				adlast.MakeAll_TypeExpr(
					adlast.Make_TypeRef_primitive(
						"Vector",
					),
					[]adlast.TypeExpr{
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"decode.test01",
									"A",
								),
							),
							[]adlast.TypeExpr{},
						),
					},
				),
				types.Make_Maybe_nothing[any](),
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{},
	)
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "MyV"),
		AST_MyV(),
	)
}

func Texpr_NoDefault() adlast.ATypeExpr[NoDefault] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "NoDefault"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[NoDefault](te)
}

func AST_NoDefault() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"NoDefault",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"a",
						"a",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"String",
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
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "NoDefault"),
		AST_NoDefault(),
	)
}

func Texpr_NullableString() adlast.ATypeExpr[NullableString] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "NullableString"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[NullableString](te)
}

func AST_NullableString() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"NullableString",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"a",
						"a",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Nullable",
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
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{},
	)
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "NullableString"),
		AST_NullableString(),
	)
}

func Texpr_SetTest() adlast.ATypeExpr[SetTest] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "SetTest"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[SetTest](te)
}

func AST_SetTest() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"SetTest",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"my_set",
						"my_set",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"sys.types",
									"Set",
								),
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
						types.Make_Maybe_just[any](
							[]interface{}{"a", "b", "z"},
						),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{},
	)
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "SetTest"),
		AST_SetTest(),
	)
}

func Texpr_StringMapString() adlast.ATypeExpr[StringMapString] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "StringMapString"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[StringMapString](te)
}

func AST_StringMapString() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"StringMapString",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"a",
						"a",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"StringMap",
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
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{},
	)
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "StringMapString"),
		AST_StringMapString(),
	)
}

func Texpr_Uint() adlast.ATypeExpr[Uint] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "Uint"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[Uint](te)
}

func AST_Uint() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"Uint",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"a",
						"a",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Word16",
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
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "Uint"),
		AST_Uint(),
	)
}

func Texpr_Unit() adlast.ATypeExpr[Unit] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "Unit"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[Unit](te)
}

func AST_Unit() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"Unit",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{},
	)
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "Unit"),
		AST_Unit(),
	)
}

func Texpr_VectorString() adlast.ATypeExpr[VectorString] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("decode.test01", "VectorString"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[VectorString](te)
}

func AST_VectorString() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"VectorString",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"a",
						"a",
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
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{},
	)
	return adlast.Make_ScopedDecl("decode.test01", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("decode.test01", "VectorString"),
		AST_VectorString(),
	)
}
