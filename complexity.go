package monday

func NewComplexity(complexityFields []ComplexityField) Query {
	if len(complexityFields) == 0 {
		return Query{
			name: "complexity",
			fields: []field{
				ComplexityAfterField().field,
				ComplexityBeforeField().field,
				ComplexityQueryField().field,
			},
		}
	}

	var fields []field
	for _, cf := range complexityFields {
		fields = append(fields, cf.field)
	}
	return Query{
		name: "complexity",
		fields: fields,
	}
}

type ComplexityField struct {
	field field
}

var (
	complexityAfterField  = ComplexityField{field{"after", nil}}
	complexityBeforeField = ComplexityField{field{"before", nil}}
	complexityQueryField  = ComplexityField{field{"query", nil}}
)

// The remainder of complexity after the query's execution.
func ComplexityAfterField() ComplexityField {
	return complexityAfterField
}

// The remainder of complexity before the query's execution.
func ComplexityBeforeField() ComplexityField {
	return complexityBeforeField
}

// The specific query's complexity.
func ComplexityQueryField() ComplexityField {
	return complexityQueryField
}
