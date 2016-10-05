package arithmetic

import (
	"reflect"
	"testing"
)

func TestShuntingYard(t *testing.T) {

	for i, c := range []struct {
		in  []interface{}
		out []interface{}
		err bool
	}{
		{
			in:  []interface{}{},
			out: nil,
			err: false,
		},
		{
			in:  []interface{}{5},
			out: []interface{}{5},
			err: false,
		},
		{
			in:  []interface{}{5, minus{}, 3},
			out: []interface{}{5, 3, minus{}},
			err: false,
		},
		{
			in:  []interface{}{5, minus{}, 3, minus{}, 4},
			out: []interface{}{5, 3, minus{}, 4, minus{}},
			err: false,
		},
	} {
		out, err := ShuntingYard(c.in)
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