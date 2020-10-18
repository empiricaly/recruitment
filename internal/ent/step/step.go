// Code generated by entc, DO NOT EDIT.

package step

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the step type in the database.
	Label = "step"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldIndex holds the string denoting the index field in the database.
	FieldIndex = "index"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldMsgArgs holds the string denoting the msgargs field in the database.
	FieldMsgArgs = "msg_args"
	// FieldHitArgs holds the string denoting the hitargs field in the database.
	FieldHitArgs = "hit_args"
	// FieldFilterArgs holds the string denoting the filterargs field in the database.
	FieldFilterArgs = "filter_args"

	// EdgeStepRun holds the string denoting the steprun edge name in mutations.
	EdgeStepRun = "stepRun"
	// EdgeTemplate holds the string denoting the template edge name in mutations.
	EdgeTemplate = "template"

	// Table holds the table name of the step in the database.
	Table = "steps"
	// StepRunTable is the table the holds the stepRun relation/edge.
	StepRunTable = "steps"
	// StepRunInverseTable is the table name for the StepRun entity.
	// It exists in this package in order to avoid circular dependency with the "steprun" package.
	StepRunInverseTable = "step_runs"
	// StepRunColumn is the table column denoting the stepRun relation/edge.
	StepRunColumn = "step_run_step"
	// TemplateTable is the table the holds the template relation/edge.
	TemplateTable = "steps"
	// TemplateInverseTable is the table name for the Template entity.
	// It exists in this package in order to avoid circular dependency with the "template" package.
	TemplateInverseTable = "templates"
	// TemplateColumn is the table column denoting the template relation/edge.
	TemplateColumn = "template_steps"
)

// Columns holds all SQL columns for step fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldType,
	FieldIndex,
	FieldDuration,
	FieldMsgArgs,
	FieldHitArgs,
	FieldFilterArgs,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Step type.
var ForeignKeys = []string{
	"step_run_step",
	"template_steps",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// Type defines the type for the type enum field.
type Type string

// Type values.
const (
	TypeMTURK_HIT          Type = "MTURK_HIT"
	TypeMTURK_MESSAGE      Type = "MTURK_MESSAGE"
	TypePARTICIPANT_FILTER Type = "PARTICIPANT_FILTER"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeMTURK_HIT, TypeMTURK_MESSAGE, TypePARTICIPANT_FILTER:
		return nil
	default:
		return fmt.Errorf("step: invalid enum value for type field: %q", _type)
	}
}
