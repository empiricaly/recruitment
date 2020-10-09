package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Participation holds the schema definition for the Participation entity.
type Participation struct {
	ent.Schema
}

// Mixin of the Participation.
func (Participation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Participation.
func (Participation) Fields() []ent.Field {
	return []ent.Field{
		// Fields from MTurk ExternalQuestion params:
		// https://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_ExternalQuestionArticle.html#ApiReference_ExternalQuestionArticle-the-external-form
		field.String("mturkWorkerId"),
		field.String("mturkAssignmentID"),
		field.String("mturkHitID"),
		field.String("mturkTurkSubmitTo"),
	}
}

// Edges of the Participation.
func (Participation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("stepRun", StepRun.Type).
			Ref("participations").
			Unique(),
		edge.From("participant", Participant.Type).
			Ref("participations").
			Unique(),
	}
}
