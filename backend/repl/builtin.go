package repl

import (
	"strconv"

	"main.go/value"
)

type BuiltInFunction struct {
	Name string
	Exec func(context *ReplContext, args []*Argument) (value.IVOR, bool, string)
}

// implementing ivor

func (b BuiltInFunction) Type() string {
	return value.IVOR_BUILTIN_FUNCTION
}

func (b BuiltInFunction) Value() interface{} {
	return b
}

func (b BuiltInFunction) Copy() value.IVOR {
	return b
}

// * Print Function
func Print(context *ReplContext, args []*Argument) (value.IVOR, bool, string) {
	return PrintCore(context, args, false)
}

func PrintLn(context *ReplContext, args []*Argument) (value.IVOR, bool, string) {
	return PrintCore(context, args, true)
}

func PrintCore(context *ReplContext, args []*Argument, newLine bool) (value.IVOR, bool, string) {

	var output string

	for i, arg := range args {

		if !value.IsPrimitiveType(arg.Value.Type()) {
			return value.DefaultNilValue, false, "La función print solo acepta tipos primitivos"
		}

		switch arg.Value.Type() {

		case value.IVOR_BOOL:
			output += strconv.FormatBool(arg.Value.Value().(bool))
		case value.IVOR_INT:
			output += strconv.Itoa(arg.Value.Value().(int))
		case value.IVOR_FLOAT:
			output += strconv.FormatFloat(arg.Value.Value().(float64), 'f', 4, 64)
		case value.IVOR_STRING:
			output += arg.Value.Value().(string)
		case value.IVOR_CHARACTER:
			output += arg.Value.Value().(string)
		case value.IVOR_NIL:
			output += "nil"
		}

		if i < len(args)-1 {
			output += " "
		}
	}

	if newLine {
		context.Console.Print(output + "\n") // println agrega doble salto si así lo deseas
	} else {
		context.Console.Print(output)
	}

	return value.DefaultNilValue, true, ""
}

/*
func Print(context *ReplContext, args []*Argument) (value.IVOR, bool, string) {

	var output string

	for i, arg := range args {

		if !value.IsPrimitiveType(arg.Value.Type()) {
			return value.DefaultNilValue, false, "La función print solo acepta tipos primitivos"
		}

		switch arg.Value.Type() {

		case value.IVOR_BOOL:
			output += strconv.FormatBool(arg.Value.Value().(bool))
		case value.IVOR_INT:
			output += strconv.Itoa(arg.Value.Value().(int))
		case value.IVOR_FLOAT:
			output += strconv.FormatFloat(arg.Value.Value().(float64), 'f', 4, 64) // 4 digits of precision
		case value.IVOR_STRING:
			output += arg.Value.Value().(string)
		case value.IVOR_CHARACTER:
			output += arg.Value.Value().(string)
		case value.IVOR_NIL:
			output += "nil"
		}

		// Add a space between each argument
		if i < len(args)-1 {
			output += " "
		}
	}

	context.Console.Print(output + "\n")

	return value.DefaultNilValue, true, ""
}

*/

// * Int Function

func Atoi(context *ReplContext, args []*Argument) (value.IVOR, bool, string) {

	if len(args) != 1 {
		return value.DefaultNilValue, false, "La función int solo acepta un argumento"
	}

	argValue := args[0].Value

	if !(argValue.Type() == value.IVOR_STRING || argValue.Type() == value.IVOR_FLOAT) {
		return value.DefaultNilValue, false, "La función Int solo acepta un argumento de tipo string o float"
	}

	if argValue.Type() == value.IVOR_STRING {
		floatValue, err := strconv.ParseFloat(argValue.Value().(string), 64)

		if err != nil {
			return value.DefaultNilValue, false, "No se pudo convertir el valor a int"
		}

		return &value.IntValue{
			InternalValue: int(floatValue),
		}, true, ""
	}

	if argValue.Type() == value.IVOR_FLOAT {
		// truncate the float

		floatValue := argValue.Value().(float64)

		return &value.IntValue{
			InternalValue: int(floatValue),
		}, true, ""
	}

	return value.DefaultNilValue, false, "No se pudo convertir el valor a int"
}

// * Float Function

func ParseFloat(context *ReplContext, args []*Argument) (value.IVOR, bool, string) {

	if len(args) != 1 {
		return value.DefaultNilValue, false, "La función float solo acepta un argumento"
	}

	argValue := args[0].Value

	if !(argValue.Type() == value.IVOR_STRING) {
		return value.DefaultNilValue, false, "La función float solo acepta un argumento de tipo string"
	}

	floatValue, err := strconv.ParseFloat(argValue.Value().(string), 64)

	if err != nil {
		return value.DefaultNilValue, false, "No se pudo convertir el valor a float"
	}

	return &value.FloatValue{
		InternalValue: floatValue,
	}, true, ""
}

// * String Function

func String(context *ReplContext, args []*Argument) (value.IVOR, bool, string) {

	if len(args) != 1 {
		return value.DefaultNilValue, false, "La función string solo acepta un argumento"
	}

	argValue := args[0].Value

	if !(argValue.Type() == value.IVOR_INT || argValue.Type() == value.IVOR_FLOAT || argValue.Type() == value.IVOR_BOOL) {
		return value.DefaultNilValue, false, "La función string solo acepta un argumento de tipo int, float o bool"
	}

	if argValue.Type() == value.IVOR_INT {
		stringValue := strconv.Itoa(argValue.Value().(int))

		return &value.StringValue{
			InternalValue: stringValue,
		}, true, ""
	}

	if argValue.Type() == value.IVOR_FLOAT {
		stringValue := strconv.FormatFloat(argValue.Value().(float64), 'f', 4, 64)

		return &value.StringValue{
			InternalValue: stringValue,
		}, true, ""
	}

	if argValue.Type() == value.IVOR_BOOL {
		stringValue := strconv.FormatBool(argValue.Value().(bool))

		return &value.StringValue{
			InternalValue: stringValue,
		}, true, ""
	}

	return value.DefaultNilValue, false, "No se pudo convertir el valor a string"
}

func TypeOf(context *ReplContext, args []*Argument) (value.IVOR, bool, string) {

	if len(args) != 1 {
		return value.DefaultNilValue, false, "La función typeOf solo acepta un argumento"
	}

	argValue := args[0].Value

	// Obtenemos el tipo directamente (ya es el nombre del tipo)
	typeName := argValue.Type()

	// Debug: imprimir el tipo en la consola interna del intérprete
	context.Console.Print("[DEBUG] TypeOf: " + typeName)

	return &value.StringValue{
		InternalValue: typeName,
	}, true, ""
}

var DefaultBuiltInFunctions = map[string]*BuiltInFunction{
	"print": {
		Name: "print",
		Exec: Print,
	},
	"println": {
		Name: "println",
		Exec: PrintLn,
	},
	"String": {
		Name: "String",
		Exec: String,
	},
	"Atoi": {
		Name: "Atoi",
		Exec: Atoi,
	},
	"parseFloat": {
		Name: "parseFloat",
		Exec: ParseFloat,
	},
	"typeOf": {
		Name: "typeOf",
		Exec: TypeOf,
	},
}
