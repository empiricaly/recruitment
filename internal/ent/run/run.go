// Code generated by entc, DO NOT EDIT.

package run

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the run type in the database.
	Label = "run"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldStartAt holds the string denoting the startat field in the database.
	FieldStartAt = "start_at"
	// FieldStartedAt holds the string denoting the startedat field in the database.
	FieldStartedAt = "started_at"
	// FieldEndedAt holds the string denoting the endedat field in the database.
	FieldEndedAt = "ended_at"
	// FieldError holds the string denoting the error field in the database.
	FieldError = "error"

	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeTemplate holds the string denoting the template edge name in mutations.
	EdgeTemplate = "template"
	// EdgeSteps holds the string denoting the steps edge name in mutations.
	EdgeSteps = "steps"

	// Table holds the table name of the run in the database.
	Table = "runs"
	// ProjectTable is the table the holds the project relation/edge.
	ProjectTable = "runs"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_runs"
	// TemplateTable is the table the holds the template relation/edge.
	TemplateTable = "templates"
	// TemplateInverseTable is the table name for the Template entity.
	// It exists in this package in order to avoid circular dependency with the "template" package.
	TemplateInverseTable = "templates"
	// TemplateColumn is the table column denoting the template relation/edge.
	TemplateColumn = "run_template"
	// StepsTable is the table the holds the steps relation/edge.
	StepsTable = "step_runs"
	// StepsInverseTable is the table name for the StepRun entity.
	// It exists in this package in order to avoid circular dependency with the "steprun" package.
	StepsInverseTable = "step_runs"
	// StepsColumn is the table column denoting the steps relation/edge.
	StepsColumn = "run_steps"
)

// Columns holds all SQL columns for run fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldStatus,
	FieldStartAt,
	FieldStartedAt,
	FieldEndedAt,
	FieldError,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Run type.
var ForeignKeys = []string{
	"project_runs",
}

var (
	// DefaultCreatedAt holds the default value on creation for the createdAt field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	UpdateDefaultUpdatedAt func() time.Time
)

// Status defines the type for the status enum field.
type Status string

// Status values.
const (
	StatusCREATED    Status = "CREATED"
	StatusRUNNING    Status = "RUNNING"
	StatusPAUSED     Status = "PAUSED"
	StatusDONE       Status = "DONE"
	StatusTERMINATED Status = "TERMINATED"
	StatusFAILED     Status = "FAILED"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusCREATED, StatusRUNNING, StatusPAUSED, StatusDONE, StatusTERMINATED, StatusFAILED:
		return nil
	default:
		return fmt.Errorf("run: invalid enum value for status field: %q", s)
	}
}
