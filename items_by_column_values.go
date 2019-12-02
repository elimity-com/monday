package monday

func NewItemsByColumnValues(columnValuesFields []ItemsByColumnValuesField) Query {
	if len(columnValuesFields) == 0 {
		return Query{
			name: "items_by_column_values",
			fields: []field{
				ItemsByColumnValuesIDField().field,
			},
		}
	}

	var fields []field
	for _, vf := range columnValuesFields {
		fields = append(fields, vf.field)
	}
	return Query{
		name:   "items_by_column_values",
		fields: fields,
	}
}

func NewItemsByColumnValuesWithArguments(columnValuesFields []ItemsByColumnValuesField, columnValuesArgs []ItemsByColumnValuesArgument) Query {
	columnValues := NewItemsByColumnValues(columnValuesFields)
	var args []argument
	for _, va := range columnValuesArgs {
		args = append(args, va.arg)
	}
	columnValues.args = args
	return columnValues
}

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
func NewItemsByColumnValuesBoardField(boardFields []BoardsField, boardArguments []BoardsArgument) ItemsByColumnValuesField {
	board := NewBoardsWithArguments(boardFields, boardArguments)
	board.name = "board"
	return ItemsByColumnValuesField{field{"boards", &board}}
}

// The item's column values.
func NewItemsByColumnValuesColumnValuesField(valuesFields []ColumnValuesField, valuesArguments []ColumnValuesArgument) ItemsByColumnValuesField {
	values := newColumnValuesWithArguments(valuesFields, valuesArguments)
	return ItemsByColumnValuesField{field{"column_values", &values}}
}

// The item's create date.
func ItemsByColumnValuesCreatedAtField() ItemsByColumnValuesField {
	return itemsByColumnValuesCreatedAtField
}

// The item's creator.
func NewItemsByColumnValuesCreatorField(creatorFields []UsersField, creatorArguments []UsersArgument) ItemsByColumnValuesField {
	creator := NewUsersWithArguments(creatorFields, creatorArguments)
	creator.name = "creator"
	return ItemsByColumnValuesField{field{"creator", &creator}}
}

// The unique identifier of the item creator.
func ItemsByColumnValuesCreatorIDField() ItemsByColumnValuesField {
	return itemsByColumnValuesCreatorIDField
}

// The group that contains this item.
func NewItemsByColumnValuesGroupField(groupFields []GroupsField, groupArguments []GroupsArgument) ItemsByColumnValuesField {
	group := NewGroupWithArguments(groupFields, groupArguments)
	group.name = "group"
	return ItemsByColumnValuesField{field{"groups", &group}}
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
func NewItemsByColumnValuesSubscribersField(subscribersFields []UsersField, subscribersArguments []UsersArgument) ItemsByColumnValuesField {
	subscribers := NewUsersWithArguments(subscribersFields, subscribersArguments)
	subscribers.name = "subscribers"
	return ItemsByColumnValuesField{field{"subscribers", &subscribers}}
}

// The item's last update date.
func ItemsByColumnValuesUpdatedAtField() ItemsByColumnValuesField {
	return itemsByColumnValuesUpdatedAtField
}

// The item's updates.
func NewItemsByColumnValuesUpdatesField(updatesFields []UpdatesField, updatesArguments []UpdatesArgument) ItemsByColumnValuesField {
	updates := NewUpdatesWithArguments(updatesFields, updatesArguments)
	return ItemsByColumnValuesField{field{"updates", &updates}}
}

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
func NewItemsByColumnValuesBoardIDArgument(value int) ItemsByColumnValuesArgument {
	return ItemsByColumnValuesArgument{argument{"board_id", value}}
}

// The column's unique identifier.
func NewItemsByColumnValuesColumnIDArgument(value string) ItemsByColumnValuesArgument {
	return ItemsByColumnValuesArgument{argument{"column_id", value}}
}

// The column value to search items by.
func NewItemsByColumnValuesColumnValueArgument(value string) ItemsByColumnValuesArgument {
	return ItemsByColumnValuesArgument{argument{"column_value", value}}
}

// The column type.
func NewItemsByColumnValuesColumnTypeArgument(value string) ItemsByColumnValuesArgument {
	return ItemsByColumnValuesArgument{argument{"column_type", value}}
}

// The state of the item (all / active / archived / deleted), the default is active.
func NewItemsByColumnValuesStateArgument(state State) ItemsByColumnValuesArgument {
	return ItemsByColumnValuesArgument{argument{"state", state.state}}
}
