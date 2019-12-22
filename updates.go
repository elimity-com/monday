package monday

// UpdateService handles all the update related methods of the Monday API.
// Updates are additional notes and information added to items outside of the structure of the board.
// The main form of communication within the platform takes place in the updates section.
type UpdateService service

// Create returns a mutation that allows you to add an update to a item.
// After the mutation runs you can query back all the board data.
// - id: the item's unique identifier.
// - body: the update text.
//
// DOCS: https://monday.com/developers/v2#mutations-section-updates
func (*UpdateService) Create(id int, body string, updatesFields []UpdatesField) Mutation {
	if len(updatesFields) == 0 {
		updatesFields = append(updatesFields, updatesIDField)
	}

	var fields []field
	for _, uf := range updatesFields {
		fields = append(fields, uf.field)
	}
	return Mutation{
		name:   "create_update",
		fields: fields,
		args: []argument{
			{"item_id", id},
			{"body", body},
		},
	}
}

// List returns a query that gets one or a collection of updates.
//
// DOCS: https://monday.com/developers/v2#queries-section-updates
func (*UpdateService) List(updatesFields []UpdatesField, updatesArgs ...UpdatesArgument) Query {
	if len(updatesFields) == 0 {
		return Query{
			name: "updates",
			fields: []field{
				UpdatesIDField().field,
			},
		}
	}

	var fields []field
	for _, uf := range updatesFields {
		fields = append(fields, uf.field)
	}
	var args []argument
	for _, ua := range updatesArgs {
		args = append(args, ua.arg)
	}
	return Query{
		name:   "updates",
		fields: fields,
		args:   args,
	}
}

// The update's graphql field(s).
type UpdatesField struct {
	field field
}

var (
	updatesBodyField      = UpdatesField{field{"body", nil}}
	updatesCreatedAtField = UpdatesField{field{"created_at", nil}}
	updatesCreatorIDField = UpdatesField{field{"creator_id", nil}}
	updatesIDField        = UpdatesField{field{"id", nil}}
	updatesItemIDField    = UpdatesField{field{"item_id", nil}}
	updatesTextBodyField  = UpdatesField{field{"text_body", nil}}
	updatesUpdatedAtField = UpdatesField{field{"updated_at", nil}}
)

// The update's html formatted body.
func UpdatesBodyField() UpdatesField {
	return updatesBodyField
}

// The update's creation date.
func UpdatesCreatedAtField() UpdatesField {
	return updatesCreatedAtField
}

// The update's creator.
func NewUpdatesCreatorField(creatorFields []UsersField, creatorArgs ...UsersArgument) UpdatesField {
	creator := Users.List(creatorFields, creatorArgs...)
	creator.name = "creator"
	return UpdatesField{field{creator.name, &creator}}
}

// The unique identifier of the update creator.
func UpdatesCreatorIDField() UpdatesField {
	return updatesCreatorIDField
}

// The update's unique identifier.
func UpdatesIDField() UpdatesField {
	return updatesIDField
}

// The update's item ID.
func UpdatesItemIDField() UpdatesField {
	return updatesItemIDField
}

// The update's replies.
func NewUpdatesRepliesField(repliesFields []RepliesField) UpdatesField {
	var fields []field
	for _, rf := range repliesFields {
		fields = append(fields, rf.field)
	}
	if len(fields) == 0 {
		fields = append(fields, RepliesIDField().field)
	}
	replies := Query{
		name:   "replies",
		fields: fields,
	}
	return UpdatesField{field{"replies", &replies}}
}

// The update's text body.
func UpdatesTextBodyField() UpdatesField {
	return updatesTextBodyField
}

// The update's last edit date.
func UpdatesUpdatedAtField() UpdatesField {
	return updatesUpdatedAtField
}

// The update's graphql argument(s).
type UpdatesArgument struct {
	arg argument
}

// Number of items to get, the default is 25.
func NewUpdatesLimitArgument(value int) UpdatesArgument {
	return UpdatesArgument{argument{"limit", value}}
}

// Page number to get, starting at 1.
func NewUpdatesPageArgument(value int) UpdatesArgument {
	return UpdatesArgument{argument{"page", value}}
}
