package monday

// ItemsByColumnValuesService handles all the items by column values related methods of the Monday API.
// Just like in a table, the different values of each item are stored in columns.
type ItemsByColumnValuesService service

// List return a query that searches for items based on their column values and returns data about those specific items.
// - boardID: the board's unique identifier.
// - columnID: the column's unique identifier.
// - value: the column value to search items by.
//
// DOCS: https://monday.com/developers/v2#queries-section-items-by-column-values
func (*ItemsByColumnValuesService) List(boardID int, columnID string, value ColumnValue,
	columnValuesFields []ItemsByColumnValuesField, columnValuesArgs ...ItemsByColumnValuesArgument) Query {
	args := []argument{
		newItemsByColumnValuesBoardIDArgument(boardID).arg,
		newItemsByColumnValuesColumnIDArgument(columnID).arg,
		newItemsByColumnValuesColumnValueArgument(value).arg,
	}

	if len(columnValuesFields) == 0 {
		return Query{
			name: "items_by_column_values",
			fields: []field{
				ItemsByColumnValuesIDField().field,
			},
			args: args,
		}
	}

	var fields []field
	for _, vf := range columnValuesFields {
		fields = append(fields, vf.field)
	}
	for _, va := range columnValuesArgs {
		args = append(args, va.arg)
	}
	return Query{
		name:   "items_by_column_values",
		fields: fields,
		args:   args,
	}
}

// The items by column values' graphql field(s).
type ItemsByColumnValuesField struct {
	field field
}

var (
	itemsByColumnValuesCreatedAtField = ItemsByColumnValuesField{field{"created_at", nil}}
	itemsByColumnValuesCreatorIDField = ItemsByColumnValuesField{field{"creator_id", nil}}
	itemsByColumnValuesIDField        = ItemsByColumnValuesField{field{"id", nil}}
	itemsByColumnValuesNameField      = ItemsByColumnValuesField{field{"name", nil}}
	itemsByColumnValuesStateField     = ItemsByColumnValuesField{field{"state", nil}}
	itemsByColumnValuesUpdatedAtField = ItemsByColumnValuesField{field{"updated_at", nil}}
)

// The board that contains this item.
func NewItemsByColumnValuesBoardField(boardFields []BoardsField, boardArgs []BoardsArgument) ItemsByColumnValuesField {
	board := Boards.List(boardFields, boardArgs...)
	board.name = "board"
	return ItemsByColumnValuesField{field{board.name, &board}}
}

// The item's column values.
func NewItemsByColumnValuesColumnValuesField(valuesFields []ColumnValuesField, valuesArgs []ColumnValuesArgument) ItemsByColumnValuesField {
	values := listColumnValues(valuesFields, valuesArgs)
	return ItemsByColumnValuesField{field{"column_values", &values}}
}

// The item's create date.
func ItemsByColumnValuesCreatedAtField() ItemsByColumnValuesField {
	return itemsByColumnValuesCreatedAtField
}

// The item's creator.
func NewItemsByColumnValuesCreatorField(creatorFields []UsersField, creatorArgs []UsersArgument) ItemsByColumnValuesField {
	creator := Users.List(creatorFields, creatorArgs...)
	creator.name = "creator"
	return ItemsByColumnValuesField{field{creator.name, &creator}}
}

// The unique identifier of the item creator.
func ItemsByColumnValuesCreatorIDField() ItemsByColumnValuesField {
	return itemsByColumnValuesCreatorIDField
}

// The group that contains this item.
func NewItemsByColumnValuesGroupField(groupFields []GroupsField, groupArgs []GroupsArgument) ItemsByColumnValuesField {
	group := Groups.list(groupFields, groupArgs...)
	group.name = "group"
	return ItemsByColumnValuesField{field{group.name, &group}}
}

// The item's unique identifier.
func ItemsByColumnValuesIDField() ItemsByColumnValuesField {
	return itemsByColumnValuesIDField
}

// The item's name.
func ItemsByColumnValuesNameField() ItemsByColumnValuesField {
	return itemsByColumnValuesNameField
}

// The board's state (all / active / archived / deleted).
func ItemsByColumnValuesStateField() ItemsByColumnValuesField {
	return itemsByColumnValuesStateField
}

// The pulses's subscribers.
func NewItemsByColumnValuesSubscribersField(subscribersFields []UsersField, subscribersArgs []UsersArgument) ItemsByColumnValuesField {
	subscribers := Users.List(subscribersFields, subscribersArgs...)
	subscribers.name = "subscribers"
	return ItemsByColumnValuesField{field{subscribers.name, &subscribers}}
}

// The item's last update date.
func ItemsByColumnValuesUpdatedAtField() ItemsByColumnValuesField {
	return itemsByColumnValuesUpdatedAtField
}

// The item's updates.
func NewItemsByColumnValuesUpdatesField(updatesFields []UpdatesField, updatesArgs []UpdatesArgument) ItemsByColumnValuesField {
	updates := Updates.List(updatesFields, updatesArgs...)
	return ItemsByColumnValuesField{field{"updates", &updates}}
}

// The items by column values' graphql argument(s).
type ItemsByColumnValuesArgument struct {
	arg argument
}

// Number of items to get, the default is 25.
func NewItemsByColumnValuesLimitArgument(value int) ItemsByColumnValuesArgument {
	return ItemsByColumnValuesArgument{argument{"limit", value}}
}

// Page number to get, starting at 1.
func NewItemsByColumnValuesPageArgument(value int) ItemsByColumnValuesArgument {
	return ItemsByColumnValuesArgument{argument{"page", value}}
}

// The board's unique identifier.
func newItemsByColumnValuesBoardIDArgument(value int) ItemsByColumnValuesArgument {
	return ItemsByColumnValuesArgument{argument{"board_id", value}}
}

// The column's unique identifier.
func newItemsByColumnValuesColumnIDArgument(value string) ItemsByColumnValuesArgument {
	return ItemsByColumnValuesArgument{argument{"column_id", value}}
}

// The column value to search items by.
func newItemsByColumnValuesColumnValueArgument(value ColumnValue) ItemsByColumnValuesArgument {
	return ItemsByColumnValuesArgument{argument{"column_value", value.value}}
}

// The column type.
func NewItemsByColumnValuesColumnTypeArgument(value string) ItemsByColumnValuesArgument {
	return ItemsByColumnValuesArgument{argument{"column_type", value}}
}

// The state of the item (all / active / archived / deleted), the default is active.
func NewItemsByColumnValuesStateArgument(state State) ItemsByColumnValuesArgument {
	return ItemsByColumnValuesArgument{argument{"state", state.state}}
}
