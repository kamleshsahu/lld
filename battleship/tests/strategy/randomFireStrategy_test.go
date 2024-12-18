package strategy

import (
	"lld/battleship/entity"
	"lld/battleship/strategy/fireStrategy"
	"testing"
)

func TestRandomFireStrategy_GetFireLocation(t *testing.T) {
	// Setup mock data
	field := &entity.Field{
		Cells: []entity.Cell{
			*entity.NewCell(0, 0),
			*entity.NewCell(1, 1),
			*entity.NewCell(2, 2),
			*entity.NewCell(4, 4),
			*entity.NewCell(5, 5),
			*entity.NewCell(6, 6),
			*entity.NewCell(7, 7),
		},
	}

	strategy := fireStrategy.NewRandomFireStrategy()
	strategy.Init([]*entity.Field{field})

	n := len(field.Cells)
	usedCells := make(map[string]bool)
	// Call GetFireLocation multiple times and ensure correct behavior
	for i := 0; i < n; i++ {
		cell, err := strategy.GetFireLocation(0)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Check if the returned cell is valid
		if cell == nil {
			t.Errorf("Expected a valid cell, got nil")
		}
		t.Log("generated cell: ", cell.ToString())

		if usedCells[cell.ToString()] {
			t.Errorf("Expected a unused cell id")
		}
		usedCells[cell.ToString()] = true
	}

	// should throw error, as not cell left
	_, err := strategy.GetFireLocation(0)
	expected := entity.ErrNoCellLeft(0)
	if err.Error() != expected.Error() {
		t.Errorf("Expected error %v, got %v", entity.ErrNoCellLeft(0), err)
	}

}
