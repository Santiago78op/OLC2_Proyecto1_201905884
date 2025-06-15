package repl

import (
	"github.com/antlr4-go/antlr/v4"
	"main.go/value"
)

// Variable representa una variable en el entorno REPL.
type Variable struct {
	Name     string      // Nombre de la variable
	Value    value.IVOR  // Valor de la variable
	Type     string      // Tipo de la variable
	IsConst  bool        // Indica si la variable es constante
	AllowNil bool        // Indica si la variable permite valores nulos
	Token    antlr.Token // Token asociado a la variable
	isProp   bool        // Indica si la variable es una propiedad
}

func (v *Variable) TypeValidation() (bool, string) {

	// Verifica sel valor sea igual a default uninitialized value
	// y si es asi, retorna true y un mensaje indicando que la variable no ha sido inicializada
	// Ejemplo: mut x int;
	if v.Value == value.DefaultUnInitializedValue {
		v.Value = value.DefaultValue(v.Type, v.Value)
		return true, ""
	}

	// Verifica si el valor de la variable es nulo
	// y si es asi, retorna true y un mensaje indicando que la variable es nula
	// Ejemplo: mut x = nil;
	if v.Value == value.DefaultNilValue {
		if v.AllowNil {
			return true, ""
		}
	}

	if v.Type != v.Value.Type() {

		// Falta validar Vectores
		// vector type validation
		if IsVectorType(v.Type) && IsVectorType(v.Value.Type()) {

			// empty vector cast
			if v.Value.Type() == "[]" {

				// modify vector type
				v.Value.(*VectorValue).ItemType = RemoveBrackets(v.Type)
				v.Value.(*VectorValue).FullType = v.Type

				return true, ""
			}

			// implicit vector conversion

			targetType := RemoveBrackets(v.Type) // inner type
			newConvertedItems := make([]value.IVOR, 0)

			for _, item := range v.Value.(*VectorValue).InternalValue {
				convertedValue, ok := value.ImplicitCast(targetType, item)

				if !ok {
					break
				}
				newConvertedItems = append(newConvertedItems, convertedValue)
			}

			if len(newConvertedItems) == len(v.Value.(*VectorValue).InternalValue) {
				return true, ""
			}

			msg := "No se puede asignar un vector de tipo " + v.Value.Type() + " a una vector de tipo " + v.Type
			v.Value = value.DefaultNilValue
			return false, msg
		}

		// Trata de hacer una conversión implícita del valor al tipo de la variable
		convertedValue, ok := value.ImplicitCast(v.Type, v.Value)

		if !ok {
			// Si la conversión falla, retorna false y un mensaje de error
			// La variable obtine el valor de nil a no saber que tipo convertir
			msg := "Type mismatch: No se puede asignar un valor al tipo: " + v.Value.Type() + " a una variable de tipo: " + v.Type
			v.Value = value.DefaultNilValue
			return false, msg
		}

		// Si la conversión es exitosa, actualiza el valor de la variable
		v.Value = convertedValue
	}

	// Si todo está bien, retorna True y un mensaje de éxito
	return true, ""
}

func (v *Variable) AssignValue(val value.IVOR, isMutatingContext bool) (bool, string) {
	// Si la variable es constante y se intenta modificar, retorna un error
	if v.IsConst {
		msg := "No se puede asignar un valor a una variable constante: " + v.Name
		return false, msg
	}

	// Si la variable es una propiedad, no se puede asignar un valor directamente
	if v.isProp && !isMutatingContext {
		msg := "No se puede asignar valor a una propiedad fuera de contexto mutable: " + v.Name
		return false, msg
	}

	// Asigna el valor a la variable
	v.Value = val

	// Si la validación de tipo es exitosa, retorna v.Value.Type() == v.Type, "Variable assigned successfully"
	return v.TypeValidation()
}
