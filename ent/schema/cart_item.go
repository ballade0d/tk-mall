package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CartItem holds the schema definition for the CartItem entity.
type CartItem struct {
	ent.Schema
}

// Fields of the CartItem.
func (CartItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantity").Default(1).Min(1),
	}
}

// Edges of the CartItem.
func (CartItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("cart", Cart.Type).Ref("items").Unique().Required(),
		edge.To("item", Item.Type).Unique().Required(),
	}
}
