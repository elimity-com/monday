package pdq

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	. "github.com/di-wu/monday"
)

type Board struct {
	Id, Name string
}

func (b Board) ID() int {
	id, _ := strconv.Atoi(b.Id)
	return id
}

func (b Board) equals(other Board) bool {
	if b.Id != other.Id || b.Name != other.Name {
		return false
	}
	return true
}

// EnsureBoard creates a public board with the given name if it not already exists.
func (c SimpleClient) EnsureBoard(name string) (Board, bool, error) {
	boards, err := c.GetBoards()
	if err != nil {
		return Board{}, false, err
	}
	var hit bool
	var board Board
	for _, b := range boards {
		if b.Name == name {
			hit = true
			board = b
			break
		}
	}
	if hit {
		return board, false, nil
	}
	board, err = c.CreateBoard(name)
	if err != nil {
		return Board{}, false, err
	}
	return board, true, nil
}

// CreateBoard creates a public board with the given name.
func (c SimpleClient) CreateBoard(name string) (Board, error) {
	resp, err := c.Exec(context.Background(), NewMutationPayload(
		Boards.Create(name, BoardsKindPublic(), []BoardsField{
			BoardsIDField(),
			BoardsNameField(),
		}),
	))
	if err != nil {
		return Board{}, err
	}
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Board{}, err
	}
	var body struct {
		Data struct {
			Board Board `json:"create_board"`
		}
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return Board{}, err
	}
	return body.Data.Board, nil
}

// GetBoardWithID returns the board with given identifier.
func (c SimpleClient) GetBoardWithID(id int) (Board, error) {
	resp, err := c.Exec(context.Background(), NewQueryPayload(
		Boards.List(
			[]BoardsField{
				BoardsIDField(),
				BoardsNameField(),
				BoardsDescriptionField(),
			},
			NewBoardsIDsArgument([]int{id}),
		),
	))
	if err != nil {
		return Board{}, err
	}
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Board{}, err
	}
	var body struct {
		Data struct {
			Boards []Board
		}
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return Board{}, err
	}

	if len(body.Data.Boards) != 1 {
		return Board{}, fmt.Errorf("no boards returned for id %d: %s", id, string(raw))
	}
	return body.Data.Boards[0], nil
}

// GetBoards returns all the boards.
func (c SimpleClient) GetBoards() ([]Board, error) {
	resp, err := c.Exec(context.Background(), NewQueryPayload(
		Boards.List(
			[]BoardsField{
				BoardsIDField(),
				BoardsNameField(),
				BoardsDescriptionField(),
			},
		),
	))
	if err != nil {
		return nil, err
	}
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var body struct {
		Data struct {
			Boards []Board
		}
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return nil, err
	}
	return body.Data.Boards, nil
}
