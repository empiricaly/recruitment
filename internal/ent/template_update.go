// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/empiricaly/recruitment/internal/ent/admin"
	"github.com/empiricaly/recruitment/internal/ent/predicate"
	"github.com/empiricaly/recruitment/internal/ent/project"
	"github.com/empiricaly/recruitment/internal/ent/run"
	"github.com/empiricaly/recruitment/internal/ent/step"
	"github.com/empiricaly/recruitment/internal/ent/template"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// TemplateUpdate is the builder for updating Template entities.
type TemplateUpdate struct {
	config
	hooks      []Hook
	mutation   *TemplateMutation
	predicates []predicate.Template
}

// Where adds a new predicate for the builder.
func (tu *TemplateUpdate) Where(ps ...predicate.Template) *TemplateUpdate {
	tu.predicates = append(tu.predicates, ps...)
	return tu
}

// SetUpdatedAt sets the updated_at field.
func (tu *TemplateUpdate) SetUpdatedAt(t time.Time) *TemplateUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetName sets the name field.
func (tu *TemplateUpdate) SetName(s string) *TemplateUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetSelectionType sets the selectionType field.
func (tu *TemplateUpdate) SetSelectionType(tt template.SelectionType) *TemplateUpdate {
	tu.mutation.SetSelectionType(tt)
	return tu
}

// SetParticipantCount sets the participantCount field.
func (tu *TemplateUpdate) SetParticipantCount(i int) *TemplateUpdate {
	tu.mutation.ResetParticipantCount()
	tu.mutation.SetParticipantCount(i)
	return tu
}

// SetNillableParticipantCount sets the participantCount field if the given value is not nil.
func (tu *TemplateUpdate) SetNillableParticipantCount(i *int) *TemplateUpdate {
	if i != nil {
		tu.SetParticipantCount(*i)
	}
	return tu
}

// AddParticipantCount adds i to participantCount.
func (tu *TemplateUpdate) AddParticipantCount(i int) *TemplateUpdate {
	tu.mutation.AddParticipantCount(i)
	return tu
}

// SetInternalCriteria sets the internalCriteria field.
func (tu *TemplateUpdate) SetInternalCriteria(b []byte) *TemplateUpdate {
	tu.mutation.SetInternalCriteria(b)
	return tu
}

// SetMturkCriteria sets the mturkCriteria field.
func (tu *TemplateUpdate) SetMturkCriteria(b []byte) *TemplateUpdate {
	tu.mutation.SetMturkCriteria(b)
	return tu
}

// SetAdult sets the adult field.
func (tu *TemplateUpdate) SetAdult(b bool) *TemplateUpdate {
	tu.mutation.SetAdult(b)
	return tu
}

// SetNillableAdult sets the adult field if the given value is not nil.
func (tu *TemplateUpdate) SetNillableAdult(b *bool) *TemplateUpdate {
	if b != nil {
		tu.SetAdult(*b)
	}
	return tu
}

// SetSandbox sets the sandbox field.
func (tu *TemplateUpdate) SetSandbox(b bool) *TemplateUpdate {
	tu.mutation.SetSandbox(b)
	return tu
}

// SetNillableSandbox sets the sandbox field if the given value is not nil.
func (tu *TemplateUpdate) SetNillableSandbox(b *bool) *TemplateUpdate {
	if b != nil {
		tu.SetSandbox(*b)
	}
	return tu
}

// AddStepIDs adds the steps edge to Step by ids.
func (tu *TemplateUpdate) AddStepIDs(ids ...string) *TemplateUpdate {
	tu.mutation.AddStepIDs(ids...)
	return tu
}

// AddSteps adds the steps edges to Step.
func (tu *TemplateUpdate) AddSteps(s ...*Step) *TemplateUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tu.AddStepIDs(ids...)
}

// SetProjectID sets the project edge to Project by id.
func (tu *TemplateUpdate) SetProjectID(id string) *TemplateUpdate {
	tu.mutation.SetProjectID(id)
	return tu
}

// SetNillableProjectID sets the project edge to Project by id if the given value is not nil.
func (tu *TemplateUpdate) SetNillableProjectID(id *string) *TemplateUpdate {
	if id != nil {
		tu = tu.SetProjectID(*id)
	}
	return tu
}

// SetProject sets the project edge to Project.
func (tu *TemplateUpdate) SetProject(p *Project) *TemplateUpdate {
	return tu.SetProjectID(p.ID)
}

