package monday

import "testing"

func TestBoardsFields(t *testing.T) {
	tests := []struct {
		boards Boards
		str    string
	}{
		{ // empty boards
			boards: Boards{},
			str:    ``,
		},
		{ // empty boards
			boards: NewBoard(nil),
			str:    `boards{id}`,
		},
		{ // string fields
			boards: Boards{
				fields: []BoardsField{
					BoardsKindField(),
					BoardsDescriptionField(),
				},
			},
			str: `boards{board_kind description}`,
		},
		{ // columns field
			boards: Boards{
				fields: []BoardsField{
					NewBoardsColumnField(NewColumns(nil)),
				},
			},
			str: `boards{columns{id}}`,
		},
	}

	for _, test := range tests {
		if str := test.boards.stringify(); str != test.str {
			t.Errorf("got: %s, expected: %s \n", str, test.str)
		}
	}
}

func TestBoardsArgs(t *testing.T) {
	tests := []struct {
		boards Boards
		str    string
	}{
		{ // empty boards
			boards: NewBoardsWithArguments(nil, nil),
			str:    `boards{id}`,
		},
		{ // single id arg
			boards: NewBoardsWithArguments(nil, []BoardsArgument{
				NewIDsBoardsArg([]int{0}),
			}),
			str: `boards(ids:0){id}`,
		},
		{ // multiple ids arg
			boards: NewBoardsWithArguments(nil, []BoardsArgument{
				NewIDsBoardsArg([]int{0, 1, 2}),
			}),
			str: `boards(ids:[0,1,2]){id}`,
		},
		{ // integer fields args
			boards: NewBoardsWithArguments(nil, []BoardsArgument{
				NewLimitBoardsArg(1),
				NewPageBoardsArg(2),
			}),
			str: `boards(limit:1,page:2){id}`,
		},
		{ // enum fields args
			boards: NewBoardsWithArguments(nil, []BoardsArgument{
				NewKindBoardsArg(PublicBoardsKind()),
				NewStateBoardsArg(AllBoardsState()),
			}),
			str: `boards(board_kind:public,state:all){id}`,
		},
		{ // newest first arg
			boards: NewBoardsWithArguments(nil, []BoardsArgument{
				NewNewestFirstBoardsArg(false),
			}),
			str: `boards(newest_first:false){id}`,
		},
	}

	for _, test := range tests {
		if str := test.boards.stringify(); str != test.str {
			t.Errorf("got: %s, expected: %s \n", str, test.str)
		}
	}
}

func TestBoardFieldsArgs(t *testing.T) {
	tests := []struct {
		boards Boards
		str    string
	}{
		{ // simple boards
			boards: NewBoardsWithArguments(
				[]BoardsField{
					BoardsNameField(),
				},
				[]BoardsArgument{
					NewLimitBoardsArg(1),
				},
			),
			str: `boards(limit:1){name}`,
		},
		{ // complex boards
			boards: NewBoardsWithArguments(
				[]BoardsField{
					NewBoardsColumnField(NewColumns([]ColumnsField{
						ColumnsTitleField(),
					})),
				},
				[]BoardsArgument{
					NewLimitBoardsArg(1),
				},
			),
			str: `boards(limit:1){columns{title}}`,
		},
	}

	for _, test := range tests {
		if str := test.boards.stringify(); str != test.str {
			t.Errorf("got: %s, expected: %s \n", str, test.str)
		}
	}
}
