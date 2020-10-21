// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/empiricaly/recruitment/internal/ent/datum"
	"github.com/empiricaly/recruitment/internal/ent/participant"
	"github.com/empiricaly/recruitment/internal/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// DatumUpdate is the builder for updating Datum entities.
type DatumUpdate struct {
	config
	hooks      []Hook
	mutation   *DatumMutation
	predicates []predicate.Datum
}

// Where adds a new predicate for the builder.
func (du *DatumUpdate) Where(ps ...predicate.Datum) *DatumUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetUpdatedAt sets the updated_at field.
func (du *DatumUpdate) SetUpdatedAt(t time.Time) *DatumUpdate {
	du.mutation.SetUpdatedAt(t)
	return du
}

// SetKey sets the key field.
func (du *DatumUpdate) SetKey(s string) *DatumUpdate {
	du.mutation.SetKey(s)
	return du
}

// SetVal sets the val field.
func (du *DatumUpdate) SetVal(s string) *DatumUpdate {
	du.mutation.SetVal(s)
	return du
}

// SetIndex sets the index field.
func (du *DatumUpdate) SetIndex(i int) *DatumUpdate {
	du.mutation.ResetIndex()
	du.mutation.SetIndex(i)
	return du
}

// SetNillableIndex sets the index field if the given value is not nil.
func (du *DatumUpdate) SetNillableIndex(i *int) *DatumUpdate {
	if i != nil {
		du.SetIndex(*i)
	}
	return du
}

// AddIndex adds i to index.
func (du *DatumUpdate) AddIndex(i int) *DatumUpdate {
	du.mutation.AddIndex(i)
	return du
}

// SetCurrent sets the current field.
func (du *DatumUpdate) SetCurrent(b bool) *DatumUpdate {
	du.mutation.SetCurrent(b)
	return du
}

// SetNillableCurrent sets the current field if the given value is not nil.
func (du *DatumUpdate) SetNillableCurrent(b *bool) *DatumUpdate {
	if b != nil {
		du.SetCurrent(*b)
	}
	return du
}

// SetVersion sets the version field.
func (du *DatumUpdate) SetVersion(i int) *DatumUpdate {
	du.mutation.ResetVersion()
	du.mutation.SetVersion(i)
	return du
}

// SetNillableVersion sets the version field if the given value is not nil.
func (du *DatumUpdate) SetNillableVersion(i *int) *DatumUpdate {
	if i != nil {
		du.SetVersion(*i)
	}
	return du
}

// AddVersion adds i to version.
func (du *DatumUpdate) AddVersion(i int) *DatumUpdate {
	du.mutation.AddVersion(i)
	return du
}

// SetDeletedAt sets the deletedAt field.
func (du *DatumUpdate) SetDeletedAt(t time.Time) *DatumUpdate {
	du.mutation.SetDeletedAt(t)
	return du
}

// SetNillableDeletedAt sets the deletedAt field if the given value is not nil.
func (du *DatumUpdate) SetNillableDeletedAt(t *time.Time) *DatumUpdate {
	if t != nil {
		du.SetDeletedAt(*t)
	}
	return du
}

// ClearDeletedAt clears the value of deletedAt.
func (du *DatumUpdate) ClearDeletedAt() *DatumUpdate {
	du.mutation.ClearDeletedAt()
	return du
}

// SetParticipantID sets the participant edge to Participant by id.
func (du *DatumUpdate) SetParticipantID(id string) *DatumUpdate {
	du.mutation.SetParticipantID(id)
	return du
}

// SetParticipant sets the participant edge to Participant.
func (du *DatumUpdate) SetParticipant(p *Participant) *DatumUpdate {
	return du.SetParticipantID(p.ID)
}

// Mutation returns the DatumMutation object of the builder.
func (du *DatumUpdate) Mutation() *DatumMutation {
	return du.mutation
}

