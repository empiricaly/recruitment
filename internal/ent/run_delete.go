// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/empiricaly/recruitment/internal/ent/predicate"
	"github.com/empiricaly/recruitment/internal/ent/run"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// RunDelete is the builder for deleting a Run entity.
type RunDelete struct {
	config
	hooks      []Hook
	mutation   *RunMutation
	predicates []predicate.Run
}

// Where adds a new predicate to the delete builder.
func (rd *RunDelete) Where(ps ...predicate.Run) *RunDelete {
	rd.predicates = append(rd.predicates, ps...)
	return rd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rd *RunDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rd.hooks) == 0 {
		affected, err = rd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RunMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rd.mutation = mutation
			affected, err = rd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rd.hooks) - 1; i >= 0; i-- {
			mut = rd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rd *RunDelete) ExecX(ctx context.Context) int {
	n, err := rd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rd *RunDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: run.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: run.FieldID,
			},
		},
	}
	if ps := rd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rd.driver, _spec)
}

// RunDeleteOne is the builder for deleting a single Run entity.
type RunDeleteOne struct {
	rd *RunDelete
}

// Exec executes the deletion query.
func (rdo *RunDeleteOne) Exec(ctx context.Context) error {
	n, err := rdo.rd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{run.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdo *RunDeleteOne) ExecX(ctx context.Context) {
	rdo.rd.ExecX(ctx)
}
