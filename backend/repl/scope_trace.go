package repl

import (
	"github.com/antlr4-go/antlr/v4"
	"main.go/value"
)

// BaseScopeTrace representa un ámbito en la traza de ejecución del REPL.
// Contiene información sobre el nombre del ámbito, su padre, sus hijos,
// las variables y funciones definidas en él. Esta estructura es útil para
// rastrear el contexto de ejecución y la visibilidad de las variables y funciones
// en diferentes niveles de anidamiento dentro del REPL.
type BaseScopeTrace struct {
	name      string                 // Nombre del ámbito
	parent    *BaseScopeTrace        // Ámbito padre
	children  []BaseScopeTrace       // Ámbitos hijos
	variables map[string]*Variable   // Variables en el ámbito
	functions map[string]*value.IVOR // Funciones en el ámbito
}

// Name devuelve el nombre del ámbito actual.
func (s *BaseScopeTrace) Name() string {
	return s.name
}

// Parent devuelve el ámbito padre del ámbito actual.
func (s *BaseScopeTrace) Parent() *BaseScopeTrace {
	return s.parent
}

// Children devuelve los ámbitos hijos del ámbito actual.
func (s *BaseScopeTrace) Children() []BaseScopeTrace {
	return s.children
}

// Valida si el tipo de dato es válido para el ámbito actual.
func (s *BaseScopeTrace) ValidType(_type string) bool {
	// Falta hacer la validación de Structs

	// Verifica si el tipo es válido, aca estan inmersos los vectores
	return value.IsPrimitiveType(_type)
}

func (s *BaseScopeTrace) AddChild(child BaseScopeTrace) {
	// Agrega un ámbito hijo al ámbito actual
	s.children = append(s.children, child)
	// Establece el ámbito padre del hijo como el ámbito actual
	child.parent = s
}

func (s *BaseScopeTrace) variableExists(variable *Variable) bool {
	// Verifica si una variable con el nombre dado ya existe en el ámbito
	if _, exists := s.variables[variable.Name]; exists {
		return true
	}

	// Si no existe, retorna false
	return false
}

func (s *BaseScopeTrace) AddVariable(name string, varType string, value value.IVOR, isConst bool, allowNil bool, token antlr.Token) (*Variable, string) {
	// Crea una nueva variable con los parámetros dados
	variable := &Variable{
		Name:     name,
		Type:     varType,
		Value:    value,
		IsConst:  isConst,
		AllowNil: allowNil,
		Token:    token,
	}

	// Verifica si la variable ya existe en el ámbito
	if s.variableExists(variable) {
		msg := "La variable '" + name + "' ya existe en el ámbito actual"
		return nil, msg
	}

	// Valida el tipo de la variable
	typesExists, msg := variable.TypeValidation()

	// Agrega la variable al mapa de variables del ámbito
	s.variables[name] = variable

	if !typesExists {
		// Si la validación falla, retorna nil y el mensaje de error
		return nil, msg
	}

	// Si la validación es exitosa, agrega la variable al ámbito
	return variable, ""
}
