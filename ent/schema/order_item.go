package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OrderItem holds the schema definition for the OrderItem entity.
type OrderItem struct {
	ent.Schema
}

// Fields of the CartItem.
func (OrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("quantity").Default(1).Min(1),
		field.Float32("price").Positive(),
	}
}

// Edges of the CartItem.
func (OrderItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).Ref("items").Unique().Required(),
		edge.To("item", Item.Type).Unique().Required(),
	}
}
