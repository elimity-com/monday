package monday

func CreateColumn(boardID int, title string, columnsType ColumnsType, columnsFields []ColumnsField) Mutation {
	if len(columnsFields) == 0 {
		columnsFields = append(columnsFields, columnsIDField)
	}

	var fields []field
	for _, cf := range columnsFields {
		fields = append(fields, cf.field)
	}
	return Mutation{
		name:   "create_column",
		fields: fields,
		args: []argument{
			{"board_id", boardID},
			{"title", title},
			{"column_type", columnsType.typ},
		},
	}
}

func ChangeColumnValue(itemID int, columnID string, boardID int, value string, itemsFields []ItemsField) Mutation {
	if len(itemsFields) == 0 {
		itemsFields = append(itemsFields, itemsIDField)
	}

	var fields []field
	for _, i := range itemsFields {
		fields = append(fields, i.field)
	}
	return Mutation{
		name:   "change_column_value",
		fields: fields,
		args: []argument{
			{"item_id", itemID},
			{"column_id", columnID},
			{"board_id", boardID},
			{"value", value},
		},
	}
}

func ChangeMultipleColumnValues(itemID int, boardID int, values string, itemsFields []ItemsField) Mutation {
	var fields []field
	for _, i := range itemsFields {
		fields = append(fields, i.field)
	}
	return Mutation{
		name:   "change_multiple_column_values",
		fields: fields,
		args: []argument{
			{"item_id", itemID},
			{"board_id", boardID},
			{"column_values", values},
		},
	}
}

func NewColumns(columnFields []ColumnsField) Query {
	if len(columnFields) == 0 {
		return Query{
			name: "columns",
			fields: []field{
				ColumnsIDField().field,
			},
		}
	}

	var fields []field
	for _, cf := range columnFields {
		fields = append(fields, cf.field)
	}
	return Query{
		name:   "columns",
		fields: fields,
	}
}

type ColumnsField struct {
	field field
}

var (
	columnsArchivedField    = ColumnsField{field{"archived", nil}}
	columnsIDField          = ColumnsField{field{"id", nil}}
	columnsSettingsStrField = ColumnsField{field{"settings_str", nil}}
	columnsTitleField       = ColumnsField{field{"title", nil}}
	columnsTypeField        = ColumnsField{field{"type", nil}}
	columnsWidthField       = ColumnsField{field{"width", nil}}
)

// Is the column archived or not.
func ColumnsArchivedField() ColumnsField {
	return columnsArchivedField
}

// The column's unique identifier.
func ColumnsIDField() ColumnsField {
	return columnsIDField
}

// The column's settings in a string form.
func ColumnsSettingsStrField() ColumnsField {
	return columnsSettingsStrField
}

// The column's title.
func ColumnsTitleField() ColumnsField {
	return columnsTitleField
}

// The column's type.
func ColumnsTypeField() ColumnsField {
	return columnsTypeField
}

// The column's width.
func ColumnsWidthField() ColumnsField {
	return columnsWidthField
}

type ColumnsType struct {
	typ string
}

var (
	columnsTypeAutoNumber   = ColumnsType{"auto_number"}
	columnsTypeCheckbox     = ColumnsType{"checkbox"}
	columnsTypeCountry      = ColumnsType{"country"}
	columnsTypeColorPicker  = ColumnsType{"color_picker"}
	columnsTypeCreationLog  = ColumnsType{"creation_log"}
	columnsTypeDate         = ColumnsType{"date"}
	columnsTypeDropdown     = ColumnsType{"dropdown"}
	columnsTypeEmail        = ColumnsType{"email"}
	columnsTypeHour         = ColumnsType{"hour"}
	columnsTypeItemID       = ColumnsType{"item_id"}
	columnsTypeLastUpdated  = ColumnsType{"last_updated"}
	columnsTypeLink         = ColumnsType{"link"}
	columnsTypeLocation     = ColumnsType{"location"}
	columnTypeLongText      = ColumnsType{"long_text"}
	columnsTypeNumbers      = ColumnsType{"numbers"}
	columnsTypePeople       = ColumnsType{"people"}
	columnsTypePhone        = ColumnsType{"phone"}
	columnsTypeProgress     = ColumnsType{"progress"}
	columnsTypeRating       = ColumnsType{"rating"}
	columnsTypeStatus       = ColumnsType{"status"}
	columnsTypeTeam         = ColumnsType{"team"}
	columnsTypeTags         = ColumnsType{"tags"}
	columnsTypeText         = ColumnsType{"text"}
	columnsTypeTimeline     = ColumnsType{"timeline"}
	columnsTypeTimeTracking = ColumnsType{"time_tracking"}
	columnsTypeVote         = ColumnsType{"vote"}
	columnsTypeWeek         = ColumnsType{"week"}
	columnsTypeWorldClock   = ColumnsType{"world_clock"}
)

