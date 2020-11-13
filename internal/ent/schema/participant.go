package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Participant holds the schema definition for the Participant entity.
type Participant struct {
	ent.Schema
}

// Mixin of the Participant.
func (Participant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Participant.
func (Participant) Fields() []ent.Field {
	return []ent.Field{
		field.String("mturkWorkerID").Optional().Nillable(),
		field.Bool("uninitialized").Optional().Nillable(),
	}
}

// Edges of the Participant.
func (Participant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("data", Datum.Type),
		edge.To("providerIDs", ProviderID.Type),
		edge.To("participations", Participation.Type),
		edge.From("createdBy", StepRun.Type).
			Ref("createdParticipants").
			Unique(),
		edge.From("steps", StepRun.Type).
			Ref("participants"),
		edge.From("projects", Project.Type).
			Ref("participants"),
		edge.From("importedBy", Admin.Type).Ref("importedParticipants"),
	}
}
