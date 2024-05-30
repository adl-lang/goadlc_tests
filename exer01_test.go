package out_test

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"adl_testing/exer01/simple_union"
	"adl_testing/exer01/struct01"

	goadl "github.com/adl-lang/goadl_rt/v3"
)

func TestEnum(t *testing.T) {
	x := simple_union.Make_UnionOfVoids_g(struct{}{})
	out := &bytes.Buffer{}
	enc := goadl.CreateJsonEncodeBinding(simple_union.Texpr_UnionOfVoids(), goadl.RESOLVER)
	err := enc.Encode(out, x)
	if err != nil {
		t.Error(err)
	}
	dec := goadl.CreateJsonDecodeBinding(simple_union.Texpr_UnionOfVoids(), goadl.RESOLVER)
	var y simple_union.UnionOfVoids
	err = dec.Decode(out, &y)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(x, y) {
		t.Errorf("not equal")
	}
}

func TestUnion(t *testing.T) {
	x := simple_union.Make_UnionOfPrimitives_A(42)
	out := &bytes.Buffer{}
	enc := goadl.CreateJsonEncodeBinding(simple_union.Texpr_UnionOfPrimitives(), goadl.RESOLVER)
	err := enc.Encode(out, x)
	if err != nil {
		t.Error(err)
	}

	if `{"A":42}` != out.String() {
		t.Error(`{"A":42} != out.String()`)
	}
	dec := goadl.CreateJsonDecodeBinding(simple_union.Texpr_UnionOfPrimitives(), goadl.RESOLVER)
	var y simple_union.UnionOfPrimitives
	err = dec.Decode(out, &y)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(x, y) {
		t.Errorf("not equal")
	}
}

func TestUnions(t *testing.T) {
	xs := []simple_union.UnionOfPrimitives{
		simple_union.Make_UnionOfPrimitives_A(42),
		simple_union.Make_UnionOfPrimitives_B(41),
		simple_union.Make_UnionOfPrimitives_c(true),
		simple_union.Make_UnionOfPrimitives_d(41.01),
		simple_union.Make_UnionOfPrimitives_e("sdfadf"),
		simple_union.Make_UnionOfPrimitives_f([]string{"a", "b", "v"}),
		simple_union.Make_UnionOfPrimitives_g(struct{}{}),
	}
	out := &bytes.Buffer{}
	te := goadl.Texpr_Vector(simple_union.Texpr_UnionOfPrimitives())
	enc := goadl.CreateJsonEncodeBinding(te, goadl.RESOLVER)
	enc.Encode(out, xs)
	if `[{"A":42},{"B":41},{"c":true},{"d":41.01},{"e":"sdfadf"},{"f":["a","b","v"]},{"g":null}]` != out.String() {
		t.Error("json str !=")
	}
	dec := goadl.CreateJsonDecodeBinding(te, goadl.RESOLVER)
	ys := []simple_union.UnionOfPrimitives{}
	err := dec.Decode(out, &ys)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(xs, ys) {
		t.Errorf("not equal")
	}
}

func TestStruct01(t *testing.T) {
	a := "a"
	x := struct01.Struct01{
		A: struct{}{},
		B: 41,
		C: "",
		D: map[string]any{
			"a": float64(123456),
		},
		E: []string{"a", "b", "c"},
		F: map[string][]string{"a": {"z"}, "b": {"x"}, "sc": {"y"}},
		I: &a,
		J: struct01.B{
			A: "sfd",
		},
	}

	out := &bytes.Buffer{}
	enc := goadl.CreateJsonEncodeBinding[struct01.Struct01](struct01.Texpr_Struct01(), goadl.RESOLVER)
	enc.Encode(out, x)
	o1 := out.String()

	dec := goadl.CreateJsonDecodeBinding(struct01.Texpr_Struct01(), goadl.RESOLVER)
	y := struct01.Struct01{}
	err := dec.Decode(out, &y)
	if err != nil {
		t.Error(err)
	}
	// if !reflect.DeepEqual(x, y) {
	// 	t.Errorf("not equal\n%+v\n%+v\n", x, y)
	// }
	sb := strings.Builder{}
	enc2 := goadl.CreateJsonEncodeBinding[struct01.Struct01](struct01.Texpr_Struct01(), goadl.RESOLVER)
	enc2.Encode(&sb, x)
	o2 := sb.String()
	if o1 != o2 {
		t.Errorf("json !=")
	}
}
