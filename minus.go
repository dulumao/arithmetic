package arithmetic

type minus struct{}

func (o minus) String() string {
	return "-"
}

func (o minus) precedence() uint8 {
	return precedencePlus
}

func (o minus) solve(st *stack) (interface{}, error) {
	o2, err := st.popFloat()
	if err != nil {
		return nil, rightError(o)
	}

	o1, err := st.popFloat()
	if err != nil {
		return nil, leftError(o, o2)
	}

	return o1 - o2, nil
}

type unaryMinus struct{}

func (o unaryMinus) String() string {
	return "-"
}

func (o unaryMinus) precedence() uint8 {
	return precedenceUnary
}

func (o unaryMinus) solve(st *stack) (interface{}, error) {
	right, err := st.popFloat()
	if err != nil {
		return nil, rightError(o)
	}

	return -right, nil
}
