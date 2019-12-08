package monday

func CreateItem(boardID int, groupID, name string, values ColumnValues, itemsFields []ItemsField) Mutation {
	if len(itemsFields) == 0 {
		itemsFields = append(itemsFields, itemsIDField)
	}

	var fields []field
	for _, i := range itemsFields {
		fields = append(fields, i.field)
	}
	args := []argument{
		{"board_id", boardID},
		{"group_id", groupID},
		{"item_name", name},
	}
	columnValues := values.join()
	if columnValues != "" {
		args = append(args, argument{"column_values", columnValues})
	}
	return Mutation{
		name:   "create_item",
		fields: fields,
		args:   args,
	}
}

func MoveItemToGroup(itemID int, groupID string, itemsFields []ItemsField) Mutation {
	if len(itemsFields) == 0 {
		itemsFields = append(itemsFields, itemsIDField)
	}

	var fields []field
	for _, i := range itemsFields {
		fields = append(fields, i.field)
	}
	return Mutation{
		name:   "move_item_to_group",
		fields: fields,
		args: []argument{
			{"item_id", itemID},
			{"group_id", groupID},
		},
	}
}

func ArchiveItem(itemID int, itemsFields []ItemsField) Mutation {
	if len(itemsFields) == 0 {
		itemsFields = append(itemsFields, itemsIDField)
	}

	var fields []field
	for _, i := range itemsFields {
		fields = append(fields, i.field)
	}
	return Mutation{
		name:   "archive_item",
		fields: fields,
		args: []argument{
			{"item_id", itemID},
		},
	}
}

func DeleteItem(itemID int, itemsFields []ItemsField) Mutation {
	if len(itemsFields) == 0 {
		itemsFields = append(itemsFields, itemsIDField)
	}

	var fields []field
	for _, i := range itemsFields {
		fields = append(fields, i.field)
	}
	return Mutation{
		name:   "delete_item",
		fields: fields,
		args: []argument{
			{"item_id", itemID},
		},
	}
}

func NewItems(itemsFields []ItemsField) Query {
	if len(itemsFields) == 0 {
		return Query{
			name: "items",
			fields: []field{
				ItemsIDField().field,
			},
		}
	}

	var fields []field
	for _, i := range itemsFields {
		fields = append(fields, i.field)
	}
	return Query{
		name:   "items",
		fields: fields,
	}
}

func NewItemsWithArguments(itemsFields []ItemsField, itemsArgs []ItemsArgument) Query {
	items := NewItems(itemsFields)
	var args []argument
	for _, ia := range itemsArgs {
		args = append(args, ia.arg)
	}
	items.args = args
	return items
}

type ItemsField struct {
	field field
}

var (
	itemsCreatedAtField = ItemsField{field{"created_at", nil}}
	itemsCreatorIDField = ItemsField{field{"creator_id", nil}}
	itemsIDField        = ItemsField{field{"id", nil}}
	itemsNameField      = ItemsField{field{"name", nil}}
	itemsStateField     = ItemsField{field{"state", nil}}
	itemsUpdatedAtField = ItemsField{field{"updated_at", nil}}
)

// The board that contains this item.
func NewItemsBoardField(boardsFields []BoardsField, boardsArguments []BoardsArgument) ItemsField {
	board := NewBoardsWithArguments(boardsFields, boardsArguments)
	board.name = "board"
	return ItemsField{field{"boards", &board}}
}

// The item's column values.
func NewItemsColumnValuesField(valuesFields []ColumnValuesField, valuesArguments []ColumnValuesArgument) ItemsField {
	values := newColumnValuesWithArguments(valuesFields, valuesArguments)
	return ItemsField{field{"column_values", &values}}
}

// The item's create date.
func ItemsCreatedAtField() ItemsField {
	return itemsCreatedAtField
}

// The item's creator.
func NewItemsCreatorField(creatorFields []UsersField, creatorArguments []UsersArgument) ItemsField {
	creator := NewUsersWithArguments(creatorFields, creatorArguments)
	creator.name = "creator"
	return ItemsField{field{"creator", &creator}}
}

// The unique identifier of the item creator.
func ItemsCreatorIDField() ItemsField {
	return itemsCreatorIDField
}

// The group that contains this item.
func NewItemsGroupField(groupsFields []GroupsField, groupsArguments []GroupsArgument) ItemsField {
	group := NewGroupsWithArguments(groupsFields, groupsArguments)
	group.name = "group"
	return ItemsField{field{"groups", &group}}
}

// The item's unique identifier.
func ItemsIDField() ItemsField {
	return itemsIDField
}

// The item's name.
func ItemsNameField() ItemsField {
	return itemsNameField
}

// The board's state (all / active / archived / deleted).
func ItemsStateField() ItemsField {
	return itemsStateField
}

// The pulses's subscribers.
func NewItemsSubscribersField(subscribersFields []UsersField, subscribersArguments []UsersArgument) ItemsField {
	subscribers := NewUsersWithArguments(subscribersFields, subscribersArguments)
	subscribers.name = "subscribers"
	return ItemsField{field{"subscribers", &subscribers}}
}

// The item's last update date.
func ItemsUpdatedAtField() ItemsField {
	return itemsUpdatedAtField
}

// The item's updates.
func NewItemsUpdatesField(updatesFields []UpdatesField, updatesArguments []UpdatesArgument) ItemsField {
	updates := NewUpdatesWithArguments(updatesFields, updatesArguments)
	return ItemsField{field{"updates", &updates}}
}

type ItemsArgument struct {
	arg argument
}

// Number of items to get, the default is 25.
func NewItemsLimitArgument(value int) ItemsArgument {
	return ItemsArgument{argument{"limit", value}}
}

// Page number to get, starting at 1.
func NewItemsPageArgument(value int) ItemsArgument {
	return ItemsArgument{argument{"page", value}}
}

// A list of items unique identifiers.
func NewItemsIDsArgument(ids []int) ItemsArgument {
	return ItemsArgument{argument{"ids", ids}}
}

// Get the recently created items at the top of the list.
func NewItemsNewestFirst(value bool) ItemsArgument {
	return ItemsArgument{argument{"newest_first", value}}
}
