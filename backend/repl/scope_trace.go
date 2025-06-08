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
	name       string                 // Nombre del ámbito
	parent     *BaseScopeTrace        // Ámbito padre
	children   []*BaseScopeTrace      // Ámbitos hijos
	variables  map[string]*Variable   // Variables en el ámbito
	functions  map[string]*value.IVOR // Funciones en el ámbito
	IsMutating bool                   // Indica si el ámbito actual está en modo de mutación
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
func (s *BaseScopeTrace) Children() []*BaseScopeTrace {
	return s.children
}

// Valida si el tipo de dato es válido para el ámbito actual.
func (s *BaseScopeTrace) ValidType(_type string) bool {
	// Falta hacer la validación de Structs

	// Verifica si el tipo es válido, aca estan inmersos los vectores
	return value.IsPrimitiveType(_type)
}

func (s *BaseScopeTrace) AddChild(child *BaseScopeTrace) {
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

// GetVariable busca una variable por su nombre en el ámbito actual.
// Si la variable existe, retorna un puntero a la variable y true.
func (s *BaseScopeTrace) GetVariable(name string) *Variable {

	// Si el nombre contiene un punto, se asume que es una variable de objeto
	/*
		if strings.Contains(name, ".") {
			return s.searchObjectVariable(name, nil)
		}
	*/
	// Busca la variable en el ámbito actual
	initScope := s

	for {
		if variable, exists := initScope.variables[name]; exists {

			/*
				if variable.Type == value.IVOR_POINTER {
					// Si la variable es un puntero, retorna el valor apuntado
					return variable.Value.(*PointerValue).AssocVariable // Retorna el valor apuntado por el puntero
				}
			*/
			// Si la variable existe, retorna un puntero a la variable
			return variable
		}

		if initScope.parent == nil {
			break // Si no hay un ámbito padre, termina la búsqueda
		}

		// Si no se encuentra la variable, sube al ámbito padre
		initScope = initScope.parent
	}

	// Si no se encuentra la variable, retorna nil
	return nil
}

/*
// searchObjectVariable busca una variable dentro de un objeto en el ámbito actual.
func (s *BaseScopeTrace) searchObjectVariable(name string, lastObj value.IVOR) *Variable {
	// Divide el nombre de la variable en partes usando el punto como separador
	parts := strings.Split(name, ".")

	// Si no hay partes, retorna un error
	if len(parts) == 0 {
		log.Fatal("Error: No se puede buscar una variable sin partes en el nombre")
		return nil // Si no hay partes, retorna nil
	}

	// Si hay una sola parte, busca la variable en el ámbito actual
	if len(parts) == 1 {
		object, exist := lastObj.(*ObjectValue)

		if exist {
			return object.InternalScope.GetVariable(name)
		}

		log.Fatal("Error: No se puede buscar una variable sin un objeto asociado")
		return nil // Si no hay un objeto asociado, retorna nil
	}

	// Si hay más de una parte.

	if lastObj == nil {
		// Si no hay un objeto asociado, busca la variable en el ámbito actual
		variable := s.GetVariable(parts[0])

		// Si la variable no existe, retorna nil
		if variable == nil {
			return nil // Si no hay un objeto asociado, retorna nil
		}

		// Asigna el valor de la variable como el objeto actual
		object := variable.Value

		switch object := object.(type) {
		case *ObjectValue:
			lastObj = object
		case *VectorValue:
			lastObj = object.ObjectValue
			default:
				return nil // Si el objeto no es un ObjectValue o VectorValue, retorna nil
		}

		return s.searchObjectVariable(strings.Join(parts[1:], "."), lastObj)
	}

	object, exist := lastObj.(*ObjectValue)

	if exist {
		lastObj = object.InternalScope.GetVariable(parts[0]).Value

		return s.searchObjectVariable(strings.Join(parts[1:], "."), lastObj)
	} else {
		log.Fatal("Error: No se puede buscar una variable sin un objeto asociado")
		return nil // Si no hay un objeto asociado, retorna nil
	}
}
*/

// Reset reinicializa el ámbito actual, eliminando todas las variables y funciones definidas en él.
func (s *BaseScopeTrace) Reset() {
	s.variables = make(map[string]*Variable)   // Reinicializa el mapa de variables
	s.functions = make(map[string]*value.IVOR) // Reinicializa el mapa de funciones
	s.children = make([]*BaseScopeTrace, 0)    // Reinicializa los hijos del ámbito
}

// IsMutatingScope verifica si el ámbito actual o alguno de sus padres está en modo de mutación.
func (s *BaseScopeTrace) IsMutatingScope() bool {
	temp := s

	for {
		if temp.IsMutating {
			return true
		}

		if temp.parent == nil {
			break
		}

		temp = temp.parent
	}

	return false
}

// NewGlobalScopeTrace crea un nuevo ámbito global para el REPL.
// Este ámbito global es el punto de partida para todas las ejecuciones en el REPL.
// Se inicializa con un nombre específico y un mapa vacío de variables y funciones.
// Este ámbito es utilizado para almacenar variables y funciones globales que pueden ser accedidas desde cualquier parte del REPL.
func NewGlobalScopeTrace() *BaseScopeTrace {

	// Falta registra la contruccion de funciones

	// Crea un nuevo ámbito global con un nombre específico
	return &BaseScopeTrace{
		name:      "Global",
		parent:    nil,
		children:  make([]*BaseScopeTrace, 0),
		variables: make(map[string]*Variable),
	}
}

// NewLocalScopeTrace crea un nuevo ámbito local para el REPL.
// Este ámbito local es utilizado para almacenar variables y funciones que son
// específicas de una ejecución o contexto particular dentro del REPL.
// Se inicializa con un nombre específico y un mapa vacío de variables y funciones.
func NewLocalScopeTrace(name string) *BaseScopeTrace {
	return &BaseScopeTrace{
		name:      name,
		parent:    nil,
		children:  make([]*BaseScopeTrace, 0),
		variables: make(map[string]*Variable),
		functions: make(map[string]*value.IVOR),
	}
}

// ScopeTrace representa la traza de ejecución del REPL, que incluye el ámbito global y el ámbito local.
type ScopeTrace struct {
	GlobalScope *BaseScopeTrace // Ámbito global del REPL
	LocalScope  *BaseScopeTrace // Ámbito local del REPL
}

// PushScope crea un nuevo ámbito local dentro de la traza de ejecución del REPL.
func (s *ScopeTrace) PushScope(name string) *BaseScopeTrace {

	newScope := NewLocalScopeTrace(name)
	s.LocalScope.AddChild(newScope)
	s.LocalScope = newScope

	return s.LocalScope
}

// PopScope elimina el ámbito local actual de la traza de ejecución del REPL,
func (s *ScopeTrace) PopScope() {
	s.LocalScope = s.LocalScope.Parent()
}

// Reset reinicializa el ámbito local actual, estableciendo el ámbito local al ámbito global.
func (s *ScopeTrace) Reset() {
	s.LocalScope = s.GlobalScope
}

// AddVariable agrega una nueva variable al ámbito local actual de la traza de ejecución del REPL.
func (s *ScopeTrace) AddVariable(name string, varType string, value value.IVOR, isConst bool, allowNil bool, token antlr.Token) (*Variable, string) {
	return s.LocalScope.AddVariable(name, varType, value, isConst, allowNil, token)
}

// GetVariable busca una variable por su nombre en el ámbito local actual de la traza de ejecución del REPL.
func (s *ScopeTrace) GetVariable(name string) *Variable {
	return s.LocalScope.GetVariable(name)
}

// IsMutatingEnvironment verifica si el ámbito local actual o alguno de sus padres está en modo de mutación.
func (s *ScopeTrace) IsMutatingEnvironment() bool {
	return s.LocalScope.IsMutatingScope()
}

// NewScopeTrace crea una nueva traza de ejecución del REPL con un ámbito global y un ámbito local inicial.
// Este ámbito local es el punto de partida para todas las ejecuciones en el REPL.
func NewScopeTrace() *ScopeTrace {
	globalScope := NewGlobalScopeTrace()
	return &ScopeTrace{
		GlobalScope: globalScope,
		LocalScope:  globalScope,
	}
}
