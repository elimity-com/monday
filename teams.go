package monday

// TeamsService handles all the team related methods of the Monday API.
// Teams are the most efficient way to manage groups of users,
// Every team is comprised of one or multiple users, and every user can be a part of multiple teams (or none).
type TeamsService service

// List returns a query that gets one or several of teams.
//
// DOCS: https://monday.com/developers/v2#queries-section-teams
func (*TeamsService) List(teamsFields []TeamsField, teamsArgs ...TeamsArgument) Query {
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
	var args []argument
	for _, ta := range teamsArgs {
		args = append(args, ta.arg)
	}
	return Query{
		name:   "teams",
		fields: fields,
		args:   args,
	}
}

// The team's graphql field(s).
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
func NewTeamsUsersField(usersFields []UsersField, usersArgs []UsersArgument) TeamsField {
	users := Users.List(usersFields, usersArgs...)
	return TeamsField{field{"users", &users}}
}

// The team's graphql argument(s).
type TeamsArgument struct {
	arg argument
}

// A list of teams unique identifiers.
func NewTeamsIDsArgument(ids []int) TeamsArgument {
	return TeamsArgument{argument{"ids", ids}}
}
