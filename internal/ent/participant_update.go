// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/empiricaly/recruitment/internal/ent/datum"
	"github.com/empiricaly/recruitment/internal/ent/participant"
	"github.com/empiricaly/recruitment/internal/ent/participation"
	"github.com/empiricaly/recruitment/internal/ent/predicate"
	"github.com/empiricaly/recruitment/internal/ent/project"
	"github.com/empiricaly/recruitment/internal/ent/providerid"
	"github.com/empiricaly/recruitment/internal/ent/steprun"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// ParticipantUpdate is the builder for updating Participant entities.
type ParticipantUpdate struct {
	config
	hooks      []Hook
	mutation   *ParticipantMutation
	predicates []predicate.Participant
}

// Where adds a new predicate for the builder.
func (pu *ParticipantUpdate) Where(ps ...predicate.Participant) *ParticipantUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetUpdatedAt sets the updated_at field.
func (pu *ParticipantUpdate) SetUpdatedAt(t time.Time) *ParticipantUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetMturkWorkerID sets the mturkWorkerID field.
func (pu *ParticipantUpdate) SetMturkWorkerID(s string) *ParticipantUpdate {
	pu.mutation.SetMturkWorkerID(s)
	return pu
}

// SetNillableMturkWorkerID sets the mturkWorkerID field if the given value is not nil.
func (pu *ParticipantUpdate) SetNillableMturkWorkerID(s *string) *ParticipantUpdate {
	if s != nil {
		pu.SetMturkWorkerID(*s)
	}
	return pu
}

// ClearMturkWorkerID clears the value of mturkWorkerID.
func (pu *ParticipantUpdate) ClearMturkWorkerID() *ParticipantUpdate {
	pu.mutation.ClearMturkWorkerID()
	return pu
}

// AddDatumIDs adds the data edge to Datum by ids.
func (pu *ParticipantUpdate) AddDatumIDs(ids ...string) *ParticipantUpdate {
	pu.mutation.AddDatumIDs(ids...)
	return pu
}

// AddData adds the data edges to Datum.
func (pu *ParticipantUpdate) AddData(d ...*Datum) *ParticipantUpdate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pu.AddDatumIDs(ids...)
}

// AddProviderIDIDs adds the providerIDs edge to ProviderID by ids.
func (pu *ParticipantUpdate) AddProviderIDIDs(ids ...string) *ParticipantUpdate {
	pu.mutation.AddProviderIDIDs(ids...)
	return pu
}

// AddProviderIDs adds the providerIDs edges to ProviderID.
func (pu *ParticipantUpdate) AddProviderIDs(p ...*ProviderID) *ParticipantUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddProviderIDIDs(ids...)
}

// AddParticipationIDs adds the participations edge to Participation by ids.
func (pu *ParticipantUpdate) AddParticipationIDs(ids ...string) *ParticipantUpdate {
	pu.mutation.AddParticipationIDs(ids...)
	return pu
}

// AddParticipations adds the participations edges to Participation.
func (pu *ParticipantUpdate) AddParticipations(p ...*Participation) *ParticipantUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddParticipationIDs(ids...)
}

// SetCreatedByID sets the createdBy edge to StepRun by id.
func (pu *ParticipantUpdate) SetCreatedByID(id string) *ParticipantUpdate {
	pu.mutation.SetCreatedByID(id)
	return pu
}

// SetNillableCreatedByID sets the createdBy edge to StepRun by id if the given value is not nil.
func (pu *ParticipantUpdate) SetNillableCreatedByID(id *string) *ParticipantUpdate {
	if id != nil {
		pu = pu.SetCreatedByID(*id)
	}
	return pu
}

// SetCreatedBy sets the createdBy edge to StepRun.
func (pu *ParticipantUpdate) SetCreatedBy(s *StepRun) *ParticipantUpdate {
	return pu.SetCreatedByID(s.ID)
}

