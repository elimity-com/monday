package monday

import (
	"fmt"
	"strings"
)

type Tags struct {
	fields []TagsField
	args   []TagsArgument
}

func (t Tags) stringify() string {
	fields := make([]string, 0)
	for _, field := range t.fields {
		fields = append(fields, field.stringify())
	}
	args := make([]string, 0)
	for _, arg := range t.args {
		args = append(args, arg.stringify())
	}
	if len(fields) == 0 {
		return ``
	}
	if len(args) == 0 {
		return fmt.Sprintf(`tags{%s}`, strings.Join(fields, " "))
	}
	return fmt.Sprintf(`tags(%s){%s}`, strings.Join(args, ","), strings.Join(fields, " "))
}

func NewTags(fields []TagsField) Tags {
	if len(fields) == 0 {
		return Tags{
			fields: []TagsField{
				TagsIDField(),
			},
		}
	}

	return Tags{
		fields: fields,
	}
}

func NewTagsWithArguments(fields []TagsField, args []TagsArgument) Tags {
	tags := NewTags(fields)
	tags.args = args
	return tags
}

type TagsField struct {
	field string
}

var (
	tagsColorField = TagsField{"color"}
	tagsIDField    = TagsField{"id"}
	tagsNameField  = TagsField{"name"}
)

func (f TagsField) stringify() string {
	return fmt.Sprint(f.field)
}

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
	argument string
	value    interface{}
}

func (a TagsArgument) stringify() string {
	switch a.argument {
	case "ids":
		switch ids := a.value.([]int); {
		case len(ids) == 1:
			return fmt.Sprintf("ids:%d", ids[0])
		case len(ids) > 1:
			return fmt.Sprintf("ids:%s", strings.Replace(fmt.Sprint(ids), " ", ",", -1))
		default:
			return ""
		}
	default:
		return fmt.Sprintf("%s:%v", a.argument, a.value)
	}
}

// A list of tags unique identifiers.
func NewIDsTagsArg(ids []int) TagsArgument {
	return TagsArgument{
		argument: "ids",
		value:    ids,
	}
}
