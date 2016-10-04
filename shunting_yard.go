package arithmetic

import (
	"fmt"
)

func ShuntingYard(input []Token) ([]Token, error) {
	
	st := &OperatorStack{}
	as := &ArityStack{}
	var output []Token

	for i, t := range input {
		
		_, op := t.Value()

		if op == nil {
			output = append(output, t)
			continue
		}

		if op.Kind() == KindFunction {
			as.Push(1)
			st.Push(op)
			continue
		}

		if op.Kind() == KindLeftParenthesis {
			st.Push(op)
			continue
		}

		if op.Kind() == KindOperation {
			for {
				v, ok := st.Pop()
				if !ok {
					break
				}

				if v.Precedence() > op.Precedence() {
					if v.Kind() == KindFunction {
						output = append(output, Number(as.Pop()))
					}
					output = append(output, v.(Token))
					continue
				}

				st.Push(v)
				break
			}

			st.Push(op)
			continue
		}

		if op.Kind() == KindComma {
			as.Inc()
			for {
				v, ok := st.Pop()
				if !ok {
					return nil, fmt.Errorf("invalid expression at position %d: %s...", i+1, tokensToString(input[:i+1]))
				}

				if v.Kind() == KindLeftParenthesis {
					st.Push(v)
					break
				}

				if v.Kind() == KindFunction {
					output = append(output, Number(as.Pop()))
				}
				output = append(output, v.(Token))
			}
			continue
		}

		if op.Kind() == KindRightParenthesis {
			for {
				v, ok := st.Pop()
				if !ok {
					return nil, fmt.Errorf("invalid expression at position %d: %s...", i+1, tokensToString(input[:i+1]))
				}

				if v.Kind() == KindLeftParenthesis {
					
					v, ok := st.Pop()
					if !ok {
						break
					}

					if v.Kind() == KindFunction {
						output = append(output, Number(as.Pop()), v.(Token))
						break
					}
					
					st.Push(v)
					break
				}
				output = append(output, v.(Token))
			}
			continue
		}
	}

	for {

		v, ok := st.Pop()
		if !ok {
			break
		}
		if v.Kind() == KindLeftParenthesis {
			return nil, fmt.Errorf("mismatched parenthesis: %s", tokensToString(input))
		}
		if v.Kind() == KindFunction {
			output = append(output, Number(as.Pop()))
		}
		output = append(output, v.(Token))
	}
	
	return output, nil
}
