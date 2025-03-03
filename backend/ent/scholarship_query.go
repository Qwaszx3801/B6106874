// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/Varissara/app/ent/educationlevel"
	"github.com/Varissara/app/ent/predicate"
	"github.com/Varissara/app/ent/scholarship"
	"github.com/Varissara/app/ent/scholarshiptype"
	"github.com/Varissara/app/ent/semester"
	"github.com/Varissara/app/ent/studyplan"
	"github.com/Varissara/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ScholarshipQuery is the builder for querying Scholarship entities.
type ScholarshipQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Scholarship
	// eager-loading edges.
	withScholarshiptypeID *ScholarshiptypeQuery
	withEducationlevelID  *EducationlevelQuery
	withStudyplanID       *StudyplanQuery
	withSemesterID        *SemesterQuery
	withUserID            *UserQuery
	withFKs               bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (sq *ScholarshipQuery) Where(ps ...predicate.Scholarship) *ScholarshipQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit adds a limit step to the query.
func (sq *ScholarshipQuery) Limit(limit int) *ScholarshipQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *ScholarshipQuery) Offset(offset int) *ScholarshipQuery {
	sq.offset = &offset
	return sq
}

// Order adds an order step to the query.
func (sq *ScholarshipQuery) Order(o ...OrderFunc) *ScholarshipQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryScholarshiptypeID chains the current query on the ScholarshiptypeID edge.
func (sq *ScholarshipQuery) QueryScholarshiptypeID() *ScholarshiptypeQuery {
	query := &ScholarshiptypeQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scholarship.Table, scholarship.FieldID, sq.sqlQuery()),
			sqlgraph.To(scholarshiptype.Table, scholarshiptype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, scholarship.ScholarshiptypeIDTable, scholarship.ScholarshiptypeIDColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEducationlevelID chains the current query on the EducationlevelID edge.
func (sq *ScholarshipQuery) QueryEducationlevelID() *EducationlevelQuery {
	query := &EducationlevelQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scholarship.Table, scholarship.FieldID, sq.sqlQuery()),
			sqlgraph.To(educationlevel.Table, educationlevel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, scholarship.EducationlevelIDTable, scholarship.EducationlevelIDColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStudyplanID chains the current query on the StudyplanID edge.
func (sq *ScholarshipQuery) QueryStudyplanID() *StudyplanQuery {
	query := &StudyplanQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scholarship.Table, scholarship.FieldID, sq.sqlQuery()),
			sqlgraph.To(studyplan.Table, studyplan.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, scholarship.StudyplanIDTable, scholarship.StudyplanIDColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySemesterID chains the current query on the SemesterID edge.
func (sq *ScholarshipQuery) QuerySemesterID() *SemesterQuery {
	query := &SemesterQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scholarship.Table, scholarship.FieldID, sq.sqlQuery()),
			sqlgraph.To(semester.Table, semester.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, scholarship.SemesterIDTable, scholarship.SemesterIDColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUserID chains the current query on the UserID edge.
func (sq *ScholarshipQuery) QueryUserID() *UserQuery {
	query := &UserQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scholarship.Table, scholarship.FieldID, sq.sqlQuery()),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, scholarship.UserIDTable, scholarship.UserIDColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Scholarship entity in the query. Returns *NotFoundError when no scholarship was found.
func (sq *ScholarshipQuery) First(ctx context.Context) (*Scholarship, error) {
	sSlice, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(sSlice) == 0 {
		return nil, &NotFoundError{scholarship.Label}
	}
	return sSlice[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *ScholarshipQuery) FirstX(ctx context.Context) *Scholarship {
	s, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return s
}

// FirstID returns the first Scholarship id in the query. Returns *NotFoundError when no id was found.
func (sq *ScholarshipQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{scholarship.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (sq *ScholarshipQuery) FirstXID(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Scholarship entity in the query, returns an error if not exactly one entity was returned.
func (sq *ScholarshipQuery) Only(ctx context.Context) (*Scholarship, error) {
	sSlice, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(sSlice) {
	case 1:
		return sSlice[0], nil
	case 0:
		return nil, &NotFoundError{scholarship.Label}
	default:
		return nil, &NotSingularError{scholarship.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *ScholarshipQuery) OnlyX(ctx context.Context) *Scholarship {
	s, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// OnlyID returns the only Scholarship id in the query, returns an error if not exactly one id was returned.
func (sq *ScholarshipQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{scholarship.Label}
	default:
		err = &NotSingularError{scholarship.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *ScholarshipQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Scholarships.
func (sq *ScholarshipQuery) All(ctx context.Context) ([]*Scholarship, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *ScholarshipQuery) AllX(ctx context.Context) []*Scholarship {
	sSlice, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return sSlice
}

// IDs executes the query and returns a list of Scholarship ids.
func (sq *ScholarshipQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := sq.Select(scholarship.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *ScholarshipQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *ScholarshipQuery) Count(ctx context.Context) (int, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *ScholarshipQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *ScholarshipQuery) Exist(ctx context.Context) (bool, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *ScholarshipQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *ScholarshipQuery) Clone() *ScholarshipQuery {
	return &ScholarshipQuery{
		config:     sq.config,
		limit:      sq.limit,
		offset:     sq.offset,
		order:      append([]OrderFunc{}, sq.order...),
		unique:     append([]string{}, sq.unique...),
		predicates: append([]predicate.Scholarship{}, sq.predicates...),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

//  WithScholarshiptypeID tells the query-builder to eager-loads the nodes that are connected to
// the "ScholarshiptypeID" edge. The optional arguments used to configure the query builder of the edge.
func (sq *ScholarshipQuery) WithScholarshiptypeID(opts ...func(*ScholarshiptypeQuery)) *ScholarshipQuery {
	query := &ScholarshiptypeQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withScholarshiptypeID = query
	return sq
}

//  WithEducationlevelID tells the query-builder to eager-loads the nodes that are connected to
// the "EducationlevelID" edge. The optional arguments used to configure the query builder of the edge.
func (sq *ScholarshipQuery) WithEducationlevelID(opts ...func(*EducationlevelQuery)) *ScholarshipQuery {
	query := &EducationlevelQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withEducationlevelID = query
	return sq
}

//  WithStudyplanID tells the query-builder to eager-loads the nodes that are connected to
// the "StudyplanID" edge. The optional arguments used to configure the query builder of the edge.
func (sq *ScholarshipQuery) WithStudyplanID(opts ...func(*StudyplanQuery)) *ScholarshipQuery {
	query := &StudyplanQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withStudyplanID = query
	return sq
}

//  WithSemesterID tells the query-builder to eager-loads the nodes that are connected to
// the "SemesterID" edge. The optional arguments used to configure the query builder of the edge.
func (sq *ScholarshipQuery) WithSemesterID(opts ...func(*SemesterQuery)) *ScholarshipQuery {
	query := &SemesterQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withSemesterID = query
	return sq
}

//  WithUserID tells the query-builder to eager-loads the nodes that are connected to
// the "UserID" edge. The optional arguments used to configure the query builder of the edge.
func (sq *ScholarshipQuery) WithUserID(opts ...func(*UserQuery)) *ScholarshipQuery {
	query := &UserQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withUserID = query
	return sq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Schlolarshipname string `json:"Schlolarshipname,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Scholarship.Query().
//		GroupBy(scholarship.FieldSchlolarshipname).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sq *ScholarshipQuery) GroupBy(field string, fields ...string) *ScholarshipGroupBy {
	group := &ScholarshipGroupBy{config: sq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Schlolarshipname string `json:"Schlolarshipname,omitempty"`
//	}
//
//	client.Scholarship.Query().
//		Select(scholarship.FieldSchlolarshipname).
//		Scan(ctx, &v)
//
func (sq *ScholarshipQuery) Select(field string, fields ...string) *ScholarshipSelect {
	selector := &ScholarshipSelect{config: sq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sq.sqlQuery(), nil
	}
	return selector
}

func (sq *ScholarshipQuery) prepareQuery(ctx context.Context) error {
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *ScholarshipQuery) sqlAll(ctx context.Context) ([]*Scholarship, error) {
	var (
		nodes       = []*Scholarship{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [5]bool{
			sq.withScholarshiptypeID != nil,
			sq.withEducationlevelID != nil,
			sq.withStudyplanID != nil,
			sq.withSemesterID != nil,
			sq.withUserID != nil,
		}
	)
	if sq.withScholarshiptypeID != nil || sq.withEducationlevelID != nil || sq.withStudyplanID != nil || sq.withSemesterID != nil || sq.withUserID != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, scholarship.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Scholarship{config: sq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
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
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := sq.withScholarshiptypeID; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Scholarship)
		for i := range nodes {
			if fk := nodes[i].scholarshiptype_scholarshiptype_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(scholarshiptype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "scholarshiptype_scholarshiptype_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.ScholarshiptypeID = n
			}
		}
	}

	if query := sq.withEducationlevelID; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Scholarship)
		for i := range nodes {
			if fk := nodes[i].educationlevel_educationlevel_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(educationlevel.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "educationlevel_educationlevel_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.EducationlevelID = n
			}
		}
	}

	if query := sq.withStudyplanID; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Scholarship)
		for i := range nodes {
			if fk := nodes[i].studyplan_studyplan_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(studyplan.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "studyplan_studyplan_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.StudyplanID = n
			}
		}
	}

	if query := sq.withSemesterID; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Scholarship)
		for i := range nodes {
			if fk := nodes[i].semester_semester_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(semester.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "semester_semester_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.SemesterID = n
			}
		}
	}

	if query := sq.withUserID; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Scholarship)
		for i := range nodes {
			if fk := nodes[i].user_user_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_user_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.UserID = n
			}
		}
	}

	return nodes, nil
}

func (sq *ScholarshipQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *ScholarshipQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (sq *ScholarshipQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   scholarship.Table,
			Columns: scholarship.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: scholarship.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *ScholarshipQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(scholarship.Table)
	selector := builder.Select(t1.Columns(scholarship.Columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(scholarship.Columns...)...)
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ScholarshipGroupBy is the builder for group-by Scholarship entities.
type ScholarshipGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *ScholarshipGroupBy) Aggregate(fns ...AggregateFunc) *ScholarshipGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scan the result into the given value.
func (sgb *ScholarshipGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sgb.path(ctx)
	if err != nil {
		return err
	}
	sgb.sql = query
	return sgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sgb *ScholarshipGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (sgb *ScholarshipGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: ScholarshipGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sgb *ScholarshipGroupBy) StringsX(ctx context.Context) []string {
	v, err := sgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (sgb *ScholarshipGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scholarship.Label}
	default:
		err = fmt.Errorf("ent: ScholarshipGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sgb *ScholarshipGroupBy) StringX(ctx context.Context) string {
	v, err := sgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (sgb *ScholarshipGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: ScholarshipGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sgb *ScholarshipGroupBy) IntsX(ctx context.Context) []int {
	v, err := sgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (sgb *ScholarshipGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scholarship.Label}
	default:
		err = fmt.Errorf("ent: ScholarshipGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sgb *ScholarshipGroupBy) IntX(ctx context.Context) int {
	v, err := sgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (sgb *ScholarshipGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: ScholarshipGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sgb *ScholarshipGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (sgb *ScholarshipGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scholarship.Label}
	default:
		err = fmt.Errorf("ent: ScholarshipGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sgb *ScholarshipGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (sgb *ScholarshipGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: ScholarshipGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sgb *ScholarshipGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (sgb *ScholarshipGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scholarship.Label}
	default:
		err = fmt.Errorf("ent: ScholarshipGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sgb *ScholarshipGroupBy) BoolX(ctx context.Context) bool {
	v, err := sgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sgb *ScholarshipGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sgb.sqlQuery().Query()
	if err := sgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgb *ScholarshipGroupBy) sqlQuery() *sql.Selector {
	selector := sgb.sql
	columns := make([]string, 0, len(sgb.fields)+len(sgb.fns))
	columns = append(columns, sgb.fields...)
	for _, fn := range sgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(sgb.fields...)
}

// ScholarshipSelect is the builder for select fields of Scholarship entities.
type ScholarshipSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ss *ScholarshipSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ss.path(ctx)
	if err != nil {
		return err
	}
	ss.sql = query
	return ss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ss *ScholarshipSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ss *ScholarshipSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: ScholarshipSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ss *ScholarshipSelect) StringsX(ctx context.Context) []string {
	v, err := ss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ss *ScholarshipSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ss.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scholarship.Label}
	default:
		err = fmt.Errorf("ent: ScholarshipSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ss *ScholarshipSelect) StringX(ctx context.Context) string {
	v, err := ss.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ss *ScholarshipSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: ScholarshipSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ss *ScholarshipSelect) IntsX(ctx context.Context) []int {
	v, err := ss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ss *ScholarshipSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ss.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scholarship.Label}
	default:
		err = fmt.Errorf("ent: ScholarshipSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ss *ScholarshipSelect) IntX(ctx context.Context) int {
	v, err := ss.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ss *ScholarshipSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: ScholarshipSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ss *ScholarshipSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ss *ScholarshipSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ss.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scholarship.Label}
	default:
		err = fmt.Errorf("ent: ScholarshipSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ss *ScholarshipSelect) Float64X(ctx context.Context) float64 {
	v, err := ss.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ss *ScholarshipSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: ScholarshipSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ss *ScholarshipSelect) BoolsX(ctx context.Context) []bool {
	v, err := ss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ss *ScholarshipSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ss.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scholarship.Label}
	default:
		err = fmt.Errorf("ent: ScholarshipSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ss *ScholarshipSelect) BoolX(ctx context.Context) bool {
	v, err := ss.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ss *ScholarshipSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ss.sqlQuery().Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ss *ScholarshipSelect) sqlQuery() sql.Querier {
	selector := ss.sql
	selector.Select(selector.Columns(ss.fields...)...)
	return selector
}
