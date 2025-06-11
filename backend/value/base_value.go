package value

// Tipos de Datos IVOR (Internal Value Object Representation)
// Representa los tipos de datos que pueden ser utilizados en el lenguaje
const (
	IVOR_INT              = "Int"
	IVOR_FLOAT            = "Float64"
	IVOR_STRING           = "String"
	IVOR_BOOL             = "Bool"
	IVOR_CHARACTER        = "Rune"
	IVOR_NIL              = "Nil"
	IVOR_BUILTIN_FUNCTION = "BuiltinFunction"
	IVOR_FUNCTION         = "Function"
	IVOR_VECTOR           = "Vector"
	IVOR_OBJECT           = "Object"
	IVOR_ANY              = "Any"
	IVOR_POINTER          = "Pointer"
	IVOR_MATRIX           = "Matrix"
	IVOR_SELF             = "Self"
	IVOR_UNINITIALIZED    = "Uninitialized"
)

// IVOR Es la Representación Interna de Objeto de Valor
// Representa un valor en el lenguaje, puede ser un número, cadena, booleano, etc.
type IVOR interface {
	Value() interface{}
	Type() string
	Copy() IVOR
}
