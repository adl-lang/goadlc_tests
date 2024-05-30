package out_test

import (
	"bytes"
	"reflect"
	"testing"

	"adl_testing/exer03/generics"

	"adl_testing/diff"

	goadl "github.com/adl-lang/goadl_rt/v3"
	"github.com/davecgh/go-spew/spew"
)

func TestGenericEncode(t *testing.T) {
	x := generics.Make_Abc[int64, string](
		1, []string{"a"}, generics.Make_Def_a[int64, string](3),
	)

	x.Kids = []generics.Abc[int64, string]{x}

	// x := generics.Abc[int64, string]{
	// 	A: 1,
	// 	B: []string{"a"},
	// 	C: 2,
	// 	D: generics.Make_Def_a[int64, string](3),
	// }
	out := &bytes.Buffer{}
	enc := goadl.CreateJsonEncodeBinding(generics.Texpr_Abc(goadl.Texpr_Int64(), goadl.Texpr_String()), goadl.RESOLVER)
	err := enc.Encode(out, x)
	if err != nil {
		t.Errorf("%v", err)
	}
	dec := goadl.CreateJsonDecodeBinding(generics.Texpr_Abc(goadl.Texpr_Int64(), goadl.Texpr_String()), goadl.RESOLVER)
	var y generics.Abc[int64, string]
	err = dec.Decode(out, &y)
	if err != nil {
		t.Errorf("%v", err)
	}
	if !reflect.DeepEqual(x, y) {
		have := spew.Sdump(x)
		want := spew.Sdump(y)
		d := diff.Diff("have", []byte(have), "want", []byte(want))
		t.Errorf("!=\n%v\n", string(d))

	}
	// out2 := bytes.Buffer{}
	// json.Indent(&out2, out.Bytes(), "", "  ")
	// fmt.Printf("%s\n", out2.String())
	// o2, _ := json.Marshal(x)
	// fmt.Printf("%s\n", string(o2))

}
