package monday

import (
	"fmt"
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
	boardsPositionField    = BoardsField{"pos", nil}
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
	case "owner":
		owner := f.value.(Users)
		owner.alt = "owner"
		return owner.stringify()
	case "subscribers":
		subscribers := f.value.(Users)
		subscribers.alt = "subscribers"
		return subscribers.stringify()
	case "tags":
		return f.value.(Tags).stringify()
	case "updates":
		return f.value.(Updates).stringify()
	default:
		return fmt.Sprint(f.field)
	}
}

// The board's folder unique identifier.
func BoardsFolderIDField() BoardsField {
	return boardsFolderIDField
}

// The board's kind (public / private / share).
func BoardsKindField() BoardsField {
	return boardsKindField
}

// The board's visible columns.
func NewBoardsColumnField(columns Columns) BoardsField {
	return BoardsField{field: "columns", value: columns}
}

// The board's description.
func BoardsDescriptionField() BoardsField {
	return boardsDescField
}

// The board's visible groups.
func NewBoardsGroupsFields(groups Groups) BoardsField {
	return BoardsField{field: "groups", value: groups}
}

// The unique identifier of the board.
func BoardsIDField() BoardsField {
	return boardsIDField
}

// The board's items (rows).
func NewBoardsItemsFields(items Items) BoardsField {
	return BoardsField{field: "items", value: items}
}

// The board's name.
func BoardsNameField() BoardsField {
	return boardsNameField
}

// The owner of the board.
func NewBoardsOwnerField(owner Users) BoardsField {
	return BoardsField{field: "owner", value: owner}
}

// The board's permissions.
func BoardsPermissionsField() BoardsField {
	return boardsPermissionsField
}

// The board's position.
func BoardsPositionField() BoardsField {
	return boardsPositionField
}

// The board's state (all / active / archived / deleted).
func BoardsStateField() BoardsField {
	return boardsStateField
}

// The board's subscribers.
func NewBoardsSubscribersField(subscribers Users) BoardsField {
	return BoardsField{field: "subscribers", value: subscribers}
}

// The board's specific tags
func NewBoardsTagsField(tags Tags) BoardsField {
	return BoardsField{field: "tags", value: tags}
}

// The board's updates.
func NewBoardsUpdatesField(updates Updates) BoardsField {
	return BoardsField{field: "updates", value: updates}
}

type BoardsArgument struct {
	argument string
	value    interface{}
}

func (a BoardsArgument) stringify() string {
	switch a.argument {
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
		return fmt.Sprintf("%s:%v", a.argument, a.value)
	}
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

// The boards's kind (public / private / share).
func NewKindBoardsArg(kind BoardsKind) BoardsArgument {
	return BoardsArgument{
		argument: "board_kind",
		value:    kind.kind,
	}
}

// The state of the boards (all / active / archived / deleted), the default is active.
func NewStateBoardsArg(state State) BoardsArgument {
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
