package monday

func newColumnValuesWithArguments(valuesFields []ColumnValuesField, valuesArguments []ColumnValuesArgument) Query {
	if len(valuesFields) == 0 {
		valuesFields = append(valuesFields, ColumnValuesIDField())
	}
	var fields []field
	for _, vf := range valuesFields {
		fields = append(fields, vf.field)
	}
	var args []argument
	for _, va := range valuesArguments {
		args = append(args, va.arg)
	}
	return Query{
		name:   "column_values",
		fields: fields,
		args:   args,
	}
}

type ColumnValuesField struct {
	field field
}

var (
	columnValuesAdditionalInfoField = ColumnValuesField{field{"additional_info", nil}}
	columnValuesIDField             = ColumnValuesField{field{"id", nil}}
	columnValuesTextField           = ColumnValuesField{field{"text", nil}}
	columnValuesTitleField          = ColumnValuesField{field{"title", nil}}
	columnValuesValueField          = ColumnValuesField{field{"value", nil}}
)

// The column value's additional information.
func ColumnValuesAdditionalInfoField() ColumnValuesField {
	return columnValuesAdditionalInfoField
}

// The column's unique identifier.
func ColumnValuesIDField() ColumnValuesField {
	return columnValuesIDField
}

// The column's textual value in string form.
func ColumnValuesTextField() ColumnValuesField {
	return columnValuesTextField
}

// The column's title.
func ColumnValuesTitleField() ColumnValuesField {
	return columnValuesTitleField
}

// The column's value in json format.
func ColumnValuesValueField() ColumnValuesField {
	return columnValuesValueField
}

type ColumnValuesArgument struct {
	arg argument
}

// A list of column ids to return.
func NewColumnValuesIDsArgument(ids []string) ColumnValuesArgument {
	return ColumnValuesArgument{argument{"ids", ids}}
}
