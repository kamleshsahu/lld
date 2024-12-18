package entity

type Field struct {
	Cells []Cell
}

func (f *Field) Copy() Field {
	fieldCopy := Field{
		Cells: make([]Cell, len(f.Cells)),
	}
	for i, cell := range f.Cells {
		fieldCopy.Cells[i] = cell.Copy()
	}
	return fieldCopy
}
