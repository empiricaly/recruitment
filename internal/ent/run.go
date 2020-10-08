// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/empiricaly/recruitment/internal/ent/project"
	"github.com/empiricaly/recruitment/internal/ent/run"
	"github.com/empiricaly/recruitment/internal/ent/template"
	"github.com/facebook/ent/dialect/sql"
)

// Run is the model entity for the Run schema.
type Run struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Status holds the value of the "status" field.
	Status run.Status `json:"status,omitempty"`
	// StartAt holds the value of the "startAt" field.
	StartAt time.Time `json:"startAt,omitempty"`
	// StartedAt holds the value of the "startedAt" field.
	StartedAt time.Time `json:"startedAt,omitempty"`
	// EndedAt holds the value of the "endedAt" field.
	EndedAt time.Time `json:"endedAt,omitempty"`
	// Error holds the value of the "error" field.
	Error string `json:"error,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RunQuery when eager-loading is set.
	Edges        RunEdges `json:"edges"`
	project_runs *string
}

// RunEdges holds the relations/edges for other nodes in the graph.
type RunEdges struct {
	// Project holds the value of the project edge.
	Project *Project
	// Template holds the value of the template edge.
	Template *Template
	// Steps holds the value of the steps edge.
	Steps []*StepRun
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RunEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[0] {
		if e.Project == nil {
			// The edge project was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// TemplateOrErr returns the Template value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RunEdges) TemplateOrErr() (*Template, error) {
	if e.loadedTypes[1] {
		if e.Template == nil {
			// The edge template was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: template.Label}
		}
		return e.Template, nil
	}
	return nil, &NotLoadedError{edge: "template"}
}

// StepsOrErr returns the Steps value or an error if the edge
// was not loaded in eager-loading.
func (e RunEdges) StepsOrErr() ([]*StepRun, error) {
	if e.loadedTypes[2] {
		return e.Steps, nil
	}
	return nil, &NotLoadedError{edge: "steps"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Run) scanValues() []interface{} {
	return []interface{}{
		&sql.NullString{}, // id
		&sql.NullTime{},   // createdAt
		&sql.NullTime{},   // updatedAt
		&sql.NullString{}, // name
		&sql.NullString{}, // status
		&sql.NullTime{},   // startAt
		&sql.NullTime{},   // startedAt
		&sql.NullTime{},   // endedAt
		&sql.NullString{}, // error
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Run) fkValues() []interface{} {
	return []interface{}{
		&sql.NullString{}, // project_runs
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Run fields.
func (r *Run) assignValues(values ...interface{}) error {
	if m, n := len(values), len(run.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value.Valid {
		r.ID = value.String
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field createdAt", values[0])
	} else if value.Valid {
		r.CreatedAt = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field updatedAt", values[1])
	} else if value.Valid {
		r.UpdatedAt = value.Time
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[2])
	} else if value.Valid {
		r.Name = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field status", values[3])
	} else if value.Valid {
		r.Status = run.Status(value.String)
	}
	if value, ok := values[4].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field startAt", values[4])
	} else if value.Valid {
		r.StartAt = value.Time
	}
	if value, ok := values[5].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field startedAt", values[5])
	} else if value.Valid {
		r.StartedAt = value.Time
	}
	if value, ok := values[6].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field endedAt", values[6])
	} else if value.Valid {
		r.EndedAt = value.Time
	}
	if value, ok := values[7].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field error", values[7])
	} else if value.Valid {
		r.Error = value.String
	}
	values = values[8:]
	if len(values) == len(run.ForeignKeys) {
		if value, ok := values[0].(*sql.NullString); !ok {
			return fmt.Errorf("unexpected type %T for field project_runs", values[0])
		} else if value.Valid {
			r.project_runs = new(string)
			*r.project_runs = value.String
		}
	}
	return nil
}

// QueryProject queries the project edge of the Run.
func (r *Run) QueryProject() *ProjectQuery {
	return (&RunClient{config: r.config}).QueryProject(r)
}

// QueryTemplate queries the template edge of the Run.
func (r *Run) QueryTemplate() *TemplateQuery {
	return (&RunClient{config: r.config}).QueryTemplate(r)
}

// QuerySteps queries the steps edge of the Run.
func (r *Run) QuerySteps() *StepRunQuery {
	return (&RunClient{config: r.config}).QuerySteps(r)
}

// Update returns a builder for updating this Run.
// Note that, you need to call Run.Unwrap() before calling this method, if this Run
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Run) Update() *RunUpdateOne {
	return (&RunClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (r *Run) Unwrap() *Run {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Run is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Run) String() string {
	var builder strings.Builder
	builder.WriteString("Run(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", createdAt=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", name=")
	builder.WriteString(r.Name)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", r.Status))
	builder.WriteString(", startAt=")
	builder.WriteString(r.StartAt.Format(time.ANSIC))
	builder.WriteString(", startedAt=")
	builder.WriteString(r.StartedAt.Format(time.ANSIC))
	builder.WriteString(", endedAt=")
	builder.WriteString(r.EndedAt.Format(time.ANSIC))
	builder.WriteString(", error=")
	builder.WriteString(r.Error)
	builder.WriteByte(')')
	return builder.String()
}

// Runs is a parsable slice of Run.
type Runs []*Run

func (r Runs) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
