package repl

import "main.go/value"

type VectorValue struct {
	*ObjectValue
	InternalValue []value.IVOR
	CurrentIndex  int
	ItemType      string
	FullType      string
	SizeValue     *value.IntValue
	IsEmpty       *value.BoolValue
}
