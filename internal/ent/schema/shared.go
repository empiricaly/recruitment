package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

var commonFields = []ent.Field{
	field.String("id"),
	field.Time("createdAt").Immutable().Default(time.Now),
	field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
}

var statusField = field.Enum("status").Values("CREATED", "RUNNING", "PAUSED", "DONE", "TERMINATED", "FAILED")
