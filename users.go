package monday

import (
	"fmt"
	"strings"
)

type Users struct {
	alt    string
	fields []UsersField
	args   []UsersArgument
}

func (u Users) stringify() string {
	prefix := "users"
	if u.alt != "" {
		prefix = u.alt
	}
	fields := make([]string, 0)
	for _, field := range u.fields {
		fields = append(fields, field.stringify())
	}
	args := make([]string, 0)
	for _, arg := range u.args {
		args = append(args, arg.stringify())
	}
	if len(fields) == 0 {
		return ``
	}
	if len(args) == 0 {
		return fmt.Sprintf(`%s{%s}`, prefix, strings.Join(fields, " "))
	}
	return fmt.Sprintf(`%s(%s){%s}`, prefix, strings.Join(args, ","), strings.Join(fields, " "))
}

func NewUsers(fields []UsersField) Users {
	if len(fields) == 0 {
		return Users{
			fields: []UsersField{
				UsersIDField(),
			},
		}
	}

	return Users{
		fields: fields,
	}
}

func NewUsersWithArguments(fields []UsersField, args []UsersArgument) Users {
	users := NewUsers(fields)
	users.args = args
	return users
}

type UsersField struct {
	field string
	value interface{}
}

var (
	usersBirthdayField           = UsersField{"birthday", nil}
	usersCountryCodeField        = UsersField{"country_code", nil}
	usersCreatedAtField          = UsersField{"created_at", nil}
	usersEmailField              = UsersField{"email", nil}
	usersEnabledField            = UsersField{"enabled", nil}
	usersIDField                 = UsersField{"id", nil}
	usersIsGuestField            = UsersField{"is_guest", nil}
	usersIsPendingField          = UsersField{"is_pending", nil}
	usersJoinDateField           = UsersField{"join_date", nil}
	usersLocationField           = UsersField{"location", nil}
	usersMobilePhoneField        = UsersField{"mobile_phone", nil}
	usersNameField               = UsersField{"name", nil}
	usersPhoneField              = UsersField{"phone", nil}
	usersPhotoOriginalField      = UsersField{"photo_original", nil}
	usersPhotoSmallField         = UsersField{"photo_small", nil}
	usersPhotoThumbField         = UsersField{"photo_thumb", nil}
	usersPhotoThumbSmallField    = UsersField{"photo_thumb_small", nil}
	usersPhotoTinyField          = UsersField{"photo_tiny", nil}
	usersTimeZoneIdentifierField = UsersField{"time_zone_identifier", nil}
	usersTitleField              = UsersField{"title", nil}
	usersURLField                = UsersField{"url", nil}
	usersUTCHoursDifference      = UsersField{"utc_hours_diff", nil}
)

func (f UsersField) stringify() string {
	switch f.field {
	case "teams":
		return f.value.(Teams).stringify()
	default:
		return fmt.Sprint(f.field)
	}
}

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
func NewUsersTeamsField(teams Teams) UsersField {
	return UsersField{field: "teams", value: teams}
}

// The user's time zone identifier.
func UsersTimeZoneIdentifier() UsersField {
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
	argument string
	value    interface{}
}

func (a UsersArgument) stringify() string {
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
func NewIDsUsersArg(ids []int) UsersArgument {
	return UsersArgument{
		argument: "ids",
		value:    ids,
	}
}

// The kind to search users by (all / non_guests / guests / non_pending).
func NewKindUsersArg(kind UsersKind) UsersArgument {
	return UsersArgument{
		argument: "kind",
		value:    kind.kind,
	}
}

// Get the recently created users at the top of the list.
func NewNewestFirstUsersArg(first bool) UsersArgument {
	return UsersArgument{
		argument: "newest_first",
		value:    first,
	}
}

// Number of users to get.
func NewLimitUsersArg(value int) UsersArgument {
	return UsersArgument{
		argument: "limit",
		value:    value,
	}
}
