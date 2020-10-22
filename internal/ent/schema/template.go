package schema

import (
	"github.com/empiricaly/recruitment/internal/model"
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Template holds the schema definition for the Run entity.
type Template struct {
	ent.Schema
}

// Mixin of the Template.
func (Template) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Template.
func (Template) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).MinLen(1),
		field.Enum("selectionType").Values("INTERNAL_DB", "MTURK_QUALIFICATIONS"),
		field.Int("participantCount").NonNegative().Default(0),
		field.JSON("internalCriteria", &model.InternalCriteria{}),
		field.JSON("mturkCriteria", &model.MTurkCriteria{}),
		field.Bool("adult").Default(false),
		field.Bool("sandbox").Default(false),
	}
}

// Edges of the Template.
func (Template) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("steps", Step.Type),
		edge.From("project", Project.Type).
			Ref("templates").
			Unique(),
		edge.From("creator", Admin.Type).
			Ref("templates").
			Unique(),
		edge.From("run", Run.Type).
			Ref("template").
			Unique(),
	}
}
