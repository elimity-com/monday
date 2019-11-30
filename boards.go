package monday

import (
	"fmt"
	"log"
	"strings"
)

type Boards struct {
	fields []BoardsField
	args   []BoardsArgument
}

func (b Boards) stringify() string {
	fields := make([]string, 0)
	for _, field := range b.fields {
		fields = append(fields, field.stringify())
	}
	args := make([]string, 0)
	for _, arg := range b.args {
		args = append(args, arg.stringify())
	}
	if len(fields) == 0 {
		return ``
	}
	if len(args) == 0 {
		return fmt.Sprintf(`boards{%s}`, strings.Join(fields, " "))
	}
	return fmt.Sprintf(`boards(%s){%s}`, strings.Join(args, ","), strings.Join(fields, " "))
}

func NewBoards(fields []BoardsField) Boards {
	if len(fields) == 0 {
		return Boards{
			fields: []BoardsField{
				BoardsIDField(),
			},
		}
	}

	return Boards{
		fields: fields,
	}
}

func NewBoardsWithArguments(fields []BoardsField, args []BoardsArgument) Boards {
	boards := NewBoards(fields)
	boards.args = args
	return boards
}

type BoardsField struct {
	field string
	value interface{}
}

var (
	boardsFolderIDField    = BoardsField{"board_folder_id", nil}
	boardsKindField        = BoardsField{"board_kind", nil}
	boardsIDField          = BoardsField{"id", nil}
	boardsDescField        = BoardsField{"description", nil}
	boardsNameField        = BoardsField{"name", nil}
	boardsPermissionsField = BoardsField{"permissions", nil}
	boardsPositionField         = BoardsField{"pos", nil}
	boardsStateField       = BoardsField{"state", nil}
)

func (f BoardsField) stringify() string {
	switch f.field {
	case "columns":
		return f.value.(Columns).stringify()
	case "groups":
		return f.value.(Groups).stringify()
	case "items":
		return f.value.(Items).stringify()
	case "owner", "subscribers":
		return f.value.(Users).stringify()
	case "tags":
		return f.value.(Tags).stringify()
	default:
		return fmt.Sprint(f.field)
	}
}

func BoardsFolderIDField() BoardsField {
	return boardsFolderIDField
}

func BoardsKindField() BoardsField {
	return boardsKindField
}

func NewBoardsColumnField(columns Columns) BoardsField {
	return BoardsField{field: "columns", value: columns}
}

func BoardsDescriptionField() BoardsField {
	return boardsDescField
}

func NewBoardsGroupsFields(groups Groups) BoardsField {
	return BoardsField{field: "groups", value: groups}
}

func BoardsIDField() BoardsField {
	return boardsIDField
}

func NewBoardsItemsFields(items Items) BoardsField {
	return BoardsField{field: "items", value: items}
}

func BoardsNameField() BoardsField {
	return boardsNameField
}

func NewBoardsOwnerField(owner Users) BoardsField {
	return BoardsField{field: "owner", value: owner}
}

func BoardsPermissionsField() BoardsField {
	return boardsPermissionsField
}

func BoardsPositionField() BoardsField {
	return boardsPositionField
}

func BoardsStateField() BoardsField {
	return boardsStateField
}

func NewBoardsSubscribersField(subscribers Users) BoardsField {
	return BoardsField{field: "subscribers", value: subscribers}
}

func NewBoardsTagsField(tags Tags) BoardsField {
	return BoardsField{field: "tags", value: tags}
}

type BoardsArgument struct {
	argument string
	value    interface{}
}

func (a BoardsArgument) stringify() string {
	switch a.argument {
	case "limit", "page", "board_kind", "state", "newest_first":
		return fmt.Sprintf("%s:%v", a.argument, a.value)
	case "ids":
		switch ids := a.value.([]int); {
		case len(ids) == 1:
			return fmt.Sprintf("ids:%d", ids[0])
		case len(ids) > 1:
			return fmt.Sprintf("ids:%s", strings.Replace(fmt.Sprint(ids), " ", ",", -1))
		default:
			return ""
		}
	default:
		log.Fatalln("unreachable boards argument")
	}
	return ""
}

type BoardsKind struct {
	kind string
}

var (
	publicBoardsKind  = BoardsKind{"public"}
	privateBoardsKind = BoardsKind{"private"}
	shareBoardsKind   = BoardsKind{"share"}
)

func PublicBoardsKind() BoardsKind {
	return publicBoardsKind
}

func PrivateBoardsKind() BoardsKind {
	return privateBoardsKind
}

func ShareBoardsKind() BoardsKind {
	return shareBoardsKind
}

type BoardsState struct {
	state string
}

var (
	allBoardsState      = BoardsState{"all"}
	activeBoardsState   = BoardsState{"active"}
	archivedBoardsState = BoardsState{"archived"}
	deletedBoardsState  = BoardsState{"deleted"}
)

func AllBoardsState() BoardsState {
	return allBoardsState
}

func ActiveBoardsState() BoardsState {
	return activeBoardsState
}

func ArchivedBoardsState() BoardsState {
	return archivedBoardsState
}

func DeletedBoardsState() BoardsState {
	return deletedBoardsState
}

// Number of items to get, the default is 25.
func NewLimitBoardsArg(value int) BoardsArgument {
	return BoardsArgument{
		argument: "limit",
		value:    value,
	}
}

// Page number to get, starting at 1.
func NewPageBoardsArg(value int) BoardsArgument {
	return BoardsArgument{
		argument: "page",
		value:    value,
	}
}

// A list of boards unique identifiers.
func NewIDsBoardsArg(ids []int) BoardsArgument {
	return BoardsArgument{
		argument: "ids",
		value:    ids,
	}
}

// The boards's kind (public / private / share)?
func NewKindBoardsArg(kind BoardsKind) BoardsArgument {
	return BoardsArgument{
		argument: "board_kind",
		value:    kind.kind,
	}
}

// The state of the boards (all / active / archived / deleted), the default is active.
func NewStateBoardsArg(state BoardsState) BoardsArgument {
	return BoardsArgument{
		argument: "state",
		value:    state.state,
	}
}

// Get the recently created boards at the top of the list.
func NewNewestFirstBoardsArg(first bool) BoardsArgument {
	return BoardsArgument{
		argument: "newest_first",
		value:    first,
	}
}
