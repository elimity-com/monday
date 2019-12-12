package pdq

import (
	"fmt"
	"testing"
)

// WARNING: this board will get modified/cleared of all its data!
const testBoardName = "Test Board"

func TestBoards(t *testing.T) {
	board, _, err := c.EnsureBoard(testBoardName)
	if err != nil {
		t.Error(err)
	}

	boards, err := c.GetBoards()
	if err != nil {
		t.Error(err)
	}
	if len(boards) < 1 {
		t.Error(fmt.Errorf("no boards found"))
		return
	}

	get, err := c.GetBoardWithID(board.ID())
	if err != nil {
		t.Error(err)
	}
	if !get.equals(board) {
		t.Errorf("got %v, expected %v", get, board)
	}
}
