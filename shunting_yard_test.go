package arithmetic

import (
	"math"
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
		{
			in:  []interface{}{5, multiply{}, 3},
			out: []interface{}{5, 3, multiply{}},
			err: false,
		},
		{
			in:  []interface{}{5, multiply{}, 3},
			out: []interface{}{5, 3, multiply{}},
			err: false,
		},
		{
			in:  []interface{}{1, multiply{}, 2, minus{}, 3, multiply{}, 4},
			out: []interface{}{1, 2, multiply{}, 3, 4, multiply{}, minus{}},
			err: false,
		},
		{
			in:  []interface{}{1, minus{}, 2, multiply{}, 3, minus{}, 4},
			out: []interface{}{1, 2, 3, multiply{}, minus{}, 4, minus{}},
			err: false,
		},
		{
			in:  []interface{}{unaryMinus{}, 1},
			out: []interface{}{1, unaryMinus{}},
			err: false,
		},
		{
			in:  []interface{}{unaryMinus{}, 1, multiply{}, 2},
			out: []interface{}{1, unaryMinus{}, 2, multiply{}},
			err: false,
		},
		{
			in:  []interface{}{2.0, multiply{}, leftParenthesis{}, unaryMinus{}, 3.0, minus{}, 5.0, rightParenthesis{}},
			out: []interface{}{2.0, 3.0, unaryMinus{}, 5.0, minus{}, multiply{}},
			err: false,
		},
		{
			in:  []interface{}{2.0, rightParenthesis{}},
			out: nil,
			err: true,
		},
		{
			in:  []interface{}{2.0, leftParenthesis{}},
			out: nil,
			err: true,
		},
		{
			in:  []interface{}{leftParenthesis{}, 2.0, rightParenthesis{}},
			out: []interface{}{2.0},
			err: false,
		},
		{
			in:  []interface{}{2.0, multiply{}, variable{"e", math.E}},
			out: []interface{}{2.0, variable{"e", math.E}, multiply{}},
			err: false,
		},
		{
			in:  []interface{}{f("max"), leftParenthesis{}, 3.0, comma{}, 1.0, rightParenthesis{}},
			out: []interface{}{3.0, 1.0, 2, f("max")},
			err: false,
		},
		{
			in:  []interface{}{"random", "string"},
			out: []interface{}{"random", "string"},
			err: false,
		},
		{
			in:  []interface{}{f("if"), leftParenthesis{}, true, comma{}, 1.0, comma{}, 0.0, rightParenthesis{}},
			out: []interface{}{true, 1.0, 0.0, 3, f("if")},
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
			t.Logf("want: %v\n", c.out)
			t.Logf("got:  %v\n", out)
			t.Fail()
		}
	}
}
