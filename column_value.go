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

func addQuotes(id string, value interface{}) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`"%v"`, value)}
}

func newIndex(id string, value int) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"index":%d}`, value)}
}

func newLabel(id, value string) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"label":%q}`, value)}
}

func NewItemNameValue(id, value string) ColumnValue {
	return addQuotes(id, value)
}

func NewTextValue(id, value string) ColumnValue {
	return addQuotes(id, value)
}

func NewLongTextValue(id, value string) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"text": %q}`, value)}
}

func NewNumberValue(id string, value int) ColumnValue {
	return addQuotes(id, value)
}

func NewStatusIndexValue(id string, value int) ColumnValue {
	return newIndex(id, value)
}

func NewStatusLabelValue(id, value string) ColumnValue {
	return newLabel(id, value)
}

func NewDropdownIndexValue(id string, value int) ColumnValue {
	return newIndex(id, value)
}

func NewDropdownLabelValue(id, value string) ColumnValue {
	return newLabel(id, value)
}

func NewPersonValue(id string, value int) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"id":"%d"}`, value)}
}

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

func NewPeopleValue(id string, value []People) ColumnValue {
	var str []string
	for _, v := range value {
		str = append(str, fmt.Sprintf(`{"id":"%d","kind",%q}`, v.ID, v.Kind.kind))
	}
	return ColumnValue{id, fmt.Sprintf(`{"personsAndTeams":[%s]}`, strings.Join(str, ","))}
}

func NewWorldClockValue(id, value string) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"timezone":%q}`, value)}
}

func NewCountryValue(id, code, name string) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"countryCode":%q,"countryName":%q}`, code, name)}
}

func NewEmailValue(id, email, text string) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"email":%q,"text":%q}`, email, text)}
}

func NewPhoneValue(id string, number int, code string) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"phone":"%d","countryShortName":%q}`, number, code)}
}

func NewURLValue(id, url, text string) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"url":%q,"text":%q}`, url, text)}
}

func NewDateValue(id string, value time.Time) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"date":%q,"time":%q}`, value.Format(dateFormat), value.Format(timeFormat))}
}

func NewTimelineValue(id string, from, to time.Time) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"from":%q,"to":%q}`, from.Format(dateFormat), to.Format(dateFormat))}
}

func NewTagsValue(id string, ids []int) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"tags_ids":%q}`, strings.Join(strings.Split(fmt.Sprint(ids), " "), ","))}
}

func NewHourValue(id string, value time.Time) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"hour":%d,"minute":%d}`, value.Hour(), value.Minute())}
}

func NewWeekValue(id string, start, end time.Time) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"week":{"startDate":%q,"endDate":%s}}`, start.Format(dateFormat), end.Format(dateFormat))}
}

func NewCheckboxValue(id string, value bool) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"checked":"%t"}`, value)}
}

func NewRatingValue(id string, value int) ColumnValue {
	return ColumnValue{id, fmt.Sprintf(`{"rating":%d}`, value)}
}

func RemoveValue(id string) ColumnValue {
	return ColumnValue{id, "{}"}
}
