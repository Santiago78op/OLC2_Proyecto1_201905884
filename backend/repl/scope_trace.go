package repl

import (
	"log"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"main.go/value"
)

/*
BaseScope ES una estructura que representa un ámbito base en el REPL (Read-Eval-Print Loop).
// Contiene información sobre el nombre del ámbito, su ámbito padre, sus hijos, variables y funciones.
// Esta estructura es utilizada para gestionar el contexto de ejecución y el alcance de las variables y funciones
// dentro del REPL, permitiendo la creación de ámbitos anidados y la resolución de nombres de variables y funciones.
*/
type BaseScope struct {
	name      string
	parent    *BaseScope
	children  []*BaseScope
	variables map[string]*Variable
	functions map[string]value.IVOR
}

// Name es un método que devuelve el nombre del ámbito actual.
func (s *BaseScope) Name() string {
	return s.name
}

// Parent es un método que devuelve el ámbito padre del ámbito actual.
func (s *BaseScope) Parent() *BaseScope {
	return s.parent
}

// Children es un método que devuelve los hijos del ámbito actual.
func (s *BaseScope) Children() []*BaseScope {
	return s.children
}

// ValidType es un método que verifica si un tipo es válido dentro del ámbito actual.
func (s *BaseScope) ValidType(_type string) bool {
	return value.IsPrimitiveType(_type)
}

// addChild es un método que agrega un hijo al ámbito actual.
func (s *BaseScope) AddChild(child *BaseScope) {
	s.children = append(s.children, child)
	child.parent = s
}

// variableExists es un método que verifica si una variable ya existe en el ámbito actual.
func (s *BaseScope) variableExists(variable *Variable) bool {

	if _, ok := s.variables[variable.Name]; ok {
		return true
	}

	return false

}

// AddVariable es un método que agrega una variable al ámbito actual.
func (s *BaseScope) AddVariable(name string, varType string, value value.IVOR, isConst bool, allowNil bool, token antlr.Token) (*Variable, string) {

	variable := &Variable{
		Name:     name,
		Value:    value,
		Type:     varType,
		IsConst:  isConst,
		AllowNil: allowNil,
		Token:    token,
	}

	if s.variableExists(variable) {
		return nil, "La variable " + name + " ya existe"
	}

	typesOk, msg := variable.TypeValidation()

	// even if the variable is not valid, we add it to the scope, (internally it will be nil)
	s.variables[name] = variable

	if !typesOk {
		// report error
		return nil, msg
	}

	return variable, ""
}

// GetVariable es un método que busca una variable por su nombre en el ámbito actual y sus ámbitos padres.
func (s *BaseScope) GetVariable(name string) *Variable {
	// verify if is refering to and object/struct function
	if strings.Contains(name, ".") {
		return s.searchObjectVariable(name, nil)
	}

	initialScope := s

	for {
		if variable, ok := initialScope.variables[name]; ok {

			// verify if is refering to a pointer
			if variable.Type == value.IVOR_POINTER {
				return variable.Value.(*PointerValue).AssocVariable // pointer of a pointer ?
			}

			return variable
		}

		if initialScope.parent == nil {
			break
		}

		initialScope = initialScope.parent
	}

	return nil
}

// GetFunction es un método que busca una función por su nombre en el ámbito actual y sus ámbitos padres.
func (s *BaseScope) searchObjectVariable(name string, lastObj value.IVOR) *Variable {

	// split name by dot
	parts := strings.Split(name, ".")

	if len(parts) == 0 {
		log.Fatal("idk what u did, cant split by dot")
		return nil
	}

	if len(parts) == 1 {
		obj, ok := lastObj.(*ObjectValue)

		if ok {
			return obj.InternalScope.GetVariable(name)
		}

		log.Fatal("idk what u did, cant convert to object")
		return nil
	}

	// then parts should be 2 or more

	if lastObj == nil {
		variable := s.GetVariable(parts[0])

		if variable == nil {
			return nil
		}

		obj := variable.Value

		// obj must be an object/struct or vector
		switch obj := obj.(type) {
		case *ObjectValue:
			lastObj = obj
		case *VectorValue:
			lastObj = obj.ObjectValue
		default:
			return nil
		}

		return s.searchObjectVariable(strings.Join(parts[1:], "."), lastObj)
	}

	obj, ok := lastObj.(*ObjectValue)

	if ok {
		lastObj = obj.InternalScope.GetVariable(parts[0]).Value

		return s.searchObjectVariable(strings.Join(parts[1:], "."), lastObj)
	} else {
		log.Fatal("idk what u did, cant convert to object")
		return nil
	}
}

