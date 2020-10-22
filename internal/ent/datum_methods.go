package ent

import (
	"context"

	"github.com/empiricaly/recruitment/internal/ent/datum"
	"github.com/pkg/errors"
	"github.com/rs/xid"
)

// SetDatum sets a new datum version respecting all the rules.
func SetDatum(ctx context.Context, tx *Tx, p *Participant, key string, val string, isAppend bool) (*Datum, error) {
	existing, err := p.QueryData().
		Where(datum.And(datum.Current(true), datum.KeyEQ(key))).
		Order(Asc(datum.FieldIndex)).
		All(ctx)
	if err != nil && !IsNotFound(err) {
		return nil, errors.Wrap(err, "get previous datum")
	}

	var version int
	var index int
	if len(existing) > 0 {
		if isAppend {
			version = existing[0].Index
		} else {
			version = existing[0].Index + 1
			index = existing[len(existing)-1].Index + 1

			ids := make([]string, len(existing))
			for i, d := range existing {
				ids[i] = d.ID
			}
			_, err = tx.Datum.
				Update().
				Where(datum.IDIn(ids...)).
				SetCurrent(false).
				Save(ctx)
			if err != nil {
				return nil, errors.Wrap(err, "update previous data")
			}
		}
	}

	newDatum, err := tx.Datum.Create().
		SetID(xid.New().String()).
		SetCurrent(true).
		SetVersion(version).
		SetIndex(index).
		SetKey(key).
		SetVal(val).
		SetParticipant(p).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "insert new datum")
	}

	return newDatum, nil
}
