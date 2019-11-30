package monday

import (
	"fmt"
	"strings"
)

type ColumnValues struct {
	fields []ColumnValuesField
	args   []ColumnValuesArgument
}

func (v ColumnValues) stringify() string {
	fields := make([]string, 0)
	for _, field := range v.fields {
		fields = append(fields, field.stringify())
	}
	args := make([]string, 0)
	for _, arg := range v.args {
		args = append(args, arg.stringify())
	}
	if len(fields) == 0 {
		return ``
	}
	if len(args) == 0 {
		return fmt.Sprintf(`items_by_column_values {%s}`, strings.Join(fields, " "))
	}
	return fmt.Sprintf(`items_by_column_values (%s){%s}`, strings.Join(args, ","), strings.Join(fields, " "))

}

func NewColumnValues(fields []ColumnValuesField) ColumnValues {
	if len(fields) == 0 {
		return ColumnValues{
			fields: []ColumnValuesField{
				ColumnValuesIDField(),
			},
		}
	}

	return ColumnValues{
		fields: fields,
	}
}

func NewColumnValuesWithArguments(fields []ColumnValuesField, args []ColumnValuesArgument) ColumnValues {
	columnValues := NewColumnValues(fields)
	columnValues.args = args
	return columnValues
}

type ColumnValuesField struct {
	field string
	value interface{}
}

var (
	columnValuesCreatedAtField = ColumnValuesField{"created_at", nil}
	columnValuesCreatorIDField = ColumnValuesField{"creator_id", nil}
	columnValuesIDField        = ColumnValuesField{"id", nil}
	columnValuesNameField      = ColumnValuesField{"name", nil}
	columnValuesStateField     = ColumnValuesField{"state", nil}
	columnValuesUpdatedAtField = ColumnValuesField{"updated_at", nil}
)

func (f ColumnValuesField) stringify() string {
	switch f.field {
	case "boards":
		return f.value.(Boards).stringify()
	case "column_values":
		return f.value.(ColumnValues).stringify()
	case "creator":
		creator := f.value.(Users)
		creator.alt = "creator"
		return creator.stringify()
	default:
		return fmt.Sprint(f.field)
	}
}

// The board that contains this item.
func NewColumnValuesBoardsField(boards Boards) ColumnValuesField {
	return ColumnValuesField{field: "boards", value: boards}
}

// The item's column values.
func NewColumnValuesColumnValuesField(values ColumnValues) ColumnValuesField {
	return ColumnValuesField{field: "column_values", value: values}
}

// The item's create date.
func ColumnValuesCreatedAtField() ColumnValuesField {
	return columnValuesCreatedAtField
}

// The item's creator.
func NewColumnValuesCreatorField(creator Users) ColumnValuesField {
	return ColumnValuesField{field: "creator", value: creator}
}

// The unique identifier of the item creator.
func ColumnValuesCreatorIDField() ColumnValuesField {
	return columnValuesCreatorIDField
}

// The group that contains this item.
func NewColumnValuesGroupsField(groups Groups) ColumnValuesField {
	return ColumnValuesField{field: "groups", value: groups}
}

// The item's unique identifier.
func ColumnValuesIDField() ColumnValuesField {
	return columnValuesIDField
}

// The item's name.
func ColumnValuesNameField() ColumnValuesField {
	return columnValuesNameField
}

// The board's state (all / active / archived / deleted).
func ColumnValuesStateField() ColumnValuesField {
	return columnValuesStateField
}

// The pulses's subscribers.
func NewColumnValuesSubscribersField(subscribers Users) ColumnValuesField {
	return ColumnValuesField{field: "subscribers", value: subscribers}
}

// The item's last update date.
func ColumnValuesUpdatedAtField() ColumnValuesField {
	return columnValuesUpdatedAtField
}

// The item's updates.
func NewColumnValuesUpdatesField(updates Updates) ColumnValuesField {
	return ColumnValuesField{field: "updates", value: updates}
}

type ColumnValuesArgument struct {
	argument string
	value    interface{}
}

func (a ColumnValuesArgument) stringify() string {
	return fmt.Sprintf("%s:%v", a.argument, a.value)
}

// Number of items to get, the default is 25.
func NewLimitColumnValuesArg(value int) ColumnValuesArgument {
	return ColumnValuesArgument{
		argument: "limit",
		value:    value,
	}
}

// Page number to get, starting at 1.
func NewPageColumnValuesArg(value int) ColumnValuesArgument {
	return ColumnValuesArgument{
		argument: "page",
		value:    value,
	}
}

// The board's unique identifier.
func NewBoardIDColumnValuesArg(value int) ColumnValuesArgument {
	return ColumnValuesArgument{
		argument: "board_id",
		value:    value,
	}
}

// The column's unique identifier.
func NewColumnIDColumnValuesArg(value string) ColumnValuesArgument {
	return ColumnValuesArgument{
		argument: "column_id",
		value:    value,
	}
}

// The column value to search items by.
func NewColumnValueColumnValuesArg(value string) ColumnValuesArgument {
	return ColumnValuesArgument{
		argument: "column_value",
		value:    value,
	}
}

// The column type.
func NewColumnTypeColumnValuesArg(value string) ColumnValuesArgument {
	return ColumnValuesArgument{
		argument: "column_type",
		value:    value,
	}
}

// The state of the item (all / active / archived / deleted), the default is active.
func NewStateColumnValuesArg(state State) ColumnValuesArgument {
	return ColumnValuesArgument{
		argument: "state",
		value:    state,
	}
}
