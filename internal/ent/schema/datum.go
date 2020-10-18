package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Datum holds the schema definition for the Datum entity.
type Datum struct {
	ent.Schema
}

// Mixin of the Datum.
func (Datum) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Datum.
func (Datum) Fields() []ent.Field {
	return []ent.Field{
		field.String("key"),
		field.Bytes("val"),
		field.Int("index").Default(0),
		// Versioning the current version is marked current == true. All inserts
		// must be done with transaction to make sure only one key per
		// participant is marked current.
		// Version is an increasing value starting from 0.
		field.Bool("current").Default(true),
		field.Int("version").Default(0),
		// DeletedAt should be set when the value is deleted and it should no
		// longer be returned for normal data queries
		field.Time("deletedAt").Optional().Nillable(),
	}
}

// Edges of the Datum.
func (Datum) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("participant", Participant.Type).
			Ref("data").
			Unique().
			Required(),
	}
}
