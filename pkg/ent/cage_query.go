// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent/cage"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent/dinosaur"
	"github.com/samsonannan/prizepicks-assessment/pkg/ent/predicate"
)

// CageQuery is the builder for querying Cage entities.
type CageQuery struct {
	config
	ctx           *QueryContext
	order         []cage.OrderOption
	inters        []Interceptor
	predicates    []predicate.Cage
	withDinosaurs *DinosaurQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CageQuery builder.
func (cq *CageQuery) Where(ps ...predicate.Cage) *CageQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CageQuery) Limit(limit int) *CageQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CageQuery) Offset(offset int) *CageQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CageQuery) Unique(unique bool) *CageQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CageQuery) Order(o ...cage.OrderOption) *CageQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryDinosaurs chains the current query on the "dinosaurs" edge.
func (cq *CageQuery) QueryDinosaurs() *DinosaurQuery {
	query := (&DinosaurClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cage.Table, cage.FieldID, selector),
			sqlgraph.To(dinosaur.Table, dinosaur.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, cage.DinosaursTable, cage.DinosaursColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Cage entity from the query.
// Returns a *NotFoundError when no Cage was found.
func (cq *CageQuery) First(ctx context.Context) (*Cage, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CageQuery) FirstX(ctx context.Context) *Cage {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Cage ID from the query.
// Returns a *NotFoundError when no Cage ID was found.
func (cq *CageQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CageQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Cage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Cage entity is found.
// Returns a *NotFoundError when no Cage entities are found.
func (cq *CageQuery) Only(ctx context.Context) (*Cage, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cage.Label}
	default:
		return nil, &NotSingularError{cage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CageQuery) OnlyX(ctx context.Context) *Cage {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Cage ID in the query.
// Returns a *NotSingularError when more than one Cage ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CageQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cage.Label}
	default:
		err = &NotSingularError{cage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CageQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Cages.
func (cq *CageQuery) All(ctx context.Context) ([]*Cage, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Cage, *CageQuery]()
	return withInterceptors[[]*Cage](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CageQuery) AllX(ctx context.Context) []*Cage {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Cage IDs.
func (cq *CageQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(cage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CageQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CageQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CageQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CageQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CageQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CageQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CageQuery) Clone() *CageQuery {
	if cq == nil {
		return nil
	}
	return &CageQuery{
		config:        cq.config,
		ctx:           cq.ctx.Clone(),
		order:         append([]cage.OrderOption{}, cq.order...),
		inters:        append([]Interceptor{}, cq.inters...),
		predicates:    append([]predicate.Cage{}, cq.predicates...),
		withDinosaurs: cq.withDinosaurs.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithDinosaurs tells the query-builder to eager-load the nodes that are connected to
// the "dinosaurs" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CageQuery) WithDinosaurs(opts ...func(*DinosaurQuery)) *CageQuery {
	query := (&DinosaurClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withDinosaurs = query
	return cq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Cage.Query().
//		GroupBy(cage.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CageQuery) GroupBy(field string, fields ...string) *CageGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CageGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = cage.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Cage.Query().
//		Select(cage.FieldCreatedAt).
//		Scan(ctx, &v)
func (cq *CageQuery) Select(fields ...string) *CageSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CageSelect{CageQuery: cq}
	sbuild.label = cage.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CageSelect configured with the given aggregations.
func (cq *CageQuery) Aggregate(fns ...AggregateFunc) *CageSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CageQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !cage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Cage, error) {
	var (
		nodes       = []*Cage{}
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withDinosaurs != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Cage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Cage{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withDinosaurs; query != nil {
		if err := cq.loadDinosaurs(ctx, query, nodes,
			func(n *Cage) { n.Edges.Dinosaurs = []*Dinosaur{} },
			func(n *Cage, e *Dinosaur) { n.Edges.Dinosaurs = append(n.Edges.Dinosaurs, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CageQuery) loadDinosaurs(ctx context.Context, query *DinosaurQuery, nodes []*Cage, init func(*Cage), assign func(*Cage, *Dinosaur)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Cage)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Dinosaur(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(cage.DinosaursColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.cage_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "cage_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "cage_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *CageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(cage.Table, cage.Columns, sqlgraph.NewFieldSpec(cage.FieldID, field.TypeUUID))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cage.FieldID)
		for i := range fields {
			if fields[i] != cage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(cage.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = cage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CageGroupBy is the group-by builder for Cage entities.
type CageGroupBy struct {
	selector
	build *CageQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CageGroupBy) Aggregate(fns ...AggregateFunc) *CageGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CageGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CageQuery, *CageGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CageGroupBy) sqlScan(ctx context.Context, root *CageQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CageSelect is the builder for selecting fields of Cage entities.
type CageSelect struct {
	*CageQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CageSelect) Aggregate(fns ...AggregateFunc) *CageSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CageSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CageQuery, *CageSelect](ctx, cs.CageQuery, cs, cs.inters, v)
}

func (cs *CageSelect) sqlScan(ctx context.Context, root *CageQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}