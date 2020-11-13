package ent

import (
	"context"
	"strings"

	"github.com/empiricaly/recruitment/internal/ent/datum"
	participantModel "github.com/empiricaly/recruitment/internal/ent/participant"
	"github.com/empiricaly/recruitment/internal/ent/predicate"
	projectModel "github.com/empiricaly/recruitment/internal/ent/project"
	"github.com/empiricaly/recruitment/internal/model"
	"github.com/pkg/errors"
	errs "github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
)

// AddParticipantFromImport will add partcipant from file
func AddParticipantFromImport(ctx context.Context, tx *Tx, importedParticipant *model.ImportedParticipant, admin *Admin, project *Project) error {
	p, err := tx.Participant.Query().
		Where(
			participantModel.And(
				participantModel.MturkWorkerID(importedParticipant.MturkWorkerID),
				participantModel.HasProjectsWith(projectModel.ID(project.ID)),
			)).
		First(ctx)
	if err != nil {
		log.Warn().Msgf("could not get participant with workerID %s", importedParticipant.MturkWorkerID)
	}

	if p == nil {
		p, err = tx.Participant.Create().
			SetID(xid.New().String()).
			SetMturkWorkerID(importedParticipant.MturkWorkerID).
			SetUninitialized(true).
			AddImportedBy(admin).
			AddProjects(project).
			Save(ctx)
		if err != nil {
			log.Error().Msgf("could not add participant with workerID %s", importedParticipant.MturkWorkerID)
			return errs.Wrap(err, "addParticipant: create participant")
		}
	}

	if importedParticipant.Data == nil {
		return nil
	}

	for _, data := range importedParticipant.Data {
		_, err := SetDatum(ctx, tx, p, data.Key, data.Val, false)
		if err != nil {
			log.Error().Msgf("error inserting datum %s", importedParticipant.MturkWorkerID)
		}
	}

	return nil
}

// MatchCondition returns true if participant matches cond Condition.
// Data associated on Participant is assumed to have been loaded in full (latest
// version) or not at all.
func (p *Participant) MatchCondition(ctx context.Context, cond *model.Condition) (bool, error) {
	data, err := p.CurrentData(ctx, nil, nil)
	if err != nil {
		return false, errors.Wrap(err, "get current data for participant")
	}

	return matchCondition(data, cond), nil
}

func matchCondition(data []*Datum, cond *model.Condition) bool {
	if len(cond.And) > 0 {
		for _, and := range cond.And {
			if !matchCondition(data, and) {
				return false
			}
		}
		return true
	}

	if len(cond.Or) > 0 {
		for _, or := range cond.Or {
			if matchCondition(data, or) {
				return true
			}
		}
		return false
	}

	// If there is no key or comparator, there's nothing to not match against.
	if cond.Key == nil || cond.Comparator == nil || string(*cond.Comparator) == "" {
		return true
	}

	key := *cond.Key
	comp := *cond.Comparator

	var datum *Datum
	for _, d := range data {
		if d.Key == key {
			datum = d
			break
		}
	}
	if datum == nil {
		if comp == model.ComparatorDoesNotExist {
			return true
		}
		return false
	}

	if comp == model.ComparatorExists {
		return true
	}

	if len(cond.Values) == 0 {
		r := gjson.Get(datum.Val, "@this")
		return r.Type != gjson.Null
	}

	if len(cond.Values) > 1 {

	}

	switch comp {
	case model.ComparatorEqualTo:
		return isEqual(datum, cond.Values)
	case model.ComparatorNotEqualTo:
		return !isEqual(datum, cond.Values)
	case model.ComparatorLessThan:
		if len(cond.Values) > 1 {
			return false
		}
		r := gjson.Get(datum.Val, "@this")
		if r.Type != gjson.Number {
			return false
		}
		val := cond.Values[0]
		if (val.Int != nil && int64(*val.Int) > r.Int()) || (val.Float != nil && *val.Float > r.Float()) {
			return true
		} else {
			return false
		}
	case model.ComparatorLessThanOrEqualTo:
		if len(cond.Values) > 1 {
			return false
		}
		r := gjson.Get(datum.Val, "@this")
		if r.Type != gjson.Number {
			return false
		}
		val := cond.Values[0]
		if (val.Int != nil && int64(*val.Int) >= r.Int()) || (val.Float != nil && *val.Float >= r.Float()) {
			return true
		} else {
			return false
		}
	case model.ComparatorGreaterThan:
		if len(cond.Values) > 1 {
			return false
		}
		r := gjson.Get(datum.Val, "@this")
		if r.Type != gjson.Number {
			return false
		}
		val := cond.Values[0]
		if (val.Int != nil && int64(*val.Int) < r.Int()) || (val.Float != nil && *val.Float < r.Float()) {
			return true
		} else {
			return false
		}
	case model.ComparatorGreaterThanOrEqualTo:
		if len(cond.Values) > 1 {
			return false
		}
		r := gjson.Get(datum.Val, "@this")
		if r.Type != gjson.Number {
			return false
		}
		val := cond.Values[0]
		if (val.Int != nil && int64(*val.Int) <= r.Int()) || (val.Float != nil && *val.Float <= r.Float()) {
			return true
		} else {
			return false
		}
	case model.ComparatorIn:
		r := gjson.Get(datum.Val, "@this")
		if r.Type != gjson.Number {
			return false
		}

		for _, v := range cond.Values {
			if isEqualValue(r, v) {
				return true
			}
		}
		return false
	case model.ComparatorNotIn:
		r := gjson.Get(datum.Val, "@this")
		if r.Type != gjson.Number {
			return false
		}

		for _, v := range cond.Values {
			if isEqualValue(r, v) {
				return false
			}
		}
		return true
	default:
		log.Warn().Msgf("could not parse datum %#v", datum.Val)
	}

	return false
}