// SetCreatorID sets the creator edge to Admin by id.
func (tu *TemplateUpdate) SetCreatorID(id string) *TemplateUpdate {
	tu.mutation.SetCreatorID(id)
	return tu
}

// SetNillableCreatorID sets the creator edge to Admin by id if the given value is not nil.
func (tu *TemplateUpdate) SetNillableCreatorID(id *string) *TemplateUpdate {
	if id != nil {
		tu = tu.SetCreatorID(*id)
	}
	return tu
}

// SetCreator sets the creator edge to Admin.
func (tu *TemplateUpdate) SetCreator(a *Admin) *TemplateUpdate {
	return tu.SetCreatorID(a.ID)
}

// SetRunID sets the run edge to Run by id.
func (tu *TemplateUpdate) SetRunID(id string) *TemplateUpdate {
	tu.mutation.SetRunID(id)
	return tu
}

// SetNillableRunID sets the run edge to Run by id if the given value is not nil.
func (tu *TemplateUpdate) SetNillableRunID(id *string) *TemplateUpdate {
	if id != nil {
		tu = tu.SetRunID(*id)
	}
	return tu
}

// SetRun sets the run edge to Run.
func (tu *TemplateUpdate) SetRun(r *Run) *TemplateUpdate {
	return tu.SetRunID(r.ID)
}

// Mutation returns the TemplateMutation object of the builder.
func (tu *TemplateUpdate) Mutation() *TemplateMutation {
	return tu.mutation
}

// ClearSteps clears all "steps" edges to type Step.
func (tu *TemplateUpdate) ClearSteps() *TemplateUpdate {
	tu.mutation.ClearSteps()
	return tu
}

// RemoveStepIDs removes the steps edge to Step by ids.
func (tu *TemplateUpdate) RemoveStepIDs(ids ...string) *TemplateUpdate {
	tu.mutation.RemoveStepIDs(ids...)
	return tu
}

// RemoveSteps removes steps edges to Step.
func (tu *TemplateUpdate) RemoveSteps(s ...*Step) *TemplateUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tu.RemoveStepIDs(ids...)
}

// ClearProject clears the "project" edge to type Project.
func (tu *TemplateUpdate) ClearProject() *TemplateUpdate {
	tu.mutation.ClearProject()
	return tu
}

// ClearCreator clears the "creator" edge to type Admin.
func (tu *TemplateUpdate) ClearCreator() *TemplateUpdate {
	tu.mutation.ClearCreator()
	return tu
}

