package monday

func newPlan(planFields []PlanField) Query {
	if len(planFields) == 0 {
		return Query{
			name: "plan",
			fields: []field{
				PlanMaxUsersField().field,
				PlanPeriodField().field,
				PlanTierField().field,
				PlanVersionField().field,
			},
		}
	}

	var fields []field
	for _, pf := range planFields {
		fields = append(fields, pf.field)
	}
	return Query{
		name:   "plan",
		fields: fields,
	}
}

type PlanField struct {
	field field
}

var (
	planMaxUsersField = PlanField{field{"max_users", nil}}
	planPeriodField   = PlanField{field{"period", nil}}
	planTierField     = PlanField{field{"tier", nil}}
	planVersionField  = PlanField{field{"version", nil}}
)

// The maximum users allowed in the plan.
func PlanMaxUsersField() PlanField {
	return planMaxUsersField
}

// The plan's time period.
func PlanPeriodField() PlanField {
	return planPeriodField
}

// The plan's tier.
func PlanTierField() PlanField {
	return planTierField
}

// The plan's version.
func PlanVersionField() PlanField {
	return planVersionField
}
