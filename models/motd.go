package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// MOTD holds the schemas definition for the MOTD entity.
type MOTD struct {
	ent.Schema
}

// Fields of the MOTD.
func (MOTD) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now),
		field.String("message").NotEmpty(),
		field.Bool("active").Default(false),
	}
}

// Edges of the MOTD.
func (MOTD) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("creator", User.Type),
	}
}

// Indexes of the MOTD
func (MOTD) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}
