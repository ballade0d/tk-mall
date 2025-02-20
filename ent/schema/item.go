package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.String("name"),
		field.Text("description"),
		field.Float32("price"),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return nil
}
