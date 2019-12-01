package monday

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
