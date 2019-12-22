package monday

// BoardsService handles all the board related methods of the Monday API.
// The board’s structure is composed of rows (called items), groups of rows (called groups), and columns.
// The data of the board is stored in the items of the board and in the updates sections of each item.
// Each board has one or more owners and subscribers.
// Additionally, there are three different board types (main, shareable, private) and each board can have different sets of permissions
type BoardsService service

// Create returns a mutation that allows you to create a new board.
// - name: the board's name.
// - kind: the board's kind (public/private/share).
//
// DOCS: https://monday.com/developers/v2#mutations-section-boards-create
func (*BoardsService) Create(name string, kind BoardsKind, boardsFields []BoardsField) Mutation {
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
			{"board_kind", kind},
		},
	}
}

// CreateFromTemplate returns a mutation that allows you to create a new board from a template.
// - name: the board's name.
// - kind: the board's kind (public/private/share).
// - templateID: board template id.
// To see all the available template ID's go to monday.labs and activate the "Developer mode" feature.
// You will then be able to see the template ID's in the create board from template screen.
//
// DOCS: https://monday.com/developers/v2#mutations-section-boards-create
func (*BoardsService) CreateFromTemplate(name string, kind BoardsKind, templateID int, boardsFields []BoardsField) Mutation {
	board := Boards.Create(name, kind, boardsFields)
	board.args = append(board.args, argument{"template_id", templateID})
	return board
}

// Archive returns a mutation that allows one to archive a single board.
//
// DOCS: https://monday.com/developers/v2#mutations-section-boards-archiving
func (*BoardsService) Archive(id int, boardsFields []BoardsField) Mutation {
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
			{"board_id", id},
		},
	}
}

// List returns a query that gets one board or a collection of boards.
//
// DOCS: https://monday.com/developers/v2#queries-section-boards
func (*BoardsService) List(boardsFields []BoardsField, boardsArgs ...BoardsArgument) Query {
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
	var args []argument
	for _, ta := range boardsArgs {
		args = append(args, ta.arg)
	}
	return Query{
		name:   "boards",
		fields: fields,
		args: args,
	}
}

// The board's graphql field(s).
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
	columns := Columns.List(columnsFields)
	return BoardsField{field{"columns", &columns}}
}

// The board's description.
func BoardsDescriptionField() BoardsField {
	return boardsDescField
}

// The board's visible groups.
func NewBoardsGroupsFields(groupsFields []GroupsField, groupsArgs []GroupsArgument) BoardsField {
	groups := Groups.list(groupsFields, groupsArgs...)
	return BoardsField{field{"groups", &groups}}
}

// The unique identifier of the board.
func BoardsIDField() BoardsField {
	return boardsIDField
}

// The board's items (rows).
func NewBoardsItemsFields(itemsFields []ItemsField, itemsArgs []ItemsArgument) BoardsField {
	items := Items.List(itemsFields, itemsArgs...)
	return BoardsField{field{"items", &items}}
}

// The board's name.
func BoardsNameField() BoardsField {
	return boardsNameField
}

// The owner of the board.
func NewBoardsOwnerField(ownerFields []UsersField, ownerArgs []UsersArgument) BoardsField {
	owner := Users.List(ownerFields, ownerArgs...)
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
func NewBoardsSubscribersField(subscribersFields []UsersField, subscribersArgs []UsersArgument) BoardsField {
	subscribers := Users.List(subscribersFields, subscribersArgs...)
	subscribers.name = "subscribers"
	return BoardsField{field{"subscribers", &subscribers}}
}

// The board's specific tags
func NewBoardsTagsField(tagsFields []TagsField, tagsArgs []TagsArgument) BoardsField {
	tags := Tags.List(tagsFields, tagsArgs...)
	return BoardsField{field{"tags", &tags}}
}

// The board's updates.
func NewBoardsUpdatesField(updatesFields []UpdatesField, updatesArgs []UpdatesArgument) BoardsField {
	updates := Updates.List(updatesFields, updatesArgs...)
	return BoardsField{field{"updates", &updates}}
}

// The board's graphql argument(s).
type BoardsArgument struct {
	arg argument
}

// The board's kind (public/private/share).
type BoardsKind struct {
	kind string
}

var (
	boardsKindPublic  = BoardsKind{"public"}
	boardsKindPrivate = BoardsKind{"private"}
	boardsKindShare   = BoardsKind{"share"}
)

// Public boards.
func BoardsKindPublic() BoardsKind {
	return boardsKindPublic
}

// Private boards.
func BoardsKindPrivate() BoardsKind {
	return boardsKindPrivate
}

// Shareable boards.
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
	return BoardsArgument{argument{"board_kind", kind}}
}

// The state of the boards (all / active / archived / deleted), the default is active.
func NewBoardsStateArgument(state State) BoardsArgument {
	return BoardsArgument{argument{"state", state.state}}
}

// Get the recently created boards at the top of the list.
func NewBoardsNewestFirstArgument(first bool) BoardsArgument {
	return BoardsArgument{argument{"newest_first", first}}
}
