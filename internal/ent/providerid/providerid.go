// Code generated by entc, DO NOT EDIT.

package providerid

import (
	"time"
)

const (
	// Label holds the string label denoting the providerid type in the database.
	Label = "provider_id"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldMturkWorkerID holds the string denoting the mturkworkerid field in the database.
	FieldMturkWorkerID = "mturk_worker_id"

	// EdgeParticpant holds the string denoting the particpant edge name in mutations.
	EdgeParticpant = "particpant"

	// Table holds the table name of the providerid in the database.
	Table = "provider_ids"
	// ParticpantTable is the table the holds the particpant relation/edge.
	ParticpantTable = "provider_ids"
	// ParticpantInverseTable is the table name for the Participant entity.
	// It exists in this package in order to avoid circular dependency with the "participant" package.
	ParticpantInverseTable = "participants"
	// ParticpantColumn is the table column denoting the particpant relation/edge.
	ParticpantColumn = "participant_provider_ids"
)

// Columns holds all SQL columns for providerid fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldMturkWorkerID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the ProviderID type.
var ForeignKeys = []string{
	"participant_provider_ids",
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
