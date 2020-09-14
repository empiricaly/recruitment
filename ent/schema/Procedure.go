package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Procedure holds the schema definition for the Run entity.
type Procedure struct {
	ent.Schema
}

// Fields of the Procedure.
func (Procedure) Fields() []ent.Field {
	// TODO field : creator, selectionType, internalCriteria, mturkCriteria, steps
	return []ent.Field{
		field.String("id"),
		field.Time("createdAt").Immutable().Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
		field.String("name"),
		field.String("participantCount"),
		field.Bool("adult"),
	}
}

// Edges of the Procedure.
func (Procedure) Edges() []ent.Edge {
	return nil
}
