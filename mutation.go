package monday

import (
	"fmt"
	"strings"
)

func NewMutationPayload(mutations ...Mutation) Payload {
	return Payload{mutations: mutations}
}

type Mutation struct {
	name   string
	fields []field
	args   []argument
}

func (m Mutation) stringify() string {
	fields := make([]string, 0)
	for _, field := range m.fields {
		fields = append(fields, field.stringify())
	}
	args := make([]string, 0)
	for _, arg := range m.args {
		args = append(args, arg.stringify())
	}
	if len(fields) == 0 {
		return ``
	}
	if len(args) == 0 {
		return fmt.Sprintf(`%s{%s}`, m.name, strings.Join(fields, " "))
	}
	return fmt.Sprintf(`%s(%s){%s}`, m.name, strings.Join(args, ","), strings.Join(fields, " "))

}
