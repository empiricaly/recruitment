package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Run holds the schema definition for the Run entity.
type Run struct {
	ent.Schema
}

// Fields of the Run.
func (Run) Fields() []ent.Field {
	// TODO field : creator, procedure, status, steps, data, currentStep
	return []ent.Field{
		field.String("id"),
		field.Time("createdAt").Immutable().Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
		field.Time("startAt"),
		field.Time("startedAt"),
		field.Time("endedAt"),
		field.String("error"),
	}
}

// Edges of the Run.
func (Run) Edges() []ent.Edge {
	return nil
}
