// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/Varissara/app/ent/educationlevel"
	"github.com/Varissara/app/ent/predicate"
	"github.com/Varissara/app/ent/scholarship"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// EducationlevelQuery is the builder for querying Educationlevel entities.
type EducationlevelQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Educationlevel
	// eager-loading edges.
	withEducationlevelID *ScholarshipQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (eq *EducationlevelQuery) Where(ps ...predicate.Educationlevel) *EducationlevelQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit adds a limit step to the query.
func (eq *EducationlevelQuery) Limit(limit int) *EducationlevelQuery {
	eq.limit = &limit
	return eq
}

// Offset adds an offset step to the query.
func (eq *EducationlevelQuery) Offset(offset int) *EducationlevelQuery {
	eq.offset = &offset
	return eq
}

// Order adds an order step to the query.
func (eq *EducationlevelQuery) Order(o ...OrderFunc) *EducationlevelQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// QueryEducationlevelID chains the current query on the EducationlevelID edge.
func (eq *EducationlevelQuery) QueryEducationlevelID() *ScholarshipQuery {
	query := &ScholarshipQuery{config: eq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(educationlevel.Table, educationlevel.FieldID, eq.sqlQuery()),
			sqlgraph.To(scholarship.Table, scholarship.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, educationlevel.EducationlevelIDTable, educationlevel.EducationlevelIDColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Educationlevel entity in the query. Returns *NotFoundError when no educationlevel was found.
func (eq *EducationlevelQuery) First(ctx context.Context) (*Educationlevel, error) {
	es, err := eq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(es) == 0 {
		return nil, &NotFoundError{educationlevel.Label}
	}
	return es[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *EducationlevelQuery) FirstX(ctx context.Context) *Educationlevel {
	e, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return e
}

// FirstID returns the first Educationlevel id in the query. Returns *NotFoundError when no id was found.
func (eq *EducationlevelQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{educationlevel.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (eq *EducationlevelQuery) FirstXID(ctx context.Context) int {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Educationlevel entity in the query, returns an error if not exactly one entity was returned.
func (eq *EducationlevelQuery) Only(ctx context.Context) (*Educationlevel, error) {
	es, err := eq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(es) {
	case 1:
		return es[0], nil
	case 0:
		return nil, &NotFoundError{educationlevel.Label}
	default:
		return nil, &NotSingularError{educationlevel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *EducationlevelQuery) OnlyX(ctx context.Context) *Educationlevel {
	e, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return e
}

// OnlyID returns the only Educationlevel id in the query, returns an error if not exactly one id was returned.
func (eq *EducationlevelQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = eq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{educationlevel.Label}
	default:
		err = &NotSingularError{educationlevel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *EducationlevelQuery) OnlyIDX(ctx context.Context) int {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Educationlevels.
func (eq *EducationlevelQuery) All(ctx context.Context) ([]*Educationlevel, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return eq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (eq *EducationlevelQuery) AllX(ctx context.Context) []*Educationlevel {
	es, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return es
}

// IDs executes the query and returns a list of Educationlevel ids.
func (eq *EducationlevelQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := eq.Select(educationlevel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *EducationlevelQuery) IDsX(ctx context.Context) []int {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *EducationlevelQuery) Count(ctx context.Context) (int, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return eq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (eq *EducationlevelQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *EducationlevelQuery) Exist(ctx context.Context) (bool, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return eq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *EducationlevelQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *EducationlevelQuery) Clone() *EducationlevelQuery {
	return &EducationlevelQuery{
		config:     eq.config,
		limit:      eq.limit,
		offset:     eq.offset,
		order:      append([]OrderFunc{}, eq.order...),
		unique:     append([]string{}, eq.unique...),
		predicates: append([]predicate.Educationlevel{}, eq.predicates...),
		// clone intermediate query.
		sql:  eq.sql.Clone(),
		path: eq.path,
	}
}

//  WithEducationlevelID tells the query-builder to eager-loads the nodes that are connected to
// the "EducationlevelID" edge. The optional arguments used to configure the query builder of the edge.
func (eq *EducationlevelQuery) WithEducationlevelID(opts ...func(*ScholarshipQuery)) *EducationlevelQuery {
	query := &ScholarshipQuery{config: eq.config}
	for _, opt := range opts {
		opt(query)
	}
	eq.withEducationlevelID = query
	return eq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Educationlevelname string `json:"Educationlevelname,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Educationlevel.Query().
//		GroupBy(educationlevel.FieldEducationlevelname).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (eq *EducationlevelQuery) GroupBy(field string, fields ...string) *EducationlevelGroupBy {
	group := &EducationlevelGroupBy{config: eq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return eq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Educationlevelname string `json:"Educationlevelname,omitempty"`
//	}
//
//	client.Educationlevel.Query().
//		Select(educationlevel.FieldEducationlevelname).
//		Scan(ctx, &v)
//
func (eq *EducationlevelQuery) Select(field string, fields ...string) *EducationlevelSelect {
	selector := &EducationlevelSelect{config: eq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return eq.sqlQuery(), nil
	}
	return selector
}

func (eq *EducationlevelQuery) prepareQuery(ctx context.Context) error {
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *EducationlevelQuery) sqlAll(ctx context.Context) ([]*Educationlevel, error) {
	var (
		nodes       = []*Educationlevel{}
		_spec       = eq.querySpec()
		loadedTypes = [1]bool{
			eq.withEducationlevelID != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Educationlevel{config: eq.config}
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
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := eq.withEducationlevelID; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Educationlevel)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Scholarship(func(s *sql.Selector) {
			s.Where(sql.InValues(educationlevel.EducationlevelIDColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.educationlevel_educationlevel_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "educationlevel_educationlevel_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "educationlevel_educationlevel_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.EducationlevelID = append(node.Edges.EducationlevelID, n)
		}
	}

	return nodes, nil
}

func (eq *EducationlevelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *EducationlevelQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := eq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (eq *EducationlevelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   educationlevel.Table,
			Columns: educationlevel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: educationlevel.FieldID,
			},
		},
		From:   eq.sql,
		Unique: true,
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eq *EducationlevelQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(educationlevel.Table)
	selector := builder.Select(t1.Columns(educationlevel.Columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(educationlevel.Columns...)...)
	}
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector)
	}
	if offset := eq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EducationlevelGroupBy is the builder for group-by Educationlevel entities.
type EducationlevelGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *EducationlevelGroupBy) Aggregate(fns ...AggregateFunc) *EducationlevelGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the group-by query and scan the result into the given value.
func (egb *EducationlevelGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := egb.path(ctx)
	if err != nil {
		return err
	}
	egb.sql = query
	return egb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (egb *EducationlevelGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := egb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (egb *EducationlevelGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(egb.fields) > 1 {
		return nil, errors.New("ent: EducationlevelGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := egb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (egb *EducationlevelGroupBy) StringsX(ctx context.Context) []string {
	v, err := egb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (egb *EducationlevelGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = egb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{educationlevel.Label}
	default:
		err = fmt.Errorf("ent: EducationlevelGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (egb *EducationlevelGroupBy) StringX(ctx context.Context) string {
	v, err := egb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (egb *EducationlevelGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(egb.fields) > 1 {
		return nil, errors.New("ent: EducationlevelGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := egb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (egb *EducationlevelGroupBy) IntsX(ctx context.Context) []int {
	v, err := egb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (egb *EducationlevelGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = egb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{educationlevel.Label}
	default:
		err = fmt.Errorf("ent: EducationlevelGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (egb *EducationlevelGroupBy) IntX(ctx context.Context) int {
	v, err := egb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (egb *EducationlevelGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(egb.fields) > 1 {
		return nil, errors.New("ent: EducationlevelGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := egb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (egb *EducationlevelGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := egb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (egb *EducationlevelGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = egb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{educationlevel.Label}
	default:
		err = fmt.Errorf("ent: EducationlevelGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (egb *EducationlevelGroupBy) Float64X(ctx context.Context) float64 {
	v, err := egb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (egb *EducationlevelGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(egb.fields) > 1 {
		return nil, errors.New("ent: EducationlevelGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := egb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (egb *EducationlevelGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := egb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (egb *EducationlevelGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = egb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{educationlevel.Label}
	default:
		err = fmt.Errorf("ent: EducationlevelGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (egb *EducationlevelGroupBy) BoolX(ctx context.Context) bool {
	v, err := egb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (egb *EducationlevelGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := egb.sqlQuery().Query()
	if err := egb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (egb *EducationlevelGroupBy) sqlQuery() *sql.Selector {
	selector := egb.sql
	columns := make([]string, 0, len(egb.fields)+len(egb.fns))
	columns = append(columns, egb.fields...)
	for _, fn := range egb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(egb.fields...)
}

// EducationlevelSelect is the builder for select fields of Educationlevel entities.
type EducationlevelSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (es *EducationlevelSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := es.path(ctx)
	if err != nil {
		return err
	}
	es.sql = query
	return es.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (es *EducationlevelSelect) ScanX(ctx context.Context, v interface{}) {
	if err := es.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (es *EducationlevelSelect) Strings(ctx context.Context) ([]string, error) {
	if len(es.fields) > 1 {
		return nil, errors.New("ent: EducationlevelSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := es.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (es *EducationlevelSelect) StringsX(ctx context.Context) []string {
	v, err := es.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (es *EducationlevelSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = es.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{educationlevel.Label}
	default:
		err = fmt.Errorf("ent: EducationlevelSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (es *EducationlevelSelect) StringX(ctx context.Context) string {
	v, err := es.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (es *EducationlevelSelect) Ints(ctx context.Context) ([]int, error) {
	if len(es.fields) > 1 {
		return nil, errors.New("ent: EducationlevelSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := es.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (es *EducationlevelSelect) IntsX(ctx context.Context) []int {
	v, err := es.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (es *EducationlevelSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = es.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{educationlevel.Label}
	default:
		err = fmt.Errorf("ent: EducationlevelSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (es *EducationlevelSelect) IntX(ctx context.Context) int {
	v, err := es.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (es *EducationlevelSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(es.fields) > 1 {
		return nil, errors.New("ent: EducationlevelSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := es.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (es *EducationlevelSelect) Float64sX(ctx context.Context) []float64 {
	v, err := es.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (es *EducationlevelSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = es.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{educationlevel.Label}
	default:
		err = fmt.Errorf("ent: EducationlevelSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (es *EducationlevelSelect) Float64X(ctx context.Context) float64 {
	v, err := es.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (es *EducationlevelSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(es.fields) > 1 {
		return nil, errors.New("ent: EducationlevelSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := es.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (es *EducationlevelSelect) BoolsX(ctx context.Context) []bool {
	v, err := es.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (es *EducationlevelSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = es.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{educationlevel.Label}
	default:
		err = fmt.Errorf("ent: EducationlevelSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (es *EducationlevelSelect) BoolX(ctx context.Context) bool {
	v, err := es.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (es *EducationlevelSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := es.sqlQuery().Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (es *EducationlevelSelect) sqlQuery() sql.Querier {
	selector := es.sql
	selector.Select(selector.Columns(es.fields...)...)
	return selector
}
