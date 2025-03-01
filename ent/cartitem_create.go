// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mall/ent/cart"
	"mall/ent/cartitem"
	"mall/ent/item"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CartItemCreate is the builder for creating a CartItem entity.
type CartItemCreate struct {
	config
	mutation *CartItemMutation
	hooks    []Hook
}

// SetQuantity sets the "quantity" field.
func (cic *CartItemCreate) SetQuantity(i int) *CartItemCreate {
	cic.mutation.SetQuantity(i)
	return cic
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (cic *CartItemCreate) SetNillableQuantity(i *int) *CartItemCreate {
	if i != nil {
		cic.SetQuantity(*i)
	}
	return cic
}

// SetCartID sets the "cart" edge to the Cart entity by ID.
func (cic *CartItemCreate) SetCartID(id int) *CartItemCreate {
	cic.mutation.SetCartID(id)
	return cic
}

// SetCart sets the "cart" edge to the Cart entity.
func (cic *CartItemCreate) SetCart(c *Cart) *CartItemCreate {
	return cic.SetCartID(c.ID)
}

// SetItemID sets the "item" edge to the Item entity by ID.
func (cic *CartItemCreate) SetItemID(id int) *CartItemCreate {
	cic.mutation.SetItemID(id)
	return cic
}

// SetItem sets the "item" edge to the Item entity.
func (cic *CartItemCreate) SetItem(i *Item) *CartItemCreate {
	return cic.SetItemID(i.ID)
}

// Mutation returns the CartItemMutation object of the builder.
func (cic *CartItemCreate) Mutation() *CartItemMutation {
	return cic.mutation
}

// Save creates the CartItem in the database.
func (cic *CartItemCreate) Save(ctx context.Context) (*CartItem, error) {
	cic.defaults()
	return withHooks(ctx, cic.sqlSave, cic.mutation, cic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cic *CartItemCreate) SaveX(ctx context.Context) *CartItem {
	v, err := cic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cic *CartItemCreate) Exec(ctx context.Context) error {
	_, err := cic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cic *CartItemCreate) ExecX(ctx context.Context) {
	if err := cic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cic *CartItemCreate) defaults() {
	if _, ok := cic.mutation.Quantity(); !ok {
		v := cartitem.DefaultQuantity
		cic.mutation.SetQuantity(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cic *CartItemCreate) check() error {
	if _, ok := cic.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "CartItem.quantity"`)}
	}
	if v, ok := cic.mutation.Quantity(); ok {
		if err := cartitem.QuantityValidator(v); err != nil {
			return &ValidationError{Name: "quantity", err: fmt.Errorf(`ent: validator failed for field "CartItem.quantity": %w`, err)}
		}
	}
	if len(cic.mutation.CartIDs()) == 0 {
		return &ValidationError{Name: "cart", err: errors.New(`ent: missing required edge "CartItem.cart"`)}
	}
	if len(cic.mutation.ItemIDs()) == 0 {
		return &ValidationError{Name: "item", err: errors.New(`ent: missing required edge "CartItem.item"`)}
	}
	return nil
}

func (cic *CartItemCreate) sqlSave(ctx context.Context) (*CartItem, error) {
	if err := cic.check(); err != nil {
		return nil, err
	}
	_node, _spec := cic.createSpec()
	if err := sqlgraph.CreateNode(ctx, cic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cic.mutation.id = &_node.ID
	cic.mutation.done = true
	return _node, nil
}

func (cic *CartItemCreate) createSpec() (*CartItem, *sqlgraph.CreateSpec) {
	var (
		_node = &CartItem{config: cic.config}
		_spec = sqlgraph.NewCreateSpec(cartitem.Table, sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeInt))
	)
	if value, ok := cic.mutation.Quantity(); ok {
		_spec.SetField(cartitem.FieldQuantity, field.TypeInt, value)
		_node.Quantity = value
	}
	if nodes := cic.mutation.CartIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cartitem.CartTable,
			Columns: []string{cartitem.CartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.cart_items = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cic.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   cartitem.ItemTable,
			Columns: []string{cartitem.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.cart_item_item = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CartItemCreateBulk is the builder for creating many CartItem entities in bulk.
type CartItemCreateBulk struct {
	config
	err      error
	builders []*CartItemCreate
}

// Save creates the CartItem entities in the database.
func (cicb *CartItemCreateBulk) Save(ctx context.Context) ([]*CartItem, error) {
	if cicb.err != nil {
		return nil, cicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cicb.builders))
	nodes := make([]*CartItem, len(cicb.builders))
	mutators := make([]Mutator, len(cicb.builders))
	for i := range cicb.builders {
		func(i int, root context.Context) {
			builder := cicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CartItemMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cicb *CartItemCreateBulk) SaveX(ctx context.Context) []*CartItem {
	v, err := cicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cicb *CartItemCreateBulk) Exec(ctx context.Context) error {
	_, err := cicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cicb *CartItemCreateBulk) ExecX(ctx context.Context) {
	if err := cicb.Exec(ctx); err != nil {
		panic(err)
	}
}