// AddStepIDs adds the steps edge to StepRun by ids.
func (pu *ParticipantUpdate) AddStepIDs(ids ...string) *ParticipantUpdate {
	pu.mutation.AddStepIDs(ids...)
	return pu
}

// AddSteps adds the steps edges to StepRun.
func (pu *ParticipantUpdate) AddSteps(s ...*StepRun) *ParticipantUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.AddStepIDs(ids...)
}

// AddProjectIDs adds the projects edge to Project by ids.
func (pu *ParticipantUpdate) AddProjectIDs(ids ...string) *ParticipantUpdate {
	pu.mutation.AddProjectIDs(ids...)
	return pu
}

// AddProjects adds the projects edges to Project.
func (pu *ParticipantUpdate) AddProjects(p ...*Project) *ParticipantUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddProjectIDs(ids...)
}

// Mutation returns the ParticipantMutation object of the builder.
func (pu *ParticipantUpdate) Mutation() *ParticipantMutation {
	return pu.mutation
}

// ClearData clears all "data" edges to type Datum.
func (pu *ParticipantUpdate) ClearData() *ParticipantUpdate {
	pu.mutation.ClearData()
	return pu
}

// RemoveDatumIDs removes the data edge to Datum by ids.
func (pu *ParticipantUpdate) RemoveDatumIDs(ids ...string) *ParticipantUpdate {
	pu.mutation.RemoveDatumIDs(ids...)
	return pu
}

// RemoveData removes data edges to Datum.
func (pu *ParticipantUpdate) RemoveData(d ...*Datum) *ParticipantUpdate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pu.RemoveDatumIDs(ids...)
}

// ClearProviderIDs clears all "providerIDs" edges to type ProviderID.
func (pu *ParticipantUpdate) ClearProviderIDs() *ParticipantUpdate {
	pu.mutation.ClearProviderIDs()
	return pu
}

// RemoveProviderIDIDs removes the providerIDs edge to ProviderID by ids.
func (pu *ParticipantUpdate) RemoveProviderIDIDs(ids ...string) *ParticipantUpdate {
	pu.mutation.RemoveProviderIDIDs(ids...)
	return pu
}

// RemoveProviderIDs removes providerIDs edges to ProviderID.
func (pu *ParticipantUpdate) RemoveProviderIDs(p ...*ProviderID) *ParticipantUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveProviderIDIDs(ids...)
}

// ClearParticipations clears all "participations" edges to type Participation.
func (pu *ParticipantUpdate) ClearParticipations() *ParticipantUpdate {
	pu.mutation.ClearParticipations()
	return pu
}

// RemoveParticipationIDs removes the participations edge to Participation by ids.
func (pu *ParticipantUpdate) RemoveParticipationIDs(ids ...string) *ParticipantUpdate {
	pu.mutation.RemoveParticipationIDs(ids...)
	return pu
}

// RemoveParticipations removes participations edges to Participation.
func (pu *ParticipantUpdate) RemoveParticipations(p ...*Participation) *ParticipantUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveParticipationIDs(ids...)
}

// ClearCreatedBy clears the "createdBy" edge to type StepRun.
func (pu *ParticipantUpdate) ClearCreatedBy() *ParticipantUpdate {
	pu.mutation.ClearCreatedBy()
	return pu
}

// ClearSteps clears all "steps" edges to type StepRun.
func (pu *ParticipantUpdate) ClearSteps() *ParticipantUpdate {
	pu.mutation.ClearSteps()
	return pu
}

// RemoveStepIDs removes the steps edge to StepRun by ids.
func (pu *ParticipantUpdate) RemoveStepIDs(ids ...string) *ParticipantUpdate {
	pu.mutation.RemoveStepIDs(ids...)
	return pu
}

// RemoveSteps removes steps edges to StepRun.
func (pu *ParticipantUpdate) RemoveSteps(s ...*StepRun) *ParticipantUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.RemoveStepIDs(ids...)
}

