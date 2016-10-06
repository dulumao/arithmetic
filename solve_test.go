package arithmetic

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {

	for i, c := range []struct {
		in  []interface{}
		out interface{}
		err bool
	}{
		{
			in:  []interface{}{},
			out: nil,
			err: false,
		},
		{
			in:  []interface{}{minus{}},
			out: nil,
			err: true,
		},
		{
			in:  []interface{}{3.0, 2.0, minus{}, minus{}},
			out: nil,
			err: true,
		},
		{
			in:  []interface{}{3.0, 2.0, minus{}},
			out: 1.0,
			err: false,
		},
		{
			in:  []interface{}{multiply{}},
			out: nil,
			err: true,
		},
		{
			in:  []interface{}{3.0, 2.0, multiply{}, multiply{}},
			out: nil,
			err: true,
		},
		{
			in:  []interface{}{3.0, 2.0, multiply{}},
			out: 6.0,
			err: false,
		},
		{
			in:  []interface{}{3.0, 2.0, multiply{}, 3.0, minus{}},
			out: 3.0,
			err: false,
		},
		{
			in:  []interface{}{1.0, 2.0, 3.0, multiply{}, minus{}, 4.0, minus{}},
			out: -9.0,
			err: false,
		},
	} {

		out, err := Solve(c.in)
		if (err != nil) != c.err {
			t.Log("case", i+1, "unexpected error")
			t.Log("want:", c.err)
			t.Log("got: ", err)
			t.Fail()
			continue
		}

		if !reflect.DeepEqual(out, c.out) {
			t.Log("case", i+1, "unexpected output")
			t.Log("want:", c.out)
			t.Log("got: ", out)
			t.Fail()
		}
	}
}
