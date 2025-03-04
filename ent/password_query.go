// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"mall/ent/password"
	"mall/ent/predicate"
	"mall/ent/user"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PasswordQuery is the builder for querying Password entities.
type PasswordQuery struct {
	config
	ctx        *QueryContext
	order      []password.OrderOption
	inters     []Interceptor
	predicates []predicate.Password
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PasswordQuery builder.
func (pq *PasswordQuery) Where(ps ...predicate.Password) *PasswordQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PasswordQuery) Limit(limit int) *PasswordQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PasswordQuery) Offset(offset int) *PasswordQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PasswordQuery) Unique(unique bool) *PasswordQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PasswordQuery) Order(o ...password.OrderOption) *PasswordQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryUser chains the current query on the "user" edge.
func (pq *PasswordQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(password.Table, password.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, password.UserTable, password.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Password entity from the query.
// Returns a *NotFoundError when no Password was found.
func (pq *PasswordQuery) First(ctx context.Context) (*Password, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{password.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PasswordQuery) FirstX(ctx context.Context) *Password {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Password ID from the query.
// Returns a *NotFoundError when no Password ID was found.
func (pq *PasswordQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{password.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PasswordQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Password entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Password entity is found.
// Returns a *NotFoundError when no Password entities are found.
func (pq *PasswordQuery) Only(ctx context.Context) (*Password, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{password.Label}
	default:
		return nil, &NotSingularError{password.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PasswordQuery) OnlyX(ctx context.Context) *Password {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Password ID in the query.
// Returns a *NotSingularError when more than one Password ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PasswordQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{password.Label}
	default:
		err = &NotSingularError{password.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PasswordQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Passwords.
func (pq *PasswordQuery) All(ctx context.Context) ([]*Password, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryAll)
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Password, *PasswordQuery]()
	return withInterceptors[[]*Password](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PasswordQuery) AllX(ctx context.Context) []*Password {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Password IDs.
func (pq *PasswordQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryIDs)
	if err = pq.Select(password.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PasswordQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PasswordQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryCount)
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PasswordQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PasswordQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PasswordQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryExist)
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PasswordQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PasswordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PasswordQuery) Clone() *PasswordQuery {
	if pq == nil {
		return nil
	}
	return &PasswordQuery{
		config:     pq.config,
		ctx:        pq.ctx.Clone(),
		order:      append([]password.OrderOption{}, pq.order...),
		inters:     append([]Interceptor{}, pq.inters...),
		predicates: append([]predicate.Password{}, pq.predicates...),
		withUser:   pq.withUser.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PasswordQuery) WithUser(opts ...func(*UserQuery)) *PasswordQuery {
	query := (&UserClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withUser = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Password string `json:"password,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Password.Query().
//		GroupBy(password.FieldPassword).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PasswordQuery) GroupBy(field string, fields ...string) *PasswordGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PasswordGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = password.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Password string `json:"password,omitempty"`
//	}
//
//	client.Password.Query().
//		Select(password.FieldPassword).
//		Scan(ctx, &v)
func (pq *PasswordQuery) Select(fields ...string) *PasswordSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PasswordSelect{PasswordQuery: pq}
	sbuild.label = password.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PasswordSelect configured with the given aggregations.
func (pq *PasswordQuery) Aggregate(fns ...AggregateFunc) *PasswordSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PasswordQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !password.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PasswordQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Password, error) {
	var (
		nodes       = []*Password{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withUser != nil,
		}
	)
	if pq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, password.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Password).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Password{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withUser; query != nil {
		if err := pq.loadUser(ctx, query, nodes, nil,
			func(n *Password, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PasswordQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Password, init func(*Password), assign func(*Password, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Password)
	for i := range nodes {
		if nodes[i].user_password == nil {
			continue
		}
		fk := *nodes[i].user_password
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_password" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *PasswordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PasswordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(password.Table, password.Columns, sqlgraph.NewFieldSpec(password.FieldID, field.TypeInt))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, password.FieldID)
		for i := range fields {
			if fields[i] != password.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PasswordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(password.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = password.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PasswordGroupBy is the group-by builder for Password entities.
type PasswordGroupBy struct {
	selector
	build *PasswordQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PasswordGroupBy) Aggregate(fns ...AggregateFunc) *PasswordGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PasswordGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, ent.OpQueryGroupBy)
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PasswordQuery, *PasswordGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PasswordGroupBy) sqlScan(ctx context.Context, root *PasswordQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PasswordSelect is the builder for selecting fields of Password entities.
type PasswordSelect struct {
	*PasswordQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PasswordSelect) Aggregate(fns ...AggregateFunc) *PasswordSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PasswordSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, ent.OpQuerySelect)
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PasswordQuery, *PasswordSelect](ctx, ps.PasswordQuery, ps, ps.inters, v)
}

func (ps *PasswordSelect) sqlScan(ctx context.Context, root *PasswordQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
