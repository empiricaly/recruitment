// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/empiricaly/recruitment/internal/ent/admin"
	"github.com/empiricaly/recruitment/internal/ent/project"
	"github.com/empiricaly/recruitment/internal/ent/run"
	"github.com/empiricaly/recruitment/internal/ent/step"
	"github.com/empiricaly/recruitment/internal/ent/template"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// TemplateCreate is the builder for creating a Template entity.
type TemplateCreate struct {
	config
	mutation *TemplateMutation
	hooks    []Hook
}

// SetCreatedAt sets the created_at field.
func (tc *TemplateCreate) SetCreatedAt(t time.Time) *TemplateCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (tc *TemplateCreate) SetNillableCreatedAt(t *time.Time) *TemplateCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the updated_at field.
func (tc *TemplateCreate) SetUpdatedAt(t time.Time) *TemplateCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (tc *TemplateCreate) SetNillableUpdatedAt(t *time.Time) *TemplateCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetName sets the name field.
func (tc *TemplateCreate) SetName(s string) *TemplateCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetSelectionType sets the selectionType field.
func (tc *TemplateCreate) SetSelectionType(tt template.SelectionType) *TemplateCreate {
	tc.mutation.SetSelectionType(tt)
	return tc
}

// SetParticipantCount sets the participantCount field.
func (tc *TemplateCreate) SetParticipantCount(i int) *TemplateCreate {
	tc.mutation.SetParticipantCount(i)
	return tc
}

// SetNillableParticipantCount sets the participantCount field if the given value is not nil.
func (tc *TemplateCreate) SetNillableParticipantCount(i *int) *TemplateCreate {
	if i != nil {
		tc.SetParticipantCount(*i)
	}
	return tc
}

// SetInternalCriteria sets the internalCriteria field.
func (tc *TemplateCreate) SetInternalCriteria(b []byte) *TemplateCreate {
	tc.mutation.SetInternalCriteria(b)
	return tc
}

// SetMturkCriteria sets the mturkCriteria field.
func (tc *TemplateCreate) SetMturkCriteria(b []byte) *TemplateCreate {
	tc.mutation.SetMturkCriteria(b)
	return tc
}

// SetAdult sets the adult field.
func (tc *TemplateCreate) SetAdult(b bool) *TemplateCreate {
	tc.mutation.SetAdult(b)
	return tc
}

// SetNillableAdult sets the adult field if the given value is not nil.
func (tc *TemplateCreate) SetNillableAdult(b *bool) *TemplateCreate {
	if b != nil {
		tc.SetAdult(*b)
	}
	return tc
}

// SetSandbox sets the sandbox field.
func (tc *TemplateCreate) SetSandbox(b bool) *TemplateCreate {
	tc.mutation.SetSandbox(b)
	return tc
}

// SetNillableSandbox sets the sandbox field if the given value is not nil.
func (tc *TemplateCreate) SetNillableSandbox(b *bool) *TemplateCreate {
	if b != nil {
		tc.SetSandbox(*b)
	}
	return tc
}

// SetID sets the id field.
func (tc *TemplateCreate) SetID(s string) *TemplateCreate {
	tc.mutation.SetID(s)
	return tc
}

// AddStepIDs adds the steps edge to Step by ids.
func (tc *TemplateCreate) AddStepIDs(ids ...string) *TemplateCreate {
	tc.mutation.AddStepIDs(ids...)
	return tc
}

// AddSteps adds the steps edges to Step.
func (tc *TemplateCreate) AddSteps(s ...*Step) *TemplateCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tc.AddStepIDs(ids...)
}

// SetProjectID sets the project edge to Project by id.
func (tc *TemplateCreate) SetProjectID(id string) *TemplateCreate {
	tc.mutation.SetProjectID(id)
	return tc
}

// SetNillableProjectID sets the project edge to Project by id if the given value is not nil.
func (tc *TemplateCreate) SetNillableProjectID(id *string) *TemplateCreate {
	if id != nil {
		tc = tc.SetProjectID(*id)
	}
	return tc
}

// SetProject sets the project edge to Project.
func (tc *TemplateCreate) SetProject(p *Project) *TemplateCreate {
	return tc.SetProjectID(p.ID)
}

// SetCreatorID sets the creator edge to Admin by id.
func (tc *TemplateCreate) SetCreatorID(id string) *TemplateCreate {
	tc.mutation.SetCreatorID(id)
	return tc
}

// SetNillableCreatorID sets the creator edge to Admin by id if the given value is not nil.
func (tc *TemplateCreate) SetNillableCreatorID(id *string) *TemplateCreate {
	if id != nil {
		tc = tc.SetCreatorID(*id)
	}
	return tc
}

// SetCreator sets the creator edge to Admin.
func (tc *TemplateCreate) SetCreator(a *Admin) *TemplateCreate {
	return tc.SetCreatorID(a.ID)
}

// SetRunID sets the run edge to Run by id.
func (tc *TemplateCreate) SetRunID(id string) *TemplateCreate {
	tc.mutation.SetRunID(id)
	return tc
}

