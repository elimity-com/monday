package monday

import (
	"fmt"
	"strings"
)

type Payload struct {
	queries []Query
}

func NewQueryPayload(queries ...Query) Payload {
	return Payload{queries: queries}
}

type Query struct {
	name   string
	fields []field
	args   []argument
}

func (q Query) stringify() string {
	fields := make([]string, 0)
	for _, field := range q.fields {
		fields = append(fields, field.stringify())
	}
	args := make([]string, 0)
	for _, arg := range q.args {
		args = append(args, arg.stringify())
	}
	if len(fields) == 0 {
		return ``
	}
	if len(args) == 0 {
		return fmt.Sprintf(`%s{%s}`, q.name, strings.Join(fields, " "))
	}
	return fmt.Sprintf(`%s(%s){%s}`, q.name, strings.Join(args, ","), strings.Join(fields, " "))
}

type field struct {
	field string
	value *Query
}

func (f field) stringify() string {
	if f.value != nil {
		return f.value.stringify()
	}
	return fmt.Sprint(f.field)
}

type argument struct {
	argument string
	value    interface{}
}

func (a argument) stringify() string {
	switch a.argument {
	case "column_id", "column_type", "column_value":
		return fmt.Sprintf("%s:%q", a.argument, a.value)
	case "ids":
		if strs, ok := a.value.([]string); ok {
			switch {
			case len(strs) == 1:
				return fmt.Sprintf("ids:%q", strs[0])
			case len(strs) > 1:
				return fmt.Sprintf("ids:%v", strings.Replace(fmt.Sprintf("%q", strs), " ", ",", -1))
			default:
				return ""
			}
		}
		switch ids := a.value.([]int); {
		case len(ids) == 1:
			return fmt.Sprintf("ids:%v", ids[0])
		case len(ids) > 1:
			return fmt.Sprintf("ids:%v", strings.Replace(fmt.Sprint(ids), " ", ",", -1))
		default:
			return ""
		}
	default:
		return fmt.Sprintf("%s:%v", a.argument, a.value)
	}
}
