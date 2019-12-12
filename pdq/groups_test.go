package pdq

import (
	"fmt"
	"testing"
)

// WARNING: this group will get modified/deleted!
const testGroupName = "Test Group"

func TestGroups(t *testing.T) {
	board, _, _ := c.EnsureBoard(testBoardName)
	group, _, err := c.EnsureGroup(board.ID(), testGroupName)
	if err != nil {
		t.Error(err)
	}

	groups, err := c.GetGroups(board.ID())
	if err != nil {
		t.Error(err)
	}
	if len(groups) < 1 {
		t.Error(fmt.Errorf("no groups found"))
		return
	}

	get, err := c.GetGroup(board.ID(), group.Id)
	if err != nil {
		t.Error(err)
	}
	if !get.equals(group) {
		t.Errorf("got %v, expected %v", get, group)
	}
}
