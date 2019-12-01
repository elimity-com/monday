package monday

func NewUsers(usersFields []UsersField) Query {
	if len(usersFields) == 0 {
		return Query{
			name: "users",
			fields: []field{
				UsersIDField().field,
			},
		}
	}

	var fields []field
	for _, uf := range usersFields {
		fields = append(fields, uf.field)
	}
	return Query{
		name:   "users",
		fields: fields,
	}
}

func NewUsersWithArguments(usersFields []UsersField, usersArgs []UsersArgument) Query {
	users := NewUsers(usersFields)
	var args []argument
	for _, ua := range usersArgs {
		args = append(args, ua.arg)
	}
	users.args = args
	return users
}

type UsersField struct {
	field field
}

var (
	usersBirthdayField           = UsersField{field{"birthday", nil}}
	usersCountryCodeField        = UsersField{field{"country_code", nil}}
	usersCreatedAtField          = UsersField{field{"created_at", nil}}
	usersEmailField              = UsersField{field{"email", nil}}
	usersEnabledField            = UsersField{field{"enabled", nil}}
	usersIDField                 = UsersField{field{"id", nil}}
	usersIsGuestField            = UsersField{field{"is_guest", nil}}
	usersIsPendingField          = UsersField{field{"is_pending", nil}}
	usersJoinDateField           = UsersField{field{"join_date", nil}}
	usersLocationField           = UsersField{field{"location", nil}}
	usersMobilePhoneField        = UsersField{field{"mobile_phone", nil}}
	usersNameField               = UsersField{field{"name", nil}}
	usersPhoneField              = UsersField{field{"phone", nil}}
	usersPhotoOriginalField      = UsersField{field{"photo_original", nil}}
	usersPhotoSmallField         = UsersField{field{"photo_small", nil}}
	usersPhotoThumbField         = UsersField{field{"photo_thumb", nil}}
	usersPhotoThumbSmallField    = UsersField{field{"photo_thumb_small", nil}}
	usersPhotoTinyField          = UsersField{field{"photo_tiny", nil}}
	usersTimeZoneIdentifierField = UsersField{field{"time_zone_identifier", nil}}
	usersTitleField              = UsersField{field{"title", nil}}
	usersURLField                = UsersField{field{"url", nil}}
	usersUTCHoursDifference      = UsersField{field{"utc_hours_diff", nil}}
)

// TODO: account? nothing found in documentation

// The user's birthday.
func UsersBirthDayField() UsersField {
	return usersBirthdayField
}

// The user's country code.
func UsersCountryCodeField() UsersField {
	return usersCountryCodeField
}

// The user's creation date.
func UsersCreatedAtField() UsersField {
	return usersCreatedAtField
}

// The user's email.
func UsersEmailField() UsersField {
	return usersEmailField
}

// Is the user enabled or not.
func UsersEnabledField() UsersField {
	return usersEnabledField
}

// The user's unique identifier.
func UsersIDField() UsersField {
	return usersIDField
}

// Is the user a guest or not.
func UsersIsGuestField() UsersField {
	return usersIsGuestField
}

// Is the user a pending user or not.
func UsersIsPendingField() UsersField {
	return usersIsPendingField
}

// The date the user joined the account.
func UsersJoinDateField() UsersField {
	return usersJoinDateField
}

// The user's location.
func UsersLocationField() UsersField {
	return usersLocationField
}

// The user's mobile phone number.
func UsersMobilePhoneField() UsersField {
	return usersMobilePhoneField
}

// The user's name.
func UsersNameField() UsersField {
	return usersNameField
}

// The user's phone number.
func UsersPhoneField() UsersField {
	return usersPhoneField
}

// The user's photo in the original size.
func UsersPhotoOriginalField() UsersField {
	return usersPhotoOriginalField
}

// The user's photo in small size (150x150).
func UsersPhotoSmallField() UsersField {
	return usersPhotoSmallField
}

// The user's photo in thumbnail size (100x100).
func UsersPhotoThumbField() UsersField {
	return usersPhotoThumbField
}

// The user's photo in small thumbnail size (50x50).
func UsersPhotoThumbSmallField() UsersField {
	return usersPhotoThumbSmallField
}

// The user's photo in tiny size (30x30).
func UsersPhotoTinyField() UsersField {
	return usersPhotoTinyField
}

// The teams the user is a member in.
func NewUsersTeamsField(teamsFields []TeamsField, teamsArguments []TeamsArgument) UsersField {
	teams := NewTeamsWithArguments(teamsFields, teamsArguments)
	return UsersField{field{"teams", &teams}}
}

// The user's time zone identifier.
func UsersTimeZoneIdentifierField() UsersField {
	return usersTimeZoneIdentifierField
}

// The user's title.
func UsersTitleField() UsersField {
	return usersTitleField
}

// The user's profile url.
func UsersURLField() UsersField {
	return usersURLField
}

// The user’s utc hours difference.
func UsersUTCHoursDifferenceField() UsersField {
	return usersUTCHoursDifference
}

type UsersArgument struct {
	arg argument
}

type UsersKind struct {
	kind string
}

var (
	allUsersKind       = UsersKind{"all"}
	nonGuestsUsersKind = UsersKind{"non_guests"}
	guestsUsersKind    = UsersKind{"guests"}
	nonPendingKind     = UsersKind{"non_pending"}
)

func AllUsersKind() UsersKind {
	return allUsersKind
}

func NonGuestsUsersKind() UsersKind {
	return nonGuestsUsersKind
}

func GuestsUsersKind() UsersKind {
	return guestsUsersKind
}

func NonPendingUsersKind() UsersKind {
	return nonPendingKind
}

// A list of users unique identifiers.
func NewUsersIDsArgument(ids []int) UsersArgument {
	return UsersArgument{argument{"ids", ids}}
}

// The kind to search users by (all / non_guests / guests / non_pending).
func NewUsersKindArgument(kind UsersKind) UsersArgument {
	return UsersArgument{argument{"kind", kind.kind}}
}

// Get the recently created users at the top of the list.
func NewUsersNewestFirstArgument(first bool) UsersArgument {
	return UsersArgument{argument{"newest_first", first}}
}

// Number of users to get.
func NewUsersLimitArgument(value int) UsersArgument {
	return UsersArgument{argument{"limit", value}}
}