// ClearProjects clears all "projects" edges to type Project.
func (pu *ParticipantUpdate) ClearProjects() *ParticipantUpdate {
	pu.mutation.ClearProjects()
	return pu
}

// RemoveProjectIDs removes the projects edge to Project by ids.
func (pu *ParticipantUpdate) RemoveProjectIDs(ids ...string) *ParticipantUpdate {
	pu.mutation.RemoveProjectIDs(ids...)
	return pu
}

// RemoveProjects removes projects edges to Project.
func (pu *ParticipantUpdate) RemoveProjects(p ...*Project) *ParticipantUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveProjectIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *ParticipantUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pu.defaults()
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ParticipantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ParticipantUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ParticipantUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ParticipantUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ParticipantUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := participant.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

func (pu *ParticipantUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   participant.Table,
			Columns: participant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: participant.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: participant.FieldUpdatedAt,
		})
	}
	if value, ok := pu.mutation.MturkWorkerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: participant.FieldMturkWorkerID,
		})
	}
	if pu.mutation.MturkWorkerIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: participant.FieldMturkWorkerID,
		})
	}
	if pu.mutation.DataCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedDataIDs(); len(nodes) > 0 && !pu.mutation.DataCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.DataIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ProviderIDsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedProviderIDsIDs(); len(nodes) > 0 && !pu.mutation.ProviderIDsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProviderIDsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ParticipationsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedParticipationsIDs(); len(nodes) > 0 && !pu.mutation.ParticipationsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ParticipationsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.CreatedByCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CreatedByIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.StepsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedStepsIDs(); len(nodes) > 0 && !pu.mutation.StepsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.StepsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   participant.ProjectsTable,
			Columns: participant.ProjectsPrimaryKey,
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
	if nodes := pu.mutation.RemovedProjectsIDs(); len(nodes) > 0 && !pu.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   participant.ProjectsTable,
			Columns: participant.ProjectsPrimaryKey,
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
	if nodes := pu.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   participant.ProjectsTable,
			Columns: participant.ProjectsPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{participant.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ParticipantUpdateOne is the builder for updating a single Participant entity.
type ParticipantUpdateOne struct {
	config
	hooks    []Hook
	mutation *ParticipantMutation
}

// SetUpdatedAt sets the updated_at field.
func (puo *ParticipantUpdateOne) SetUpdatedAt(t time.Time) *ParticipantUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetMturkWorkerID sets the mturkWorkerID field.
func (puo *ParticipantUpdateOne) SetMturkWorkerID(s string) *ParticipantUpdateOne {
	puo.mutation.SetMturkWorkerID(s)
	return puo
}

// SetNillableMturkWorkerID sets the mturkWorkerID field if the given value is not nil.
func (puo *ParticipantUpdateOne) SetNillableMturkWorkerID(s *string) *ParticipantUpdateOne {
	if s != nil {
		puo.SetMturkWorkerID(*s)
	}
	return puo
}

// ClearMturkWorkerID clears the value of mturkWorkerID.
func (puo *ParticipantUpdateOne) ClearMturkWorkerID() *ParticipantUpdateOne {
	puo.mutation.ClearMturkWorkerID()
	return puo
}

// AddDatumIDs adds the data edge to Datum by ids.
func (puo *ParticipantUpdateOne) AddDatumIDs(ids ...string) *ParticipantUpdateOne {
	puo.mutation.AddDatumIDs(ids...)
	return puo
}

// AddData adds the data edges to Datum.
func (puo *ParticipantUpdateOne) AddData(d ...*Datum) *ParticipantUpdateOne {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return puo.AddDatumIDs(ids...)
}

// AddProviderIDIDs adds the providerIDs edge to ProviderID by ids.
func (puo *ParticipantUpdateOne) AddProviderIDIDs(ids ...string) *ParticipantUpdateOne {
	puo.mutation.AddProviderIDIDs(ids...)
	return puo
}

// AddProviderIDs adds the providerIDs edges to ProviderID.
func (puo *ParticipantUpdateOne) AddProviderIDs(p ...*ProviderID) *ParticipantUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddProviderIDIDs(ids...)
}