// ClearRun clears the "run" edge to type Run.
func (tu *TemplateUpdate) ClearRun() *TemplateUpdate {
	tu.mutation.ClearRun()
	return tu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (tu *TemplateUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tu.defaults()
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TemplateUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TemplateUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TemplateUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TemplateUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := template.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TemplateUpdate) check() error {
	if v, ok := tu.mutation.Name(); ok {
		if err := template.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := tu.mutation.SelectionType(); ok {
		if err := template.SelectionTypeValidator(v); err != nil {
			return &ValidationError{Name: "selectionType", err: fmt.Errorf("ent: validator failed for field \"selectionType\": %w", err)}
		}
	}
	if v, ok := tu.mutation.ParticipantCount(); ok {
		if err := template.ParticipantCountValidator(v); err != nil {
			return &ValidationError{Name: "participantCount", err: fmt.Errorf("ent: validator failed for field \"participantCount\": %w", err)}
		}
	}
	return nil
}

func (tu *TemplateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   template.Table,
			Columns: template.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: template.FieldID,
			},
		},
	}
	if ps := tu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: template.FieldUpdatedAt,
		})
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: template.FieldName,
		})
	}
	if value, ok := tu.mutation.SelectionType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: template.FieldSelectionType,
		})
	}
	if value, ok := tu.mutation.ParticipantCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: template.FieldParticipantCount,
		})
	}
	if value, ok := tu.mutation.AddedParticipantCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: template.FieldParticipantCount,
		})
	}
	if value, ok := tu.mutation.InternalCriteria(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: template.FieldInternalCriteria,
		})
	}
	if value, ok := tu.mutation.MturkCriteria(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: template.FieldMturkCriteria,
		})
	}
	if value, ok := tu.mutation.Adult(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: template.FieldAdult,
		})
	}
	if value, ok := tu.mutation.Sandbox(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: template.FieldSandbox,
		})
	}
	if tu.mutation.StepsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedStepsIDs(); len(nodes) > 0 && !tu.mutation.StepsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.StepsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ProjectCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ProjectIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.CreatorCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.CreatorIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.RunCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RunIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{template.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TemplateUpdateOne is the builder for updating a single Template entity.
type TemplateUpdateOne struct {
	config
	hooks    []Hook
	mutation *TemplateMutation
}

// SetUpdatedAt sets the updated_at field.
func (tuo *TemplateUpdateOne) SetUpdatedAt(t time.Time) *TemplateUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetName sets the name field.
func (tuo *TemplateUpdateOne) SetName(s string) *TemplateUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetSelectionType sets the selectionType field.
func (tuo *TemplateUpdateOne) SetSelectionType(tt template.SelectionType) *TemplateUpdateOne {
	tuo.mutation.SetSelectionType(tt)
	return tuo
}

// SetParticipantCount sets the participantCount field.
func (tuo *TemplateUpdateOne) SetParticipantCount(i int) *TemplateUpdateOne {
	tuo.mutation.ResetParticipantCount()
	tuo.mutation.SetParticipantCount(i)
	return tuo
}

// SetNillableParticipantCount sets the participantCount field if the given value is not nil.
func (tuo *TemplateUpdateOne) SetNillableParticipantCount(i *int) *TemplateUpdateOne {
	if i != nil {
		tuo.SetParticipantCount(*i)
	}
	return tuo
}

// AddParticipantCount adds i to participantCount.
func (tuo *TemplateUpdateOne) AddParticipantCount(i int) *TemplateUpdateOne {
	tuo.mutation.AddParticipantCount(i)
	return tuo
}

// SetInternalCriteria sets the internalCriteria field.
func (tuo *TemplateUpdateOne) SetInternalCriteria(b []byte) *TemplateUpdateOne {
	tuo.mutation.SetInternalCriteria(b)
	return tuo
}

// SetMturkCriteria sets the mturkCriteria field.
func (tuo *TemplateUpdateOne) SetMturkCriteria(b []byte) *TemplateUpdateOne {
	tuo.mutation.SetMturkCriteria(b)
	return tuo
}

// SetAdult sets the adult field.
func (tuo *TemplateUpdateOne) SetAdult(b bool) *TemplateUpdateOne {
	tuo.mutation.SetAdult(b)
	return tuo
}

// SetNillableAdult sets the adult field if the given value is not nil.
func (tuo *TemplateUpdateOne) SetNillableAdult(b *bool) *TemplateUpdateOne {
	if b != nil {
		tuo.SetAdult(*b)
	}
	return tuo
}

// SetSandbox sets the sandbox field.
func (tuo *TemplateUpdateOne) SetSandbox(b bool) *TemplateUpdateOne {
	tuo.mutation.SetSandbox(b)
	return tuo
}

// SetNillableSandbox sets the sandbox field if the given value is not nil.
func (tuo *TemplateUpdateOne) SetNillableSandbox(b *bool) *TemplateUpdateOne {
	if b != nil {
		tuo.SetSandbox(*b)
	}
	return tuo
}

// AddStepIDs adds the steps edge to Step by ids.
func (tuo *TemplateUpdateOne) AddStepIDs(ids ...string) *TemplateUpdateOne {
	tuo.mutation.AddStepIDs(ids...)
	return tuo
}

// AddSteps adds the steps edges to Step.
func (tuo *TemplateUpdateOne) AddSteps(s ...*Step) *TemplateUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tuo.AddStepIDs(ids...)
}

// SetProjectID sets the project edge to Project by id.
func (tuo *TemplateUpdateOne) SetProjectID(id string) *TemplateUpdateOne {
	tuo.mutation.SetProjectID(id)
	return tuo
}

// SetNillableProjectID sets the project edge to Project by id if the given value is not nil.
func (tuo *TemplateUpdateOne) SetNillableProjectID(id *string) *TemplateUpdateOne {
	if id != nil {
		tuo = tuo.SetProjectID(*id)
	}
	return tuo
}

// SetProject sets the project edge to Project.
func (tuo *TemplateUpdateOne) SetProject(p *Project) *TemplateUpdateOne {
	return tuo.SetProjectID(p.ID)
}

// SetCreatorID sets the creator edge to Admin by id.
func (tuo *TemplateUpdateOne) SetCreatorID(id string) *TemplateUpdateOne {
	tuo.mutation.SetCreatorID(id)
	return tuo
}

// SetNillableCreatorID sets the creator edge to Admin by id if the given value is not nil.
func (tuo *TemplateUpdateOne) SetNillableCreatorID(id *string) *TemplateUpdateOne {
	if id != nil {
		tuo = tuo.SetCreatorID(*id)
	}
	return tuo
}

