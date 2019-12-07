package monday

func DuplicateGroup(boardID int, groupID string, AddToTop bool, groupsFields []GroupsField) Mutation {
	if len(groupsFields) == 0 {
		groupsFields = append(groupsFields, groupsIDField)
	}

	var fields []field
	for _, gf := range groupsFields {
		fields = append(fields, gf.field)
	}
	return Mutation{
		name:   "duplicate_group",
		fields: fields,
		args: []argument{
			{"board_id", boardID},
			{"group_id", groupID},
			{"add_to_top", AddToTop},
		},
	}
}

func DuplicateGroupWidthNewTitle(boardID int, groupID string, AddToTop bool, groupTitle string, groupsFields []GroupsField) Mutation {
	group := DuplicateGroup(boardID, groupID, AddToTop, groupsFields)
	group.args = append(group.args, argument{"group_title", groupTitle})
	return group
}

func CreateGroup(boardID int, groupName string, groupsFields []GroupsField) Mutation {
	if len(groupsFields) == 0 {
		groupsFields = append(groupsFields, groupsIDField)
	}

	var fields []field
	for _, gf := range groupsFields {
		fields = append(fields, gf.field)
	}
	return Mutation{
		name:   "create_group",
		fields: fields,
		args: []argument{
			{"board_id", boardID},
			{"group_name", groupName},
		},
	}
}

func ArchiveGroup(boardID int, groupID string, groupsFields []GroupsField) Mutation {
	if len(groupsFields) == 0 {
		groupsFields = append(groupsFields, groupsIDField)
	}

	var fields []field
	for _, gf := range groupsFields {
		fields = append(fields, gf.field)
	}
	return Mutation{
		name:   "archive_group",
		fields: fields,
		args: []argument{
			{"board_id", boardID},
			{"group_id", groupID},
		},
	}
}

func DeleteGroup(boardID int, groupID string, groupsFields []GroupsField) Mutation {
	if len(groupsFields) == 0 {
		groupsFields = append(groupsFields, groupsIDField)
	}

	var fields []field
	for _, gf := range groupsFields {
		fields = append(fields, gf.field)
	}
	return Mutation{
		name:   "delete_group",
		fields: fields,
		args: []argument{
			{"board_id", boardID},
			{"group_id", groupID},
		},
	}
}

func NewGroups(groupsFields []GroupsField) Query {
	if len(groupsFields) == 0 {
		return Query{
			name: "groups",
			fields: []field{
				GroupsIDField().field,
			},
		}
	}

	var fields []field
	for _, gf := range groupsFields {
		fields = append(fields, gf.field)
	}
	return Query{
		name:   "groups",
		fields: fields,
	}
}

func NewGroupsWithArguments(groupsFields []GroupsField, groupsArgs []GroupsArgument) Query {
	groups := NewGroups(groupsFields)
	var args []argument
	for _, ga := range groupsArgs {
		args = append(args, ga.arg)
	}
	groups.args = args
	return groups
}

type GroupsField struct {
	field field
}

var (
	groupsArchivedField = GroupsField{field{"archived", nil}}
	groupsColorField    = GroupsField{field{"color", nil}}
	groupsDeletedField  = GroupsField{field{"deleted", nil}}
	groupsIDField       = GroupsField{field{"id", nil}}
	groupsPositionField = GroupsField{field{"position", nil}}
	groupsTitleField    = GroupsField{field{"title", nil}}
)

// Is the group archived or not.
func GroupsArchivedField() GroupsField {
	return groupsArchivedField
}

// The group's color.
func GroupsColorField() GroupsField {
	return groupsColorField
}

// Is the group deleted or not.
func GroupsDeletedField() GroupsField {
	return groupsDeletedField
}

// The group's unique identifier.
func GroupsIDField() GroupsField {
	return groupsIDField
}

// The items in the group.
func NewGroupsItemsField(itemsFields []ItemsField, itemsArguments []ItemsArgument) GroupsField {
	items := NewItemsWithArguments(itemsFields, itemsArguments)
	return GroupsField{field{"items", &items}}
}

// The group's position in the board.
func GroupsPositionField() GroupsField {
	return groupsPositionField
}

// The group's title.
func GroupsTitleField() GroupsField {
	return groupsTitleField
}

type GroupsArgument struct {
	arg argument
}

// A list of group unique identifiers.
func NewGroupsIDsArgument(ids []string) GroupsArgument {
	return GroupsArgument{argument{"ids", ids}}
}