// Number items according to their order in the group/board.
func ColumnsTypeAutoNumber() ColumnsType {
	return columnsTypeAutoNumber
}

// Check off items and see what's done at a glance.
func ColumnsTypeCheckBox() ColumnsType {
	return columnsTypeCheckbox
}

// Choose a country.
func ColumnsTypeCountry() ColumnsType {
	return columnsTypeCountry
}

// Manage a design system using a color palette.
func ColumnsTypeColorPicker() ColumnsType {
	return columnsTypeColorPicker
}

// Add the item creator and creation date automatically.
func ColumnsTypeCreationLog() ColumnsType {
	return columnsTypeCreationLog
}

// Add dates like deadlines to ensure you never drop the ball.
func ColumnsTypeDate() ColumnsType {
	return columnsTypeDate
}

// Create a dropdown list of options.
func ColumnsTypeDropdown() ColumnsType {
	return columnsTypeDropdown
}

// Email team members and clients directly from your board.
func ColumnsTypeEmail() ColumnsType {
	return columnsTypeEmail
}

// Add times to manage and schedule tasks, shifts and more.
func ColumnsTypeHour() ColumnsType {
	return columnsTypeHour
}

// Show a unique ID for each item.
func ColumnsTypeItemID() ColumnsType {
	return columnsTypeItemID
}

// Add the person that last updated the item and the date.
func ColumnsTypeLastUpdated() ColumnsType {
	return columnsTypeLastUpdated
}

// Simply hyperlink to any website.
func ColumnsTypeLink() ColumnsType {
	return columnsTypeLink
}

// Place multiple locations on a geographic map.
func ColumnsTypeLocation() ColumnsType {
	return columnsTypeLocation
}

// Add large amounts of text without changing column width.
func ColumnsTypeLongText() ColumnsType {
	return columnTypeLongText
}

// Add revenue, costs, time estimations and more.
func ColumnsTypeNumbers() ColumnsType {
	return columnsTypeNumbers
}

// Assign people to improve team work.
func ColumnsTypePeople() ColumnsType {
	return columnsTypePeople
}

// Call your contacts directly from monday.com.
func ColumnsTypePhone() ColumnsType {
	return columnsTypePhone
}

// Show progress by combining status columns in a battery.
func ColumnsTypeProgress() ColumnsType {
	return columnsTypeProgress
}

// Rate or rank anything visually.
func ColumnsTypeRating() ColumnsType {
	return columnsTypeRating
}

// Get an instant overview of where things stand.
func ColumnsTypeStatus() ColumnsType {
	return columnsTypeStatus
}

// Assign a full team to an item.
func ColumnsTypeTeam() ColumnsType {
	return columnsTypeTeam
}

// Add tags to categorize items across multiple boards.
func ColumnsTypeTags() ColumnsType {
	return columnsTypeTags
}

// Add textual information e.g. addresses, names or keywords.
func ColumnsTypeText() ColumnsType {
	return columnsTypeText
}

// Visually see a breakdown of your team's workload by time.
func ColumnsTypeTimeline() ColumnsType {
	return columnsTypeTimeline
}

// Basic time tracking for every item.
func ColumnsTypeTimeTracking() ColumnsType {
	return columnsTypeTimeTracking
}

// ​Vote on an item e.g. pick a new feature or a favorite lunch place.
func ColumnsTypeVote() ColumnsType {
	return columnsTypeVote
}

// Select the week on which each item should be completed.
func ColumnsTypeWeek() ColumnsType {
	return columnsTypeWeek
}

// Keep track of the time anywhere in the world.
func ColumnsTypeWorldClock() ColumnsType {
	return columnsTypeWorldClock
}
