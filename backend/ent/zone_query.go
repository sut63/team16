// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/G16/app/ent/equipment"
	"github.com/G16/app/ent/predicate"
	"github.com/G16/app/ent/zone"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ZoneQuery is the builder for querying Zone entities.
type ZoneQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Zone
	// eager-loading edges.
	withEquipment *EquipmentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (zq *ZoneQuery) Where(ps ...predicate.Zone) *ZoneQuery {
	zq.predicates = append(zq.predicates, ps...)
	return zq
}

// Limit adds a limit step to the query.
func (zq *ZoneQuery) Limit(limit int) *ZoneQuery {
	zq.limit = &limit
	return zq
}

// Offset adds an offset step to the query.
func (zq *ZoneQuery) Offset(offset int) *ZoneQuery {
	zq.offset = &offset
	return zq
}

// Order adds an order step to the query.
func (zq *ZoneQuery) Order(o ...OrderFunc) *ZoneQuery {
	zq.order = append(zq.order, o...)
	return zq
}

// QueryEquipment chains the current query on the equipment edge.
func (zq *ZoneQuery) QueryEquipment() *EquipmentQuery {
	query := &EquipmentQuery{config: zq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := zq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(zone.Table, zone.FieldID, zq.sqlQuery()),
			sqlgraph.To(equipment.Table, equipment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, zone.EquipmentTable, zone.EquipmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(zq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Zone entity in the query. Returns *NotFoundError when no zone was found.
func (zq *ZoneQuery) First(ctx context.Context) (*Zone, error) {
	zs, err := zq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(zs) == 0 {
		return nil, &NotFoundError{zone.Label}
	}
	return zs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (zq *ZoneQuery) FirstX(ctx context.Context) *Zone {
	z, err := zq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return z
}

// FirstID returns the first Zone id in the query. Returns *NotFoundError when no id was found.
func (zq *ZoneQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = zq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{zone.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (zq *ZoneQuery) FirstXID(ctx context.Context) int {
	id, err := zq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Zone entity in the query, returns an error if not exactly one entity was returned.
func (zq *ZoneQuery) Only(ctx context.Context) (*Zone, error) {
	zs, err := zq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(zs) {
	case 1:
		return zs[0], nil
	case 0:
		return nil, &NotFoundError{zone.Label}
	default:
		return nil, &NotSingularError{zone.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (zq *ZoneQuery) OnlyX(ctx context.Context) *Zone {
	z, err := zq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return z
}

// OnlyID returns the only Zone id in the query, returns an error if not exactly one id was returned.
func (zq *ZoneQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = zq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{zone.Label}
	default:
		err = &NotSingularError{zone.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (zq *ZoneQuery) OnlyIDX(ctx context.Context) int {
	id, err := zq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Zones.
func (zq *ZoneQuery) All(ctx context.Context) ([]*Zone, error) {
	if err := zq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return zq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (zq *ZoneQuery) AllX(ctx context.Context) []*Zone {
	zs, err := zq.All(ctx)
	if err != nil {
		panic(err)
	}
	return zs
}

// IDs executes the query and returns a list of Zone ids.
func (zq *ZoneQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := zq.Select(zone.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (zq *ZoneQuery) IDsX(ctx context.Context) []int {
	ids, err := zq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (zq *ZoneQuery) Count(ctx context.Context) (int, error) {
	if err := zq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return zq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (zq *ZoneQuery) CountX(ctx context.Context) int {
	count, err := zq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (zq *ZoneQuery) Exist(ctx context.Context) (bool, error) {
	if err := zq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return zq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (zq *ZoneQuery) ExistX(ctx context.Context) bool {
	exist, err := zq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (zq *ZoneQuery) Clone() *ZoneQuery {
	return &ZoneQuery{
		config:     zq.config,
		limit:      zq.limit,
		offset:     zq.offset,
		order:      append([]OrderFunc{}, zq.order...),
		unique:     append([]string{}, zq.unique...),
		predicates: append([]predicate.Zone{}, zq.predicates...),
		// clone intermediate query.
		sql:  zq.sql.Clone(),
		path: zq.path,
	}
}

//  WithEquipment tells the query-builder to eager-loads the nodes that are connected to
// the "equipment" edge. The optional arguments used to configure the query builder of the edge.
func (zq *ZoneQuery) WithEquipment(opts ...func(*EquipmentQuery)) *ZoneQuery {
	query := &EquipmentQuery{config: zq.config}
	for _, opt := range opts {
		opt(query)
	}
	zq.withEquipment = query
	return zq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EQUIPMENTZONE string `json:"EQUIPMENTZONE,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Zone.Query().
//		GroupBy(zone.FieldEQUIPMENTZONE).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (zq *ZoneQuery) GroupBy(field string, fields ...string) *ZoneGroupBy {
	group := &ZoneGroupBy{config: zq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := zq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return zq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		EQUIPMENTZONE string `json:"EQUIPMENTZONE,omitempty"`
//	}
//
//	client.Zone.Query().
//		Select(zone.FieldEQUIPMENTZONE).
//		Scan(ctx, &v)
//
func (zq *ZoneQuery) Select(field string, fields ...string) *ZoneSelect {
	selector := &ZoneSelect{config: zq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := zq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return zq.sqlQuery(), nil
	}
	return selector
}

func (zq *ZoneQuery) prepareQuery(ctx context.Context) error {
	if zq.path != nil {
		prev, err := zq.path(ctx)
		if err != nil {
			return err
		}
		zq.sql = prev
	}
	return nil
}

func (zq *ZoneQuery) sqlAll(ctx context.Context) ([]*Zone, error) {
	var (
		nodes       = []*Zone{}
		_spec       = zq.querySpec()
		loadedTypes = [1]bool{
			zq.withEquipment != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Zone{config: zq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, zq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := zq.withEquipment; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Zone)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Equipment(func(s *sql.Selector) {
			s.Where(sql.InValues(zone.EquipmentColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.zone_equipment
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "zone_equipment" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "zone_equipment" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Equipment = append(node.Edges.Equipment, n)
		}
	}

	return nodes, nil
}

func (zq *ZoneQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := zq.querySpec()
	return sqlgraph.CountNodes(ctx, zq.driver, _spec)
}

func (zq *ZoneQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := zq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (zq *ZoneQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   zone.Table,
			Columns: zone.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: zone.FieldID,
			},
		},
		From:   zq.sql,
		Unique: true,
	}
	if ps := zq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := zq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := zq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := zq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (zq *ZoneQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(zq.driver.Dialect())
	t1 := builder.Table(zone.Table)
	selector := builder.Select(t1.Columns(zone.Columns...)...).From(t1)
	if zq.sql != nil {
		selector = zq.sql
		selector.Select(selector.Columns(zone.Columns...)...)
	}
	for _, p := range zq.predicates {
		p(selector)
	}
	for _, p := range zq.order {
		p(selector)
	}
	if offset := zq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := zq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ZoneGroupBy is the builder for group-by Zone entities.
type ZoneGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (zgb *ZoneGroupBy) Aggregate(fns ...AggregateFunc) *ZoneGroupBy {
	zgb.fns = append(zgb.fns, fns...)
	return zgb
}

// Scan applies the group-by query and scan the result into the given value.
func (zgb *ZoneGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := zgb.path(ctx)
	if err != nil {
		return err
	}
	zgb.sql = query
	return zgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (zgb *ZoneGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := zgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (zgb *ZoneGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(zgb.fields) > 1 {
		return nil, errors.New("ent: ZoneGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := zgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (zgb *ZoneGroupBy) StringsX(ctx context.Context) []string {
	v, err := zgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (zgb *ZoneGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = zgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{zone.Label}
	default:
		err = fmt.Errorf("ent: ZoneGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (zgb *ZoneGroupBy) StringX(ctx context.Context) string {
	v, err := zgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (zgb *ZoneGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(zgb.fields) > 1 {
		return nil, errors.New("ent: ZoneGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := zgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (zgb *ZoneGroupBy) IntsX(ctx context.Context) []int {
	v, err := zgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (zgb *ZoneGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = zgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{zone.Label}
	default:
		err = fmt.Errorf("ent: ZoneGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (zgb *ZoneGroupBy) IntX(ctx context.Context) int {
	v, err := zgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (zgb *ZoneGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(zgb.fields) > 1 {
		return nil, errors.New("ent: ZoneGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := zgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (zgb *ZoneGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := zgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (zgb *ZoneGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = zgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{zone.Label}
	default:
		err = fmt.Errorf("ent: ZoneGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (zgb *ZoneGroupBy) Float64X(ctx context.Context) float64 {
	v, err := zgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (zgb *ZoneGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(zgb.fields) > 1 {
		return nil, errors.New("ent: ZoneGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := zgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (zgb *ZoneGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := zgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (zgb *ZoneGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = zgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{zone.Label}
	default:
		err = fmt.Errorf("ent: ZoneGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (zgb *ZoneGroupBy) BoolX(ctx context.Context) bool {
	v, err := zgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (zgb *ZoneGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := zgb.sqlQuery().Query()
	if err := zgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (zgb *ZoneGroupBy) sqlQuery() *sql.Selector {
	selector := zgb.sql
	columns := make([]string, 0, len(zgb.fields)+len(zgb.fns))
	columns = append(columns, zgb.fields...)
	for _, fn := range zgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(zgb.fields...)
}

// ZoneSelect is the builder for select fields of Zone entities.
type ZoneSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (zs *ZoneSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := zs.path(ctx)
	if err != nil {
		return err
	}
	zs.sql = query
	return zs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (zs *ZoneSelect) ScanX(ctx context.Context, v interface{}) {
	if err := zs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (zs *ZoneSelect) Strings(ctx context.Context) ([]string, error) {
	if len(zs.fields) > 1 {
		return nil, errors.New("ent: ZoneSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := zs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (zs *ZoneSelect) StringsX(ctx context.Context) []string {
	v, err := zs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (zs *ZoneSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = zs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{zone.Label}
	default:
		err = fmt.Errorf("ent: ZoneSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (zs *ZoneSelect) StringX(ctx context.Context) string {
	v, err := zs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (zs *ZoneSelect) Ints(ctx context.Context) ([]int, error) {
	if len(zs.fields) > 1 {
		return nil, errors.New("ent: ZoneSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := zs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (zs *ZoneSelect) IntsX(ctx context.Context) []int {
	v, err := zs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (zs *ZoneSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = zs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{zone.Label}
	default:
		err = fmt.Errorf("ent: ZoneSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (zs *ZoneSelect) IntX(ctx context.Context) int {
	v, err := zs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (zs *ZoneSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(zs.fields) > 1 {
		return nil, errors.New("ent: ZoneSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := zs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (zs *ZoneSelect) Float64sX(ctx context.Context) []float64 {
	v, err := zs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (zs *ZoneSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = zs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{zone.Label}
	default:
		err = fmt.Errorf("ent: ZoneSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (zs *ZoneSelect) Float64X(ctx context.Context) float64 {
	v, err := zs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (zs *ZoneSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(zs.fields) > 1 {
		return nil, errors.New("ent: ZoneSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := zs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (zs *ZoneSelect) BoolsX(ctx context.Context) []bool {
	v, err := zs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (zs *ZoneSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = zs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{zone.Label}
	default:
		err = fmt.Errorf("ent: ZoneSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (zs *ZoneSelect) BoolX(ctx context.Context) bool {
	v, err := zs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (zs *ZoneSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := zs.sqlQuery().Query()
	if err := zs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (zs *ZoneSelect) sqlQuery() sql.Querier {
	selector := zs.sql
	selector.Select(selector.Columns(zs.fields...)...)
	return selector
}
