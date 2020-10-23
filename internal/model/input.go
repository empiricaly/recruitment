package model

func FilterArgsFromInput(criti *FilterStepArgsInput) *FilterStepArgs {
	args := &FilterStepArgs{}
	args.Condition = ConditionFromInput(criti.Condition)
	args.Filter = criti.Filter
	args.Js = criti.Js
	args.Type = criti.Type

	return args
}

func MsgArgsFromInput(criti *MessageStepArgsInput) *MessageStepArgs {
	args := &MessageStepArgs{}
	args.Subject = criti.Subject
	args.URL = criti.URL

	if criti.Message != nil {
		args.Message = *criti.Message
	}

	if criti.MessageType != nil {
		args.MessageType = *criti.MessageType
	}

	args.Lobby = criti.Lobby
	args.LobbyType = criti.LobbyType
	args.LobbyExpiration = criti.LobbyExpiration

	return args
}

func HITStepArgsFromInput(criti *HITStepArgsInput) *HITStepArgs {
	args := &HITStepArgs{}

	if criti.Title != nil {
		args.Title = *criti.Title
	}

	if criti.Description != nil {
		args.Description = *criti.Description
	}

	if criti.Keywords != nil {
		args.Keywords = *criti.Keywords
	}

	if criti.Microbatch != nil {
		args.Microbatch = *criti.Microbatch
	}

	if criti.Reward != nil {
		args.Reward = *criti.Reward
	}

	if criti.Timeout != nil {
		args.Timeout = *criti.Timeout
	}

	args.Duration = criti.Duration

	if criti.WorkersCount != nil {
		args.WorkersCount = *criti.WorkersCount
	}

	return args
}

func MturkCriteriaFromInput(criti *MTurkCriteriaInput) *MTurkCriteria {
	crit := &MTurkCriteria{}
	if len(criti.Qualifications) > 0 {
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
	if len(criti.Values) > 0 {
		crit.Values = make([]int, len(criti.Values))
		for i, v := range criti.Values {
			crit.Values[i] = v
		}
	}
	if len(criti.Locales) > 0 {
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
	if condi == nil {
		return nil
	}
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
