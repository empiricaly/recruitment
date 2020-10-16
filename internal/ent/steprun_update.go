// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/empiricaly/recruitment/internal/ent/participant"
	"github.com/empiricaly/recruitment/internal/ent/participation"
	"github.com/empiricaly/recruitment/internal/ent/predicate"
	"github.com/empiricaly/recruitment/internal/ent/run"
	"github.com/empiricaly/recruitment/internal/ent/step"
	"github.com/empiricaly/recruitment/internal/ent/steprun"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// StepRunUpdate is the builder for updating StepRun entities.
type StepRunUpdate struct {
	config
	hooks      []Hook
	mutation   *StepRunMutation
	predicates []predicate.StepRun
}

// Where adds a new predicate for the builder.
func (sru *StepRunUpdate) Where(ps ...predicate.StepRun) *StepRunUpdate {
	sru.predicates = append(sru.predicates, ps...)
	return sru
}

// SetUpdatedAt sets the updated_at field.
func (sru *StepRunUpdate) SetUpdatedAt(t time.Time) *StepRunUpdate {
	sru.mutation.SetUpdatedAt(t)
	return sru
}

// SetStatus sets the status field.
func (sru *StepRunUpdate) SetStatus(s steprun.Status) *StepRunUpdate {
	sru.mutation.SetStatus(s)
	return sru
}

// SetStartedAt sets the startedAt field.
func (sru *StepRunUpdate) SetStartedAt(t time.Time) *StepRunUpdate {
	sru.mutation.SetStartedAt(t)
	return sru
}

// SetNillableStartedAt sets the startedAt field if the given value is not nil.
func (sru *StepRunUpdate) SetNillableStartedAt(t *time.Time) *StepRunUpdate {
	if t != nil {
		sru.SetStartedAt(*t)
	}
	return sru
}

// ClearStartedAt clears the value of startedAt.
func (sru *StepRunUpdate) ClearStartedAt() *StepRunUpdate {
	sru.mutation.ClearStartedAt()
	return sru
}

// SetEndedAt sets the endedAt field.
func (sru *StepRunUpdate) SetEndedAt(t time.Time) *StepRunUpdate {
	sru.mutation.SetEndedAt(t)
	return sru
}

// SetNillableEndedAt sets the endedAt field if the given value is not nil.
func (sru *StepRunUpdate) SetNillableEndedAt(t *time.Time) *StepRunUpdate {
	if t != nil {
		sru.SetEndedAt(*t)
	}
	return sru
}

// ClearEndedAt clears the value of endedAt.
func (sru *StepRunUpdate) ClearEndedAt() *StepRunUpdate {
	sru.mutation.ClearEndedAt()
	return sru
}

// SetIndex sets the index field.
func (sru *StepRunUpdate) SetIndex(i int) *StepRunUpdate {
	sru.mutation.ResetIndex()
	sru.mutation.SetIndex(i)
	return sru
}

// AddIndex adds i to index.
func (sru *StepRunUpdate) AddIndex(i int) *StepRunUpdate {
	sru.mutation.AddIndex(i)
	return sru
}

// SetParticipantsCount sets the participantsCount field.
func (sru *StepRunUpdate) SetParticipantsCount(i int) *StepRunUpdate {
	sru.mutation.ResetParticipantsCount()
	sru.mutation.SetParticipantsCount(i)
	return sru
}

// AddParticipantsCount adds i to participantsCount.
func (sru *StepRunUpdate) AddParticipantsCount(i int) *StepRunUpdate {
	sru.mutation.AddParticipantsCount(i)
	return sru
}

// SetHitID sets the hitID field.
func (sru *StepRunUpdate) SetHitID(s string) *StepRunUpdate {
	sru.mutation.SetHitID(s)
	return sru
}

// SetNillableHitID sets the hitID field if the given value is not nil.
func (sru *StepRunUpdate) SetNillableHitID(s *string) *StepRunUpdate {
	if s != nil {
		sru.SetHitID(*s)
	}
	return sru
}

// ClearHitID clears the value of hitID.
func (sru *StepRunUpdate) ClearHitID() *StepRunUpdate {
	sru.mutation.ClearHitID()
	return sru
}

