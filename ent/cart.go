// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mall/ent/cart"
	"mall/ent/user"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Cart is the model entity for the Cart schema.
type Cart struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CartQuery when eager-loading is set.
	Edges        CartEdges `json:"edges"`
	user_cart    *int
	selectValues sql.SelectValues
}

// CartEdges holds the relations/edges for other nodes in the graph.
type CartEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Items holds the value of the items edge.
	Items []*CartItem `json:"items,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CartEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// ItemsOrErr returns the Items value or an error if the edge
// was not loaded in eager-loading.
func (e CartEdges) ItemsOrErr() ([]*CartItem, error) {
	if e.loadedTypes[1] {
		return e.Items, nil
	}
	return nil, &NotLoadedError{edge: "items"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Cart) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cart.FieldID:
			values[i] = new(sql.NullInt64)
		case cart.ForeignKeys[0]: // user_cart
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Cart fields.
func (c *Cart) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cart.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case cart.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_cart", value)
			} else if value.Valid {
				c.user_cart = new(int)
				*c.user_cart = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Cart.
// This includes values selected through modifiers, order, etc.
func (c *Cart) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Cart entity.
func (c *Cart) QueryUser() *UserQuery {
	return NewCartClient(c.config).QueryUser(c)
}

// QueryItems queries the "items" edge of the Cart entity.
func (c *Cart) QueryItems() *CartItemQuery {
	return NewCartClient(c.config).QueryItems(c)
}

// Update returns a builder for updating this Cart.
// Note that you need to call Cart.Unwrap() before calling this method if this Cart
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Cart) Update() *CartUpdateOne {
	return NewCartClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Cart entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Cart) Unwrap() *Cart {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Cart is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Cart) String() string {
	var builder strings.Builder
	builder.WriteString("Cart(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Carts is a parsable slice of Cart.
type Carts []*Cart
