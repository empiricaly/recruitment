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
		// If true, the corresponding participant was added during this step participation
		field.Bool("addedParticipant").Default(false),
		// Fields from MTurk ExternalQuestion params:
		// https://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_ExternalQuestionArticle.html#ApiReference_ExternalQuestionArticle-the-external-form
		field.String("mturkWorkerID"),
		field.String("mturkAssignmentID"),
		field.String("mturkHitID"),
		// Corresponds to MTurk AcceptTime: The date and time the Worker
		// accepted the assignment.
		field.Time("mturkAcceptedAt"),
		// Corresponds to SubmitTime: If the Worker has submitted results,
		// SubmitTime is the date and time the assignment was submitted.
		// This value is omitted from the assignment if the Worker has not
		// yet submitted results.
		field.Time("mturkSubmittedAt"),
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
