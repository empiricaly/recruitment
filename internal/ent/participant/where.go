// Code generated by entc, DO NOT EDIT.

package participant

import (
	"time"

	"github.com/empiricaly/recruitment/internal/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// MturkWorkerID applies equality check predicate on the "mturkWorkerID" field. It's identical to MturkWorkerIDEQ.
func MturkWorkerID(v string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMturkWorkerID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Participant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Participant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Participant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Participant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Participant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Participant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Participant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Participant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// MturkWorkerIDEQ applies the EQ predicate on the "mturkWorkerID" field.
func MturkWorkerIDEQ(v string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMturkWorkerID), v))
	})
}

// MturkWorkerIDNEQ applies the NEQ predicate on the "mturkWorkerID" field.
func MturkWorkerIDNEQ(v string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMturkWorkerID), v))
	})
}

// MturkWorkerIDIn applies the In predicate on the "mturkWorkerID" field.
func MturkWorkerIDIn(vs ...string) predicate.Participant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Participant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMturkWorkerID), v...))
	})
}

// MturkWorkerIDNotIn applies the NotIn predicate on the "mturkWorkerID" field.
func MturkWorkerIDNotIn(vs ...string) predicate.Participant {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Participant(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMturkWorkerID), v...))
	})
}

// MturkWorkerIDGT applies the GT predicate on the "mturkWorkerID" field.
func MturkWorkerIDGT(v string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMturkWorkerID), v))
	})
}

// MturkWorkerIDGTE applies the GTE predicate on the "mturkWorkerID" field.
func MturkWorkerIDGTE(v string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMturkWorkerID), v))
	})
}

// MturkWorkerIDLT applies the LT predicate on the "mturkWorkerID" field.
func MturkWorkerIDLT(v string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMturkWorkerID), v))
	})
}

// MturkWorkerIDLTE applies the LTE predicate on the "mturkWorkerID" field.
func MturkWorkerIDLTE(v string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMturkWorkerID), v))
	})
}

// MturkWorkerIDContains applies the Contains predicate on the "mturkWorkerID" field.
func MturkWorkerIDContains(v string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMturkWorkerID), v))
	})
}

// MturkWorkerIDHasPrefix applies the HasPrefix predicate on the "mturkWorkerID" field.
func MturkWorkerIDHasPrefix(v string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMturkWorkerID), v))
	})
}

// MturkWorkerIDHasSuffix applies the HasSuffix predicate on the "mturkWorkerID" field.
func MturkWorkerIDHasSuffix(v string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMturkWorkerID), v))
	})
}

// MturkWorkerIDIsNil applies the IsNil predicate on the "mturkWorkerID" field.
func MturkWorkerIDIsNil() predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMturkWorkerID)))
	})
}

// MturkWorkerIDNotNil applies the NotNil predicate on the "mturkWorkerID" field.
func MturkWorkerIDNotNil() predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMturkWorkerID)))
	})
}

// MturkWorkerIDEqualFold applies the EqualFold predicate on the "mturkWorkerID" field.
func MturkWorkerIDEqualFold(v string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMturkWorkerID), v))
	})
}

// MturkWorkerIDContainsFold applies the ContainsFold predicate on the "mturkWorkerID" field.
func MturkWorkerIDContainsFold(v string) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMturkWorkerID), v))
	})
}

// HasProviderIDs applies the HasEdge predicate on the "providerIDs" edge.
func HasProviderIDs() predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProviderIDsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProviderIDsTable, ProviderIDsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProviderIDsWith applies the HasEdge predicate on the "providerIDs" edge with a given conditions (other predicates).
func HasProviderIDsWith(preds ...predicate.ProviderID) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProviderIDsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProviderIDsTable, ProviderIDsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParticipations applies the HasEdge predicate on the "participations" edge.
func HasParticipations() predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParticipationsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ParticipationsTable, ParticipationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParticipationsWith applies the HasEdge predicate on the "participations" edge with a given conditions (other predicates).
func HasParticipationsWith(preds ...predicate.Participation) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ParticipationsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ParticipationsTable, ParticipationsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCreatedBy applies the HasEdge predicate on the "createdBy" edge.
func HasCreatedBy() predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CreatedByTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatedByTable, CreatedByColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatedByWith applies the HasEdge predicate on the "createdBy" edge with a given conditions (other predicates).
func HasCreatedByWith(preds ...predicate.StepRun) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CreatedByInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatedByTable, CreatedByColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSteps applies the HasEdge predicate on the "steps" edge.
func HasSteps() predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StepsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, StepsTable, StepsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStepsWith applies the HasEdge predicate on the "steps" edge with a given conditions (other predicates).
func HasStepsWith(preds ...predicate.StepRun) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StepsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, StepsTable, StepsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Participant) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Participant) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Participant) predicate.Participant {
	return predicate.Participant(func(s *sql.Selector) {
		p(s.Not())
	})
}