package pdq

import (
	"testing"

	. "github.com/di-wu/monday"
)

func TestColumnValue(t *testing.T) {
	board, _, _ := c.EnsureBoard(testBoardName)
	group, _, _ := c.EnsureGroup(board.ID(), testGroupName)
	item, _, _ := c.EnsureItem(board.ID(), group.Id, testItemName)
	columns, _ := c.GetColumns(board.ID())

	var columnID string
	for _, c := range columns {
		if c.Title == testColumnTitle {
			columnID = c.Id
			break
		}
	}

	if err := c.EnsureColumnValue(board.ID(), item.ID(), NewStatusLabelValue(columnID, "Stuck")); err != nil {
		t.Error(err)
	}

	_, err := c.GetItemColumnValues(item.ID())
	if err != nil {
		t.Error()
	}
}