func isEqual(datum *Datum, values []*model.CompValue) bool {
	if len(values) == 1 {
		r := gjson.Get(datum.Val, "@this")
		return isEqualValue(r, values[0])
	}

	r := gjson.Get(datum.Val, "@this")
	switch r.Type {
	case gjson.JSON:
		if len(values) > 1 {
			if !r.IsArray() {
				return false
			}
			arr := r.Array()
			if len(arr) != len(values) {
				return false
			}

			for i, v := range arr {
				if !isEqualValue(v, values[i]) {
					return false
				}
			}

			return true
		}

		// TODO implement JSON structure comparison
		log.Warn().Msg("JSON object comparison NOT IMPLEMENTED!")
		return false
	default:
		log.Warn().Msgf("could not parse datum %#v", datum.Val)
		return false
	}
}

func isEqualValue(r gjson.Result, val *model.CompValue) bool {
	switch r.Type {
	case gjson.JSON, gjson.String:
		if val.String == nil {
			return false
		}

		if strings.TrimSpace(*val.String) == strings.TrimSpace(r.String()) {
			return true
		} else {
			return false
		}
	case gjson.Number:
		if val.Float == nil && val.Int == nil {
			return false
		}

		if (val.Int != nil && int64(*val.Int) == r.Int()) || (val.Float != nil && *val.Float == r.Float()) {
			return true
		} else {
			return false
		}
	case gjson.False:
		if val.Boolean == nil {
			return false
		}

		return *val.Boolean == false
	case gjson.True:
		if val.Boolean == nil {
			return false
		}

		return *val.Boolean == true
	case gjson.Null:
		if val.Int != nil || val.Float != nil || val.String != nil || val.Boolean != nil {
			return false
		} else {
			return true
		}
	default:
		log.Warn().Msgf("could not parse datum %#v", r)
		return false
	}
}

// CurrentData return all Datum objects that are current and not deleted for
// Participant. Keys and deleted arguments are optional.
// IF Data is already associated on Participant that will be returned.
func (p *Participant) CurrentData(ctx context.Context, keys []string, deleted *bool) ([]*Datum, error) {
	data, err := p.Edges.DataOrErr()
	if err == nil {
		return data, nil
	}

	predicates := []predicate.Datum{
		datum.Current(true),
	}

	if deleted != nil && *deleted {
		predicates = append(predicates, datum.DeletedAtNotNil())
	} else {
		predicates = append(predicates, datum.DeletedAtIsNil())
	}

	if keys != nil {
		predicates = append(predicates, datum.KeyIn(keys...))
	}

	return p.QueryData().Where(datum.And(predicates...)).
		Order(Asc(datum.FieldKey), Asc(datum.FieldIndex)).
		All(ctx)
}
