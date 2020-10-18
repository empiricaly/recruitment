// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/empiricaly/recruitment/internal/ent/datum"
	"github.com/empiricaly/recruitment/internal/ent/participant"
	"github.com/empiricaly/recruitment/internal/ent/participation"
	"github.com/empiricaly/recruitment/internal/ent/providerid"
	"github.com/empiricaly/recruitment/internal/ent/steprun"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// ParticipantCreate is the builder for creating a Participant entity.
type ParticipantCreate struct {
	config
	mutation *ParticipantMutation
	hooks    []Hook
}

// SetCreatedAt sets the created_at field.
func (pc *ParticipantCreate) SetCreatedAt(t time.Time) *ParticipantCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (pc *ParticipantCreate) SetNillableCreatedAt(t *time.Time) *ParticipantCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the updated_at field.
func (pc *ParticipantCreate) SetUpdatedAt(t time.Time) *ParticipantCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (pc *ParticipantCreate) SetNillableUpdatedAt(t *time.Time) *ParticipantCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetMturkWorkerID sets the mturkWorkerID field.
func (pc *ParticipantCreate) SetMturkWorkerID(s string) *ParticipantCreate {
	pc.mutation.SetMturkWorkerID(s)
	return pc
}

// SetNillableMturkWorkerID sets the mturkWorkerID field if the given value is not nil.
func (pc *ParticipantCreate) SetNillableMturkWorkerID(s *string) *ParticipantCreate {
	if s != nil {
		pc.SetMturkWorkerID(*s)
	}
	return pc
}

// SetID sets the id field.
func (pc *ParticipantCreate) SetID(s string) *ParticipantCreate {
	pc.mutation.SetID(s)
	return pc
}

// AddDatumIDs adds the data edge to Datum by ids.
func (pc *ParticipantCreate) AddDatumIDs(ids ...string) *ParticipantCreate {
	pc.mutation.AddDatumIDs(ids...)
	return pc
}

// AddData adds the data edges to Datum.
func (pc *ParticipantCreate) AddData(d ...*Datum) *ParticipantCreate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pc.AddDatumIDs(ids...)
}

// AddProviderIDIDs adds the providerIDs edge to ProviderID by ids.
func (pc *ParticipantCreate) AddProviderIDIDs(ids ...string) *ParticipantCreate {
	pc.mutation.AddProviderIDIDs(ids...)
	return pc
}

// AddProviderIDs adds the providerIDs edges to ProviderID.
func (pc *ParticipantCreate) AddProviderIDs(p ...*ProviderID) *ParticipantCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddProviderIDIDs(ids...)
}

// AddParticipationIDs adds the participations edge to Participation by ids.
func (pc *ParticipantCreate) AddParticipationIDs(ids ...string) *ParticipantCreate {
	pc.mutation.AddParticipationIDs(ids...)
	return pc
}

// AddParticipations adds the participations edges to Participation.
func (pc *ParticipantCreate) AddParticipations(p ...*Participation) *ParticipantCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddParticipationIDs(ids...)
}

// SetCreatedByID sets the createdBy edge to StepRun by id.
func (pc *ParticipantCreate) SetCreatedByID(id string) *ParticipantCreate {
	pc.mutation.SetCreatedByID(id)
	return pc
}

// SetNillableCreatedByID sets the createdBy edge to StepRun by id if the given value is not nil.
func (pc *ParticipantCreate) SetNillableCreatedByID(id *string) *ParticipantCreate {
	if id != nil {
		pc = pc.SetCreatedByID(*id)
	}
	return pc
}

// SetCreatedBy sets the createdBy edge to StepRun.
func (pc *ParticipantCreate) SetCreatedBy(s *StepRun) *ParticipantCreate {
	return pc.SetCreatedByID(s.ID)
}

// AddStepIDs adds the steps edge to StepRun by ids.
func (pc *ParticipantCreate) AddStepIDs(ids ...string) *ParticipantCreate {
	pc.mutation.AddStepIDs(ids...)
	return pc
}

// AddSteps adds the steps edges to StepRun.
func (pc *ParticipantCreate) AddSteps(s ...*StepRun) *ParticipantCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddStepIDs(ids...)
}

// Mutation returns the ParticipantMutation object of the builder.
func (pc *ParticipantCreate) Mutation() *ParticipantMutation {
	return pc.mutation
}

// Save creates the Participant in the database.
func (pc *ParticipantCreate) Save(ctx context.Context) (*Participant, error) {
	var (
		err  error
		node *Participant
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ParticipantMutation)
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
func (pc *ParticipantCreate) SaveX(ctx context.Context) *Participant {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pc *ParticipantCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := participant.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := participant.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ParticipantCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if v, ok := pc.mutation.ID(); ok {
		if err := participant.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	return nil
}

func (pc *ParticipantCreate) sqlSave(ctx context.Context) (*Participant, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (pc *ParticipantCreate) createSpec() (*Participant, *sqlgraph.CreateSpec) {
	var (
		_node = &Participant{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: participant.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: participant.FieldID,
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
			Column: participant.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: participant.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.MturkWorkerID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: participant.FieldMturkWorkerID,
		})
		_node.MturkWorkerID = &value
	}
	if nodes := pc.mutation.DataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   participant.DataTable,
			Columns: []string{participant.DataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: datum.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ProviderIDsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   participant.ProviderIDsTable,
			Columns: []string{participant.ProviderIDsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: providerid.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ParticipationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   participant.ParticipationsTable,
			Columns: []string{participant.ParticipationsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   participant.CreatedByTable,
			Columns: []string{participant.CreatedByColumn},
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
	if nodes := pc.mutation.StepsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   participant.StepsTable,
			Columns: participant.StepsPrimaryKey,
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
	return _node, _spec
}

// ParticipantCreateBulk is the builder for creating a bulk of Participant entities.
type ParticipantCreateBulk struct {
	config
	builders []*ParticipantCreate
}

// Save creates the Participant entities in the database.
func (pcb *ParticipantCreateBulk) Save(ctx context.Context) ([]*Participant, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Participant, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ParticipantMutation)
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
func (pcb *ParticipantCreateBulk) SaveX(ctx context.Context) []*Participant {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
