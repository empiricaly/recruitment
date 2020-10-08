package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Step holds the schema definition for the Step entity.
type Step struct {
	ent.Schema
}

// Fields of the Step.
func (Step) Fields() []ent.Field {
	return append(
		append([]ent.Field{}, commonFields...),
		field.Enum("type").Values("MTURK_HIT", "MTURK_MESSAGE", "PARTICIPANT_FILTER"),
		field.Int("index"),
		field.Int("duration"),
		field.Bytes("msgArgs").Optional(),
		field.Bytes("hitArgs").Optional(),
		field.Bytes("filterArgs").Optional(),
	)
}

// Edges of the Step.
func (Step) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("stepRun", StepRun.Type).
			Ref("step").
			Unique(),
		edge.From("template", Template.Type).
			Ref("steps").
			Unique(),
	}
}
