package repl

import (
	"fmt"
	"main.go/value"
)

// Tipos de símbolos
const (
	SYMBOL_VARIABLE  = "Variable"
	SYMBOL_FUNCTION  = "Función"
	SYMBOL_STRUCT    = "Struct"
	SYMBOL_PARAMETER = "Parámetro"
)

// Símbolo en la tabla
type Symbol struct {
	ID        string
	Type      string     // Tipo del símbolo (Variable, Función, etc.)
	DataType  string     // Tipo de dato (int, string, etc.)
	Value     value.IVOR // Valor actual
	Line      int        // Línea donde se declaró
	Column    int        // Columna donde se declaró
	Scope     string     // Ámbito donde se declaró
	IsMutable bool       // Si es mutable (para variables)
}

// Scope representa un ámbito
type Scope struct {
	Name    string
	Parent  *Scope
	Symbols map[string]*Symbol
}

// ScopeTrace maneja la pila de ámbitos
type ScopeTrace struct {
	CurrentScope *Scope
	GlobalScope  *Scope
}

// NewScope crea un nuevo ámbito
func NewScope(name string, parent *Scope) *Scope {
	return &Scope{
		Name:    name,
		Parent:  parent,
		Symbols: make(map[string]*Symbol),
	}
}

// NewScopeTrace crea un nuevo trazador de ámbitos
func NewScopeTrace() *ScopeTrace {
	global := NewScope("Global", nil)
	return &ScopeTrace{
		CurrentScope: global,
		GlobalScope:  global,
	}
}

// PushScope entra a un nuevo ámbito
func (st *ScopeTrace) PushScope(name string) {
	newScope := NewScope(name, st.CurrentScope)
	st.CurrentScope = newScope
}

// PopScope sale del ámbito actual
func (st *ScopeTrace) PopScope() {
	if st.CurrentScope.Parent != nil {
		st.CurrentScope = st.CurrentScope.Parent
	}
}

// SetVariable declara o actualiza una variable
func (st *ScopeTrace) SetVariable(id string, dataType string, val value.IVOR, line int, column int, isMutable bool) error {
	// Verificar si ya existe en el ámbito actual
	if _, exists := st.CurrentScope.Symbols[id]; exists {
		return fmt.Errorf("la variable '%s' ya existe en este ámbito", id)
	}

	symbol := &Symbol{
		ID:        id,
		Type:      SYMBOL_VARIABLE,
		DataType:  dataType,
		Value:     val,
		Line:      line,
		Column:    column,
		Scope:     st.CurrentScope.Name,
		IsMutable: isMutable,
	}

	st.CurrentScope.Symbols[id] = symbol
	return nil
}

// GetVariable busca una variable en la cadena de ámbitos
func (st *ScopeTrace) GetVariable(id string) (*Symbol, error) {
	current := st.CurrentScope
	for current != nil {
		if symbol, exists := current.Symbols[id]; exists && symbol.Type == SYMBOL_VARIABLE {
			return symbol, nil
		}
		current = current.Parent
	}
	return nil, fmt.Errorf("la variable '%s' no está definida", id)
}

// UpdateVariable actualiza el valor de una variable existente
func (st *ScopeTrace) UpdateVariable(id string, val value.IVOR) error {
	symbol, err := st.GetVariable(id)
	if err != nil {
		return err
	}

	if !symbol.IsMutable {
		return fmt.Errorf("la variable '%s' no es mutable", id)
	}

	// Verificar tipos compatibles
	if symbol.Value.Type() != val.Type() {
		// Permitir conversión int/float
		if (symbol.Value.Type() == value.IVOR_INT && val.Type() == value.IVOR_FLOAT) ||
			(symbol.Value.Type() == value.IVOR_FLOAT && val.Type() == value.IVOR_INT) {
			// Conversión automática
			if symbol.Value.Type() == value.IVOR_INT && val.Type() == value.IVOR_FLOAT {
				// Convertir float a int (truncar)
				floatVal := val.(*value.FloatValue).InternalValue
				val = value.NewIntValue(int(floatVal))
			} else if symbol.Value.Type() == value.IVOR_FLOAT && val.Type() == value.IVOR_INT {
				// Convertir int a float
				intVal := val.(*value.IntValue).InternalValue
				val = value.NewFloatValue(float64(intVal))
			}
		} else {
			return fmt.Errorf("no se puede asignar tipo '%s' a variable de tipo '%s'", val.Type(), symbol.Value.Type())
		}
	}

	symbol.Value = val
	return nil
}

// SetFunction declara una función
func (st *ScopeTrace) SetFunction(id string, parameters []Parameter, returnType string, line int, column int) error {
	// Las funciones se declaran siempre en el ámbito global
	if _, exists := st.GlobalScope.Symbols[id]; exists {
		return fmt.Errorf("la función '%s' ya existe", id)
	}

	// Convertir parámetros a value.Parameter
	var valueParams []value.Parameter
	for _, p := range parameters {
		valueParams = append(valueParams, value.Parameter{
			Name: p.Name,
			Type: p.Type,
		})
	}
	funcValue := &value.FunctionValue{
		Name:       id,
		Parameters: valueParams,
		ReturnType: returnType,
	}

	symbol := &Symbol{
		ID:       id,
		Type:     SYMBOL_FUNCTION,
		DataType: "function",
		Value:    funcValue,
		Line:     line,
		Column:   column,
		Scope:    "Global",
	}

	st.GlobalScope.Symbols[id] = symbol
	return nil
}

// GetFunction busca una función
func (st *ScopeTrace) GetFunction(id string) (*Symbol, error) {
	if symbol, exists := st.GlobalScope.Symbols[id]; exists && symbol.Type == SYMBOL_FUNCTION {
		return symbol, nil
	}
	return nil, fmt.Errorf("la función '%s' no está definida", id)
}

// SetStruct declara una estructura
func (st *ScopeTrace) SetStruct(id string, fields map[string]string, line int, column int) error {
	// Los structs se declaran siempre en el ámbito global
	if _, exists := st.GlobalScope.Symbols[id]; exists {
		return fmt.Errorf("el struct '%s' ya existe", id)
	}

	// Crear un valor de struct tipo
	structValue := &value.StructTypeValue{
		Name:   id,
		Fields: fields,
	}

	symbol := &Symbol{
		ID:       id,
		Type:     SYMBOL_STRUCT,
		DataType: "struct_type",
		Value:    structValue,
		Line:     line,
		Column:   column,
		Scope:    "Global",
	}

	st.GlobalScope.Symbols[id] = symbol
	return nil
}

// GetStruct busca un struct
func (st *ScopeTrace) GetStruct(id string) (*Symbol, error) {
	if symbol, exists := st.GlobalScope.Symbols[id]; exists && symbol.Type == SYMBOL_STRUCT {
		return symbol, nil
	}
	return nil, fmt.Errorf("el struct '%s' no está definido", id)
}

// GetAllSymbols devuelve todos los símbolos para el reporte
func (st *ScopeTrace) GetAllSymbols() []*Symbol {
	var symbols []*Symbol

	// Función recursiva para recolectar símbolos
	var collectSymbols func(*Scope)
	collectSymbols = func(scope *Scope) {
		for _, symbol := range scope.Symbols {
			symbols = append(symbols, symbol)
		}
	}

	// Recolectar de todos los ámbitos accesibles
	current := st.CurrentScope
	for current != nil {
		collectSymbols(current)
		current = current.Parent
	}

	return symbols
}

// Parameter representa un parámetro de función
type Parameter struct {
	Name string
	Type string
}
