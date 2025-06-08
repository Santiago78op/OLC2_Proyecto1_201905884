package repl

import (
	"main.go/value"
)

type evalFunc func(value.IVOR, value.IVOR) (bool, string, value.IVOR) // takes 2 values and returns a value
type conversionFunc func(value.IVOR) value.IVOR                       // takes a value and returns a value (different type)

type BinaryValidation struct {
	LeftType        string // allowed left type
	RightType       string // allowed right type
	LeftConversion  conversionFunc
	RightConversion conversionFunc
	Eval            evalFunc
}

type BinaryStrategy struct {
	Name        string
	Validations []BinaryValidation
	Viceversa   bool // if true, the validation is also performed in the opposite order
	DefaultEval evalFunc
}

func (s *BinaryStrategy) Validate(left, right value.IVOR) (bool, string, value.IVOR) {

	// nil in any side is, by default return nil

	if left.Type() == value.IVOR_NIL || right.Type() == value.IVOR_NIL {
		return false, "No es posible realizar operaciones con valores nulos", value.DefaultNilValue
	}

	for _, valid := range s.Validations {

		if valid.LeftType == left.Type() && valid.RightType == right.Type() {

			if valid.LeftConversion != nil {
				left = valid.LeftConversion(left)
			}

			if valid.RightConversion != nil {
				right = valid.RightConversion(right)
			}

			if valid.Eval != nil {
				return valid.Eval(left, right)
			}

			return s.DefaultEval(left, right)
		}

		if s.Viceversa && valid.LeftType == right.Type() && valid.RightType == left.Type() {

			if valid.LeftConversion != nil {
				right = valid.LeftConversion(right)
			}

			if valid.RightConversion != nil {
				left = valid.RightConversion(left)
			}

			if valid.Eval != nil {
				return valid.Eval(left, right)
			}

			return s.DefaultEval(left, right)
		}

	}

	msg := "No es posible realizar la operación '" + s.Name + "' con los tipos '" + left.Type() + "' y '" + right.Type() + "'"

	return false, msg, value.DefaultNilValue
}

// * arithmetic operators

// int + int; float + float; float + int (viceversa); string + string
var addStrategy = BinaryStrategy{
	Name:        "+",
	Viceversa:   true,
	DefaultEval: nil,
	Validations: []BinaryValidation{
		{
			LeftType:        value.IVOR_INT,
			RightType:       value.IVOR_INT,
			LeftConversion:  nil,
			RightConversion: nil,
			Eval: func(left, right value.IVOR) (bool, string, value.IVOR) {
				return true, "", &value.IntValue{
					InternalValue: left.(*value.IntValue).InternalValue + right.(*value.IntValue).InternalValue,
				}
			},
		},
		{
			LeftType:        value.IVOR_FLOAT,
			RightType:       value.IVOR_FLOAT,
			LeftConversion:  nil,
			RightConversion: nil,
			Eval: func(left, right value.IVOR) (bool, string, value.IVOR) {
				return true, "", &value.FloatValue{
					InternalValue: left.(*value.FloatValue).InternalValue + right.(*value.FloatValue).InternalValue,
				}
			},
		},
		{
			LeftType:       value.IVOR_FLOAT,
			RightType:      value.IVOR_INT,
			LeftConversion: nil,
			RightConversion: func(v value.IVOR) value.IVOR {
				return &value.FloatValue{
					InternalValue: float64(v.(*value.IntValue).InternalValue),
				}
			},
			Eval: func(left, right value.IVOR) (bool, string, value.IVOR) {
				return true, "", &value.FloatValue{
					InternalValue: left.(*value.FloatValue).InternalValue + right.(*value.FloatValue).InternalValue,
				}
			},
		},
		{
			LeftType:        value.IVOR_STRING,
			RightType:       value.IVOR_STRING,
			LeftConversion:  nil,
			RightConversion: nil,
			Eval: func(left, right value.IVOR) (bool, string, value.IVOR) {
				return true, "", &value.StringValue{
					InternalValue: left.(*value.StringValue).InternalValue + right.(*value.StringValue).InternalValue,
				}
			},
		},
		{
			LeftType:  value.IVOR_CHARACTER,
			RightType: value.IVOR_STRING,
			LeftConversion: func(v value.IVOR) value.IVOR {
				return &value.StringValue{
					InternalValue: string(v.(*value.CharacterValue).InternalValue),
				}
			},
			RightConversion: nil,
			Eval: func(left, right value.IVOR) (bool, string, value.IVOR) {
				return true, "", &value.StringValue{
					InternalValue: left.(*value.StringValue).InternalValue + right.(*value.StringValue).InternalValue,
				}
			},
		},
	},
}