func (s *BaseScope) AddFunction(name string, function value.IVOR) (bool, string) {
	// check if function already exists

	if _, ok := s.functions[name]; ok {
		return false, "La funcion " + name + " ya existe"
	}

	s.functions[name] = function

	return true, ""
}

func (s *BaseScope) GetFunction(name string) (value.IVOR, string) {

	// verify if is refering to and object/struct function
	if strings.Contains(name, ".") {
		return s.searchObjectFunction(name, nil)
	}

	initialScope := s

	for {
		if function, ok := initialScope.functions[name]; ok {
			return function, ""
		}

		if initialScope.parent == nil {
			break
		}

		initialScope = initialScope.parent
	}

	return nil, "La funcion " + name + " no existe"
}

// obj1.obj2.func1()

func (s *BaseScope) searchObjectFunction(name string, lastObj value.IVOR) (value.IVOR, string) {

	// split name by dot
	parts := strings.Split(name, ".")

	if len(parts) == 0 {
		log.Fatal("idk what u did, cant split by dot")
		return nil, ""
	}

	if len(parts) == 1 {
		obj, ok := lastObj.(*ObjectValue)

		if ok {
			return obj.InternalScope.GetFunction(name)
		}

		log.Fatal("idk what u did, cant convert to object")
		return nil, ""
	}
}

// Reset es un método que reinicia el ámbito actual, eliminando todas las variables, hijos y funciones.
func (s *BaseScope) Reset() {
	s.variables = make(map[string]*Variable)
	s.children = make([]*BaseScope, 0)
	s.functions = make(map[string]value.IVOR)
}

func NewGlobalScope() *BaseScope {

	// register built-in functions

	funcs := make(map[string]value.IVOR)

	for k, v := range DefaultBuiltInFunctions {
		funcs[k] = v
	}

	return &BaseScope{
		name:      "global",
		variables: make(map[string]*Variable),
		children:  make([]*BaseScope, 0),
		structs:   make(map[string]*Struct),
		parent:    nil,
		functions: funcs,
	}
}

func NewLocalScope(name string) *BaseScope {
	return &BaseScope{
		name:      name,
		variables: make(map[string]*Variable),
		functions: make(map[string]value.IVOR),
		children:  make([]*BaseScope, 0),
		parent:    nil,
	}
}

type ScopeTrace struct {
	GlobalScope  *BaseScope
	CurrentScope *BaseScope
}

func (s *ScopeTrace) PushScope(name string) *BaseScope {

	newScope := NewLocalScope(name)
	s.CurrentScope.AddChild(newScope)
	s.CurrentScope = newScope

	return s.CurrentScope
}

func (s *ScopeTrace) PopScope() {
	s.CurrentScope = s.CurrentScope.Parent()
}

func (s *ScopeTrace) Reset() {
	s.CurrentScope = s.GlobalScope
}

func (s *ScopeTrace) AddVariable(name string, varType string, value value.IVOR, isConst bool, allowNil bool, token antlr.Token) (*Variable, string) {
	return s.CurrentScope.AddVariable(name, varType, value, isConst, allowNil, token)
}

func (s *ScopeTrace) GetVariable(name string) *Variable {
	return s.CurrentScope.GetVariable(name)
}

func (s *ScopeTrace) AddFunction(name string, function value.IVOR) (bool, string) {
	return s.CurrentScope.AddFunction(name, function)
}

func (s *ScopeTrace) GetFunction(name string) (value.IVOR, string) {
	return s.CurrentScope.GetFunction(name)
}
