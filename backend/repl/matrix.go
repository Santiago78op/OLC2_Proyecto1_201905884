package repl

import "main.go/value"

type MatrixValue struct {
	Items    [][]value.IVOR // bidimensional puro
	ItemType string
	FullType string
}

func NewMatrixValue(items [][]value.IVOR, fullType string, itemType string) *MatrixValue {
	return &MatrixValue{
		Items:    items,
		ItemType: itemType,
		FullType: fullType,
	}
}

func (v MatrixValue) Value() interface{} {
	return v.Items
}

func (v MatrixValue) Type() string {
	return v.FullType
}

func (v MatrixValue) Copy() value.IVOR {
	copyItems := make([][]value.IVOR, len(v.Items))
	for i := range v.Items {
		copyItems[i] = make([]value.IVOR, len(v.Items[i]))
		for j := range v.Items[i] {
			copyItems[i][j] = v.Items[i][j].Copy()
		}
	}

	return NewMatrixValue(copyItems, v.FullType, v.ItemType)
}

func (v *MatrixValue) Set(index []int, value value.IVOR) bool {

	return false
}

func removeBuiltinsFromVector(vectorItems []value.IVOR) {
	for i := 0; i < len(vectorItems); i++ {
		if item, ok := vectorItems[i].(*VectorValue); ok {
			item.ObjectValue.InternalScope.Reset()
			// removeBuiltinsFromVector(item.InternalValue)
		} else {
			break
		}
	}
}

type MatrixItemReference struct {
	Matrix *MatrixValue
	Index  []int
	Value  value.IVOR
}
