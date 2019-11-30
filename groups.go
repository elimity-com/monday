package monday

import (
	"fmt"
	"strings"
)

type Groups struct {
	alt    string
	fields []GroupsField
	args   []GroupsArgument
}

func (g Groups) stringify() string {
	prefix := "groups"
	if g.alt != "" {
		prefix = g.alt
	}

	fields := make([]string, 0)
	for _, field := range g.fields {
		fields = append(fields, field.stringify())
	}
	args := make([]string, 0)
	for _, arg := range g.args {
		args = append(args, arg.stringify())
	}
	if len(fields) == 0 {
		return ``
	}
	if len(args) == 0 {
		return fmt.Sprintf(`%s{%s}`, prefix, strings.Join(fields, " "))
	}
	return fmt.Sprintf(`%s(%s){%s}`, prefix, strings.Join(args, ","), strings.Join(fields, " "))

}

func NewGroups(fields []GroupsField) Groups {
	if len(fields) == 0 {
		return Groups{
			fields: []GroupsField{
				GroupsIDField(),
			},
		}
	}

	return Groups{
		fields: fields,
	}
}

func NewGroupWithArguments(fields []GroupsField, args []GroupsArgument) Groups {
	groups := NewGroups(fields)
	groups.args = args
	return groups
}

type GroupsField struct {
	field string
	value interface{}
}

var (
	groupsArchivedField = GroupsField{"archived", nil}
	groupsColorField    = GroupsField{"color", nil}
	groupsDeletedField  = GroupsField{"deleted", nil}
	groupsIDField       = GroupsField{"id", nil}
	groupsPositionField = GroupsField{"position", nil}
	groupsTitleField    = GroupsField{"title", nil}
)

func (f GroupsField) stringify() string {
	switch f.field {
	case "items":
		return f.value.(Items).stringify()
	default:
		return fmt.Sprint(f.field)
	}
}

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
func NewGroupsItemsField(items Items) GroupsField {
	return GroupsField{field: "items", value: items}
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
	argument string
	value    interface{}
}

func (a GroupsArgument) stringify() string {
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

// A list of group unique identifiers.
func NewIDsGroupsArg(ids []int) GroupsArgument {
	return GroupsArgument{
		argument: "ids",
		value:    ids,
	}
}
