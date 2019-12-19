package pdq

import (
	"fmt"
	"testing"
)

// WARNING: this item will get modified/deleted!
const testItemName = "Test Item"

func TestItems(t *testing.T) {
	board, _, _ := c.EnsureBoard(testBoardName)
	group, _, _ := c.EnsureGroup(board.ID(), testGroupName)
	item, _, err := c.EnsureItem(board.ID(), group.Id, testItemName)
	if err != nil {
		t.Error(err)
	}

	items, err := c.GetItems(board.ID(), group.Id)
	if err != nil {
		t.Error(err)
	}
	if len(items) < 1 {
		t.Error(fmt.Errorf("no items found"))
		return
	}

	get, err := c.GetItemWithID(board.ID(), group.Id, item.ID())
	if err != nil {
		t.Error(err)
	}
	if !get.equals(item) {
		t.Errorf("got %v, expected %v", get, group)
	}
}
