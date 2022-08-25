package models

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// File holds the schemas definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now),
		field.String("filename"),
		field.String("cdn_file_name").Default("null"),
		field.String("original_filename"),
		field.String("mime_type"),
		field.Int64("size"),
		field.Bool("should_embed").Default(true).StructTag(`json:"should_embed"`),
		field.JSON("embed", Embed{}),
		field.String("hash").Default("Unknown"),
		field.String("user_agent").Default("Unknown"),
		field.Bool("exploding").Default(false).StructTag(`json:"exploding"`),
		field.Bool("exploded").Default(false).StructTag(`json:"exploded"`),
		field.String("deletion_key").Sensitive().Default(""),
		field.Bool("archived"), // an archived file is more than 30 days old, and stored on a different disk
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("uploader", User.Type),
	}
}

// Indexes of the File
func (File) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("filename"),
	}
}
