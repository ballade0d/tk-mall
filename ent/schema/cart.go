package schema

import "entgo.io/ent"

// Cart holds the schema definition for the Cart entity.
type Cart struct {
	ent.Schema
}

// Fields of the Cart.
func (Cart) Fields() []ent.Field {
	return nil
}

// Edges of the Cart.
func (Cart) Edges() []ent.Edge {
	return nil
}
