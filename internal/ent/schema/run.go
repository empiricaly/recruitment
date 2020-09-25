package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Run holds the schema definition for the Run entity.
type Run struct {
	ent.Schema
}

// Fields of the Run.
func (Run) Fields() []ent.Field {
	// TODO field : creator, procedure, status, steps, data, currentStep
	return append(
		append([]ent.Field{}, commonFields...),
		field.String("name"),
		field.Time("startAt"),
		field.Time("startedAt"),
		field.Time("endedAt"),
		field.String("error"),
	)
}

// Edges of the Run.
func (Run) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("procedure", Procedure.Type).Unique(),
	}
}
