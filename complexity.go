package monday

// ComplexityService handles all the complexity related methods of the Monday API.
// The amount of API calls each account can perform is limited by complexity per minute.
// Querying for complexity allows you to determine the complexity of your query and how much of your 1 minute cap will be used.
type ComplexityService service

// Simply add the complexity query to your query and you’ll receive another field which represents the complexity level.
//
// DOCS: https://monday.com/developers/v2#queries-section-complexity
func (*ComplexityService) List(complexityFields []ComplexityField) Query {
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
		name:   "complexity",
		fields: fields,
	}
}

// The complexity's graphql field(s).
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
