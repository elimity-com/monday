package monday

func CreateUpdate(itemID int, body string, updatesFields []UpdatesField) Mutation {
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
			{"item_id", itemID},
			{"body", body},
		},
	}
}

func NewUpdates(updatesFields []UpdatesField) Query {
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
	return Query{
		name:   "updates",
		fields: fields,
	}
}

func NewUpdatesWithArguments(updatesFields []UpdatesField, updatesArgs []UpdatesArgument) Query {
	updates := NewUpdates(updatesFields)
	var args []argument
	for _, ua := range updatesArgs {
		args = append(args, ua.arg)
	}
	updates.args = args
	return updates
}

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
	return updatesUpdatedAtField
}

// The update's creator.
func NewUpdatesCreatorField(creatorFields []UsersField, creatorArguments []UsersArgument) UpdatesField {
	creator := NewUsersWithArguments(creatorFields, creatorArguments)
	creator.name = "creator"
	return UpdatesField{field{"creator", &creator}}
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
	replies := newReplies(repliesFields)
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
