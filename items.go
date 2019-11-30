package monday

import (
	"fmt"
	"log"
	"strings"
)

type Items struct {
	fields []ItemsField
	args   []ItemsArgument
}

func (i Items) stringify() string {
	fields := make([]string, 0)
	for _, field := range i.fields {
		fields = append(fields, field.stringify())
	}
	args := make([]string, 0)
	for _, arg := range i.args {
		args = append(args, arg.stringify())
	}
	if len(fields) == 0 {
		return ``
	}
	if len(args) == 0 {
		return fmt.Sprintf(`items{%s}`, strings.Join(fields, " "))
	}
	return fmt.Sprintf(`items(%s){%s}`, strings.Join(args, ","), strings.Join(fields, " "))
}

func NewItems(fields []ItemsField) Items {
	if len(fields) == 0 {
		return Items{
			fields: []ItemsField{
				ItemsIDField(),
			},
		}
	}

	return Items{
		fields: fields,
	}
}

func NewItemsWithArguments(fields []ItemsField, args []ItemsArgument) Items {
	items := NewItems(fields)
	items.args = args
	return items
}

type ItemsField struct {
	field string
	value interface{}
}

var (
	itemsCreatedAtField = ItemsField{"created_at", nil}
	itemsCreatorIDField = ItemsField{"creator_id", nil}
	itemsIDField        = ItemsField{"id", nil}
	itemsNameField      = ItemsField{"name", nil}
	itemsStateField     = ItemsField{"state", nil}
	itemsUpdatedAtField = ItemsField{"updated_at", nil}
)

func (f ItemsField) stringify() string {
	switch f.field {
	case "column_values":
		return f.value.(ColumnValues).stringify()
	case "group":
		group := f.value.(Groups)
		group.alt = "group"
		return group.stringify()
	case "subscribers":
		return f.value.(Users).stringify()
	case "updates":
		return f.value.(Updates).stringify()
	default:
		return fmt.Sprint(f.field)
	}
}

// The board that contains this item.¬
func NewItemsBoardsField(boards Boards) ItemsField {
	return ItemsField{field: "boards", value: boards}
}

// The item's column values.
func NewItemsColumnValuesField(values ColumnValues) ItemsField {
	return ItemsField{field: "column_values", value: values}
}

// The item's create date.
func ItemsCreatedAtField() ItemsField {
	return itemsCreatedAtField
}

// The item's creator.
func NewItemsCreatorField(creator Users) ItemsField {
	return ItemsField{field: "creator", value: creator}
}

// The unique identifier of the item creator.
func ItemsCreatorIDField() ItemsField {
	return itemsCreatorIDField
}

// The group that contains this item.
func NewItemsGroupsField(groups Groups) ItemsField {
	return ItemsField{field: "groups", value: groups}
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
func NewItemsSubscribersField(subscribers Users) ItemsField {
	return ItemsField{field: "subscribers", value: subscribers}
}

// The item's last update date.
func ItemsUpdatedAtField() ItemsField {
	return itemsUpdatedAtField
}

// The item's updates.
func NewItemsUpdatesField(updates Updates) ItemsField {
	return ItemsField{field: "updates", value: updates}
}

type ItemsArgument struct {
	argument string
	value    interface{}
}

func (a ItemsArgument) stringify() string {
	switch a.argument {
	case "limit", "page", "newest_first":
		return fmt.Sprintf("%s:%v", a.argument, a.value)
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
		log.Fatalln("unreachable boards argument")
	}
	return ""
}

// Number of items to get, the default is 25.
func NewLimitItemsArg(value int) ItemsArgument {
	return ItemsArgument{
		argument: "limit",
		value:    value,
	}
}

// Page number to get, starting at 1.
func NewPageItemsArg(value int) ItemsArgument {
	return ItemsArgument{
		argument: "page",
		value:    value,
	}
}

// A list of items unique identifiers.
func NewIDsItemsArg(ids []int) ItemsArgument {
	return ItemsArgument{
		argument: "ids",
		value:    ids,
	}
}

// Get the recently created items at the top of the list.
func NewNewestFirstItemsArg(first bool) ItemsArgument {
	return ItemsArgument{
		argument: "newest_first",
		value:    first,
	}
}
