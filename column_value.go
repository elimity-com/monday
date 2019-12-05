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
	value string
}

type ColumnValues []ColumnValue

func (v ColumnValues) join() string {
	var str []string
	for _, v := range v {
		str = append(str, v.value)
	}
	return strings.Join(str, "")
}

func addQuotes(value interface{}) ColumnValue {
	return ColumnValue{fmt.Sprintf(`"%v"`, value)}
}

func newIndex(value int) ColumnValue {
	return ColumnValue{fmt.Sprintf(`{"index": "%d"}`, value)}
}

func newLabel(value string) ColumnValue {
	return ColumnValue{fmt.Sprintf(`{"label": %q}`, value)}
}

func NewItemNameValue(value string) ColumnValue {
	return addQuotes(value)
}

func NewTextValue(value string) ColumnValue {
	return addQuotes(value)
}

func NewLongTextValue(value string) ColumnValue {
	return ColumnValue{fmt.Sprintf(`{"text": %q}`, value)}
}

func NewNumberValue(value int) ColumnValue {
	return addQuotes(value)
}

func NewStatusIndexValue(value int) ColumnValue {
	return newIndex(value)
}

func NewStatusLabelValue(value string) ColumnValue {
	return newLabel(value)
}

func NewDropdownIndexValue(value int) ColumnValue {
	return newIndex(value)
}

func NewDropdownLabelValue(value string) ColumnValue {
	return newLabel(value)
}

func NewPersonValue(value int) ColumnValue {
	return ColumnValue{fmt.Sprintf(`{"id":"%d"}`, value)}
}

func NewTeamValue(value int) ColumnValue {
	return ColumnValue{fmt.Sprintf(`{"team_id":"%d"}`, value)}
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

func NewPeopleValue(value []People) ColumnValue {
	var str []string
	for _, v := range value {
		str = append(str, fmt.Sprintf(`{"id":"%d","kind",%q}`, v.ID, v.Kind.kind))
	}
	return ColumnValue{fmt.Sprintf(`{"personsAndTeams":[%s]}`, strings.Join(str, ","))}
}

func NewWorldClockValue(value string) ColumnValue {
	return ColumnValue{fmt.Sprintf(`{"timezone":%q}`, value)}
}

func NewCountryValue(code, name string) ColumnValue {
	return ColumnValue{fmt.Sprintf(`{"countryCode":%q,"countryName":%q}`, code, name)}
}

func NewEmailValue(email, text string) ColumnValue {
	return ColumnValue{fmt.Sprintf(`{"email":%q,"text":%q}`, email, text)}
}

func NewPhoneValue(number int, code string) ColumnValue {
	return ColumnValue{fmt.Sprintf(`{"phone":"%d","countryShortName":%q}`, number, code)}
}

func NewURLValue(url, text string) ColumnValue {
	return ColumnValue{fmt.Sprintf(`{"url":%q,"text":%q}`, url, text)}
}

func NewDateValue(value time.Time) ColumnValue {
	return ColumnValue{fmt.Sprintf(`{"date":%q,"time":%q}`, value.Format(dateFormat), value.Format(timeFormat))}
}

func NewTimelineValue(from, to time.Time) ColumnValue {
	return ColumnValue{fmt.Sprintf(`{"from":%q,"to":%q}`, from.Format(dateFormat), to.Format(dateFormat))}
}

func NewTagsValue(ids []int) ColumnValue {
	return ColumnValue{fmt.Sprintf(`{"tags_ids":%q}`, strings.Join(strings.Split(fmt.Sprint(ids), " "), ","))}
}

func NewHourValue(value time.Time) ColumnValue {
	return ColumnValue{fmt.Sprintf(`{"hour":%d,"minute":%d}`, value.Hour(), value.Minute())}
}

func NewWeekValue(start, end time.Time) ColumnValue {
	return ColumnValue{fmt.Sprintf(`{"week":{"startDate":%q,"endDate":%s}}`, start.Format(dateFormat), end.Format(dateFormat))}
}

func NewCheckboxValue(value bool) ColumnValue {
	return ColumnValue{fmt.Sprintf(`{"checked":"%t"}`, value)}
}

func NewRatingValue(value int) ColumnValue {
	return ColumnValue{fmt.Sprintf(`{"rating":%d}`, value)}
}

func RemoveValue() ColumnValue {
	return ColumnValue{"{}"}
}
