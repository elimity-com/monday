package monday

// The plan's graphql field(s).
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
