package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// StepRun holds the schema definition for the StepRun entity.
type StepRun struct {
	ent.Schema
}

// Fields of the StepRun.
func (StepRun) Fields() []ent.Field {
	// TODO field : step, status, participants,
	return []ent.Field{
		field.Time("startAt"),
		field.Time("endedAt"),
		field.Int("participantsCount"),
	}
}

// Edges of the StepRun.
func (StepRun) Edges() []ent.Edge {
	return nil
}
