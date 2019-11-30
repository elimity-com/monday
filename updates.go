package monday

import (
	"fmt"
	"strings"
)

type Updates struct {
	fields []UpdatesField
	args   []UpdatesArgument
}

func (u Updates) stringify() string {
	fields := make([]string, 0)
	for _, field := range u.fields {
		fields = append(fields, field.stringify())
	}
	args := make([]string, 0)
	for _, arg := range u.args {
		args = append(args, arg.stringify())
	}
	if len(fields) == 0 {
		return ``
	}
	if len(args) == 0 {
		return fmt.Sprintf(`updates{%s}`, strings.Join(fields, " "))
	}
	return fmt.Sprintf(`updates(%s){%s}`, strings.Join(args, ","), strings.Join(fields, " "))
}

func NewUpdates(fields []UpdatesField) Updates {
	if len(fields) == 0 {
		return Updates{
			fields: []UpdatesField{
				UpdatesIDField(),
			},
		}
	}

	return Updates{
		fields: fields,
	}
}

func NewUpdatesWithArguments(fields []UpdatesField, args []UpdatesArgument) Updates {
	updates := NewUpdates(fields)
	updates.args = args
	return updates
}

type UpdatesField struct {
	field string
	value interface{}
}

var (
	updatesBodyField      = UpdatesField{"body", nil}
	updatesCreatedAtField = UpdatesField{"created_at", nil}
	updatesCreatorIDField = UpdatesField{"creator_id", nil}
	updatesIDField        = UpdatesField{"id", nil}
	updatesItemIDField    = UpdatesField{"item_id", nil}
	// TODO: replies? nothing found in documentation
	updatesTextBodyField  = UpdatesField{"text_body", nil}
	updatesUpdatedAtField = UpdatesField{"updated_at", nil}
)

func (f UpdatesField) stringify() string {
	switch f.field {
	case "creator":
		creator := f.value.(Users)
		creator.alt = "creator"
		return creator.stringify()
	default:
		return fmt.Sprint(f.field)
	}
}

// The update's html formatted body.
func UpdatesBodyField() UpdatesField {
	return updatesBodyField
}

// The update's creation date.
func UpdatesCreatedAtField() UpdatesField {
	return updatesUpdatedAtField
}

// The update's creator.
func UpdatesCreatorField(creator Users) UpdatesField {
	return UpdatesField{field: "creator", value: creator}
}

// The unique identifier of the update creator.
func UpdatesCreatorIDField() UpdatesField {
	return updatesCreatorIDField
}

// The update's unique identifier.
func UpdatesIDField() UpdatesField {
	return updatesIDField
}

//The update's item ID.
func UpdatesItemIDField() UpdatesField {
	return updatesItemIDField
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
	argument string
	value    interface{}
}

func (a UpdatesArgument) stringify() string {
	return fmt.Sprintf("%s:%v", a.argument, a.value)
}

// Number of items to get, the default is 25.
func NewLimitUpdatesArg(value int) UpdatesArgument {
	return UpdatesArgument{
		argument: "limit",
		value:    value,
	}
}

// Page number to get, starting at 1.
func NewPageUpdatesArg(value int) UpdatesArgument {
	return UpdatesArgument{
		argument: "page",
		value:    value,
	}
}
