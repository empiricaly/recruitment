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
	return append(
		append([]ent.Field{}, commonFields...),
		field.String("name"),
		statusField,
		field.Time("startAt").Optional(),
		field.Time("startedAt").Optional(),
		field.Time("endedAt").Optional(),
		field.String("error").Optional(),
	)
}

// Edges of the Run.
func (Run) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("runs").
			Unique(),
		edge.To("template", Template.Type).
			Unique().
			Required(),
		edge.To("steps", StepRun.Type),
	}
}
