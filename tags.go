package monday

// TagsService handles all the tag related methods of the Monday API.
// Tags are objects that help you group items from different groups or different boards throughout your account by
// a consistent keyword. Tag entities are created and presented through the “Tags” column
type TagsService service

// CreateOrGet returns a mutation that allows you to create new tags or receive their data if they already exist.
// - name: the (new) tag's name.
//
// DOCS: https://monday.com/developers/v2#mutations-section-tags
func (*TagsService) CreateOrGet(name string, tagsFields []TagsField) Mutation {
	if len(tagsFields) == 0 {
		tagsFields = append(tagsFields, tagsIDField)
	}

	var fields []field
	for _, tf := range tagsFields {
		fields = append(fields, tf.field)
	}
	return Mutation{
		name:   "create_or_get_tag",
		fields: fields,
		args: []argument{
			{"tag_name", name},
		},
	}
}

// CreateOrGetByBoard returns a mutation that allows you to create new tags or receive their data if they already exist.
// Private boards require separate tags so if you wish to use your tags in a private board use the 'board_id' argument.
// No need to send this argument for public boards.
// - name: the (new) tag's name.
// - boardID: the private board id to create the tag at (not needed for public boards).
//
// DOCS: https://monday.com/developers/v2#mutations-section-tags
func (*TagsService) CreateOrGetByBoard(name string, boardID int, tagsFields []TagsField) Mutation {
	tag := Tags.CreateOrGet(name, tagsFields)
	tag.args = append(tag.args, argument{"board_id", boardID})
	return tag
}

// List returns a query that gets one or a collection of the account's public tags.
// Public tags are the tags that appear in public boards.
// If you want to find tags of private boards you can use the boards query.
//
// DOCS: https://monday.com/developers/v2#queries-section-tags
func (*TagsService) List(tagsFields []TagsField, tagsArgs ...TagsArgument) Query {
	if len(tagsFields) == 0 {
		return Query{
			name: "tags",
			fields: []field{
				TagsIDField().field,
			},
		}
	}

	var fields []field
	for _, tf := range tagsFields {
		fields = append(fields, tf.field)
	}
	var args []argument
	for _, ta := range tagsArgs {
		args = append(args, ta.arg)
	}
	return Query{
		name:   "tags",
		fields: fields,
		args:   args,
	}
}

// The tag's graphql field(s).
type TagsField struct {
	field field
}

var (
	tagsColorField = TagsField{field{"color", nil}}
	tagsIDField    = TagsField{field{"id", nil}}
	tagsNameField  = TagsField{field{"name", nil}}
)

// The tag's color.
func TagsColorField() TagsField {
	return tagsColorField
}

// The tag's unique identifier.
func TagsIDField() TagsField {
	return tagsIDField
}

// The tag's name.
func TagsNameField() TagsField {
	return tagsNameField
}

// The tag's graphql argument(s).
type TagsArgument struct {
	arg argument
}

// A list of tags unique identifiers.
func NewTagsIDsArgument(ids []int) TagsArgument {
	return TagsArgument{argument{argument: "ids", value: ids}}
}
