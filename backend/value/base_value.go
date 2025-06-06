package value

const (
	IVOR_INT              = "int"
	IVOR_FLOAT            = "float64"
	IVOR_STRING           = "string"
	IVOR_BOOL             = "bool"
	IVOR_CHARACTER        = "rune"
	IVOR_NIL              = "nil"
	IVOR_BUILTIN_FUNCTION = "builtin_function"
	IVOR_FUNCTION         = "function"
	IVOR_VECTOR           = "vector"
	IVOR_OBJECT           = "object"
	IVOR_ANY              = "any"
	IVOR_POINTER          = "pointer"
	IVOR_MATRIX           = "matrix"
	IVOR_SELF             = "self"
	IVOR_UNINITIALIZED    = "uninitialized"
)

// IVOR Es la Representación Interna de Objeto de Valor
type IVOR interface {
	Value() interface{}
	Type() string
	Copy() IVOR
}
