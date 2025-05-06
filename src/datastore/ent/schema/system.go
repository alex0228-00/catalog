package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// System holds the schema definition for the System entity.
type System struct {
	ent.Schema
}

// Fields of the System.
func (System) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique(),
		field.String("type").NotEmpty(),
		field.String("host").NotEmpty(),
		field.String("unique_identifier"),
		field.Text("credentials").NotEmpty(),

		field.String("created_by").NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.String("updated_by").NotEmpty(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("deleted_by").Optional().Nillable(),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the System.
func (System) Edges() []ent.Edge {
	return nil
}
