package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Password holds the schema definition for the Password entity.
type Password struct {
	ent.Schema
}

// Fields of the Password.
func (Password) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("password"),
	}
}

// Edges of the Password.
func (Password) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("password").Unique(),
	}
}
