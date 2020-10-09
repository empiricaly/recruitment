package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// ProviderID holds the schema definition for the ProviderID entity.
type ProviderID struct {
	ent.Schema
}

// Mixin of the ProviderID.
func (ProviderID) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the ProviderID.
func (ProviderID) Fields() []ent.Field {
	return []ent.Field{
		field.String("mturkWorkerID"),
	}
}

// Edges of the ProviderID.
func (ProviderID) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("particpant", Participant.Type).
			Ref("providerIDs").
			Unique(),
	}
}
