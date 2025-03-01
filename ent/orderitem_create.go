// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mall/ent/item"
	"mall/ent/order"
	"mall/ent/orderitem"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderItemCreate is the builder for creating a OrderItem entity.
type OrderItemCreate struct {
	config
	mutation *OrderItemMutation
	hooks    []Hook
}

// SetQuantity sets the "quantity" field.
func (oic *OrderItemCreate) SetQuantity(i int) *OrderItemCreate {
	oic.mutation.SetQuantity(i)
	return oic
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableQuantity(i *int) *OrderItemCreate {
	if i != nil {
		oic.SetQuantity(*i)
	}
	return oic
}

// SetPrice sets the "price" field.
func (oic *OrderItemCreate) SetPrice(f float32) *OrderItemCreate {
	oic.mutation.SetPrice(f)
	return oic
}

// SetOrderID sets the "order" edge to the Order entity by ID.
func (oic *OrderItemCreate) SetOrderID(id int) *OrderItemCreate {
	oic.mutation.SetOrderID(id)
	return oic
}

// SetOrder sets the "order" edge to the Order entity.
func (oic *OrderItemCreate) SetOrder(o *Order) *OrderItemCreate {
	return oic.SetOrderID(o.ID)
}

// SetItemID sets the "item" edge to the Item entity by ID.
func (oic *OrderItemCreate) SetItemID(id int) *OrderItemCreate {
	oic.mutation.SetItemID(id)
	return oic
}

// SetItem sets the "item" edge to the Item entity.
func (oic *OrderItemCreate) SetItem(i *Item) *OrderItemCreate {
	return oic.SetItemID(i.ID)
}

// Mutation returns the OrderItemMutation object of the builder.
func (oic *OrderItemCreate) Mutation() *OrderItemMutation {
	return oic.mutation
}

// Save creates the OrderItem in the database.
func (oic *OrderItemCreate) Save(ctx context.Context) (*OrderItem, error) {
	oic.defaults()
	return withHooks(ctx, oic.sqlSave, oic.mutation, oic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oic *OrderItemCreate) SaveX(ctx context.Context) *OrderItem {
	v, err := oic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oic *OrderItemCreate) Exec(ctx context.Context) error {
	_, err := oic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oic *OrderItemCreate) ExecX(ctx context.Context) {
	if err := oic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oic *OrderItemCreate) defaults() {
	if _, ok := oic.mutation.Quantity(); !ok {
		v := orderitem.DefaultQuantity
		oic.mutation.SetQuantity(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oic *OrderItemCreate) check() error {
	if _, ok := oic.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "OrderItem.quantity"`)}
	}
	if v, ok := oic.mutation.Quantity(); ok {
		if err := orderitem.QuantityValidator(v); err != nil {
			return &ValidationError{Name: "quantity", err: fmt.Errorf(`ent: validator failed for field "OrderItem.quantity": %w`, err)}
		}
	}
	if _, ok := oic.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "OrderItem.price"`)}
	}
	if v, ok := oic.mutation.Price(); ok {
		if err := orderitem.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf(`ent: validator failed for field "OrderItem.price": %w`, err)}
		}
	}
	if len(oic.mutation.OrderIDs()) == 0 {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required edge "OrderItem.order"`)}
	}
	if len(oic.mutation.ItemIDs()) == 0 {
		return &ValidationError{Name: "item", err: errors.New(`ent: missing required edge "OrderItem.item"`)}
	}
	return nil
}

func (oic *OrderItemCreate) sqlSave(ctx context.Context) (*OrderItem, error) {
	if err := oic.check(); err != nil {
		return nil, err
	}
	_node, _spec := oic.createSpec()
	if err := sqlgraph.CreateNode(ctx, oic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	oic.mutation.id = &_node.ID
	oic.mutation.done = true
	return _node, nil
}

func (oic *OrderItemCreate) createSpec() (*OrderItem, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderItem{config: oic.config}
		_spec = sqlgraph.NewCreateSpec(orderitem.Table, sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeInt))
	)
	if value, ok := oic.mutation.Quantity(); ok {
		_spec.SetField(orderitem.FieldQuantity, field.TypeInt, value)
		_node.Quantity = value
	}
	if value, ok := oic.mutation.Price(); ok {
		_spec.SetField(orderitem.FieldPrice, field.TypeFloat32, value)
		_node.Price = value
	}
	if nodes := oic.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.OrderTable,
			Columns: []string{orderitem.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.order_items = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oic.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   orderitem.ItemTable,
			Columns: []string{orderitem.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.order_item_item = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderItemCreateBulk is the builder for creating many OrderItem entities in bulk.
type OrderItemCreateBulk struct {
	config
	err      error
	builders []*OrderItemCreate
}

// Save creates the OrderItem entities in the database.
func (oicb *OrderItemCreateBulk) Save(ctx context.Context) ([]*OrderItem, error) {
	if oicb.err != nil {
		return nil, oicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(oicb.builders))
	nodes := make([]*OrderItem, len(oicb.builders))
	mutators := make([]Mutator, len(oicb.builders))
	for i := range oicb.builders {
		func(i int, root context.Context) {
			builder := oicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderItemMutation)
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
					_, err = mutators[i+1].Mutate(root, oicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oicb *OrderItemCreateBulk) SaveX(ctx context.Context) []*OrderItem {
	v, err := oicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oicb *OrderItemCreateBulk) Exec(ctx context.Context) error {
	_, err := oicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oicb *OrderItemCreateBulk) ExecX(ctx context.Context) {
	if err := oicb.Exec(ctx); err != nil {
		panic(err)
	}
}
