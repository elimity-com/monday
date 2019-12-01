package monday

func NewTeams(teamsFields []TeamsField) Query {
	if len(teamsFields) == 0 {
		return Query{
			name: "teams",
			fields: []field{
				TeamsIDField().field,
			},
		}
	}

	var fields []field
	for _, tf := range teamsFields {
		fields = append(fields, tf.field)
	}
	return Query{
		name:   "teams",
		fields: fields,
	}
}

func NewTeamsWithArguments(teamsFields []TeamsField, teamsArgs []TeamsArgument) Query {
	teams := NewTeams(teamsFields)
	var args []argument
	for _, ta := range teamsArgs {
		args = append(args, ta.arg)
	}
	teams.args = args
	return teams
}

type TeamsField struct {
	field field
}

var (
	teamsIDField         = TeamsField{field{"id", nil}}
	teamsNameField       = TeamsField{field{"name", nil}}
	teamsPictureURLField = TeamsField{field{"picture_url", nil}}
)

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
func NewTeamsUsersField(usersFields []UsersField, usersArguments []UsersArgument) TeamsField {
	users := NewUsersWithArguments(usersFields, usersArguments)
	return TeamsField{field{"users", &users}}
}

type TeamsArgument struct {
	arg argument
}

// A list of teams unique identifiers.
func NewTeamsIDsArgument(ids []int) TeamsArgument {
	return TeamsArgument{argument{"ids", ids}}
}
