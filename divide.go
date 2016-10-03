package arithmetic

import "fmt"

type Divide struct{}

func (o Divide) String() string {
	return "/"
}

func (o Divide) Value() (Operand, Operator) {
	return nil, o
}

func (o Divide) Kind() Kind {
	return KindOperation
}

func (o Divide) Precedence() uint8 {
	return 2
}

func (o Divide) Solve(st OperandStack) (Operand, error) {
	right, ok := st.Pop()
	if !ok {
		return nil, fmt.Errorf("invalid operation: \"/\" must be followed by a valid operand or expression")
	}

	r, err := ToFloat(right)
	if err != nil {
		return nil, fmt.Errorf("invalid operand: %s", err)
	}

	left, ok := st.Pop()
	if !ok {
		return nil, fmt.Errorf("invalid operation: \"/ %s\" must be preceeded by a valid operand or expression", right)
	}

	l, err := ToFloat(left)
	if err != nil {
		return nil, fmt.Errorf("invalid operand: %s", err)
	}

	return l / r, nil
}
