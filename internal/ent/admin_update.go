// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/empiricaly/recruitment/internal/ent/admin"
	"github.com/empiricaly/recruitment/internal/ent/predicate"
	"github.com/empiricaly/recruitment/internal/ent/procedure"
	"github.com/empiricaly/recruitment/internal/ent/project"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// AdminUpdate is the builder for updating Admin entities.
type AdminUpdate struct {
	config
	hooks      []Hook
	mutation   *AdminMutation
	predicates []predicate.Admin
}

// Where adds a new predicate for the builder.
func (au *AdminUpdate) Where(ps ...predicate.Admin) *AdminUpdate {
	au.predicates = append(au.predicates, ps...)
	return au
}

// SetUpdatedAt sets the updatedAt field.
func (au *AdminUpdate) SetUpdatedAt(t time.Time) *AdminUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetName sets the name field.
func (au *AdminUpdate) SetName(s string) *AdminUpdate {
	au.mutation.SetName(s)
	return au
}

// SetUsername sets the username field.
func (au *AdminUpdate) SetUsername(s string) *AdminUpdate {
	au.mutation.SetUsername(s)
	return au
}

// AddProjectIDs adds the projects edge to Project by ids.
func (au *AdminUpdate) AddProjectIDs(ids ...string) *AdminUpdate {
	au.mutation.AddProjectIDs(ids...)
	return au
}

// AddProjects adds the projects edges to Project.
func (au *AdminUpdate) AddProjects(p ...*Project) *AdminUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return au.AddProjectIDs(ids...)
}

// AddProcedureIDs adds the procedures edge to Procedure by ids.
func (au *AdminUpdate) AddProcedureIDs(ids ...string) *AdminUpdate {
	au.mutation.AddProcedureIDs(ids...)
	return au
}

// AddProcedures adds the procedures edges to Procedure.
func (au *AdminUpdate) AddProcedures(p ...*Procedure) *AdminUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return au.AddProcedureIDs(ids...)
}

// Mutation returns the AdminMutation object of the builder.
func (au *AdminUpdate) Mutation() *AdminMutation {
	return au.mutation
}

// RemoveProjectIDs removes the projects edge to Project by ids.
func (au *AdminUpdate) RemoveProjectIDs(ids ...string) *AdminUpdate {
	au.mutation.RemoveProjectIDs(ids...)
	return au
}

// RemoveProjects removes projects edges to Project.
func (au *AdminUpdate) RemoveProjects(p ...*Project) *AdminUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return au.RemoveProjectIDs(ids...)
}

// RemoveProcedureIDs removes the procedures edge to Procedure by ids.
func (au *AdminUpdate) RemoveProcedureIDs(ids ...string) *AdminUpdate {
	au.mutation.RemoveProcedureIDs(ids...)
	return au
}

// RemoveProcedures removes procedures edges to Procedure.
func (au *AdminUpdate) RemoveProcedures(p ...*Procedure) *AdminUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return au.RemoveProcedureIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (au *AdminUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := admin.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}

	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdminMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AdminUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AdminUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AdminUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AdminUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   admin.Table,
			Columns: admin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: admin.FieldID,
			},
		},
	}
	if ps := au.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: admin.FieldUpdatedAt,
		})
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admin.FieldName,
		})
	}
	if value, ok := au.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admin.FieldUsername,
		})
	}
	if nodes := au.mutation.RemovedProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.ProjectsTable,
			Columns: []string{admin.ProjectsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.ProjectsTable,
			Columns: []string{admin.ProjectsColumn},
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
	if nodes := au.mutation.RemovedProceduresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.ProceduresTable,
			Columns: []string{admin.ProceduresColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ProceduresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.ProceduresTable,
			Columns: []string{admin.ProceduresColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{admin.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AdminUpdateOne is the builder for updating a single Admin entity.
type AdminUpdateOne struct {
	config
	hooks    []Hook
	mutation *AdminMutation
}

// SetUpdatedAt sets the updatedAt field.
func (auo *AdminUpdateOne) SetUpdatedAt(t time.Time) *AdminUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetName sets the name field.
func (auo *AdminUpdateOne) SetName(s string) *AdminUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetUsername sets the username field.
func (auo *AdminUpdateOne) SetUsername(s string) *AdminUpdateOne {
	auo.mutation.SetUsername(s)
	return auo
}

// AddProjectIDs adds the projects edge to Project by ids.
func (auo *AdminUpdateOne) AddProjectIDs(ids ...string) *AdminUpdateOne {
	auo.mutation.AddProjectIDs(ids...)
	return auo
}

// AddProjects adds the projects edges to Project.
func (auo *AdminUpdateOne) AddProjects(p ...*Project) *AdminUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return auo.AddProjectIDs(ids...)
}

// AddProcedureIDs adds the procedures edge to Procedure by ids.
func (auo *AdminUpdateOne) AddProcedureIDs(ids ...string) *AdminUpdateOne {
	auo.mutation.AddProcedureIDs(ids...)
	return auo
}

// AddProcedures adds the procedures edges to Procedure.
func (auo *AdminUpdateOne) AddProcedures(p ...*Procedure) *AdminUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return auo.AddProcedureIDs(ids...)
}

// Mutation returns the AdminMutation object of the builder.
func (auo *AdminUpdateOne) Mutation() *AdminMutation {
	return auo.mutation
}

// RemoveProjectIDs removes the projects edge to Project by ids.
func (auo *AdminUpdateOne) RemoveProjectIDs(ids ...string) *AdminUpdateOne {
	auo.mutation.RemoveProjectIDs(ids...)
	return auo
}

// RemoveProjects removes projects edges to Project.
func (auo *AdminUpdateOne) RemoveProjects(p ...*Project) *AdminUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return auo.RemoveProjectIDs(ids...)
}

// RemoveProcedureIDs removes the procedures edge to Procedure by ids.
func (auo *AdminUpdateOne) RemoveProcedureIDs(ids ...string) *AdminUpdateOne {
	auo.mutation.RemoveProcedureIDs(ids...)
	return auo
}

// RemoveProcedures removes procedures edges to Procedure.
func (auo *AdminUpdateOne) RemoveProcedures(p ...*Procedure) *AdminUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return auo.RemoveProcedureIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (auo *AdminUpdateOne) Save(ctx context.Context) (*Admin, error) {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := admin.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}

	var (
		err  error
		node *Admin
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdminMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AdminUpdateOne) SaveX(ctx context.Context) *Admin {
	a, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return a
}

// Exec executes the query on the entity.
func (auo *AdminUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AdminUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AdminUpdateOne) sqlSave(ctx context.Context) (a *Admin, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   admin.Table,
			Columns: admin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: admin.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Admin.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: admin.FieldUpdatedAt,
		})
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admin.FieldName,
		})
	}
	if value, ok := auo.mutation.Username(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: admin.FieldUsername,
		})
	}
	if nodes := auo.mutation.RemovedProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.ProjectsTable,
			Columns: []string{admin.ProjectsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.ProjectsTable,
			Columns: []string{admin.ProjectsColumn},
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
	if nodes := auo.mutation.RemovedProceduresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.ProceduresTable,
			Columns: []string{admin.ProceduresColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ProceduresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   admin.ProceduresTable,
			Columns: []string{admin.ProceduresColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	a = &Admin{config: auo.config}
	_spec.Assign = a.assignValues
	_spec.ScanValues = a.scanValues()
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{admin.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return a, nil
}
