package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").NotEmpty(),
		field.Enum("status").Values("pending", "paid", "shipped", "delivered", "cancelled").Default("pending"),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("order").Unique(),
		edge.To("items", OrderItem.Type),
		edge.To("payment", Payment.Type),
	}
}
