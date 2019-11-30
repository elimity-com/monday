package monday

import (
	"fmt"
	"strings"
)

type Columns struct {
	fields []ColumnsField
}

func (c Columns) stringify() string {
	fields := make([]string, 0)
	for _, field := range c.fields {
		fields = append(fields, field.stringify())
	}
	if len(fields) == 0 {
		return ``
	}
	return fmt.Sprintf("columns{%s}", strings.Join(fields, " "))
}

func NewColumns(fields []ColumnsField) Columns {
	if len(fields) == 0 {
		return Columns{
			fields: []ColumnsField{
				ColumnsIDField(),
			},
		}
	}
	return Columns{
		fields: fields,
	}
}

type ColumnsField struct {
	field string
}

var (
	columnsArchivedField    = ColumnsField{"archived"}
	columnsIDField          = ColumnsField{"id"}
	columnsSettingsStrField = ColumnsField{"settings_str"}
	columnsTitleField       = ColumnsField{"title"}
	columnsTypeField        = ColumnsField{"type"}
	columnsWidthField       = ColumnsField{"width"}
)

func (f ColumnsField) stringify() string {
	return fmt.Sprint(f.field)
}

func ColumnsArchivedField() ColumnsField {
	return columnsArchivedField
}

func ColumnsIDField() ColumnsField {
	return columnsIDField
}

func ColumnsSettingsStrField() ColumnsField {
	return columnsSettingsStrField
}

func ColumnsTitleField() ColumnsField {
	return columnsTitleField
}

func ColumnsTypeField() ColumnsField {
	return columnsTypeField
}

func ColumnsWidthField() ColumnsField {
	return columnsWidthField
}
