// Code generated by entc, DO NOT EDIT.

package steprun

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the steprun type in the database.
	Label = "step_run"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldStartedAt holds the string denoting the startedat field in the database.
	FieldStartedAt = "started_at"
	// FieldEndedAt holds the string denoting the endedat field in the database.
	FieldEndedAt = "ended_at"
	// FieldIndex holds the string denoting the index field in the database.
	FieldIndex = "index"
	// FieldParticipantsCount holds the string denoting the participantscount field in the database.
	FieldParticipantsCount = "participants_count"
	// FieldHitID holds the string denoting the hitid field in the database.
	FieldHitID = "hit_id"
	// FieldUrlToken holds the string denoting the urltoken field in the database.
	FieldUrlToken = "url_token"

	// EdgeCreatedParticipants holds the string denoting the createdparticipants edge name in mutations.
	EdgeCreatedParticipants = "createdParticipants"
	// EdgeParticipants holds the string denoting the participants edge name in mutations.
	EdgeParticipants = "participants"
	// EdgeParticipations holds the string denoting the participations edge name in mutations.
	EdgeParticipations = "participations"
	// EdgeStep holds the string denoting the step edge name in mutations.
	EdgeStep = "step"
	// EdgeRun holds the string denoting the run edge name in mutations.
	EdgeRun = "run"

	// Table holds the table name of the steprun in the database.
	Table = "step_runs"
	// CreatedParticipantsTable is the table the holds the createdParticipants relation/edge.
	CreatedParticipantsTable = "participants"
	// CreatedParticipantsInverseTable is the table name for the Participant entity.
	// It exists in this package in order to avoid circular dependency with the "participant" package.
	CreatedParticipantsInverseTable = "participants"
	// CreatedParticipantsColumn is the table column denoting the createdParticipants relation/edge.
	CreatedParticipantsColumn = "step_run_created_participants"
	// ParticipantsTable is the table the holds the participants relation/edge. The primary key declared below.
	ParticipantsTable = "step_run_participants"
	// ParticipantsInverseTable is the table name for the Participant entity.
	// It exists in this package in order to avoid circular dependency with the "participant" package.
	ParticipantsInverseTable = "participants"
	// ParticipationsTable is the table the holds the participations relation/edge.
	ParticipationsTable = "participations"
	// ParticipationsInverseTable is the table name for the Participation entity.
	// It exists in this package in order to avoid circular dependency with the "participation" package.
	ParticipationsInverseTable = "participations"
	// ParticipationsColumn is the table column denoting the participations relation/edge.
	ParticipationsColumn = "step_run_participations"
	// StepTable is the table the holds the step relation/edge.
	StepTable = "steps"
	// StepInverseTable is the table name for the Step entity.
	// It exists in this package in order to avoid circular dependency with the "step" package.
	StepInverseTable = "steps"
	// StepColumn is the table column denoting the step relation/edge.
	StepColumn = "step_run_step"
	// RunTable is the table the holds the run relation/edge.
	RunTable = "step_runs"
	// RunInverseTable is the table name for the Run entity.
	// It exists in this package in order to avoid circular dependency with the "run" package.
	RunInverseTable = "runs"
	// RunColumn is the table column denoting the run relation/edge.
	RunColumn = "run_steps"
)

// Columns holds all SQL columns for steprun fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldStartedAt,
	FieldEndedAt,
	FieldIndex,
	FieldParticipantsCount,
	FieldHitID,
	FieldUrlToken,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the StepRun type.
var ForeignKeys = []string{
	"run_steps",
}

var (
	// ParticipantsPrimaryKey and ParticipantsColumn2 are the table columns denoting the
	// primary key for the participants relation (M2M).
	ParticipantsPrimaryKey = []string{"step_run_id", "participant_id"}
)

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
		return fmt.Errorf("steprun: invalid enum value for status field: %q", s)
	}
}
