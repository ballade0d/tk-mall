// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"mall/ent/cart"
	"mall/ent/cartitem"
	"mall/ent/item"
	"mall/ent/predicate"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CartItemQuery is the builder for querying CartItem entities.
type CartItemQuery struct {
	config
	ctx        *QueryContext
	order      []cartitem.OrderOption
	inters     []Interceptor
	predicates []predicate.CartItem
	withCart   *CartQuery
	withItem   *ItemQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CartItemQuery builder.
func (ciq *CartItemQuery) Where(ps ...predicate.CartItem) *CartItemQuery {
	ciq.predicates = append(ciq.predicates, ps...)
	return ciq
}

// Limit the number of records to be returned by this query.
func (ciq *CartItemQuery) Limit(limit int) *CartItemQuery {
	ciq.ctx.Limit = &limit
	return ciq
}

// Offset to start from.
func (ciq *CartItemQuery) Offset(offset int) *CartItemQuery {
	ciq.ctx.Offset = &offset
	return ciq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ciq *CartItemQuery) Unique(unique bool) *CartItemQuery {
	ciq.ctx.Unique = &unique
	return ciq
}

// Order specifies how the records should be ordered.
func (ciq *CartItemQuery) Order(o ...cartitem.OrderOption) *CartItemQuery {
	ciq.order = append(ciq.order, o...)
	return ciq
}

// QueryCart chains the current query on the "cart" edge.
func (ciq *CartItemQuery) QueryCart() *CartQuery {
	query := (&CartClient{config: ciq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ciq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cartitem.Table, cartitem.FieldID, selector),
			sqlgraph.To(cart.Table, cart.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cartitem.CartTable, cartitem.CartColumn),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryItem chains the current query on the "item" edge.
func (ciq *CartItemQuery) QueryItem() *ItemQuery {
	query := (&ItemClient{config: ciq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ciq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cartitem.Table, cartitem.FieldID, selector),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, cartitem.ItemTable, cartitem.ItemColumn),
		)
		fromU = sqlgraph.SetNeighbors(ciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CartItem entity from the query.
// Returns a *NotFoundError when no CartItem was found.
func (ciq *CartItemQuery) First(ctx context.Context) (*CartItem, error) {
	nodes, err := ciq.Limit(1).All(setContextOp(ctx, ciq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cartitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ciq *CartItemQuery) FirstX(ctx context.Context) *CartItem {
	node, err := ciq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CartItem ID from the query.
// Returns a *NotFoundError when no CartItem ID was found.
func (ciq *CartItemQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ciq.Limit(1).IDs(setContextOp(ctx, ciq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cartitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ciq *CartItemQuery) FirstIDX(ctx context.Context) int {
	id, err := ciq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CartItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CartItem entity is found.
// Returns a *NotFoundError when no CartItem entities are found.
func (ciq *CartItemQuery) Only(ctx context.Context) (*CartItem, error) {
	nodes, err := ciq.Limit(2).All(setContextOp(ctx, ciq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cartitem.Label}
	default:
		return nil, &NotSingularError{cartitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ciq *CartItemQuery) OnlyX(ctx context.Context) *CartItem {
	node, err := ciq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CartItem ID in the query.
// Returns a *NotSingularError when more than one CartItem ID is found.
// Returns a *NotFoundError when no entities are found.
func (ciq *CartItemQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ciq.Limit(2).IDs(setContextOp(ctx, ciq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cartitem.Label}
	default:
		err = &NotSingularError{cartitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ciq *CartItemQuery) OnlyIDX(ctx context.Context) int {
	id, err := ciq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CartItems.
func (ciq *CartItemQuery) All(ctx context.Context) ([]*CartItem, error) {
	ctx = setContextOp(ctx, ciq.ctx, ent.OpQueryAll)
	if err := ciq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CartItem, *CartItemQuery]()
	return withInterceptors[[]*CartItem](ctx, ciq, qr, ciq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ciq *CartItemQuery) AllX(ctx context.Context) []*CartItem {
	nodes, err := ciq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CartItem IDs.
func (ciq *CartItemQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ciq.ctx.Unique == nil && ciq.path != nil {
		ciq.Unique(true)
	}
	ctx = setContextOp(ctx, ciq.ctx, ent.OpQueryIDs)
	if err = ciq.Select(cartitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ciq *CartItemQuery) IDsX(ctx context.Context) []int {
	ids, err := ciq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ciq *CartItemQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ciq.ctx, ent.OpQueryCount)
	if err := ciq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ciq, querierCount[*CartItemQuery](), ciq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ciq *CartItemQuery) CountX(ctx context.Context) int {
	count, err := ciq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ciq *CartItemQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ciq.ctx, ent.OpQueryExist)
	switch _, err := ciq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ciq *CartItemQuery) ExistX(ctx context.Context) bool {
	exist, err := ciq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CartItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ciq *CartItemQuery) Clone() *CartItemQuery {
	if ciq == nil {
		return nil
	}
	return &CartItemQuery{
		config:     ciq.config,
		ctx:        ciq.ctx.Clone(),
		order:      append([]cartitem.OrderOption{}, ciq.order...),
		inters:     append([]Interceptor{}, ciq.inters...),
		predicates: append([]predicate.CartItem{}, ciq.predicates...),
		withCart:   ciq.withCart.Clone(),
		withItem:   ciq.withItem.Clone(),
		// clone intermediate query.
		sql:  ciq.sql.Clone(),
		path: ciq.path,
	}
}

// WithCart tells the query-builder to eager-load the nodes that are connected to
// the "cart" edge. The optional arguments are used to configure the query builder of the edge.
func (ciq *CartItemQuery) WithCart(opts ...func(*CartQuery)) *CartItemQuery {
	query := (&CartClient{config: ciq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ciq.withCart = query
	return ciq
}

// WithItem tells the query-builder to eager-load the nodes that are connected to
// the "item" edge. The optional arguments are used to configure the query builder of the edge.
func (ciq *CartItemQuery) WithItem(opts ...func(*ItemQuery)) *CartItemQuery {
	query := (&ItemClient{config: ciq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ciq.withItem = query
	return ciq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Quantity int `json:"quantity,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CartItem.Query().
//		GroupBy(cartitem.FieldQuantity).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ciq *CartItemQuery) GroupBy(field string, fields ...string) *CartItemGroupBy {
	ciq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CartItemGroupBy{build: ciq}
	grbuild.flds = &ciq.ctx.Fields
	grbuild.label = cartitem.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Quantity int `json:"quantity,omitempty"`
//	}
//
//	client.CartItem.Query().
//		Select(cartitem.FieldQuantity).
//		Scan(ctx, &v)
func (ciq *CartItemQuery) Select(fields ...string) *CartItemSelect {
	ciq.ctx.Fields = append(ciq.ctx.Fields, fields...)
	sbuild := &CartItemSelect{CartItemQuery: ciq}
	sbuild.label = cartitem.Label
	sbuild.flds, sbuild.scan = &ciq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CartItemSelect configured with the given aggregations.
func (ciq *CartItemQuery) Aggregate(fns ...AggregateFunc) *CartItemSelect {
	return ciq.Select().Aggregate(fns...)
}

func (ciq *CartItemQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ciq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ciq); err != nil {
				return err
			}
		}
	}
	for _, f := range ciq.ctx.Fields {
		if !cartitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ciq.path != nil {
		prev, err := ciq.path(ctx)
		if err != nil {
			return err
		}
		ciq.sql = prev
	}
	return nil
}

func (ciq *CartItemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CartItem, error) {
	var (
		nodes       = []*CartItem{}
		withFKs     = ciq.withFKs
		_spec       = ciq.querySpec()
		loadedTypes = [2]bool{
			ciq.withCart != nil,
			ciq.withItem != nil,
		}
	)
	if ciq.withCart != nil || ciq.withItem != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, cartitem.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CartItem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CartItem{config: ciq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ciq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ciq.withCart; query != nil {
		if err := ciq.loadCart(ctx, query, nodes, nil,
			func(n *CartItem, e *Cart) { n.Edges.Cart = e }); err != nil {
			return nil, err
		}
	}
	if query := ciq.withItem; query != nil {
		if err := ciq.loadItem(ctx, query, nodes, nil,
			func(n *CartItem, e *Item) { n.Edges.Item = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ciq *CartItemQuery) loadCart(ctx context.Context, query *CartQuery, nodes []*CartItem, init func(*CartItem), assign func(*CartItem, *Cart)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*CartItem)
	for i := range nodes {
		if nodes[i].cart_items == nil {
			continue
		}
		fk := *nodes[i].cart_items
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(cart.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "cart_items" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ciq *CartItemQuery) loadItem(ctx context.Context, query *ItemQuery, nodes []*CartItem, init func(*CartItem), assign func(*CartItem, *Item)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*CartItem)
	for i := range nodes {
		if nodes[i].cart_item_item == nil {
			continue
		}
		fk := *nodes[i].cart_item_item
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(item.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "cart_item_item" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ciq *CartItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ciq.querySpec()
	_spec.Node.Columns = ciq.ctx.Fields
	if len(ciq.ctx.Fields) > 0 {
		_spec.Unique = ciq.ctx.Unique != nil && *ciq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ciq.driver, _spec)
}

func (ciq *CartItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(cartitem.Table, cartitem.Columns, sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeInt))
	_spec.From = ciq.sql
	if unique := ciq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ciq.path != nil {
		_spec.Unique = true
	}
	if fields := ciq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cartitem.FieldID)
		for i := range fields {
			if fields[i] != cartitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ciq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ciq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ciq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ciq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ciq *CartItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ciq.driver.Dialect())
	t1 := builder.Table(cartitem.Table)
	columns := ciq.ctx.Fields
	if len(columns) == 0 {
		columns = cartitem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ciq.sql != nil {
		selector = ciq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ciq.ctx.Unique != nil && *ciq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ciq.predicates {
		p(selector)
	}
	for _, p := range ciq.order {
		p(selector)
	}
	if offset := ciq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ciq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CartItemGroupBy is the group-by builder for CartItem entities.
type CartItemGroupBy struct {
	selector
	build *CartItemQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cigb *CartItemGroupBy) Aggregate(fns ...AggregateFunc) *CartItemGroupBy {
	cigb.fns = append(cigb.fns, fns...)
	return cigb
}

// Scan applies the selector query and scans the result into the given value.
func (cigb *CartItemGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cigb.build.ctx, ent.OpQueryGroupBy)
	if err := cigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CartItemQuery, *CartItemGroupBy](ctx, cigb.build, cigb, cigb.build.inters, v)
}

func (cigb *CartItemGroupBy) sqlScan(ctx context.Context, root *CartItemQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cigb.fns))
	for _, fn := range cigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cigb.flds)+len(cigb.fns))
		for _, f := range *cigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CartItemSelect is the builder for selecting fields of CartItem entities.
type CartItemSelect struct {
	*CartItemQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cis *CartItemSelect) Aggregate(fns ...AggregateFunc) *CartItemSelect {
	cis.fns = append(cis.fns, fns...)
	return cis
}

// Scan applies the selector query and scans the result into the given value.
func (cis *CartItemSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cis.ctx, ent.OpQuerySelect)
	if err := cis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CartItemQuery, *CartItemSelect](ctx, cis.CartItemQuery, cis, cis.inters, v)
}

func (cis *CartItemSelect) sqlScan(ctx context.Context, root *CartItemQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cis.fns))
	for _, fn := range cis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
