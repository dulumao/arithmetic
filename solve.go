package arithmetic

func Solve(input []interface{}) (interface{}, error) {
	st := &stack{}

	for _, t := range input {

		switch v := t.(type) {
		// 		case function:
		// 			size, ok := st.popInt()
		// 			if !ok {
		// 				return nil, err
		// 			}
		//
		// 			args, err := st.slice()
		// 			if err != nil {
		// 				return nil, err
		// 			}
		//
		// 			o, err := v.solve(args...)
		// 			if err != nil {
		// 				return nil, err
		// 			}
		//
		// 			st.Push(o)

		case operator:
			o, err := v.solve(st)
			if err != nil {
				return nil, err
			}
			st.push(o)

		default:
			st.push(v)
		}
	}

	out, _ := st.pop()
	return out, nil
}
