package monday

import (
	"fmt"
	"strings"
)

type Teams struct {
	fields []TeamsField
	args   []TeamsArgument
}

func (t Teams) stringify() string {
	fields := make([]string, 0)
	for _, field := range t.fields {
		fields = append(fields, field.stringify())
	}
	args := make([]string, 0)
	for _, arg := range t.args {
		args = append(args, arg.stringify())
	}
	if len(fields) == 0 {
		return ``
	}
	if len(args) == 0 {
		return fmt.Sprintf(`teams{%s}`, strings.Join(fields, " "))
	}
	return fmt.Sprintf(`teams(%s){%s}`, strings.Join(args, ","), strings.Join(fields, " "))
}

func NewTeams(fields []TeamsField) Teams {
	if len(fields) == 0 {
		return Teams{
			fields: []TeamsField{
				TeamsIDField(),
			},
		}
	}

	return Teams{
		fields: fields,
	}
}

func NewTeamsWithArguments(fields []TeamsField, args []TeamsArgument) Teams {
	teams := NewTeams(fields)
	teams.args = args
	return teams
}

type TeamsField struct {
	field string
	value interface{}
}

var (
	teamsIDField         = TeamsField{"id", nil}
	teamsNameField       = TeamsField{"name", nil}
	teamsPictureURLField = TeamsField{"picture_url", nil}
)

func (f TeamsField) stringify() string {
	switch f.field {
	case "users":
		return f.value.(Users).stringify()
	default:
		return fmt.Sprint(f.field)
	}
}

// The team's unique identifier.
func TeamsIDField() TeamsField {
	return teamsIDField
}

// The team's name.
func TeamsNameField() TeamsField {
	return teamsNameField
}

// The team's picture url.
func TeamsPictureURLField() TeamsField {
	return teamsPictureURLField
}

// The users in the team.
func NewTeamsUsersField(users Users) TeamsField {
	return TeamsField{field: "users", value: users}
}

type TeamsArgument struct {
	argument string
	value    interface{}
}

func (a TeamsArgument) stringify() string {
	switch a.argument {
	case "ids":
		switch ids := a.value.([]int); {
		case len(ids) == 1:
			return fmt.Sprintf("ids:%d", ids[0])
		case len(ids) > 1:
			return fmt.Sprintf("ids:%s", strings.Replace(fmt.Sprint(ids), " ", ",", -1))
		default:
			return ""
		}
	default:
		return fmt.Sprintf("%s:%v", a.argument, a.value)
	}
}

// A list of teams unique identifiers.
func NewIDsTeamsArg(ids []int) TeamsArgument {
	return TeamsArgument{
		argument: "ids",
		value:    ids,
	}
}
