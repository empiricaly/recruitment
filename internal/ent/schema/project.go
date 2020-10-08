package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return append(
		append([]ent.Field{}, commonFields...),
		field.String("projectID"),
		field.String("name"),
	)
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("runs", Run.Type),
		edge.To("templates", Template.Type),
		edge.From("owner", Admin.Type).
			Ref("projects").
			Unique(),
	}
}