// int - int; float - float; float - int (viceversa)
var subStrategy = BinaryStrategy{
	Name:        "-",
	Viceversa:   true,
	DefaultEval: nil,
	Validations: []BinaryValidation{
		{
			LeftType:        value.IVOR_INT,
			RightType:       value.IVOR_INT,
			LeftConversion:  nil,
			RightConversion: nil,
			Eval: func(left, right value.IVOR) (bool, string, value.IVOR) {
				return true, "", &value.IntValue{
					InternalValue: left.(*value.IntValue).InternalValue - right.(*value.IntValue).InternalValue,
				}
			},
		},
		{
			LeftType:        value.IVOR_FLOAT,
			RightType:       value.IVOR_FLOAT,
			LeftConversion:  nil,
			RightConversion: nil,
			Eval: func(left, right value.IVOR) (bool, string, value.IVOR) {
				return true, "", &value.FloatValue{
					InternalValue: left.(*value.FloatValue).InternalValue - right.(*value.FloatValue).InternalValue,
				}
			},
		},
		{
			LeftType:       value.IVOR_FLOAT,
			RightType:      value.IVOR_INT,
			LeftConversion: nil,
			RightConversion: func(v value.IVOR) value.IVOR {
				return &value.FloatValue{
					InternalValue: float64(v.(*value.IntValue).InternalValue),
				}
			},
			Eval: func(left, right value.IVOR) (bool, string, value.IVOR) {
				return true, "", &value.FloatValue{
					InternalValue: left.(*value.FloatValue).InternalValue - right.(*value.FloatValue).InternalValue,
				}
			},
		},
	},
}

// int * int; float * float; float * int (viceversa)
var mulStrategy = BinaryStrategy{
	Name:        "*",
	Viceversa:   true,
	DefaultEval: nil,
	Validations: []BinaryValidation{
		{
			LeftType:        value.IVOR_INT,
			RightType:       value.IVOR_INT,
			LeftConversion:  nil,
			RightConversion: nil,
			Eval: func(left, right value.IVOR) (bool, string, value.IVOR) {
				return true, "", &value.IntValue{
					InternalValue: left.(*value.IntValue).InternalValue * right.(*value.IntValue).InternalValue,
				}
			},
		},
		{
			LeftType:        value.IVOR_FLOAT,
			RightType:       value.IVOR_FLOAT,
			LeftConversion:  nil,
			RightConversion: nil,
			Eval: func(left, right value.IVOR) (bool, string, value.IVOR) {
				return true, "", &value.FloatValue{
					InternalValue: left.(*value.FloatValue).InternalValue * right.(*value.FloatValue).InternalValue,
				}
			},
		},
		{
			LeftType:       value.IVOR_FLOAT,
			RightType:      value.IVOR_INT,
			LeftConversion: nil,
			RightConversion: func(v value.IVOR) value.IVOR {
				return &value.FloatValue{
					InternalValue: float64(v.(*value.IntValue).InternalValue),
				}
			},
			Eval: func(left, right value.IVOR) (bool, string, value.IVOR) {
				return true, "", &value.FloatValue{
					InternalValue: left.(*value.FloatValue).InternalValue * right.(*value.FloatValue).InternalValue,
				}
			},
		},
	},
}

// int / int; float / float; float / int (viceversa) !division by zero
var divStrategy = BinaryStrategy{
	Name:        "/",
	Viceversa:   true,
	DefaultEval: nil,
	Validations: []BinaryValidation{
		{
			LeftType:        value.IVOR_INT,
			RightType:       value.IVOR_INT,
			LeftConversion:  nil,
			RightConversion: nil,
			Eval: func(left, right value.IVOR) (bool, string, value.IVOR) {

				if right.(*value.IntValue).InternalValue == 0 {
					return false, "No se puede dividir entre cero", value.DefaultNilValue
				}

				return true, "", &value.IntValue{
					InternalValue: left.(*value.IntValue).InternalValue / right.(*value.IntValue).InternalValue,
				}
			},
		},
		{
			LeftType:        value.IVOR_FLOAT,
			RightType:       value.IVOR_FLOAT,
			LeftConversion:  nil,
			RightConversion: nil,
			Eval: func(left, right value.IVOR) (bool, string, value.IVOR) {

				if right.(*value.FloatValue).InternalValue == 0 {
					return false, "No se puede dividir entre cero", value.DefaultNilValue
				}

				return true, "", &value.FloatValue{
					InternalValue: left.(*value.FloatValue).InternalValue / right.(*value.FloatValue).InternalValue,
				}
			},
		},
		{
			LeftType:       value.IVOR_FLOAT,
			RightType:      value.IVOR_INT,
			LeftConversion: nil,
			RightConversion: func(v value.IVOR) value.IVOR {
				return &value.FloatValue{
					InternalValue: float64(v.(*value.IntValue).InternalValue),
				}
			},
			Eval: func(left, right value.IVOR) (bool, string, value.IVOR) {

				if right.(*value.FloatValue).InternalValue == 0 {
					return false, "No se puede dividir entre cero", value.DefaultNilValue
				}

				return true, "", &value.FloatValue{
					InternalValue: left.(*value.FloatValue).InternalValue / right.(*value.FloatValue).InternalValue,
				}
			},
		},
	},
}
