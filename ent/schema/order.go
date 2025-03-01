package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return nil
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Required(),
		edge.To("items", OrderItem.Type),
		edge.To("payment", Payment.Type).Unique(),
	}
}
