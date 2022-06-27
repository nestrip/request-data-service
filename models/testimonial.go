package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Testimonial holds the schemas definition for the Testimonial entity.
type Testimonial struct {
	ent.Schema
}

// Fields of the Testimonial.
func (Testimonial) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.Time("created_at").Default(time.Now),
		field.String("message"),
		field.Bool("approved").Optional().Nillable().StructTag(`json:"approved"`),
	}
}

func (Testimonial) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", User.Type).Ref("testimonial").Unique(),
	}
}
