package monday

import "fmt"

// ItemsService handles all the item related methods of the Monday API.
// Items are the objects that hold the actual data within the board, to better illustrate this,
// you can think of a board as a table and a item as a single row in that table.
type ItemsService service

// Create returns a mutation that allows you to create a new item in the different boards.
// - boardID: the board's unique identifier.
// - groupID: the group's unique identifier.
// - name: the new item's name.
// - values: the column values of the new item.
//
// DOCS: https://monday.com/developers/v2#mutations-section-items-create
func (*ItemsService) Create(boardID int, groupID, name string, values []ColumnValue, itemsFields []ItemsField) Mutation {
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
	var columnValues string
	for _, v := range values {
		columnValues += fmt.Sprintf(`{%q:%s}`, v.id, v.value)
	}
	if columnValues != "" {
		args = append(args, argument{"column_values", columnValues})
	}
	return Mutation{
		name:   "create_item",
		fields: fields,
		args:   args,
	}
}

// MoveToGroup returns a mutation that allows you to move a item between groups in the same board.
// - itemID: the item's unique identifier.
// - groupID: the group's unique identifier.
//
// DOCS: https://monday.com/developers/v2#mutations-section-items-move-item-to-group
func (*ItemsService) MoveToGroup(itemID int, groupID string, itemsFields []ItemsField) Mutation {
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

// Archive returns a mutation that allows one to archive a single item.
// - id: the item's unique identifier.
//
// DOCS: https://monday.com/developers/v2#mutations-section-items-archive
func (*ItemsService) Archive(id int, itemsFields []ItemsField) Mutation {
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
			{"item_id", id},
		},
	}
}

// Delete returns a mutation that allows one to delete a single item.
// - id: the item's unique identifier.
//
// DOCS: https://monday.com/developers/v2#mutations-section-items-archive
func (*ItemsService) Delete(id int, itemsFields []ItemsField) Mutation {
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
			{"item_id", id},
		},
	}
}

// List returns a query that gets one or a collection of items.
//
// DOCS: https://monday.com/developers/v2#queries-section-items
func (*ItemsService) List(itemsFields []ItemsField, itemsArgs ...ItemsArgument) Query {
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
	var args []argument
	for _, ia := range itemsArgs {
		args = append(args, ia.arg)
	}
	return Query{
		name:   "items",
		fields: fields,
		args:   args,
	}
}

// The item's graphql field(s).
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
func NewItemsBoardField(boardsFields []BoardsField, boardsArgs []BoardsArgument) ItemsField {
	board := Boards.List(boardsFields, boardsArgs...)
	board.name = "board"
	return ItemsField{field{"boards", &board}}
}

// The item's column values.
func NewItemsColumnValuesField(valuesFields []ColumnValuesField, valuesArgs []ColumnValuesArgument) ItemsField {
	values := listColumnValues(valuesFields, valuesArgs)
	return ItemsField{field{"column_values", &values}}
}

// The item's create date.
func ItemsCreatedAtField() ItemsField {
	return itemsCreatedAtField
}

// The item's creator.
func NewItemsCreatorField(creatorFields []UsersField, creatorArgs []UsersArgument) ItemsField {
	creator := Users.List(creatorFields, creatorArgs...)
	creator.name = "creator"
	return ItemsField{field{"creator", &creator}}
}

// The unique identifier of the item creator.
func ItemsCreatorIDField() ItemsField {
	return itemsCreatorIDField
}

// The group that contains this item.
func NewItemsGroupField(groupsFields []GroupsField, groupsArgs []GroupsArgument) ItemsField {
	group := Groups.list(groupsFields, groupsArgs...)
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
func NewItemsSubscribersField(subscribersFields []UsersField, subscribersArgs []UsersArgument) ItemsField {
	subscribers := Users.List(subscribersFields, subscribersArgs...)
	subscribers.name = "subscribers"
	return ItemsField{field{"subscribers", &subscribers}}
}

// The item's last update date.
func ItemsUpdatedAtField() ItemsField {
	return itemsUpdatedAtField
}

// The item's updates.
func NewItemsUpdatesField(updatesFields []UpdatesField, updatesArgs []UpdatesArgument) ItemsField {
	updates := Updates.List(updatesFields, updatesArgs...)
	return ItemsField{field{"updates", &updates}}
}

// The item's graphql argument(s).
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
