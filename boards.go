package monday

func CreateBoard(name string, kind BoardsKind, boardsFields []BoardsField) Mutation {
	if len(boardsFields) == 0 {
		boardsFields = append(boardsFields, boardsIDField)
	}

	var fields []field
	for _, bf := range boardsFields {
		fields = append(fields, bf.field)
	}
	return Mutation{
		name:   "create_board",
		fields: fields,
		args: []argument{
			{"board_name", name},
			{"board_kind", kind.kind},
		},
	}
}

func CreatBoardFromTemplate(name string, kind BoardsKind, templateID int, boardsFields []BoardsField) Mutation {
	board := CreateBoard(name, kind, boardsFields)
	board.args = append(board.args, argument{"template_id", templateID})
	return board
}

func ArchiveBoard(boardID int, boardsFields []BoardsField) Mutation {
	if len(boardsFields) == 0 {
		boardsFields = append(boardsFields, boardsIDField)
	}

	var fields []field
	for _, bf := range boardsFields {
		fields = append(fields, bf.field)
	}
	return Mutation{
		name:   "archive_board",
		fields: fields,
		args: []argument{
			{"board_id", boardID},
		},
	}
}

func NewBoards(boardsFields []BoardsField) Query {
	if len(boardsFields) == 0 {
		return Query{
			name: "boards",
			fields: []field{
				BoardsIDField().field,
			},
		}
	}

	var fields []field
	for _, bf := range boardsFields {
		fields = append(fields, bf.field)
	}
	return Query{
		name:   "boards",
		fields: fields,
	}
}

func NewBoardsWithArguments(boardsFields []BoardsField, boardsArgs []BoardsArgument) Query {
	boards := NewBoards(boardsFields)
	var args []argument
	for _, ta := range boardsArgs {
		args = append(args, ta.arg)
	}
	boards.args = args
	return boards
}

type BoardsField struct {
	field field
}

var (
	boardsFolderIDField    = BoardsField{field{"board_folder_id", nil}}
	boardsKindField        = BoardsField{field{"board_kind", nil}}
	boardsIDField          = BoardsField{field{"id", nil}}
	boardsDescField        = BoardsField{field{"description", nil}}
	boardsNameField        = BoardsField{field{"name", nil}}
	boardsPermissionsField = BoardsField{field{"permissions", nil}}
	boardsPositionField    = BoardsField{field{"pos", nil}}
	boardsStateField       = BoardsField{field{"state", nil}}
)

// The board's folder unique identifier.
func BoardsFolderIDField() BoardsField {
	return boardsFolderIDField
}

// The board's kind (public / private / share).
func BoardsKindField() BoardsField {
	return boardsKindField
}

// The board's visible columns.
func NewBoardsColumnField(columnsFields []ColumnsField) BoardsField {
	columns := NewColumns(columnsFields)
	return BoardsField{field{"columns", &columns}}
}

// The board's description.
func BoardsDescriptionField() BoardsField {
	return boardsDescField
}

// The board's visible groups.
func NewBoardsGroupsFields(groupsFields []GroupsField, groupsArguments []GroupsArgument) BoardsField {
	groups := NewGroupsWithArguments(groupsFields, groupsArguments)
	return BoardsField{field{"groups", &groups}}
}

// The unique identifier of the board.
func BoardsIDField() BoardsField {
	return boardsIDField
}

// The board's items (rows).
func NewBoardsItemsFields(itemsFields []ItemsField, itemsArguments []ItemsArgument) BoardsField {
	items := NewItemsWithArguments(itemsFields, itemsArguments)
	return BoardsField{field{"items", &items}}
}

// The board's name.
func BoardsNameField() BoardsField {
	return boardsNameField
}

// The owner of the board.
func NewBoardsOwnerField(ownerFields []UsersField, ownerArguments []UsersArgument) BoardsField {
	owner := NewUsersWithArguments(ownerFields, ownerArguments)
	owner.name = "owner"
	return BoardsField{field{"owner", &owner}}
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
func NewBoardsSubscribersField(subscribersFields []UsersField, subscribersArguments []UsersArgument) BoardsField {
	subscribers := NewUsersWithArguments(subscribersFields, subscribersArguments)
	subscribers.name = "subscribers"
	return BoardsField{field{"subscribers", &subscribers}}
}

// The board's specific tags
func NewBoardsTagsField(tagsFields []TagsField, tagsArguments []TagsArgument) BoardsField {
	tags := NewTagsWithArguments(tagsFields, tagsArguments)
	return BoardsField{field{"tags", &tags}}
}

// The board's updates.
func NewBoardsUpdatesField(updatesFields []UpdatesField, updatesArguments []UpdatesArgument) BoardsField {
	updates := NewUpdatesWithArguments(updatesFields, updatesArguments)
	return BoardsField{field{"updates", &updates}}
}

type BoardsArgument struct {
	arg argument
}

type BoardsKind struct {
	kind string
}

var (
	boardsKindPublic  = BoardsKind{"public"}
	boardsKindPrivate = BoardsKind{"private"}
	boardsKindShare   = BoardsKind{"share"}
)

func BoardsKindPublic() BoardsKind {
	return boardsKindPublic
}

func BoardsKindPrivate() BoardsKind {
	return boardsKindPrivate
}

func BoardsKindShare() BoardsKind {
	return boardsKindShare
}

// Number of items to get, the default is 25.
func NewBoardsLimitArgument(value int) BoardsArgument {
	return BoardsArgument{argument{"limit", value}}
}

// Page number to get, starting at 1.
func NewBoardsPageArgument(value int) BoardsArgument {
	return BoardsArgument{argument{"page", value}}
}

// A list of boards unique identifiers.
func NewBoardsIDsArgument(ids []int) BoardsArgument {
	return BoardsArgument{argument{"ids", ids}}
}

// The boards's kind (public / private / share).
func NewBoardsKindArgument(kind BoardsKind) BoardsArgument {
	return BoardsArgument{argument{"board_kind", kind.kind}}
}

// The state of the boards (all / active / archived / deleted), the default is active.
func NewBoardsStateArgument(state State) BoardsArgument {
	return BoardsArgument{argument{"state", state.state}}
}

// Get the recently created boards at the top of the list.
func NewBoardsNewestFirstArgument(first bool) BoardsArgument {
	return BoardsArgument{argument{"newest_first", first}}
}
