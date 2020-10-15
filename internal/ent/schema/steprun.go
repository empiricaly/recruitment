package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// StepRun holds the schema definition for the StepRun entity.
type StepRun struct {
	ent.Schema
}

// Mixin of the StepRun.
func (StepRun) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		StatusMixin{},
	}
}

// Fields of the StepRun.
func (StepRun) Fields() []ent.Field {
	return []ent.Field{
		field.Int("participantsCount"),
		field.String("hitID").Optional().Nillable(),
	}
}

// Edges of the StepRun.
func (StepRun) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("createdParticipants", Participant.Type),
		edge.To("participants", Participant.Type),
		edge.To("participations", Participation.Type),
		edge.To("step", Step.Type).
			Unique().
			Required(),
		edge.From("run", Run.Type).
			Ref("steps").
			Unique(),
	}
}