// AddParticipationIDs adds the participations edge to Participation by ids.
func (puo *ParticipantUpdateOne) AddParticipationIDs(ids ...string) *ParticipantUpdateOne {
	puo.mutation.AddParticipationIDs(ids...)
	return puo
}

// AddParticipations adds the participations edges to Participation.
func (puo *ParticipantUpdateOne) AddParticipations(p ...*Participation) *ParticipantUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddParticipationIDs(ids...)
}

// SetCreatedByID sets the createdBy edge to StepRun by id.
func (puo *ParticipantUpdateOne) SetCreatedByID(id string) *ParticipantUpdateOne {
	puo.mutation.SetCreatedByID(id)
	return puo
}

// SetNillableCreatedByID sets the createdBy edge to StepRun by id if the given value is not nil.
func (puo *ParticipantUpdateOne) SetNillableCreatedByID(id *string) *ParticipantUpdateOne {
	if id != nil {
		puo = puo.SetCreatedByID(*id)
	}
	return puo
}

// SetCreatedBy sets the createdBy edge to StepRun.
func (puo *ParticipantUpdateOne) SetCreatedBy(s *StepRun) *ParticipantUpdateOne {
	return puo.SetCreatedByID(s.ID)
}

// AddStepIDs adds the steps edge to StepRun by ids.
func (puo *ParticipantUpdateOne) AddStepIDs(ids ...string) *ParticipantUpdateOne {
	puo.mutation.AddStepIDs(ids...)
	return puo
}

// AddSteps adds the steps edges to StepRun.
func (puo *ParticipantUpdateOne) AddSteps(s ...*StepRun) *ParticipantUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.AddStepIDs(ids...)
}

// AddProjectIDs adds the projects edge to Project by ids.
func (puo *ParticipantUpdateOne) AddProjectIDs(ids ...string) *ParticipantUpdateOne {
	puo.mutation.AddProjectIDs(ids...)
	return puo
}

// AddProjects adds the projects edges to Project.
func (puo *ParticipantUpdateOne) AddProjects(p ...*Project) *ParticipantUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddProjectIDs(ids...)
}

// Mutation returns the ParticipantMutation object of the builder.
func (puo *ParticipantUpdateOne) Mutation() *ParticipantMutation {
	return puo.mutation
}

// ClearData clears all "data" edges to type Datum.
func (puo *ParticipantUpdateOne) ClearData() *ParticipantUpdateOne {
	puo.mutation.ClearData()
	return puo
}

// RemoveDatumIDs removes the data edge to Datum by ids.
func (puo *ParticipantUpdateOne) RemoveDatumIDs(ids ...string) *ParticipantUpdateOne {
	puo.mutation.RemoveDatumIDs(ids...)
	return puo
}

// RemoveData removes data edges to Datum.
func (puo *ParticipantUpdateOne) RemoveData(d ...*Datum) *ParticipantUpdateOne {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return puo.RemoveDatumIDs(ids...)
}

// ClearProviderIDs clears all "providerIDs" edges to type ProviderID.
func (puo *ParticipantUpdateOne) ClearProviderIDs() *ParticipantUpdateOne {
	puo.mutation.ClearProviderIDs()
	return puo
}

// RemoveProviderIDIDs removes the providerIDs edge to ProviderID by ids.
func (puo *ParticipantUpdateOne) RemoveProviderIDIDs(ids ...string) *ParticipantUpdateOne {
	puo.mutation.RemoveProviderIDIDs(ids...)
	return puo
}

// RemoveProviderIDs removes providerIDs edges to ProviderID.
func (puo *ParticipantUpdateOne) RemoveProviderIDs(p ...*ProviderID) *ParticipantUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveProviderIDIDs(ids...)
}

