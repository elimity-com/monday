package monday

import (
	"fmt"
	"strings"
)

type Complexity struct {
	fields []ComplexityField
}

func (c Complexity) stringify() string {
	fields := make([]string, 0)
	for _, field := range c.fields {
		fields = append(fields, field.stringify())
	}
	if len(fields) == 0 {
		return ``
	}
	return fmt.Sprintf(`complexity{%s}`, strings.Join(fields, " "))
}

func NewComplexity(fields []ComplexityField) Complexity {
	if len(fields) == 0 {
		return Complexity{
			fields: []ComplexityField{
				ComplexityAfterField(),
				ComplexityBeforeField(),
				ComplexityQueryField(),
			},
		}
	}

	return Complexity{
		fields: fields,
	}
}

type ComplexityField struct {
	field string
}

var (
	complexityAfterField  = ComplexityField{"after"}
	complexityBeforeField = ComplexityField{"before"}
	complexityQueryField  = ComplexityField{"query"}
)

func (f ComplexityField) stringify() string {
	return fmt.Sprint(f.field)
}

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