// SetUrlToken sets the urlToken field.
func (sru *StepRunUpdate) SetUrlToken(s string) *StepRunUpdate {
	sru.mutation.SetUrlToken(s)
	return sru
}

// AddCreatedParticipantIDs adds the createdParticipants edge to Participant by ids.
func (sru *StepRunUpdate) AddCreatedParticipantIDs(ids ...string) *StepRunUpdate {
	sru.mutation.AddCreatedParticipantIDs(ids...)
	return sru
}

// AddCreatedParticipants adds the createdParticipants edges to Participant.
func (sru *StepRunUpdate) AddCreatedParticipants(p ...*Participant) *StepRunUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sru.AddCreatedParticipantIDs(ids...)
}

// AddParticipantIDs adds the participants edge to Participant by ids.
func (sru *StepRunUpdate) AddParticipantIDs(ids ...string) *StepRunUpdate {
	sru.mutation.AddParticipantIDs(ids...)
	return sru
}

// AddParticipants adds the participants edges to Participant.
func (sru *StepRunUpdate) AddParticipants(p ...*Participant) *StepRunUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sru.AddParticipantIDs(ids...)
}

// AddParticipationIDs adds the participations edge to Participation by ids.
func (sru *StepRunUpdate) AddParticipationIDs(ids ...string) *StepRunUpdate {
	sru.mutation.AddParticipationIDs(ids...)
	return sru
}

// AddParticipations adds the participations edges to Participation.
func (sru *StepRunUpdate) AddParticipations(p ...*Participation) *StepRunUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sru.AddParticipationIDs(ids...)
}

// SetStepID sets the step edge to Step by id.
func (sru *StepRunUpdate) SetStepID(id string) *StepRunUpdate {
	sru.mutation.SetStepID(id)
	return sru
}

// SetStep sets the step edge to Step.
func (sru *StepRunUpdate) SetStep(s *Step) *StepRunUpdate {
	return sru.SetStepID(s.ID)
}

// SetRunID sets the run edge to Run by id.
func (sru *StepRunUpdate) SetRunID(id string) *StepRunUpdate {
	sru.mutation.SetRunID(id)
	return sru
}

// SetNillableRunID sets the run edge to Run by id if the given value is not nil.
func (sru *StepRunUpdate) SetNillableRunID(id *string) *StepRunUpdate {
	if id != nil {
		sru = sru.SetRunID(*id)
	}
	return sru
}

// SetRun sets the run edge to Run.
func (sru *StepRunUpdate) SetRun(r *Run) *StepRunUpdate {
	return sru.SetRunID(r.ID)
}

// Mutation returns the StepRunMutation object of the builder.
func (sru *StepRunUpdate) Mutation() *StepRunMutation {
	return sru.mutation
}

// RemoveCreatedParticipantIDs removes the createdParticipants edge to Participant by ids.
func (sru *StepRunUpdate) RemoveCreatedParticipantIDs(ids ...string) *StepRunUpdate {
	sru.mutation.RemoveCreatedParticipantIDs(ids...)
	return sru
}

// RemoveCreatedParticipants removes createdParticipants edges to Participant.
func (sru *StepRunUpdate) RemoveCreatedParticipants(p ...*Participant) *StepRunUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sru.RemoveCreatedParticipantIDs(ids...)
}

// RemoveParticipantIDs removes the participants edge to Participant by ids.
func (sru *StepRunUpdate) RemoveParticipantIDs(ids ...string) *StepRunUpdate {
	sru.mutation.RemoveParticipantIDs(ids...)
	return sru
}

// RemoveParticipants removes participants edges to Participant.
func (sru *StepRunUpdate) RemoveParticipants(p ...*Participant) *StepRunUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sru.RemoveParticipantIDs(ids...)
}

// RemoveParticipationIDs removes the participations edge to Participation by ids.
func (sru *StepRunUpdate) RemoveParticipationIDs(ids ...string) *StepRunUpdate {
	sru.mutation.RemoveParticipationIDs(ids...)
	return sru
}

// RemoveParticipations removes participations edges to Participation.
func (sru *StepRunUpdate) RemoveParticipations(p ...*Participation) *StepRunUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sru.RemoveParticipationIDs(ids...)
}

