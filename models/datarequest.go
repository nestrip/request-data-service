package models

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// DataRequest holds the schemas definition for the DataRequest entity.
type DataRequest struct {
	ent.Schema
}

// Fields of the DataRequest.
func (DataRequest) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.Time("created_at").Default(time.Now),

		field.String("cdn_name"),
		field.Bool("expired").Default(false),
	}
}

func (DataRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", User.Type).Ref("data_requests").Unique(),
	}
}
