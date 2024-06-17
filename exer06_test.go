package out_test

import (
	"bytes"
	"testing"

	"adl_testing/exer06/tttest"

	"adl_testing/diff"

	goadl "github.com/adl-lang/goadl_rt/v3"
)

func TestTypeTokenEncode(t *testing.T) {
	z := tttest.Make_Z()
	enc := goadl.CreateJsonEncodeBinding(tttest.Texpr_Z(), goadl.RESOLVER)
	buf := bytes.Buffer{}
	err := enc.Encode(&buf, z)
	if err != nil {
		t.Error(err)
	}
	a := buf.String()
	dec := goadl.CreateJsonDecodeBinding(tttest.Texpr_Z(), goadl.RESOLVER)
	var z2 tttest.Z
	err = dec.Decode(&buf, &z2)
	if err != nil {
		t.Error(err)
	}

	buf2 := bytes.Buffer{}
	err = enc.Encode(&buf2, z2)
	if err != nil {
		t.Error(err)
	}

	// note is it not true that !reflect.DeepEquals(z, z2) since TypeToken is encoded as null

	a2 := buf2.String()
	if a != a2 {
		a := []byte(a)
		a2 := []byte(a2)
		t.Errorf(string(diff.Diff("z", a, "z2", a2)))
	}
}
