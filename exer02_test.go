package out_test

import (
	"bytes"
	"testing"

	"adl_testing/exer02/a"
	"adl_testing/exer02/b"

	goadl "github.com/adl-lang/goadl_rt/v3"
)

func TestExec02Encode(t *testing.T) {
	x := a.A{
		A: b.B{},
	}
	out := &bytes.Buffer{}
	enc := goadl.CreateJsonEncodeBinding[a.A](a.Texpr_A(), goadl.RESOLVER)
	enc.Encode(out, x)
	// fmt.Printf("%s\n", string(out.Bytes()))
	// o2, _ := json.Marshal(x)
	// fmt.Printf("%s\n", string(o2))

}