// ClearParticipant clears the "participant" edge to type Participant.
func (du *DatumUpdate) ClearParticipant() *DatumUpdate {
	du.mutation.ClearParticipant()
	return du
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DatumUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	du.defaults()
	if len(du.hooks) == 0 {
		if err = du.check(); err != nil {
			return 0, err
		}
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DatumMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = du.check(); err != nil {
				return 0, err
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DatumUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DatumUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DatumUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DatumUpdate) defaults() {
	if _, ok := du.mutation.UpdatedAt(); !ok {
		v := datum.UpdateDefaultUpdatedAt()
		du.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DatumUpdate) check() error {
	if v, ok := du.mutation.Key(); ok {
		if err := datum.KeyValidator(v); err != nil {
			return &ValidationError{Name: "key", err: fmt.Errorf("ent: validator failed for field \"key\": %w", err)}
		}
	}
	if v, ok := du.mutation.Val(); ok {
		if err := datum.ValValidator(v); err != nil {
			return &ValidationError{Name: "val", err: fmt.Errorf("ent: validator failed for field \"val\": %w", err)}
		}
	}
	if _, ok := du.mutation.ParticipantID(); du.mutation.ParticipantCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"participant\"")
	}
	return nil
}

func (du *DatumUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   datum.Table,
			Columns: datum.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: datum.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: datum.FieldUpdatedAt,
		})
	}
	if value, ok := du.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datum.FieldKey,
		})
	}
	if value, ok := du.mutation.Val(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datum.FieldVal,
		})
	}
	if value, ok := du.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: datum.FieldIndex,
		})
	}
	if value, ok := du.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: datum.FieldIndex,
		})
	}
	if value, ok := du.mutation.Current(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: datum.FieldCurrent,
		})
	}
	if value, ok := du.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: datum.FieldVersion,
		})
	}
	if value, ok := du.mutation.AddedVersion(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: datum.FieldVersion,
		})
	}
	if value, ok := du.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: datum.FieldDeletedAt,
		})
	}
	if du.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: datum.FieldDeletedAt,
		})
	}
	if du.mutation.ParticipantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   datum.ParticipantTable,
			Columns: []string{datum.ParticipantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: participant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.ParticipantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   datum.ParticipantTable,
			Columns: []string{datum.ParticipantColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{datum.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DatumUpdateOne is the builder for updating a single Datum entity.
type DatumUpdateOne struct {
	config
	hooks    []Hook
	mutation *DatumMutation
}

// SetUpdatedAt sets the updated_at field.
func (duo *DatumUpdateOne) SetUpdatedAt(t time.Time) *DatumUpdateOne {
	duo.mutation.SetUpdatedAt(t)
	return duo
}

// SetKey sets the key field.
func (duo *DatumUpdateOne) SetKey(s string) *DatumUpdateOne {
	duo.mutation.SetKey(s)
	return duo
}

// SetVal sets the val field.
func (duo *DatumUpdateOne) SetVal(s string) *DatumUpdateOne {
	duo.mutation.SetVal(s)
	return duo
}

// SetIndex sets the index field.
func (duo *DatumUpdateOne) SetIndex(i int) *DatumUpdateOne {
	duo.mutation.ResetIndex()
	duo.mutation.SetIndex(i)
	return duo
}

// SetNillableIndex sets the index field if the given value is not nil.
func (duo *DatumUpdateOne) SetNillableIndex(i *int) *DatumUpdateOne {
	if i != nil {
		duo.SetIndex(*i)
	}
	return duo
}

// AddIndex adds i to index.
func (duo *DatumUpdateOne) AddIndex(i int) *DatumUpdateOne {
	duo.mutation.AddIndex(i)
	return duo
}

// SetCurrent sets the current field.
func (duo *DatumUpdateOne) SetCurrent(b bool) *DatumUpdateOne {
	duo.mutation.SetCurrent(b)
	return duo
}

// SetNillableCurrent sets the current field if the given value is not nil.
func (duo *DatumUpdateOne) SetNillableCurrent(b *bool) *DatumUpdateOne {
	if b != nil {
		duo.SetCurrent(*b)
	}
	return duo
}

// SetVersion sets the version field.
func (duo *DatumUpdateOne) SetVersion(i int) *DatumUpdateOne {
	duo.mutation.ResetVersion()
	duo.mutation.SetVersion(i)
	return duo
}

// SetNillableVersion sets the version field if the given value is not nil.
func (duo *DatumUpdateOne) SetNillableVersion(i *int) *DatumUpdateOne {
	if i != nil {
		duo.SetVersion(*i)
	}
	return duo
}

// AddVersion adds i to version.
func (duo *DatumUpdateOne) AddVersion(i int) *DatumUpdateOne {
	duo.mutation.AddVersion(i)
	return duo
}

// SetDeletedAt sets the deletedAt field.
func (duo *DatumUpdateOne) SetDeletedAt(t time.Time) *DatumUpdateOne {
	duo.mutation.SetDeletedAt(t)
	return duo
}

// SetNillableDeletedAt sets the deletedAt field if the given value is not nil.
func (duo *DatumUpdateOne) SetNillableDeletedAt(t *time.Time) *DatumUpdateOne {
	if t != nil {
		duo.SetDeletedAt(*t)
	}
	return duo
}

// ClearDeletedAt clears the value of deletedAt.
func (duo *DatumUpdateOne) ClearDeletedAt() *DatumUpdateOne {
	duo.mutation.ClearDeletedAt()
	return duo
}

// SetParticipantID sets the participant edge to Participant by id.
func (duo *DatumUpdateOne) SetParticipantID(id string) *DatumUpdateOne {
	duo.mutation.SetParticipantID(id)
	return duo
}

// SetParticipant sets the participant edge to Participant.
func (duo *DatumUpdateOne) SetParticipant(p *Participant) *DatumUpdateOne {
	return duo.SetParticipantID(p.ID)
}

// Mutation returns the DatumMutation object of the builder.
func (duo *DatumUpdateOne) Mutation() *DatumMutation {
	return duo.mutation
}

// ClearParticipant clears the "participant" edge to type Participant.
func (duo *DatumUpdateOne) ClearParticipant() *DatumUpdateOne {
	duo.mutation.ClearParticipant()
	return duo
}

// Save executes the query and returns the updated entity.
func (duo *DatumUpdateOne) Save(ctx context.Context) (*Datum, error) {
	var (
		err  error
		node *Datum
	)
	duo.defaults()
	if len(duo.hooks) == 0 {
		if err = duo.check(); err != nil {
			return nil, err
		}
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DatumMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = duo.check(); err != nil {
				return nil, err
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DatumUpdateOne) SaveX(ctx context.Context) *Datum {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DatumUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DatumUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DatumUpdateOne) defaults() {
	if _, ok := duo.mutation.UpdatedAt(); !ok {
		v := datum.UpdateDefaultUpdatedAt()
		duo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DatumUpdateOne) check() error {
	if v, ok := duo.mutation.Key(); ok {
		if err := datum.KeyValidator(v); err != nil {
			return &ValidationError{Name: "key", err: fmt.Errorf("ent: validator failed for field \"key\": %w", err)}
		}
	}
	if v, ok := duo.mutation.Val(); ok {
		if err := datum.ValValidator(v); err != nil {
			return &ValidationError{Name: "val", err: fmt.Errorf("ent: validator failed for field \"val\": %w", err)}
		}
	}
	if _, ok := duo.mutation.ParticipantID(); duo.mutation.ParticipantCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"participant\"")
	}
	return nil
}

func (duo *DatumUpdateOne) sqlSave(ctx context.Context) (_node *Datum, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   datum.Table,
			Columns: datum.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: datum.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Datum.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: datum.FieldUpdatedAt,
		})
	}
	if value, ok := duo.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datum.FieldKey,
		})
	}
	if value, ok := duo.mutation.Val(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: datum.FieldVal,
		})
	}
	if value, ok := duo.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: datum.FieldIndex,
		})
	}
	if value, ok := duo.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: datum.FieldIndex,
		})
	}
	if value, ok := duo.mutation.Current(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: datum.FieldCurrent,
		})
	}
	if value, ok := duo.mutation.Version(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: datum.FieldVersion,
		})
	}
	if value, ok := duo.mutation.AddedVersion(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: datum.FieldVersion,
		})
	}
	if value, ok := duo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: datum.FieldDeletedAt,
		})
	}
	if duo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: datum.FieldDeletedAt,
		})
	}
	if duo.mutation.ParticipantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   datum.ParticipantTable,
			Columns: []string{datum.ParticipantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: participant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.ParticipantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   datum.ParticipantTable,
			Columns: []string{datum.ParticipantColumn},
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
	_node = &Datum{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{datum.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
