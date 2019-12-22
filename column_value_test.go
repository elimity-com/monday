package monday

import (
	"testing"
	"time"
)

// DOCS: https://monday.com/developers/v2#column-values-section
func TestColumnValue(t *testing.T) {
	for _, test := range []struct {
		column ColumnValue
		str    string
	}{
		{
			column: NewItemNameValue("", "My item"),
			str:    `"My item"`,
		},
		{
			column: NewTextValue("", "Sample text"),
			str:    `"Sample text"`,
		},
		{
			column: NewLongTextValue("", "Sample text"),
			str:    `{"text":"Sample text"}`,
		},
		{
			column: NewNumberValue("", 3),
			str:    `"3"`,
		},
		{
			column: NewStatusIndexValue("", 0),
			str:    `{"index":0}`,
		},
		{
			column: NewStatusLabelValue("", "Done"),
			str:    `{"label":"Done"}`,
		},
		{
			column: NewDropdownIndexValue("", []int{1}),
			str:    `{"ids":[1]}`,
		},
		{
			column: NewDropdownLabelValue("", []string{"My label"}),
			str:    `{"labels":["My label"]}`,
		},
		{
			column: NewPersonValue("", 235326),
			str:    `{"id":235326}`,
		},
		{
			column: NewTeamValue("", 51166),
			str:    `{"team_id":51166}`,
		},
		{
			column: NewPeopleValue("", []People{
				{4616627, PeopleKindPerson()},
				{4616666, PeopleKindPerson()},
				{51166, PeopleKindTeam()},
			}),
			str: `{"personsAndTeams":[{"id":4616627,"kind":"person"},{"id":4616666,"kind":"person"},{"id":51166,"kind":"team"}]}`,
		},
		{
			column: NewWorldClockValue("", "Europe/London"),
			str:    `{"timezone":"Europe/London"}`,
		},
		{
			column: NewCountryValue("", "US", "United States"),
			str:    `{"countryCode":"US","countryName":"United States"}`,
		},
		{
			column: NewEmailValue("", "itsmyemail@mailserver.com", "my email"),
			str:    `{"email":"itsmyemail@mailserver.com","text":"my email"}`,
		},
		{
			column: NewPhoneValue("", 11231234567, "US"),
			str:    `{"phone":"11231234567","countryShortName":"US"}`,
		},
		{
			column: NewURLValue("", "http://monday.com", "go to monday!"),
			str:    `{"url":"http://monday.com","text":"go to monday!"}`,
		},
		{
			column: NewDateValue("", mustParse("2006-01-02 15:04:05", "2019-06-03 13:25:00")),
			str:    `{"date":"2019-06-03","time":"13:25:00"}`,
		},
		{
			column: NewTimelineValue("", mustParse(dateFormat, "2019-06-03"), mustParse(dateFormat, "2019-06-07")),
			str:    `{"from":"2019-06-03","to":"2019-06-07"}`,
		},
		{
			column: NewTagsValue("", []int{295026, 295064}),
			str:    `{"tag_ids":[295026,295064]}`,
		},
		{
			column: NewHourValue("", mustParse("15:04", "16:42")),
			str:    `{"hour":16,"minute":42}`,
		},
		{
			column: NewWeekValue("", mustParse(dateFormat, "2019-06-10"), mustParse(dateFormat, "2019-06-16")),
			str:    `{"week":{"startDate":"2019-06-10","endDate":"2019-06-16"}}`,
		},
		{
			column: NewCheckboxValue("", true),
			str:    `{"checked":"true"}`,
		},
		{
			column: NewRatingValue("", 5),
			str:    `{"rating":5}`,
		},
		{
			column: RemoveValue(""),
			str:    `{}`,
		},
	} {
		if value := test.column.value; value != test.str {
			t.Errorf("got: %s, expected: %s", value, test.str)
		}
	}
}

func mustParse(layout, value string) time.Time {
	t, _ := time.Parse(layout, value)
	return t
}