// SetCreator sets the creator edge to Admin.
func (tuo *TemplateUpdateOne) SetCreator(a *Admin) *TemplateUpdateOne {
	return tuo.SetCreatorID(a.ID)
}

// SetRunID sets the run edge to Run by id.
func (tuo *TemplateUpdateOne) SetRunID(id string) *TemplateUpdateOne {
	tuo.mutation.SetRunID(id)
	return tuo
}

// SetNillableRunID sets the run edge to Run by id if the given value is not nil.
func (tuo *TemplateUpdateOne) SetNillableRunID(id *string) *TemplateUpdateOne {
	if id != nil {
		tuo = tuo.SetRunID(*id)
	}
	return tuo
}

// SetRun sets the run edge to Run.
func (tuo *TemplateUpdateOne) SetRun(r *Run) *TemplateUpdateOne {
	return tuo.SetRunID(r.ID)
}

// Mutation returns the TemplateMutation object of the builder.
func (tuo *TemplateUpdateOne) Mutation() *TemplateMutation {
	return tuo.mutation
}

// ClearSteps clears all "steps" edges to type Step.
func (tuo *TemplateUpdateOne) ClearSteps() *TemplateUpdateOne {
	tuo.mutation.ClearSteps()
	return tuo
}

// RemoveStepIDs removes the steps edge to Step by ids.
func (tuo *TemplateUpdateOne) RemoveStepIDs(ids ...string) *TemplateUpdateOne {
	tuo.mutation.RemoveStepIDs(ids...)
	return tuo
}

// RemoveSteps removes steps edges to Step.
func (tuo *TemplateUpdateOne) RemoveSteps(s ...*Step) *TemplateUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tuo.RemoveStepIDs(ids...)
}

// ClearProject clears the "project" edge to type Project.
func (tuo *TemplateUpdateOne) ClearProject() *TemplateUpdateOne {
	tuo.mutation.ClearProject()
	return tuo
}

// ClearCreator clears the "creator" edge to type Admin.
func (tuo *TemplateUpdateOne) ClearCreator() *TemplateUpdateOne {
	tuo.mutation.ClearCreator()
	return tuo
}

// ClearRun clears the "run" edge to type Run.
func (tuo *TemplateUpdateOne) ClearRun() *TemplateUpdateOne {
	tuo.mutation.ClearRun()
	return tuo
}

// Save executes the query and returns the updated entity.
func (tuo *TemplateUpdateOne) Save(ctx context.Context) (*Template, error) {
	var (
		err  error
		node *Template
	)
	tuo.defaults()
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TemplateUpdateOne) SaveX(ctx context.Context) *Template {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TemplateUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TemplateUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TemplateUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := template.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TemplateUpdateOne) check() error {
	if v, ok := tuo.mutation.Name(); ok {
		if err := template.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.SelectionType(); ok {
		if err := template.SelectionTypeValidator(v); err != nil {
			return &ValidationError{Name: "selectionType", err: fmt.Errorf("ent: validator failed for field \"selectionType\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.ParticipantCount(); ok {
		if err := template.ParticipantCountValidator(v); err != nil {
			return &ValidationError{Name: "participantCount", err: fmt.Errorf("ent: validator failed for field \"participantCount\": %w", err)}
		}
	}
	return nil
}

func (tuo *TemplateUpdateOne) sqlSave(ctx context.Context) (_node *Template, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   template.Table,
			Columns: template.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: template.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Template.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: template.FieldUpdatedAt,
		})
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: template.FieldName,
		})
	}
	if value, ok := tuo.mutation.SelectionType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: template.FieldSelectionType,
		})
	}
	if value, ok := tuo.mutation.ParticipantCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: template.FieldParticipantCount,
		})
	}
	if value, ok := tuo.mutation.AddedParticipantCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: template.FieldParticipantCount,
		})
	}
	if value, ok := tuo.mutation.InternalCriteria(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: template.FieldInternalCriteria,
		})
	}
	if value, ok := tuo.mutation.MturkCriteria(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: template.FieldMturkCriteria,
		})
	}
	if value, ok := tuo.mutation.Adult(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: template.FieldAdult,
		})
	}
	if value, ok := tuo.mutation.Sandbox(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: template.FieldSandbox,
		})
	}
	if tuo.mutation.StepsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedStepsIDs(); len(nodes) > 0 && !tuo.mutation.StepsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.StepsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ProjectCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ProjectIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.CreatorCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.CreatorIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.RunCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RunIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Template{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{template.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
