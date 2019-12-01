package monday

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

// TODO: column_values

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
	group := NewGroupWithArguments(groupsFields, groupsArguments)
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

// TODO: 'items' doesn't accept argument 'newest_first'?
