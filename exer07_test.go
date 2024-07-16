package out_test

import (
	"strings"
	"testing"

	"adl_testing/exer07/lifting"

	goadl "github.com/adl-lang/goadl_rt/v3"
)

func TestLifting(t *testing.T) {
	jv1 := strings.NewReader(`{"a":"abc","b":42}`)
	dec := goadl.CreateJsonDecodeBinding(lifting.Texpr_Lifted(), goadl.RESOLVER)
	l1 := lifting.Lifted{}
	err := dec.Decode(jv1, &l1)
	if err != nil {
		t.Error(err)
	}
	_, is := l1.Cast_org_field()
	if !is {
		t.Errorf("l1 is org field type")
	}

}
