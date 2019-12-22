package monday

// GroupsService handles all the group related methods of the Monday API.
// Items are grouped together in units called groups.
// Each board contains one or multiple groups, and each group can hold one or many items.
type GroupsService service

// Duplicate returns a mutation that duplicates a group with all of its items.
// - boardID: the board's unique identifier.
// - groupID: the group's unique identifier.
// - top: should the new group be added to the top?
//
// DOCS: https://monday.com/developers/v2#mutations-section-groups-duplicate
func (*GroupsService) Duplicate(boardID int, groupID string, top bool, groupsFields []GroupsField) Mutation {
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
			{"add_to_top", top},
		},
	}
}

// DuplicateWithNewTitle returns a mutation that duplicates a group with all of its items.
// - boardID: the board's unique identifier.
// - groupID: the group's unique identifier.
// - top: should the new group be added to the top?
// - title: the group's title.
//
// DOCS: https://monday.com/developers/v2#mutations-section-groups-duplicate
func (*GroupsService) DuplicateWithNewTitle(boardID int, groupID string, top bool, title string, groupsFields []GroupsField) Mutation {
	group := Groups.Duplicate(boardID, groupID, top, groupsFields)
	group.args = append(group.args, argument{"group_title", title})
	return group
}

// Create returns a mutation that creates a new empty group.
// - boardID: the board's unique identifier.
// - groupID: The group's unique identifier.
//
// DOCS: https://monday.com/developers/v2#mutations-section-groups-create
func (*GroupsService) Create(boardID int, groupName string, groupsFields []GroupsField) Mutation {
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

// Archive returns a mutation that archives a group with all of its items.
// - boardID: the board's unique identifier.
// - groupID: The group's unique identifier.
//
// DOCS: https://monday.com/developers/v2#mutations-section-groups-archive
func (*GroupsService) Archive(boardID int, groupID string, groupsFields []GroupsField) Mutation {
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

// Delete returns a mutation that deletes a group with all of its items.
// - boardID: the board's unique identifier.
// - groupID: The group's unique identifier.
//
// DOCS: https://monday.com/developers/v2#mutations-section-groups-delete
func (*GroupsService) Delete(boardID int, groupID string, groupsFields []GroupsField) Mutation {
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

// List returns a query that gets one group or a collection of groups in a specific board.
//
// DOCS: https://monday.com/developers/v2#queries-section-groups
func (*GroupsService) list(groupsFields []GroupsField, groupsArgs ...GroupsArgument) Query {
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
	var args []argument
	for _, ga := range groupsArgs {
		args = append(args, ga.arg)
	}
	return Query{
		name:   "groups",
		fields: fields,
		args:   args,
	}
}

// The group's graphql field(s).
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
func NewGroupsItemsField(itemsFields []ItemsField, itemsArgs []ItemsArgument) GroupsField {
	items := Items.List(itemsFields, itemsArgs...)
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

// The group's graphql argument(s).
type GroupsArgument struct {
	arg argument
}

// A list of group unique identifiers.
func NewGroupsIDsArgument(ids []string) GroupsArgument {
	return GroupsArgument{argument{"ids", ids}}
}