// SetNillableRunID sets the run edge to Run by id if the given value is not nil.
func (tc *TemplateCreate) SetNillableRunID(id *string) *TemplateCreate {
	if id != nil {
		tc = tc.SetRunID(*id)
	}
	return tc
}

// SetRun sets the run edge to Run.
func (tc *TemplateCreate) SetRun(r *Run) *TemplateCreate {
	return tc.SetRunID(r.ID)
}

// Mutation returns the TemplateMutation object of the builder.
func (tc *TemplateCreate) Mutation() *TemplateMutation {
	return tc.mutation
}

// Save creates the Template in the database.
func (tc *TemplateCreate) Save(ctx context.Context) (*Template, error) {
	var (
		err  error
		node *Template
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TemplateCreate) SaveX(ctx context.Context) *Template {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (tc *TemplateCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := template.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := template.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tc.mutation.ParticipantCount(); !ok {
		v := template.DefaultParticipantCount
		tc.mutation.SetParticipantCount(v)
	}
	if _, ok := tc.mutation.Adult(); !ok {
		v := template.DefaultAdult
		tc.mutation.SetAdult(v)
	}
	if _, ok := tc.mutation.Sandbox(); !ok {
		v := template.DefaultSandbox
		tc.mutation.SetSandbox(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TemplateCreate) check() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := tc.mutation.Name(); ok {
		if err := template.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := tc.mutation.SelectionType(); !ok {
		return &ValidationError{Name: "selectionType", err: errors.New("ent: missing required field \"selectionType\"")}
	}
	if v, ok := tc.mutation.SelectionType(); ok {
		if err := template.SelectionTypeValidator(v); err != nil {
			return &ValidationError{Name: "selectionType", err: fmt.Errorf("ent: validator failed for field \"selectionType\": %w", err)}
		}
	}
	if _, ok := tc.mutation.ParticipantCount(); !ok {
		return &ValidationError{Name: "participantCount", err: errors.New("ent: missing required field \"participantCount\"")}
	}
	if v, ok := tc.mutation.ParticipantCount(); ok {
		if err := template.ParticipantCountValidator(v); err != nil {
			return &ValidationError{Name: "participantCount", err: fmt.Errorf("ent: validator failed for field \"participantCount\": %w", err)}
		}
	}
	if _, ok := tc.mutation.InternalCriteria(); !ok {
		return &ValidationError{Name: "internalCriteria", err: errors.New("ent: missing required field \"internalCriteria\"")}
	}
	if _, ok := tc.mutation.MturkCriteria(); !ok {
		return &ValidationError{Name: "mturkCriteria", err: errors.New("ent: missing required field \"mturkCriteria\"")}
	}
	if _, ok := tc.mutation.Adult(); !ok {
		return &ValidationError{Name: "adult", err: errors.New("ent: missing required field \"adult\"")}
	}
	if _, ok := tc.mutation.Sandbox(); !ok {
		return &ValidationError{Name: "sandbox", err: errors.New("ent: missing required field \"sandbox\"")}
	}
	if v, ok := tc.mutation.ID(); ok {
		if err := template.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf("ent: validator failed for field \"id\": %w", err)}
		}
	}
	return nil
}

func (tc *TemplateCreate) sqlSave(ctx context.Context) (*Template, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (tc *TemplateCreate) createSpec() (*Template, *sqlgraph.CreateSpec) {
	var (
		_node = &Template{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: template.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: template.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: template.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: template.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: template.FieldName,
		})
		_node.Name = value
	}
	if value, ok := tc.mutation.SelectionType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: template.FieldSelectionType,
		})
		_node.SelectionType = value
	}
	if value, ok := tc.mutation.ParticipantCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: template.FieldParticipantCount,
		})
		_node.ParticipantCount = value
	}
	if value, ok := tc.mutation.InternalCriteria(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: template.FieldInternalCriteria,
		})
		_node.InternalCriteria = value
	}
	if value, ok := tc.mutation.MturkCriteria(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: template.FieldMturkCriteria,
		})
		_node.MturkCriteria = value
	}
	if value, ok := tc.mutation.Adult(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: template.FieldAdult,
		})
		_node.Adult = value
	}
	if value, ok := tc.mutation.Sandbox(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: template.FieldSandbox,
		})
		_node.Sandbox = value
	}
	if nodes := tc.mutation.StepsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   template.StepsTable,
			Columns: []string{template.StepsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   template.ProjectTable,
			Columns: []string{template.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   template.CreatorTable,
			Columns: []string{template.CreatorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: admin.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.RunIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   template.RunTable,
			Columns: []string{template.RunColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TemplateCreateBulk is the builder for creating a bulk of Template entities.
type TemplateCreateBulk struct {
	config
	builders []*TemplateCreate
}

// Save creates the Template entities in the database.
func (tcb *TemplateCreateBulk) Save(ctx context.Context) ([]*Template, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Template, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TemplateMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (tcb *TemplateCreateBulk) SaveX(ctx context.Context) []*Template {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
