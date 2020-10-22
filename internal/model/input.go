package model

func MturkCriteriaFromInput(criti *MTurkCriteriaInput) *MTurkCriteria {
	crit := &MTurkCriteria{}
	if len(crit.Qualifications) > 0 {
		crit.Qualifications = make([]*MTurkQualificationCriteria, len(criti.Qualifications))
		for i, q := range criti.Qualifications {
			crit.Qualifications[i] = MTurkQualificationFromInput(q)
		}
	}

	return crit
}

func MTurkQualificationFromInput(criti *MTurkQualificationCriteriaInput) *MTurkQualificationCriteria {
	crit := &MTurkQualificationCriteria{
		ID:         criti.ID,
		Comparator: *criti.Comparator,
	}
	if len(crit.Values) > 0 {
		crit.Values = make([]int, len(criti.Values))
		for i, v := range criti.Values {
			crit.Values[i] = v
		}
	}
	if len(crit.Locales) > 0 {
		crit.Locales = make([]*MTurkLocale, len(criti.Locales))
		for i, l := range criti.Locales {
			crit.Locales[i] = &MTurkLocale{
				Subdivision: l.Subdivision,
				Country:     l.Country,
			}
		}
	}

	return crit
}

func InternalCriteriaFromInput(criti *InternalCriteriaInput) *InternalCriteria {
	crit := &InternalCriteria{}
	crit.All = criti.All
	crit.Condition = ConditionFromInput(criti.Condition)

	return crit
}

func ConditionFromInput(condi *ConditionInput) *Condition {
	cond := &Condition{}
	if len(condi.And) > 0 {
		cond.And = make([]*Condition, len(condi.And))
		for i, c := range condi.And {
			cond.And[i] = ConditionFromInput(c)
		}
	}
	if len(condi.Or) > 0 {
		cond.Or = make([]*Condition, len(condi.Or))
		for i, c := range condi.Or {
			cond.Or[i] = ConditionFromInput(c)
		}
	}
	if len(condi.Values) > 0 {
		cond.Values = make([]*CompValue, len(condi.Values))
		for i, c := range condi.Values {
			cond.Values[i] = &CompValue{
				Int:     c.Int,
				Float:   c.Float,
				String:  c.String,
				Boolean: c.Boolean,
			}
		}
	}
	cond.Comparator = condi.Comparator
	cond.Key = condi.Key

	return cond
}