// ClearParticipations clears all "participations" edges to type Participation.
func (puo *ParticipantUpdateOne) ClearParticipations() *ParticipantUpdateOne {
	puo.mutation.ClearParticipations()
	return puo
}

// RemoveParticipationIDs removes the participations edge to Participation by ids.
func (puo *ParticipantUpdateOne) RemoveParticipationIDs(ids ...string) *ParticipantUpdateOne {
	puo.mutation.RemoveParticipationIDs(ids...)
	return puo
}

// RemoveParticipations removes participations edges to Participation.
func (puo *ParticipantUpdateOne) RemoveParticipations(p ...*Participation) *ParticipantUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveParticipationIDs(ids...)
}

// ClearCreatedBy clears the "createdBy" edge to type StepRun.
func (puo *ParticipantUpdateOne) ClearCreatedBy() *ParticipantUpdateOne {
	puo.mutation.ClearCreatedBy()
	return puo
}

// ClearSteps clears all "steps" edges to type StepRun.
func (puo *ParticipantUpdateOne) ClearSteps() *ParticipantUpdateOne {
	puo.mutation.ClearSteps()
	return puo
}

// RemoveStepIDs removes the steps edge to StepRun by ids.
func (puo *ParticipantUpdateOne) RemoveStepIDs(ids ...string) *ParticipantUpdateOne {
	puo.mutation.RemoveStepIDs(ids...)
	return puo
}

// RemoveSteps removes steps edges to StepRun.
func (puo *ParticipantUpdateOne) RemoveSteps(s ...*StepRun) *ParticipantUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.RemoveStepIDs(ids...)
}

// ClearProjects clears all "projects" edges to type Project.
func (puo *ParticipantUpdateOne) ClearProjects() *ParticipantUpdateOne {
	puo.mutation.ClearProjects()
	return puo
}

// RemoveProjectIDs removes the projects edge to Project by ids.
func (puo *ParticipantUpdateOne) RemoveProjectIDs(ids ...string) *ParticipantUpdateOne {
	puo.mutation.RemoveProjectIDs(ids...)
	return puo
}

// RemoveProjects removes projects edges to Project.
func (puo *ParticipantUpdateOne) RemoveProjects(p ...*Project) *ParticipantUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveProjectIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *ParticipantUpdateOne) Save(ctx context.Context) (*Participant, error) {
	var (
		err  error
		node *Participant
	)
	puo.defaults()
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ParticipantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ParticipantUpdateOne) SaveX(ctx context.Context) *Participant {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ParticipantUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ParticipantUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ParticipantUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := participant.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

func (puo *ParticipantUpdateOne) sqlSave(ctx context.Context) (_node *Participant, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   participant.Table,
			Columns: participant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: participant.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Participant.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: participant.FieldUpdatedAt,
		})
	}
	if value, ok := puo.mutation.MturkWorkerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: participant.FieldMturkWorkerID,
		})
	}
	if puo.mutation.MturkWorkerIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: participant.FieldMturkWorkerID,
		})
	}
	if puo.mutation.DataCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedDataIDs(); len(nodes) > 0 && !puo.mutation.DataCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.DataIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ProviderIDsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedProviderIDsIDs(); len(nodes) > 0 && !puo.mutation.ProviderIDsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProviderIDsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ParticipationsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedParticipationsIDs(); len(nodes) > 0 && !puo.mutation.ParticipationsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ParticipationsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.CreatedByCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CreatedByIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.StepsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedStepsIDs(); len(nodes) > 0 && !puo.mutation.StepsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.StepsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   participant.ProjectsTable,
			Columns: participant.ProjectsPrimaryKey,
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
	if nodes := puo.mutation.RemovedProjectsIDs(); len(nodes) > 0 && !puo.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   participant.ProjectsTable,
			Columns: participant.ProjectsPrimaryKey,
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
	if nodes := puo.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   participant.ProjectsTable,
			Columns: participant.ProjectsPrimaryKey,
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
	_node = &Participant{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{participant.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
