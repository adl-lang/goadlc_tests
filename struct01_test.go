package out_test

import (
	"adl_testing/exer01/struct01"
	"bytes"
	"testing"

	goadl "github.com/adl-lang/goadl_rt/v3"
)

func TestXxx(t *testing.T) {
	a := "a"
	x := struct01.Make_Struct01(
		41,
		"",
		map[string]any{
			"a": 1234567890,
		},
		[]string{"a", "b", "c"},
		map[string][]string{"a": {"z"}, "b": {"x"}, "c": {"y"}},
		map[string]map[string]string{},
		map[string]map[string]*string{},
		&a,
		struct01.Make_B(
			"sfd",
		),
	)

	// v := reflect.ValueOf(x)
	// f0 := v.Field(0)
	// fmt.Printf("%v\n", f0.IsZero())
	// f2 := v.Field(2)
	// fmt.Printf("%v\n", f2.IsZero())

	out := &bytes.Buffer{}
	enc := goadl.CreateJsonEncodeBinding[struct01.Struct01](struct01.Texpr_Struct01(), goadl.RESOLVER)
	enc.Encode(out, x)
	// fmt.Printf("%s\n", string(out.Bytes()))
}
