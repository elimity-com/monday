package monday

func NewAccount(accountFields []AccountField) Query {
	if len(accountFields) == 0 {
		return Query{
			name: "account",
			fields: []field{
				AccountIDField().field,
			},
		}
	}

	var fields []field
	for _, af := range accountFields {
		fields = append(fields, af.field)
	}
	return Query{
		name:   "account",
		fields: fields,
	}
}

type AccountField struct {
	field field
}

var (
	accountFirstDayOfTheWeekField    = AccountField{field{"first_day_of_the_week", nil}}
	accountIDField                   = AccountField{field{"id", nil}}
	accountLogoField                 = AccountField{field{"logo", nil}}
	accountNameField                 = AccountField{field{"name", nil}}
	accountShowTimelineWeekendsField = AccountField{field{"show_timeline_weekends", nil}}
	accountSlugField                 = AccountField{field{"slug", nil}}
)

// 	The first day of the week for the account (sunday / monday)
func AccountFirstDayOfTheWeekField() AccountField {
	return accountFirstDayOfTheWeekField
}

// The account's unique identifier.
func AccountIDField() AccountField {
	return accountIDField
}

// The account's logo.
func AccountLogoField() AccountField {
	return accountLogoField
}

// The account's name.
func AccountNameField() AccountField {
	return accountNameField
}

// The account's payment plan.
func AccountPlanField(planFields []PlanField) AccountField {
	plan := newPlan(planFields)
	return AccountField{field{"plan", &plan}}
}

// Show weekends in timeline.
func AccountShowTimelineWeekendsField() AccountField {
	return accountShowTimelineWeekendsField
}

// The account's slug.
func AccountSlugField() AccountField {
	return accountSlugField
}
