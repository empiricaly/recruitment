package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Procedure holds the schema definition for the Run entity.
type Procedure struct {
	ent.Schema
}

// Fields of the Procedure.
func (Procedure) Fields() []ent.Field {
	// TODO field : creator, selectionType, internalCriteria, mturkCriteria, steps
	return append(
		append([]ent.Field{}, commonFields...),
		field.String("name").MaxLen(255).MinLen(1),
		field.String("selectionType"),
		field.Int("participantCount").NonNegative().Default(0),
		field.Bool("adult").Default(false),
	)
}

// Edges of the Procedure.
func (Procedure) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("procedures").
			Unique(),
		edge.From("owner", Admin.Type).
			Ref("procedures").
			Unique(),
		edge.From("run", Run.Type).
			Ref("procedure").
			Unique().
			Required(),
	}
}
