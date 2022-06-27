package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// Domain holds the schemas definition for the Domain entity.
type Domain struct {
	ent.Schema
}

// Fields of the Domain.
func (Domain) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now),

		field.String("domain").Unique().Immutable(),
		field.String("cloudflare_id").Unique().Immutable(),
		field.Time("expires_at").Optional().Nillable(),

		field.Strings("nameservers"),

		field.Bool("pending").Default(true).StructTag(`json:"pending"`),
		field.Bool("private").Default(false).StructTag(`json:"private"`),
		field.Bool("approved").Default(false).StructTag(`json:"approved"`),
	}
}

// Edges of the Domain.
func (Domain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("creator", User.Type),
		edge.To("usable_by", User.Type),
	}
}

// Indexes of the Domain
func (Domain) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("domain"),
	}
}
