package monday

func NewTags(tagsFields []TagsField) Query {
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
	return Query{
		name:   "tags",
		fields: fields,
	}
}

func NewTagsWithArguments(tagsFields []TagsField, tagsArgs []TagsArgument) Query {
	tags := NewTags(tagsFields)
	var args []argument
	for _, ta := range tagsArgs {
		args = append(args, ta.arg)
	}
	tags.args = args
	return tags
}

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

type TagsArgument struct {
	arg argument
}

// A list of tags unique identifiers.
func NewTagsIDsArgument(ids []int) TagsArgument {
	return TagsArgument{argument{argument: "ids", value: ids}}
}
