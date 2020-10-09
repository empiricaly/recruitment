package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

// BaseMixin implements the ent.Mixin for sharing
// time fields and the ID with package schemas.
type BaseMixin struct {
	mixin.Schema
}

// Fields for BaseMixin.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MinLen(20).
			MaxLen(20).
			NotEmpty().
			Unique().
			Immutable(),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// StatusMixin implements the ent.Mixin for sharing
// status fields with package schemas.
type StatusMixin struct {
	mixin.Schema
}

// Fields for BaseMixin.
func (StatusMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").Values(
			"CREATED",
			"RUNNING",
			"PAUSED",
			"DONE",
			"TERMINATED",
			"FAILED",
		),
		field.Time("startedAt").Optional().Nillable(),
		field.Time("endedAt").Optional().Nillable(),
	}
}
