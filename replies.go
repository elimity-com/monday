package monday

// The reply's graphql field(s).
type RepliesField struct {
	field field
}

var (
	replyBodyField      = RepliesField{field{"body", nil}}
	replyCreatedAtField = RepliesField{field{"created_at", nil}}
	replyCreatorIDField = RepliesField{field{"creator_id", nil}}
	replyIDField        = RepliesField{field{"id", nil}}
	replyTextBodyField  = RepliesField{field{"text_body", nil}}
	replyUpdatedAtField = RepliesField{field{"updated_at", nil}}
)

// The reply's html formatted body.
func RepliesBodyField() RepliesField {
	return replyBodyField
}

// The reply's creation date.
func RepliesCreatedAtField() RepliesField {
	return replyCreatedAtField
}

// The reply's creator.
func NewRepliesCreatorField(creatorFields []UsersField, creatorArgs ...UsersArgument) RepliesField {
	creator := Users.List(creatorFields, creatorArgs...)
	creator.name = "creator"
	return RepliesField{field{creator.name, &creator}}
}

// The unique identifier of the reply creator.
func RepliesCreatorIDField() RepliesField {
	return replyCreatorIDField
}

// The reply's unique identifier.
func RepliesIDField() RepliesField {
	return replyIDField
}

// The reply's text body.
func RepliesTextBodyField() RepliesField {
	return replyTextBodyField
}

// The reply's last edit date.
func RepliesUpdatedAtField() RepliesField {
	return replyUpdatedAtField
}
