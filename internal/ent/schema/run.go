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

// Mixin of the Run.
func (Run) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		StatusMixin{},
	}
}

// Fields of the Run.
func (Run) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("startAt").Optional().Nillable(),
		field.String("error").Optional().Nillable(),
	}
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
		edge.To("currentStep", StepRun.Type).
			Unique(),
		edge.To("steps", StepRun.Type),
	}
}
