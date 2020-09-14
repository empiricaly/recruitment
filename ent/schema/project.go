package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	// TODO field : creator, procedure, status, steps, data
	return []ent.Field{
		field.String("id"),
		field.Time("createdAt").Immutable().Default(time.Now),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
		field.Time("startAt"),
		field.Time("startedAt"),
		field.Time("endedAt"),
		field.String("error"),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return nil
}
