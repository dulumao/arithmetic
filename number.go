package arithmetic

import (
	"fmt"
	"strconv"
)

type Number float64

func (o Number) String() string {
	return strconv.FormatFloat(float64(o), 'f', 2, 64)
}

func (o Number) Value() (Operand, Operator) {
	return o, nil
}

func ToFloat(o Operand) (float64, error) {
	v, ok := o.(Number)
	if !ok {
		return 0, fmt.Errorf("expecing float, having %v(%T)", o, o)
	}

	return float64(v), nil
}

func ToInt(o Operand) (int, error) {
	v, err := ToFloat(o)
	if err != nil {
		return 0, err
	}

	return int(v), nil
}
