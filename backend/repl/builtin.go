package repl

import (
	"strconv"
	"strings"

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

		// Verificar si es un tipo primitivo O un vector
		if !value.IsPrimitiveType(arg.Value.Type()) && !IsVectorType(arg.Value.Type()) {
			return value.DefaultNilValue, false, "La función print solo acepta tipos primitivos y vectores"
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
		default:
			// Si no es un tipo primitivo, verificar si es un vector
			if IsVectorType(arg.Value.Type()) {
				vectorOutput := formatVector(arg.Value.(*VectorValue))
				output += vectorOutput
			} else {
				return value.DefaultNilValue, false, "Tipo no soportado para print: " + arg.Value.Type()
			}
		}

		// Add a space between each argument
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

// Función auxiliar para formatear vectores
func formatVector(vector *VectorValue) string {
	if len(vector.InternalValue) == 0 {
		return "[ ]"
	}

	var result strings.Builder
	result.WriteString("[ ")

	for i, item := range vector.InternalValue {
		switch item.Type() {
		case value.IVOR_BOOL:
			result.WriteString(strconv.FormatBool(item.Value().(bool)))
		case value.IVOR_INT:
			result.WriteString(strconv.Itoa(item.Value().(int)))
		case value.IVOR_FLOAT:
			result.WriteString(strconv.FormatFloat(item.Value().(float64), 'f', 4, 64))
		case value.IVOR_STRING:
			result.WriteString(item.Value().(string))
		case value.IVOR_CHARACTER:
			result.WriteString(item.Value().(string))
		case value.IVOR_NIL:
			result.WriteString("nil")
		default:
			// Para vectores anidados u otros tipos
			if IsVectorType(item.Type()) {
				result.WriteString(formatVector(item.(*VectorValue)))
			} else {
				result.WriteString(item.Type()) // Mostrar el tipo si no se puede formatear
			}
		}

		// Agregar espacio entre elementos (excepto el último)
		if i < len(vector.InternalValue)-1 {
			result.WriteString(" ")
		}
	}

	result.WriteString(" ]")
	return result.String()
}

// * Atoi Function

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

// * TypeOf Function
func TypeOf(context *ReplContext, args []*Argument) (value.IVOR, bool, string) {

	if len(args) != 1 {
		return value.DefaultNilValue, false, "La función typeOf solo acepta un argumento"
	}

	argValue := args[0].Value

	// Obtenemos el tipo directamente (ya es el nombre del tipo)
	typeName := argValue.Type()

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
