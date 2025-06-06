package value

import (
	"fmt"
	"strings"
)

// IntValue representa un valor entero
type IntValue struct {
	InternalValue int
}

func (i IntValue) Value() interface{} {
	return i.InternalValue
}

func (i IntValue) Type() string {
	return IVOR_INT
}

func (i IntValue) Copy() IVOR {
	return &IntValue{i.InternalValue}
}

// FloatValue representa un valor flotante
type FloatValue struct {
	InternalValue float64
}

func (f FloatValue) Value() interface{} {
	return f.InternalValue
}

func (f FloatValue) Type() string {
	return IVOR_FLOAT
}

func (f FloatValue) Copy() IVOR {
	return &FloatValue{f.InternalValue}
}

// StringValue representa una cadena
type StringValue struct {
	InternalValue string
}

func (s StringValue) Value() interface{} {
	return s.InternalValue
}

func (s StringValue) Type() string {
	return IVOR_STRING
}

func (s StringValue) Copy() IVOR {
	return &StringValue{s.InternalValue}
}

// BoolValue representa un valor booleano
type BoolValue struct {
	InternalValue bool
}

func (b BoolValue) Value() interface{} {
	return b.InternalValue
}

func (b BoolValue) Type() string {
	return IVOR_BOOL
}

func (b BoolValue) Copy() IVOR {
	return &BoolValue{b.InternalValue}
}

// RuneValue representa un carácter
type RuneValue struct {
	InternalValue rune
}

func (r RuneValue) Value() interface{} {
	return r.InternalValue
}

func (r RuneValue) Type() string {
	return IVOR_CHARACTER
}

func (r RuneValue) Copy() IVOR {
	return &RuneValue{r.InternalValue}
}

// NilValue representa un valor nulo
type NilValue struct{}

func (n NilValue) Value() interface{} {
	return nil
}

func (n NilValue) Type() string {
	return IVOR_NIL
}

func (n NilValue) Copy() IVOR {
	return &NilValue{}
}

// SliceValue representa un slice/array
type SliceValue struct {
	Elements    []IVOR
	ElementType string
}

func (v SliceValue) Value() interface{} {
	return v.Elements
}

func (v SliceValue) Type() string {
	return IVOR_VECTOR
}

func (v SliceValue) Copy() IVOR {
	newElements := make([]IVOR, len(v.Elements))
	for i, elem := range v.Elements {
		newElements[i] = elem.Copy()
	}
	return &SliceValue{
		Elements:    newElements,
		ElementType: v.ElementType,
	}
}

func (v SliceValue) String() string {
	var elements []string
	for _, elem := range v.Elements {
		elements = append(elements, fmt.Sprintf("%v", elem.Value()))
	}
	return "[" + strings.Join(elements, " ") + "]"
}

// StructValue representa una instancia de struct
type StructValue struct {
	StructType string
	Fields     map[string]IVOR
}

func (s StructValue) Value() interface{} {
	return s.Fields
}

func (s StructValue) Type() string {
	return IVOR_OBJECT
}

func (s StructValue) Copy() IVOR {
	newFields := make(map[string]IVOR)
	for key, value := range s.Fields {
		newFields[key] = value.Copy()
	}
	return &StructValue{
		StructType: s.StructType,
		Fields:     newFields,
	}
}

func (s StructValue) String() string {
	var fields []string
	for key, value := range s.Fields {
		fields = append(fields, fmt.Sprintf("%s: %v", key, value.Value()))
	}
	return s.StructType + "{" + strings.Join(fields, ", ") + "}"
}

// FunctionValue representa una función
type FunctionValue struct {
	Name       string
	Parameters []Parameter
	ReturnType string
}

func (f FunctionValue) Value() interface{} {
	return f.Name
}

func (f FunctionValue) Type() string {
	return IVOR_FUNCTION
}

func (f FunctionValue) Copy() IVOR {
	return &FunctionValue{
		Name:       f.Name,
		Parameters: f.Parameters,
		ReturnType: f.ReturnType,
	}
}

// StructTypeValue representa la definición de un tipo struct
type StructTypeValue struct {
	Name   string
	Fields map[string]string // nombre -> tipo
}

func (s StructTypeValue) Value() interface{} {
	return s.Fields
}

func (s StructTypeValue) Type() string {
	return "struct_type"
}

func (s StructTypeValue) Copy() IVOR {
	newFields := make(map[string]string)
	for key, value := range s.Fields {
		newFields[key] = value
	}
	return &StructTypeValue{
		Name:   s.Name,
		Fields: newFields,
	}
}

// Parameter representa un parámetro de función
type Parameter struct {
	Name string
	Type string
}

// Funciones de utilidad para crear valores

func NewIntValue(val int) IVOR {
	return &IntValue{InternalValue: val}
}

func NewFloatValue(val float64) IVOR {
	return &FloatValue{InternalValue: val}
}

func NewStringValue(val string) IVOR {
	return &StringValue{InternalValue: val}
}

func NewBoolValue(val bool) IVOR {
	return &BoolValue{InternalValue: val}
}

func NewRuneValue(val rune) IVOR {
	return &RuneValue{InternalValue: val}
}

func NewNilValue() IVOR {
	return &NilValue{}
}

func NewSliceValue(elementType string) IVOR {
	return &SliceValue{
		Elements:    make([]IVOR, 0),
		ElementType: elementType,
	}
}

func NewStructValue(structType string, fields map[string]IVOR) IVOR {
	return &StructValue{
		StructType: structType,
		Fields:     fields,
	}
}

// GetDefaultValue devuelve el valor por defecto según el tipo
func GetDefaultValue(dataType string) IVOR {
	switch dataType {
	case IVOR_INT:
		return NewIntValue(0)
	case IVOR_FLOAT:
		return NewFloatValue(0.0)
	case IVOR_STRING:
		return NewStringValue("")
	case IVOR_BOOL:
		return NewBoolValue(false)
	case IVOR_CHARACTER:
		return NewRuneValue(0)
	default:
		return NewNilValue()
	}
}
