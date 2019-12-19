package pdq

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	. "github.com/di-wu/monday"
)

type Column struct {
	Id, Title, Type string
	SettingsStr      string `json:"settings_str"`
}

func (c Column) equals(other Column) bool {
	if c.Id != other.Id || c.Title != other.Title {
		return false
	}
	return true
}

// EnsureGroup creates a column with the given title if it not already exists.
func (c SimpleClient) EnsureColumn(boardID int, title string, columnType ColumnsType) (Column, bool, error) {
	columns, err := c.GetColumns(boardID)
	if err != nil {
		return Column{}, false, err
	}
	var hit bool
	var column Column
	for _, v := range columns {
		if v.Title == title {
			hit = true
			column = v
			break
		}
	}
	if hit {
		return column, false, nil
	}
	column, err = c.CreateColumn(boardID, title, columnType)
	if err != nil {
		return Column{}, false, err
	}
	return column, true, nil
}

// EnsureGroup creates a status column with the given title if it not already exists.
func (c SimpleClient) EnsureStatusColumn(boardID int, title string, values []string) (Column, bool, error) {
	columns, err := c.GetColumns(boardID)
	if err != nil {
		return Column{}, false, err
	}
	var hit bool
	var column Column
	for _, v := range columns {
		if v.Title == title {
			hit = true
			column = v
			break
		}
	}
	if hit {
		return column, false, nil
	}
	column, err = c.CreateStatusColumn(boardID, title, values)
	if err != nil {
		return Column{}, false, err
	}
	return column, true, nil
}

// CreateColumn creates a column of the specified type with given title.
func (c SimpleClient) CreateColumn(boardID int, title string, columnType ColumnsType) (Column, error) {
	resp, err := c.Exec(context.Background(), NewMutationPayload(
		CreateColumn(boardID, title, columnType,
			[]ColumnsField{
				ColumnsIDField(),
				ColumnsTitleField(),
				ColumnsTypeField(),
			},
		),
	))
	if err != nil {
		return Column{}, err
	}
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Column{}, err
	}
	var body struct {
		Data struct {
			Column Column `json:"create_column"`
		}
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return Column{}, err
	}
	return body.Data.Column, nil
}

// CreateStatusColumn creates a status column with given title and default values.
func (c SimpleClient) CreateStatusColumn(boardID int, title string, values []string) (Column, error) {
	resp, err := c.Exec(context.Background(), NewMutationPayload(
		CreateColumnWithDefaults(boardID, title, ColumnsTypeStatus(),
			fmt.Sprintf(`{"labels": ["%s"]}`, strings.Join(values, `", "`)),
			[]ColumnsField{
				ColumnsIDField(),
				ColumnsTitleField(),
			},
		),
	))
	if err != nil {
		return Column{}, err
	}
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Column{}, err
	}
	var body struct {
		Data struct {
			Column Column `json:"create_column"`
		}
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return Column{}, err
	}
	return body.Data.Column, nil
}

func (c SimpleClient) GetColumnWithID(boardID int, columnID string) (Column, error) {
	columns, err := c.GetColumns(boardID)
	if err != nil {
		return Column{}, err
	}
	for _, column := range columns {
		if column.Id == columnID {
			return column, nil
		}
	}
	return Column{}, fmt.Errorf("no column returned for id %s", columnID)
}

// GetColumns returns all the columns.
func (c SimpleClient) GetColumns(boardID int) ([]Column, error) {
	resp, err := c.Exec(context.Background(), NewQueryPayload(
		NewBoardsWithArguments(
			[]BoardsField{
				NewBoardsColumnField(
					[]ColumnsField{
						ColumnsIDField(),
						ColumnsTitleField(),
						ColumnsTypeField(),
						ColumnsSettingsStrField(),
					}),
			},
			[]BoardsArgument{
				NewBoardsIDsArgument([]int{boardID}),
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
			Boards []struct {
				Columns []Column
			}
		}
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return nil, err
	}
	if len(body.Data.Boards) != 1 {
		return nil, fmt.Errorf("no boards returned for id %d: %s", boardID, string(raw))
	}
	return body.Data.Boards[0].Columns, nil
}
