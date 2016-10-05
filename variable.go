package arithmetic

// import (
// 	"fmt"
// 	"math"
// 	"strings"
// )
//
// func init() {
// 	RegisterVariable("e", math.E)
// 	RegisterVariable("pi", math.Pi)
// 	RegisterVariable("phi", math.Phi)
// 	RegisterVariable("sqrt2", math.Sqrt2)
// 	RegisterVariable("sqrte", math.SqrtE)
// 	RegisterVariable("sqrtpi", math.SqrtPi)
// 	RegisterVariable("sqrtphi", math.SqrtPhi)
// 	RegisterVariable("ln2", math.Ln2)
// 	RegisterVariable("log2e", math.Log2E)
// 	RegisterVariable("ln10", math.Ln10)
// 	RegisterVariable("ln10e", math.Log10E)
// 	RegisterVariable("true", true)
// 	RegisterVariable("false", false)
// }
//
// var variables = map[string]variable{}
//
// func RegisterVariable(label string, value interface{}) {
//
// 	label = strings.ToLower(label)
//
// 	if _, ok := functions[label]; ok {
// 		panic(fmt.Sprintf("%s already defined as function", label))
// 	}
//
// 	if _, ok := variables[label]; ok {
// 		panic(fmt.Sprintf("%s already defined as variable", label))
// 	}
//
// 	var v variable
// 	v.label = label
//
// 	switch t := value.(type) {
// 	case float64:
// 		v.operand = Number(t)
// 	case bool:
// 		v.operand = Boolean(t)
// 	default:
// 		panic(fmt.Sprintf("forbidden variable %s type %T", label, value))
// 	}
//
// 	variables[label] = v
// }
//
// type variable struct {
// 	label   string
// 	operand Operand
// }
//
// func (v variable) String() string {
// 	return v.label
// }
//
// func (v variable) Value() (Operand, Operator) {
// 	return v.operand, nil
// }
