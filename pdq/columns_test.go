package pdq

import (
	"fmt"
	"testing"

	. "github.com/di-wu/monday"
)

const testColumnTitle = "Test Title"

func TestColumns(t *testing.T) {
	board, _, _ := c.EnsureBoard(testBoardName)
	column, _, err := c.EnsureColumn(board.ID(), testColumnTitle, ColumnsTypeStatus())
	if err != nil {
		t.Error(err)
	}

	columns, err := c.GetColumns(board.ID())
	if err != nil {
		t.Error(err)
	}
	if len(columns) < 1 {
		t.Error(fmt.Errorf("no columns found"))
		return
	}

	get, err := c.GetColumnWithID(board.ID(), column.Id)
	if err != nil {
		t.Error(err)
	}
	if !get.equals(column) {
		t.Errorf("got %v, expected %v", get, column)
	}
}
