package monday

import (
	"fmt"
	"strings"
	"time"
)

const (
	dateFormat = "2006-01-02"
	timeFormat = "15:04:05"
)

type ColumnValue struct {
	id, value string
}

func (v ColumnValue) ID() string {
	return v.id
}

func (v ColumnValue) Value() string {
	return v.value
}

func addQuotes(id string, value interface{}) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`"%v"`, value)}
}

func newIndex(id string, value int) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"index":%d}`, value)}
}

func newLabel(id, value string) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"label":%q}`, value)}
}

// To update the item's name, send a string of 1 to 255 characters.
func NewItemNameValue(id, value string) ColumnValue {
	return addQuotes(id, value)
}

// To update the text column send a string.
func NewTextValue(id, value string) ColumnValue {
	return addQuotes(id, value)
}

// To update the long text column, send a string up to 2000 characters.
func NewLongTextValue(id, value string) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"text": %q}`, value)}
}

// To update the number column send a string containing a float or int.
func NewNumberValue(id string, value int) ColumnValue {
	return addQuotes(id, value)
}

// To update a status column, send the index of the status you want to select.
func NewStatusIndexValue(id string, value int) ColumnValue {
	return newIndex(id, value)
}

// To update a status column, send the label of the status you want to select.
func NewStatusLabelValue(id, value string) ColumnValue {
	return newLabel(id, value)
}

// To update a dropdown column, send the id of the label you want to select.
func NewDropdownIndexValue(id string, value int) ColumnValue {
	return newIndex(id, value)
}

// To update a dropdown column, send the label you want to select.
func NewDropdownLabelValue(id, value string) ColumnValue {
	return newLabel(id, value)
}

// To update a person column, send the ID of the user.
func NewPersonValue(id string, value int) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"id":"%d"}`, value)}
}

// To update a team column send the ID of the team.
// The ID of a specific team can be found by using the teams query,
// checking which teams a particular user is a part of (with the User object).
func NewTeamValue(id string, value int) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"team_id":"%d"}`, value)}
}

type People struct {
	ID   int
	Kind PeopleKind
}

type PeopleKind struct {
	kind string
}

var (
	peopleKindPerson = PeopleKind{"person"}
	peopleKindTeam   = PeopleKind{"team"}
)

func PeopleKindPerson() PeopleKind {
	return peopleKindPerson
}

func PeopleKindTeam() PeopleKind {
	return peopleKindTeam
}

// To update a people column, send an array with the people or teams you want to add to the column.
// Each item in the array should include the ID of the person/team, and a string signifying whether the item is a person or a team.
func NewPeopleValue(id string, value []People) ColumnValue {
	var str []string
	for _, v := range value {
		str = append(str, fmt.Sprintf(`{"id":"%d","kind",%q}`, v.ID, v.Kind.kind))
	}
	return ColumnValue{id, fmt.Sprintf(`{"personsAndTeams":[%s]}`, strings.Join(str, ","))}
}

// To update a world clock column, send the timezone of the user as a string in continent/city form.
// You can get the list of available timezones here: http://www.worldtimezone.com
func NewWorldClockValue(id, value string) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"timezone":%q}`, value)}
}

// To update a country column send the iso-2 country code (2 letter code) and the country name.
// You can get the list of available countries here: http://country.io/names.json
func NewCountryValue(id, code, name string) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"countryCode":%q,"countryName":%q}`, code, name)}
}

// To update an email column, send the email address in the email field.
// You can also pass display text in the text field.
func NewEmailValue(id, email, text string) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"email":%q,"text":%q}`, email, text)}
}

// To update a phone column send the phone number (digits only) in a string and the iso-2 country code (2 letter code).
// You can get the list of available countries here: http://country.io/names.json
func NewPhoneValue(id string, number int, code string) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"phone":"%d","countryShortName":%q}`, number, code)}
}

// To update a link column, write the URL (including http/https) in the url field.
// You can also pass display text in the text field.
func NewURLValue(id, url, text string) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"url":%q,"text":%q}`, url, text)}
}

// To update a date column, send the date as a string in a YYYY-MM-DD format.
// You can also add a time by passing a “time” field in HH:MM:SS format.
func NewDateValue(id string, value time.Time) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"date":%q,"time":%q}`, value.Format(dateFormat), value.Format(timeFormat))}
}

// To update a timeline column, send the start and end dates in a YYYY-MM-DD format,
// where the start date is “from” and the end date is “to”.
func NewTimelineValue(id string, from, to time.Time) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"from":%q,"to":%q}`, from.Format(dateFormat), to.Format(dateFormat))}
}

// To update a tags column, send the tag ID’s in an array.
func NewTagsValue(id string, ids []int) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"tags_ids":%q}`, strings.Join(strings.Split(fmt.Sprint(ids), " "), ","))}
}

// To update an hour column, send the hour and minute in 24-hour format.
func NewHourValue(id string, value time.Time) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"hour":%d,"minute":%d}`, value.Hour(), value.Minute())}
}

// To update a week column send the start and end dates in a YYYY-MM-DD format.
// Date must be 7 days apart (inclusive of the first and last date) and
// start at the beginning of the work week defined in the account.
func NewWeekValue(id string, start, end time.Time) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"week":{"startDate":%q,"endDate":%s}}`, start.Format(dateFormat), end.Format(dateFormat))}
}

// To check the box in the checkbox column, send a 'checked' field with true.
// To uncheck the box remove the column value.
func NewCheckboxValue(id string, value bool) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"checked":"%t"}`, value)}
}

// To update a rating column send a number between 1 and your rating scale.
func NewRatingValue(id string, value int) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"rating":%d}`, value)}
}

// To remove a column value (delete the existing value) send an empty string.
func RemoveValue(id string) ColumnValue {
	return ColumnValue{id, "{}"}
}
