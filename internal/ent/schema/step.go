package schema

import (
	"github.com/empiricaly/recruitment/internal/model"
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Step holds the schema definition for the Step entity.
type Step struct {
	ent.Schema
}

// Mixin of the Step.
func (Step) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Step.
func (Step) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values("MTURK_HIT", "MTURK_MESSAGE", "PARTICIPANT_FILTER", "WAIT"),
		field.Int("index"),
		field.Int("duration"),
		field.JSON("msgArgs", &model.MessageStepArgs{}).Optional(),
		field.JSON("hitArgs", &model.HITStepArgs{}).Optional(),
		field.JSON("filterArgs", &model.FilterStepArgs{}).Optional(),
	}
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