// ClearStep clears the step edge to Step.
func (sru *StepRunUpdate) ClearStep() *StepRunUpdate {
	sru.mutation.ClearStep()
	return sru
}

// ClearRun clears the run edge to Run.
func (sru *StepRunUpdate) ClearRun() *StepRunUpdate {
	sru.mutation.ClearRun()
	return sru
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (sru *StepRunUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := sru.mutation.UpdatedAt(); !ok {
		v := steprun.UpdateDefaultUpdatedAt()
		sru.mutation.SetUpdatedAt(v)
	}
	if v, ok := sru.mutation.Status(); ok {
		if err := steprun.StatusValidator(v); err != nil {
			return 0, &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}

	if _, ok := sru.mutation.StepID(); sru.mutation.StepCleared() && !ok {
		return 0, errors.New("ent: clearing a unique edge \"step\"")
	}

	var (
		err      error
		affected int
	)
	if len(sru.hooks) == 0 {
		affected, err = sru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StepRunMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sru.mutation = mutation
			affected, err = sru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sru.hooks) - 1; i >= 0; i-- {
			mut = sru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sru *StepRunUpdate) SaveX(ctx context.Context) int {
	affected, err := sru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sru *StepRunUpdate) Exec(ctx context.Context) error {
	_, err := sru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sru *StepRunUpdate) ExecX(ctx context.Context) {
	if err := sru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sru *StepRunUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   steprun.Table,
			Columns: steprun.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: steprun.FieldID,
			},
		},
	}
	if ps := sru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: steprun.FieldUpdatedAt,
		})
	}
	if value, ok := sru.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: steprun.FieldStatus,
		})
	}
	if value, ok := sru.mutation.StartedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: steprun.FieldStartedAt,
		})
	}
	if sru.mutation.StartedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: steprun.FieldStartedAt,
		})
	}
	if value, ok := sru.mutation.EndedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: steprun.FieldEndedAt,
		})
	}
	if sru.mutation.EndedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: steprun.FieldEndedAt,
		})
	}
	if value, ok := sru.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: steprun.FieldIndex,
		})
	}
	if value, ok := sru.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: steprun.FieldIndex,
		})
	}
	if value, ok := sru.mutation.ParticipantsCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: steprun.FieldParticipantsCount,
		})
	}
	if value, ok := sru.mutation.AddedParticipantsCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: steprun.FieldParticipantsCount,
		})
	}
	if value, ok := sru.mutation.HitID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: steprun.FieldHitID,
		})
	}
	if sru.mutation.HitIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: steprun.FieldHitID,
		})
	}
	if value, ok := sru.mutation.UrlToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: steprun.FieldUrlToken,
		})
	}
	if nodes := sru.mutation.RemovedCreatedParticipantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   steprun.CreatedParticipantsTable,
			Columns: []string{steprun.CreatedParticipantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: participant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.CreatedParticipantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   steprun.CreatedParticipantsTable,
			Columns: []string{steprun.CreatedParticipantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: participant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := sru.mutation.RemovedParticipantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   steprun.ParticipantsTable,
			Columns: steprun.ParticipantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: participant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.ParticipantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   steprun.ParticipantsTable,
			Columns: steprun.ParticipantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: participant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := sru.mutation.RemovedParticipationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   steprun.ParticipationsTable,
			Columns: []string{steprun.ParticipationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: participation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.ParticipationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   steprun.ParticipationsTable,
			Columns: []string{steprun.ParticipationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: participation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sru.mutation.StepCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   steprun.StepTable,
			Columns: []string{steprun.StepColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: step.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.StepIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   steprun.StepTable,
			Columns: []string{steprun.StepColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: step.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sru.mutation.RunCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   steprun.RunTable,
			Columns: []string{steprun.RunColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: run.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sru.mutation.RunIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   steprun.RunTable,
			Columns: []string{steprun.RunColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: run.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{steprun.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// StepRunUpdateOne is the builder for updating a single StepRun entity.
type StepRunUpdateOne struct {
	config
	hooks    []Hook
	mutation *StepRunMutation
}

// SetUpdatedAt sets the updated_at field.
func (sruo *StepRunUpdateOne) SetUpdatedAt(t time.Time) *StepRunUpdateOne {
	sruo.mutation.SetUpdatedAt(t)
	return sruo
}

// SetStatus sets the status field.
func (sruo *StepRunUpdateOne) SetStatus(s steprun.Status) *StepRunUpdateOne {
	sruo.mutation.SetStatus(s)
	return sruo
}

// SetStartedAt sets the startedAt field.
func (sruo *StepRunUpdateOne) SetStartedAt(t time.Time) *StepRunUpdateOne {
	sruo.mutation.SetStartedAt(t)
	return sruo
}

// SetNillableStartedAt sets the startedAt field if the given value is not nil.
func (sruo *StepRunUpdateOne) SetNillableStartedAt(t *time.Time) *StepRunUpdateOne {
	if t != nil {
		sruo.SetStartedAt(*t)
	}
	return sruo
}

// ClearStartedAt clears the value of startedAt.
func (sruo *StepRunUpdateOne) ClearStartedAt() *StepRunUpdateOne {
	sruo.mutation.ClearStartedAt()
	return sruo
}

// SetEndedAt sets the endedAt field.
func (sruo *StepRunUpdateOne) SetEndedAt(t time.Time) *StepRunUpdateOne {
	sruo.mutation.SetEndedAt(t)
	return sruo
}

// SetNillableEndedAt sets the endedAt field if the given value is not nil.
func (sruo *StepRunUpdateOne) SetNillableEndedAt(t *time.Time) *StepRunUpdateOne {
	if t != nil {
		sruo.SetEndedAt(*t)
	}
	return sruo
}

// ClearEndedAt clears the value of endedAt.
func (sruo *StepRunUpdateOne) ClearEndedAt() *StepRunUpdateOne {
	sruo.mutation.ClearEndedAt()
	return sruo
}

// SetIndex sets the index field.
func (sruo *StepRunUpdateOne) SetIndex(i int) *StepRunUpdateOne {
	sruo.mutation.ResetIndex()
	sruo.mutation.SetIndex(i)
	return sruo
}

// AddIndex adds i to index.
func (sruo *StepRunUpdateOne) AddIndex(i int) *StepRunUpdateOne {
	sruo.mutation.AddIndex(i)
	return sruo
}

// SetParticipantsCount sets the participantsCount field.
func (sruo *StepRunUpdateOne) SetParticipantsCount(i int) *StepRunUpdateOne {
	sruo.mutation.ResetParticipantsCount()
	sruo.mutation.SetParticipantsCount(i)
	return sruo
}

// AddParticipantsCount adds i to participantsCount.
func (sruo *StepRunUpdateOne) AddParticipantsCount(i int) *StepRunUpdateOne {
	sruo.mutation.AddParticipantsCount(i)
	return sruo
}

// SetHitID sets the hitID field.
func (sruo *StepRunUpdateOne) SetHitID(s string) *StepRunUpdateOne {
	sruo.mutation.SetHitID(s)
	return sruo
}

// SetNillableHitID sets the hitID field if the given value is not nil.
func (sruo *StepRunUpdateOne) SetNillableHitID(s *string) *StepRunUpdateOne {
	if s != nil {
		sruo.SetHitID(*s)
	}
	return sruo
}

// ClearHitID clears the value of hitID.
func (sruo *StepRunUpdateOne) ClearHitID() *StepRunUpdateOne {
	sruo.mutation.ClearHitID()
	return sruo
}

// SetUrlToken sets the urlToken field.
func (sruo *StepRunUpdateOne) SetUrlToken(s string) *StepRunUpdateOne {
	sruo.mutation.SetUrlToken(s)
	return sruo
}

// AddCreatedParticipantIDs adds the createdParticipants edge to Participant by ids.
func (sruo *StepRunUpdateOne) AddCreatedParticipantIDs(ids ...string) *StepRunUpdateOne {
	sruo.mutation.AddCreatedParticipantIDs(ids...)
	return sruo
}

// AddCreatedParticipants adds the createdParticipants edges to Participant.
func (sruo *StepRunUpdateOne) AddCreatedParticipants(p ...*Participant) *StepRunUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sruo.AddCreatedParticipantIDs(ids...)
}

// AddParticipantIDs adds the participants edge to Participant by ids.
func (sruo *StepRunUpdateOne) AddParticipantIDs(ids ...string) *StepRunUpdateOne {
	sruo.mutation.AddParticipantIDs(ids...)
	return sruo
}

// AddParticipants adds the participants edges to Participant.
func (sruo *StepRunUpdateOne) AddParticipants(p ...*Participant) *StepRunUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sruo.AddParticipantIDs(ids...)
}

// AddParticipationIDs adds the participations edge to Participation by ids.
func (sruo *StepRunUpdateOne) AddParticipationIDs(ids ...string) *StepRunUpdateOne {
	sruo.mutation.AddParticipationIDs(ids...)
	return sruo
}

// AddParticipations adds the participations edges to Participation.
func (sruo *StepRunUpdateOne) AddParticipations(p ...*Participation) *StepRunUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sruo.AddParticipationIDs(ids...)
}

// SetStepID sets the step edge to Step by id.
func (sruo *StepRunUpdateOne) SetStepID(id string) *StepRunUpdateOne {
	sruo.mutation.SetStepID(id)
	return sruo
}

// SetStep sets the step edge to Step.
func (sruo *StepRunUpdateOne) SetStep(s *Step) *StepRunUpdateOne {
	return sruo.SetStepID(s.ID)
}

// SetRunID sets the run edge to Run by id.
func (sruo *StepRunUpdateOne) SetRunID(id string) *StepRunUpdateOne {
	sruo.mutation.SetRunID(id)
	return sruo
}

// SetNillableRunID sets the run edge to Run by id if the given value is not nil.
func (sruo *StepRunUpdateOne) SetNillableRunID(id *string) *StepRunUpdateOne {
	if id != nil {
		sruo = sruo.SetRunID(*id)
	}
	return sruo
}

// SetRun sets the run edge to Run.
func (sruo *StepRunUpdateOne) SetRun(r *Run) *StepRunUpdateOne {
	return sruo.SetRunID(r.ID)
}

// Mutation returns the StepRunMutation object of the builder.
func (sruo *StepRunUpdateOne) Mutation() *StepRunMutation {
	return sruo.mutation
}

// RemoveCreatedParticipantIDs removes the createdParticipants edge to Participant by ids.
func (sruo *StepRunUpdateOne) RemoveCreatedParticipantIDs(ids ...string) *StepRunUpdateOne {
	sruo.mutation.RemoveCreatedParticipantIDs(ids...)
	return sruo
}

// RemoveCreatedParticipants removes createdParticipants edges to Participant.
func (sruo *StepRunUpdateOne) RemoveCreatedParticipants(p ...*Participant) *StepRunUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sruo.RemoveCreatedParticipantIDs(ids...)
}

// RemoveParticipantIDs removes the participants edge to Participant by ids.
func (sruo *StepRunUpdateOne) RemoveParticipantIDs(ids ...string) *StepRunUpdateOne {
	sruo.mutation.RemoveParticipantIDs(ids...)
	return sruo
}

// RemoveParticipants removes participants edges to Participant.
func (sruo *StepRunUpdateOne) RemoveParticipants(p ...*Participant) *StepRunUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sruo.RemoveParticipantIDs(ids...)
}

// RemoveParticipationIDs removes the participations edge to Participation by ids.
func (sruo *StepRunUpdateOne) RemoveParticipationIDs(ids ...string) *StepRunUpdateOne {
	sruo.mutation.RemoveParticipationIDs(ids...)
	return sruo
}

// RemoveParticipations removes participations edges to Participation.
func (sruo *StepRunUpdateOne) RemoveParticipations(p ...*Participation) *StepRunUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sruo.RemoveParticipationIDs(ids...)
}

// ClearStep clears the step edge to Step.
func (sruo *StepRunUpdateOne) ClearStep() *StepRunUpdateOne {
	sruo.mutation.ClearStep()
	return sruo
}

// ClearRun clears the run edge to Run.
func (sruo *StepRunUpdateOne) ClearRun() *StepRunUpdateOne {
	sruo.mutation.ClearRun()
	return sruo
}

// Save executes the query and returns the updated entity.
func (sruo *StepRunUpdateOne) Save(ctx context.Context) (*StepRun, error) {
	if _, ok := sruo.mutation.UpdatedAt(); !ok {
		v := steprun.UpdateDefaultUpdatedAt()
		sruo.mutation.SetUpdatedAt(v)
	}
	if v, ok := sruo.mutation.Status(); ok {
		if err := steprun.StatusValidator(v); err != nil {
			return nil, &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}

	if _, ok := sruo.mutation.StepID(); sruo.mutation.StepCleared() && !ok {
		return nil, errors.New("ent: clearing a unique edge \"step\"")
	}

	var (
		err  error
		node *StepRun
	)
	if len(sruo.hooks) == 0 {
		node, err = sruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StepRunMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sruo.mutation = mutation
			node, err = sruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sruo.hooks) - 1; i >= 0; i-- {
			mut = sruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sruo *StepRunUpdateOne) SaveX(ctx context.Context) *StepRun {
	sr, err := sruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return sr
}

// Exec executes the query on the entity.
func (sruo *StepRunUpdateOne) Exec(ctx context.Context) error {
	_, err := sruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sruo *StepRunUpdateOne) ExecX(ctx context.Context) {
	if err := sruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sruo *StepRunUpdateOne) sqlSave(ctx context.Context) (sr *StepRun, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   steprun.Table,
			Columns: steprun.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: steprun.FieldID,
			},
		},
	}
	id, ok := sruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing StepRun.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := sruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: steprun.FieldUpdatedAt,
		})
	}
	if value, ok := sruo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: steprun.FieldStatus,
		})
	}
	if value, ok := sruo.mutation.StartedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: steprun.FieldStartedAt,
		})
	}
	if sruo.mutation.StartedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: steprun.FieldStartedAt,
		})
	}
	if value, ok := sruo.mutation.EndedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: steprun.FieldEndedAt,
		})
	}
	if sruo.mutation.EndedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: steprun.FieldEndedAt,
		})
	}
	if value, ok := sruo.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: steprun.FieldIndex,
		})
	}
	if value, ok := sruo.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: steprun.FieldIndex,
		})
	}
	if value, ok := sruo.mutation.ParticipantsCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: steprun.FieldParticipantsCount,
		})
	}
	if value, ok := sruo.mutation.AddedParticipantsCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: steprun.FieldParticipantsCount,
		})
	}
	if value, ok := sruo.mutation.HitID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: steprun.FieldHitID,
		})
	}
	if sruo.mutation.HitIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: steprun.FieldHitID,
		})
	}
	if value, ok := sruo.mutation.UrlToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: steprun.FieldUrlToken,
		})
	}
	if nodes := sruo.mutation.RemovedCreatedParticipantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   steprun.CreatedParticipantsTable,
			Columns: []string{steprun.CreatedParticipantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: participant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.CreatedParticipantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   steprun.CreatedParticipantsTable,
			Columns: []string{steprun.CreatedParticipantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: participant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := sruo.mutation.RemovedParticipantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   steprun.ParticipantsTable,
			Columns: steprun.ParticipantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: participant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.ParticipantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   steprun.ParticipantsTable,
			Columns: steprun.ParticipantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: participant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := sruo.mutation.RemovedParticipationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   steprun.ParticipationsTable,
			Columns: []string{steprun.ParticipationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: participation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.ParticipationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   steprun.ParticipationsTable,
			Columns: []string{steprun.ParticipationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: participation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sruo.mutation.StepCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   steprun.StepTable,
			Columns: []string{steprun.StepColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: step.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.StepIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   steprun.StepTable,
			Columns: []string{steprun.StepColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: step.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sruo.mutation.RunCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   steprun.RunTable,
			Columns: []string{steprun.RunColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: run.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sruo.mutation.RunIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   steprun.RunTable,
			Columns: []string{steprun.RunColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: run.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	sr = &StepRun{config: sruo.config}
	_spec.Assign = sr.assignValues
	_spec.ScanValues = sr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, sruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{steprun.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return sr, nil
}
