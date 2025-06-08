package repl

import (
	"fmt"
)

/*
Console es una estructura que simula un entorno de consola para mostrar mensajes.
*/
type Console struct {
	output string
}

/*
Print es un método de la estructura Console que agrega un mensaje a la salida.
Este método simula el comportamiento de imprimir en una consola.
*/
func (c *Console) Print(s string) {
	c.output += s + "\n"
}

/*
Show es un método de la estructura Console que muestra el contenido actual de la salida.
*/
func (c *Console) Show() {
	fmt.Println(c.output)
}

/*
Clear es un método de la estructura Console que limpia el contenido de la salida.
*/
func (c *Console) Clear() {
	c.output = ""
}

/*
NewConsole es una función que crea una nueva instancia de Console.
*/
func NewConsole() *Console {
	return &Console{}
}

/*
GetOutput es un método de la estructura Console que devuelve el contenido actual de la salida.
*/
func (c *Console) GetOutput() string {
	return c.output
}
