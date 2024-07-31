package out_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	T31 "adl_testing/test31/test31"

	goadl "github.com/adl-lang/goadl_rt/v3"
	"github.com/adl-lang/goadl_rt/v3/sys/adlast"
	"github.com/adl-lang/goadl_rt/v3/sys/annotations"
	"github.com/adl-lang/goadl_rt/v3/sys/types"
)

func TestTypeDiscrimination(t *testing.T) {
	type TypeDiscTest struct {
		name    string
		dstZero any
		texpr   adlast.TypeExpr
		json    string
		want    any
		wantErr *string
	}
	tests := []TypeDiscTest{
		{
			name:    "MyStructV2",
			texpr:   T31.Texpr_MyStructV2().Value,
			json:    `{"quantity": 42}`,
			dstZero: &T31.MyStructV2{},
			want:    T31.Make_MyStructV2(T31.Make_Measure_count(42)),
		},
		{
			name:    "count into measure",
			texpr:   T31.Texpr_Measure().Value,
			json:    `42`,
			dstZero: &T31.Measure{},
			want:    T31.Make_Measure_count(42),
		},
		{
			name:    "s1 into structtest",
			texpr:   T31.Texpr_StructTest().Value,
			json:    `{"quant": 2, "value": 3}`,
			dstZero: &T31.StructTest{},
			want:    T31.Make_StructTest_abc(T31.Make_S1(2, 3.0)),
		},
		{
			name:    "void error test",
			texpr:   T31.Texpr_VoidTest().Value,
			json:    `null`,
			want:    nil,
			wantErr: goadl.Addr(`cannot use Json or Void as a type discriminator`),
		},
		{
			name:    "UnionOfLiftedUnion type mismatch error test",
			texpr:   T31.Texpr_UnionOfLiftedUnion().Value,
			dstZero: &T31.UnionOfLiftedUnion{},
			json:    `null`,
			want:    nil,
			wantErr: goadl.Addr(`primitive type mismatch. expected "Nullable" received`),
		},
		{
			name:    "UnionOfLiftedUnion error test",
			texpr:   T31.Texpr_UnionOfLiftedUnion().Value,
			dstZero: &T31.UnionOfLiftedUnion{},
			json:    `{"def": {"quant": 2, "value": 3}}`,
			want:    nil,
			wantErr: goadl.Addr(`ambiguous matching type discriminators `),
		},
		{
			name:    "UnionUnion a",
			texpr:   T31.Texpr_UnionUnion().Value,
			dstZero: &T31.UnionUnion{},
			json:    `{"a": "is an a"}`,
			want:    T31.Make_UnionUnion_abc(T31.Make_U1_a("is an a")),
		},
		{
			name:    "UnionUnion b",
			texpr:   T31.Texpr_UnionUnion().Value,
			dstZero: &T31.UnionUnion{},
			json:    `{"b": 99}`,
			want:    T31.Make_UnionUnion_abc(T31.Make_U1_b(99)),
		},
		{
			name:    "UnionUnion def",
			texpr:   T31.Texpr_UnionUnion().Value,
			dstZero: &T31.UnionUnion{},
			json:    `{"def": {"quant": 2, "value": 3}}`,
			want:    T31.Make_UnionUnion_def(T31.Make_S1(2, 3.0)),
		},
		{
			name:    "Struct02Test u1",
			texpr:   T31.Texpr_Struct02Test().Value,
			dstZero: &T31.Struct02Test{},
			json:    `{"a": "is an a"}`,
			want:    T31.Make_Struct02Test_u1(T31.Make_U1_a("is an a")),
		},
		{
			name:    "Struct02Test u1:b",
			texpr:   T31.Texpr_Struct02Test().Value,
			dstZero: &T31.Struct02Test{},
			json:    `{"b": 99}`,
			want:    T31.Make_Struct02Test_u1(T31.Make_U1_b(99)),
		},
		{
			name:    "Struct02Test s1",
			texpr:   T31.Texpr_Struct02Test().Value,
			dstZero: &T31.Struct02Test{},
			json:    `{"quant": 2, "value": 3}`,
			want:    T31.Make_Struct02Test_s1(T31.Make_S1(2, 3.0)),
		},
		{
			name:    "Struct02Test s1 - no lifting",
			texpr:   T31.Texpr_Struct02Test().Value,
			dstZero: &T31.Struct02Test{},
			json:    `{"s1":{"quant": 2, "value": 3}}`,
			want:    T31.Make_Struct02Test_s1(T31.Make_S1(2, 3.0)),
		},
		{
			name:    "NullableTest - a",
			texpr:   T31.Texpr_NullableTest().Value,
			dstZero: &T31.NullableTest{},
			json:    `null`,
			want:    T31.Make_NullableTest_a(nil),
		},
		{
			name:    "NullableTest - a:sdf",
			texpr:   T31.Texpr_NullableTest().Value,
			dstZero: &T31.NullableTest{},
			json:    `"sdf"`,
			want:    T31.Make_NullableTest_a(goadl.Addr("sdf")),
		},
		{
			name:    "VectorTest - a",
			texpr:   T31.Texpr_VectorTest().Value,
			dstZero: &T31.VectorTest{},
			json:    `[]`,
			want:    T31.Make_VectorTest_a([]string{}),
		},
		{
			name:    "VectorErrorTest",
			texpr:   T31.Texpr_VectorErrorTest().Value,
			dstZero: &T31.VectorErrorTest{},
			json:    `[]`,
			want:    T31.Make_VectorErrorTest_a([]string{}),
			wantErr: goadl.Addr(`ambiguous matching type discriminators`),
		},
		{
			name:    "VectorOfTypeDiscriminationUnionTest",
			texpr:   T31.Texpr_VectorOfTypeDiscriminationUnionTest().Value,
			dstZero: &T31.VectorOfTypeDiscriminationUnionTest{},
			json:    `{"b":[{"quant": 2, "value": 3},{"a": "is an a"},{"b": 99},{"s1":{"quant": 2, "value": 3}}]}`,
			want: T31.Make_VectorOfTypeDiscriminationUnionTest_b([]T31.Struct02Test{
				T31.Make_Struct02Test_s1(T31.Make_S1(2, 3.0)),
				T31.Make_Struct02Test_u1(T31.Make_U1_a("is an a")),
				T31.Make_Struct02Test_u1(T31.Make_U1_b(99)),
				T31.Make_Struct02Test_s1(T31.Make_S1(2, 3.0)),
			}),
		},
		{
			name:    "UnionOfVectorOfTypeDiscriminationUnionTest",
			texpr:   T31.Texpr_UnionOfVectorOfTypeDiscriminationUnionTest().Value,
			dstZero: &T31.UnionOfVectorOfTypeDiscriminationUnionTest{},
			json:    `[]`,
			want:    T31.Make_UnionOfVectorOfTypeDiscriminationUnionTest_b([]T31.Struct02Test{}),
			// wantErr: `union of union containing TypeDiscrimination branches not supported`,
		},
		{
			name:    "Ua",
			texpr:   T31.Texpr_Ua().Value,
			dstZero: &T31.Ua{},
			json:    `{"u_b":[{"u_c":[{"a":["def"]}]}]}`,
			want: T31.Make_Ua_u_b([]T31.Ub{
				T31.Make_Ub_u_c([]T31.Uc{
					T31.Make_Uc_a([]string{"def"}),
				}),
			}),
		},
		{
			name:    "Ua - with lifting",
			texpr:   T31.Texpr_Ua().Value,
			dstZero: &T31.Ua{},
			json:    `{"u_b":[{"u_c":[["def"]]}]}`,
			want: T31.Make_Ua_u_b([]T31.Ub{
				T31.Make_Ub_u_c([]T31.Uc{
					T31.Make_Uc_a([]string{"def"}),
				}),
			}),
		},
		{
			name:    "UofMaybe",
			texpr:   T31.Texpr_UofMaybe().Value,
			dstZero: &T31.UofMaybe{},
			json:    `{"a":{"just":{"abc":{"a":"a string"}}}}`,
			want:    T31.Make_UofMaybe_a(types.Make_Maybe_just(T31.Make_UnionUnion_abc(T31.Make_U1_a("a string")))),
		},
		{
			name:    "UofMaybe - lifting",
			texpr:   T31.Texpr_UofMaybe().Value,
			dstZero: &T31.UofMaybe{},
			json:    `{"a":{"just":{"a":"a string"}}}`,
			want:    T31.Make_UofMaybe_a(types.Make_Maybe_just(T31.Make_UnionUnion_abc(T31.Make_U1_a("a string")))),
		},
		{
			name:    "UTypeDiscriminationofMaybe - lifting U1",
			texpr:   T31.Texpr_UTypeDiscriminationofMaybe().Value,
			dstZero: &T31.UTypeDiscriminationofMaybe{},
			json:    `{"a":{"just":{"a":"a string"}}}`,
			want:    T31.Make_UTypeDiscriminationofMaybe_a(types.Make_Maybe_just(T31.Make_UnionUnion_abc(T31.Make_U1_a("a string")))),
		},
		{
			name:    "UTypeDiscriminationofMaybe - lifting U1 & top",
			texpr:   T31.Texpr_UTypeDiscriminationofMaybe().Value,
			dstZero: &T31.UTypeDiscriminationofMaybe{},
			json:    `{"just":{"a":"a string"}}`,
			want:    T31.Make_UTypeDiscriminationofMaybe_a(types.Make_Maybe_just(T31.Make_UnionUnion_abc(T31.Make_U1_a("a string")))),
		},
		{
			name:    "Maybe-lifting",
			texpr:   goadl.Texpr_Maybe(T31.Texpr_U1()).Value,
			dstZero: &types.Maybe[T31.U1]{},
			json:    `{"just":"a string"}`,
			want:    types.Make_Maybe_just(T31.Make_U1_a("a string")),
		},
		{
			name:    "U1-lifting ",
			texpr:   T31.Texpr_U1().Value,
			dstZero: &T31.U1{},
			json:    `"a string"`,
			want:    T31.Make_U1_a("a string"),
		},
		// {
		//   name: "A2",
		//   texpr: T31.Texpr_A2().Value,
		// dstZero: & T31.A2{},
		//   json: `{"a":"a string"}`,
		//   want: T31.Make_A2("a", T31.Make_A1_a( "a string")),
		// },
		// {
		//   name: "A2 - 2",
		//   texpr: T31.Texpr_A2().Value,
		// dstZero: & T31.A2{},
		//   json: `"a string"`,
		//   want: T31.Make_A2("a", T31.Make_A1_a( "a string")),
		//   wantErr: "primitive type mismatch. expected String received"
		// },
		{
			name:    "Deep",
			texpr:   T31.Texpr_A1().Value,
			dstZero: &T31.A1{},
			json:    `"sdaf"`,
			want:    T31.Make_A1_a(T31.Make_A2_b(T31.Make_A3_c(T31.Make_A4_d("sdaf")))),
		},
		{
			name:    "Deep-lift a4",
			texpr:   T31.Texpr_A1().Value,
			dstZero: &T31.A1{},
			json:    `{"d": "sdaf", "@v": 9}`,
			want:    T31.Make_A1_a(T31.Make_A2_b(T31.Make_A3_c(T31.Make_A4_d("sdaf")))),
		},
		{
			name:    "Deep-lift a3",
			texpr:   T31.Texpr_A1().Value,
			dstZero: &T31.A1{},
			json:    `{"c": {"d": "sdaf", "@v": 9}, "@v": 8}`,
			want:    T31.Make_A1_a(T31.Make_A2_b(T31.Make_A3_c(T31.Make_A4_d("sdaf")))),
		},
		{
			name:    "A3-lift a3.a4.d",
			texpr:   T31.Texpr_A3().Value,
			dstZero: &T31.A3{},
			json:    `{"c": "sdaf", "@v": 8}`,
			want:    T31.Make_A3_c(T31.Make_A4_d("sdaf")),
		},
		{
			name:    "Deep-lift a3.a4.d",
			texpr:   T31.Texpr_A1().Value,
			dstZero: &T31.A1{},
			json:    `{"c": "sdaf", "@v": 8}`,
			want:    T31.Make_A1_a(T31.Make_A2_b(T31.Make_A3_c(T31.Make_A4_d("sdaf")))),
		},
		{
			name:    "UnionUnion-lifting all",
			texpr:   T31.Texpr_UnionUnion().Value,
			dstZero: &T31.UnionUnion{},
			json:    `"a string"`,
			// want: T31.Make_UnionUnion("a", "a string"),
			want: T31.Make_UnionUnion_abc(T31.Make_U1_a("a string")),
		},
		{
			name:    "UTypeDiscriminationofMaybe-lifting all ",
			texpr:   T31.Texpr_UTypeDiscriminationofMaybe().Value,
			dstZero: &T31.UTypeDiscriminationofMaybe{},
			json:    `{"just":"a string"}`,
			want:    T31.Make_UTypeDiscriminationofMaybe_a(types.Make_Maybe_just(T31.Make_UnionUnion_abc(T31.Make_U1_a("a string")))),
		},
		{
			name:    "B1 f1",
			texpr:   T31.Texpr_B1().Value,
			dstZero: &T31.B1{},
			json:    `"a string"`,
			want:    T31.Make_B1_f1("a string"),
		},
		{
			name:    "B1 f2",
			texpr:   T31.Texpr_B1().Value,
			dstZero: &T31.B1{},
			json:    `["a string"]`,
			want:    T31.Make_B1_f2([]string{"a string"}),
		},
		{
			name:    "B2 f3",
			texpr:   T31.Texpr_B2().Value,
			dstZero: &T31.B2{},
			json:    `["a string"]`,
			want:    T31.Make_B2_f3([]string{"a string"}),
		},
		{
			name:    "B3 f3",
			texpr:   T31.Texpr_B3().Value,
			dstZero: &T31.B3{},
			json:    `["a string"]`,
			want:    T31.Make_B3_f3([]string{"a string"}),
		},
		{
			name:    "B4 a",
			texpr:   T31.Texpr_B4().Value,
			dstZero: &T31.B4{},
			json:    `{"a":"asdfghjkl"}`,
			want:    T31.Make_B4_b(T31.Make_GenU_a[string]("asdfghjkl")),
		},
		{
			name:    "B4 b",
			texpr:   T31.Texpr_B4().Value,
			dstZero: &T31.B4{},
			json:    `{"b":{"a":"asdfghjkl"}}`,
			want:    T31.Make_B4_b(T31.Make_GenU_a[string]("asdfghjkl")),
		},
		{
			name:    "B5 a",
			texpr:   T31.Texpr_B5().Value,
			dstZero: &T31.B5{},
			json:    `{"a":"asdfghjkl"}`,
			want:    T31.Make_B5_b(T31.Make_GenU_a[string]("asdfghjkl")),
		},
	}
	for i, tt := range tests {
		_ = i
		// if i > 2 {
		// 	break
		// }
		t.Run(tt.name, func(t *testing.T) {
			lter := annotations.CreateLifter(goadl.RESOLVER, tt.texpr)
			var jv any
			err := json.Unmarshal([]byte(tt.json), &jv)
			if err != nil {
				t.Errorf("input json error %v", err)
			}
			jv2, err := lter(jv)
			// json2, _ := json.Marshal(jv2)
			// fmt.Printf("%s\n%v\n%v\n", tt.name, tt.json, string(json2))
			if tt.wantErr != nil {
				if err != nil {
					if !strings.HasPrefix(err.Error(), *tt.wantErr) {
						t.Errorf("%s\nwant: %s\nhave: %s", tt.name, *tt.wantErr, err.Error())
					}
					return
				}
			}
			if err != nil {
				t.Errorf("lifter error :%v", err)
			}
			jdb := goadl.CreateUncheckedJsonDecodeBinding(tt.texpr, goadl.RESOLVER)
			err = jdb.DecodeFromAny(jv2, tt.dstZero)
			if err != nil {
				t.Errorf("DecodeFromAny error %v", err)
			}
			dst := reflect.ValueOf(tt.dstZero).Elem().Interface()
			if !reflect.DeepEqual(tt.want, dst) {
				json2, _ := json.Marshal(jv2)
				// fmt.Printf("%s\n%v\n%v\n", tt.name, tt.json, string(json2))
				t.Errorf("%s\n%+v\n%+v\n%s\n%s", tt.name, tt.want, dst, tt.json, string(json2))
			}
		})
	}
}
