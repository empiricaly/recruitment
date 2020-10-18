// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/empiricaly/recruitment/internal/ent/participant"
	"github.com/empiricaly/recruitment/internal/ent/participation"
	"github.com/empiricaly/recruitment/internal/ent/steprun"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// ParticipationCreate is the builder for creating a Participation entity.
type ParticipationCreate struct {
	config
	mutation *ParticipationMutation
	hooks    []Hook
}

// SetCreatedAt sets the created_at field.
func (pc *ParticipationCreate) SetCreatedAt(t time.Time) *ParticipationCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (pc *ParticipationCreate) SetNillableCreatedAt(t *time.Time) *ParticipationCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the updated_at field.
func (pc *ParticipationCreate) SetUpdatedAt(t time.Time) *ParticipationCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (pc *ParticipationCreate) SetNillableUpdatedAt(t *time.Time) *ParticipationCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetMturkWorkerId sets the mturkWorkerId field.
func (pc *ParticipationCreate) SetMturkWorkerId(s string) *ParticipationCreate {
	pc.mutation.SetMturkWorkerId(s)
	return pc
}

// SetMturkAssignmentID sets the mturkAssignmentID field.
func (pc *ParticipationCreate) SetMturkAssignmentID(s string) *ParticipationCreate {
	pc.mutation.SetMturkAssignmentID(s)
	return pc
}

// SetMturkHitID sets the mturkHitID field.
func (pc *ParticipationCreate) SetMturkHitID(s string) *ParticipationCreate {
	pc.mutation.SetMturkHitID(s)
	return pc
}

// SetMturkAcceptedAt sets the mturkAcceptedAt field.
func (pc *ParticipationCreate) SetMturkAcceptedAt(t time.Time) *ParticipationCreate {
	pc.mutation.SetMturkAcceptedAt(t)
	return pc
}

// SetMturkSubmittedAt sets the mturkSubmittedAt field.
func (pc *ParticipationCreate) SetMturkSubmittedAt(t time.Time) *ParticipationCreate {
	pc.mutation.SetMturkSubmittedAt(t)
	return pc
}

// SetID sets the id field.
func (pc *ParticipationCreate) SetID(s string) *ParticipationCreate {
	pc.mutation.SetID(s)
	return pc
}

// SetStepRunID sets the stepRun edge to StepRun by id.
func (pc *ParticipationCreate) SetStepRunID(id string) *ParticipationCreate {
	pc.mutation.SetStepRunID(id)
	return pc
}

// SetNillableStepRunID sets the stepRun edge to StepRun by id if the given value is not nil.
func (pc *ParticipationCreate) SetNillableStepRunID(id *string) *ParticipationCreate {
	if id != nil {
		pc = pc.SetStepRunID(*id)
	}
	return pc
}

// SetStepRun sets the stepRun edge to StepRun.
func (pc *ParticipationCreate) SetStepRun(s *StepRun) *ParticipationCreate {
	return pc.SetStepRunID(s.ID)
}

// SetParticipantID sets the participant edge to Participant by id.
func (pc *ParticipationCreate) SetParticipantID(id string) *ParticipationCreate {
	pc.mutation.SetParticipantID(id)
	return pc
}

// SetNillableParticipantID sets the participant edge to Participant by id if the given value is not nil.
func (pc *ParticipationCreate) SetNillableParticipantID(id *string) *ParticipationCreate {
	if id != nil {
		pc = pc.SetParticipantID(*id)
	}
	return pc
}

// SetParticipant sets the participant edge to Participant.
func (pc *ParticipationCreate) SetParticipant(p *Participant) *ParticipationCreate {
	return pc.SetParticipantID(p.ID)
}

// Mutation returns the ParticipationMutation object of the builder.
func (pc *ParticipationCreate) Mutation() *ParticipationMutation {
	return pc.mutation
}

// Save creates the Participation in the database.
func (pc *ParticipationCreate) Save(ctx context.Context) (*Participation, error) {
	var (
		err  error
		node *Participation
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ParticipationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ParticipationCreate) SaveX(ctx context.Context) *Participation {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pc *ParticipationCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := participation.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := participation.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ParticipationCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := pc.mutation.MturkWorkerId(); !ok {
		return &ValidationError{Name: "mturkWorkerId", err: errors.New("ent: missing required field \"mturkWorkerId\"")}
	}
	if _, ok := pc.mutation.MturkAssignmentID(); !ok {
		return &ValidationError{Name: "mturkAssignmentID", err: errors.New("ent: missing required field \"mturkAssignmentID\"")}
	}
	if _, ok := pc.mutation.MturkHitID(); !ok {
		return &ValidationError{Name: "mturkHitID", err: errors.New("ent: missing required field \"mturkHitID\"")}
	}
	if _, ok := pc.mutation.MturkAcceptedAt(); !ok {
		return &ValidationError{Name: "mturkAcceptedAt", err: errors.New("ent: missing required field \"mturkAcceptedAt\"")}
	}
	if _, ok := pc.mutation.MturkSubmittedAt(); !ok {
		return &ValidationError{Name: "mturkSubmittedAt", err: errors.New("ent: missing required field \"mturkSubmittedAt\"")}
	}
	if v, ok := pc.mutation.ID(); ok {
		if err := participation.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	return nil
}

func (pc *ParticipationCreate) sqlSave(ctx context.Context) (*Participation, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (pc *ParticipationCreate) createSpec() (*Participation, *sqlgraph.CreateSpec) {
	var (
		_node = &Participation{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: participation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: participation.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: participation.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: participation.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.MturkWorkerId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: participation.FieldMturkWorkerId,
		})
		_node.MturkWorkerId = value
	}
	if value, ok := pc.mutation.MturkAssignmentID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: participation.FieldMturkAssignmentID,
		})
		_node.MturkAssignmentID = value
	}
	if value, ok := pc.mutation.MturkHitID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: participation.FieldMturkHitID,
		})
		_node.MturkHitID = value
	}
	if value, ok := pc.mutation.MturkAcceptedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: participation.FieldMturkAcceptedAt,
		})
		_node.MturkAcceptedAt = value
	}
	if value, ok := pc.mutation.MturkSubmittedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: participation.FieldMturkSubmittedAt,
		})
		_node.MturkSubmittedAt = value
	}
	if nodes := pc.mutation.StepRunIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   participation.StepRunTable,
			Columns: []string{participation.StepRunColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: steprun.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ParticipantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   participation.ParticipantTable,
			Columns: []string{participation.ParticipantColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ParticipationCreateBulk is the builder for creating a bulk of Participation entities.
type ParticipationCreateBulk struct {
	config
	builders []*ParticipationCreate
}

// Save creates the Participation entities in the database.
func (pcb *ParticipationCreateBulk) Save(ctx context.Context) ([]*Participation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Participation, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ParticipationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (pcb *ParticipationCreateBulk) SaveX(ctx context.Context) []*Participation {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
