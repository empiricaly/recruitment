// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/empiricaly/recruitment/internal/ent/procedure"
	"github.com/empiricaly/recruitment/internal/ent/run"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// RunCreate is the builder for creating a Run entity.
type RunCreate struct {
	config
	mutation *RunMutation
	hooks    []Hook
}

// SetCreatedAt sets the createdAt field.
func (rc *RunCreate) SetCreatedAt(t time.Time) *RunCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the createdAt field if the given value is not nil.
func (rc *RunCreate) SetNillableCreatedAt(t *time.Time) *RunCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the updatedAt field.
func (rc *RunCreate) SetUpdatedAt(t time.Time) *RunCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the updatedAt field if the given value is not nil.
func (rc *RunCreate) SetNillableUpdatedAt(t *time.Time) *RunCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetName sets the name field.
func (rc *RunCreate) SetName(s string) *RunCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetStartAt sets the startAt field.
func (rc *RunCreate) SetStartAt(t time.Time) *RunCreate {
	rc.mutation.SetStartAt(t)
	return rc
}

// SetStartedAt sets the startedAt field.
func (rc *RunCreate) SetStartedAt(t time.Time) *RunCreate {
	rc.mutation.SetStartedAt(t)
	return rc
}

// SetEndedAt sets the endedAt field.
func (rc *RunCreate) SetEndedAt(t time.Time) *RunCreate {
	rc.mutation.SetEndedAt(t)
	return rc
}

// SetError sets the error field.
func (rc *RunCreate) SetError(s string) *RunCreate {
	rc.mutation.SetError(s)
	return rc
}

// SetID sets the id field.
func (rc *RunCreate) SetID(s string) *RunCreate {
	rc.mutation.SetID(s)
	return rc
}

// SetProcedureID sets the procedure edge to Procedure by id.
func (rc *RunCreate) SetProcedureID(id string) *RunCreate {
	rc.mutation.SetProcedureID(id)
	return rc
}

// SetNillableProcedureID sets the procedure edge to Procedure by id if the given value is not nil.
func (rc *RunCreate) SetNillableProcedureID(id *string) *RunCreate {
	if id != nil {
		rc = rc.SetProcedureID(*id)
	}
	return rc
}

// SetProcedure sets the procedure edge to Procedure.
func (rc *RunCreate) SetProcedure(p *Procedure) *RunCreate {
	return rc.SetProcedureID(p.ID)
}

// Mutation returns the RunMutation object of the builder.
func (rc *RunCreate) Mutation() *RunMutation {
	return rc.mutation
}

// Save creates the Run in the database.
func (rc *RunCreate) Save(ctx context.Context) (*Run, error) {
	if err := rc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *Run
	)
	if len(rc.hooks) == 0 {
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RunMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rc.mutation = mutation
			node, err = rc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RunCreate) SaveX(ctx context.Context) *Run {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rc *RunCreate) preSave() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := run.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := run.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := rc.mutation.StartAt(); !ok {
		return &ValidationError{Name: "startAt", err: errors.New("ent: missing required field \"startAt\"")}
	}
	if _, ok := rc.mutation.StartedAt(); !ok {
		return &ValidationError{Name: "startedAt", err: errors.New("ent: missing required field \"startedAt\"")}
	}
	if _, ok := rc.mutation.EndedAt(); !ok {
		return &ValidationError{Name: "endedAt", err: errors.New("ent: missing required field \"endedAt\"")}
	}
	if _, ok := rc.mutation.Error(); !ok {
		return &ValidationError{Name: "error", err: errors.New("ent: missing required field \"error\"")}
	}
	return nil
}

func (rc *RunCreate) sqlSave(ctx context.Context) (*Run, error) {
	r, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return r, nil
}

func (rc *RunCreate) createSpec() (*Run, *sqlgraph.CreateSpec) {
	var (
		r     = &Run{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: run.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: run.FieldID,
			},
		}
	)
	if id, ok := rc.mutation.ID(); ok {
		r.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: run.FieldCreatedAt,
		})
		r.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: run.FieldUpdatedAt,
		})
		r.UpdatedAt = value
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: run.FieldName,
		})
		r.Name = value
	}
	if value, ok := rc.mutation.StartAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: run.FieldStartAt,
		})
		r.StartAt = value
	}
	if value, ok := rc.mutation.StartedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: run.FieldStartedAt,
		})
		r.StartedAt = value
	}
	if value, ok := rc.mutation.EndedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: run.FieldEndedAt,
		})
		r.EndedAt = value
	}
	if value, ok := rc.mutation.Error(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: run.FieldError,
		})
		r.Error = value
	}
	if nodes := rc.mutation.ProcedureIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   run.ProcedureTable,
			Columns: []string{run.ProcedureColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: procedure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return r, _spec
}

// RunCreateBulk is the builder for creating a bulk of Run entities.
type RunCreateBulk struct {
	config
	builders []*RunCreate
}

// Save creates the Run entities in the database.
func (rcb *RunCreateBulk) Save(ctx context.Context) ([]*Run, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Run, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*RunMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (rcb *RunCreateBulk) SaveX(ctx context.Context) []*Run {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
