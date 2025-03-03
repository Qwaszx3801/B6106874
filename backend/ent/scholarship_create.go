// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/Varissara/app/ent/educationlevel"
	"github.com/Varissara/app/ent/scholarship"
	"github.com/Varissara/app/ent/scholarshiptype"
	"github.com/Varissara/app/ent/semester"
	"github.com/Varissara/app/ent/studyplan"
	"github.com/Varissara/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ScholarshipCreate is the builder for creating a Scholarship entity.
type ScholarshipCreate struct {
	config
	mutation *ScholarshipMutation
	hooks    []Hook
}

// SetSchlolarshipname sets the Schlolarshipname field.
func (sc *ScholarshipCreate) SetSchlolarshipname(s string) *ScholarshipCreate {
	sc.mutation.SetSchlolarshipname(s)
	return sc
}

// SetOrganization sets the organization field.
func (sc *ScholarshipCreate) SetOrganization(s string) *ScholarshipCreate {
	sc.mutation.SetOrganization(s)
	return sc
}

// SetScholarshiptypeIDID sets the ScholarshiptypeID edge to Scholarshiptype by id.
func (sc *ScholarshipCreate) SetScholarshiptypeIDID(id int) *ScholarshipCreate {
	sc.mutation.SetScholarshiptypeIDID(id)
	return sc
}

// SetNillableScholarshiptypeIDID sets the ScholarshiptypeID edge to Scholarshiptype by id if the given value is not nil.
func (sc *ScholarshipCreate) SetNillableScholarshiptypeIDID(id *int) *ScholarshipCreate {
	if id != nil {
		sc = sc.SetScholarshiptypeIDID(*id)
	}
	return sc
}

// SetScholarshiptypeID sets the ScholarshiptypeID edge to Scholarshiptype.
func (sc *ScholarshipCreate) SetScholarshiptypeID(s *Scholarshiptype) *ScholarshipCreate {
	return sc.SetScholarshiptypeIDID(s.ID)
}

// SetEducationlevelIDID sets the EducationlevelID edge to Educationlevel by id.
func (sc *ScholarshipCreate) SetEducationlevelIDID(id int) *ScholarshipCreate {
	sc.mutation.SetEducationlevelIDID(id)
	return sc
}

// SetNillableEducationlevelIDID sets the EducationlevelID edge to Educationlevel by id if the given value is not nil.
func (sc *ScholarshipCreate) SetNillableEducationlevelIDID(id *int) *ScholarshipCreate {
	if id != nil {
		sc = sc.SetEducationlevelIDID(*id)
	}
	return sc
}

// SetEducationlevelID sets the EducationlevelID edge to Educationlevel.
func (sc *ScholarshipCreate) SetEducationlevelID(e *Educationlevel) *ScholarshipCreate {
	return sc.SetEducationlevelIDID(e.ID)
}

// SetStudyplanIDID sets the StudyplanID edge to Studyplan by id.
func (sc *ScholarshipCreate) SetStudyplanIDID(id int) *ScholarshipCreate {
	sc.mutation.SetStudyplanIDID(id)
	return sc
}

// SetNillableStudyplanIDID sets the StudyplanID edge to Studyplan by id if the given value is not nil.
func (sc *ScholarshipCreate) SetNillableStudyplanIDID(id *int) *ScholarshipCreate {
	if id != nil {
		sc = sc.SetStudyplanIDID(*id)
	}
	return sc
}

// SetStudyplanID sets the StudyplanID edge to Studyplan.
func (sc *ScholarshipCreate) SetStudyplanID(s *Studyplan) *ScholarshipCreate {
	return sc.SetStudyplanIDID(s.ID)
}

// SetSemesterIDID sets the SemesterID edge to Semester by id.
func (sc *ScholarshipCreate) SetSemesterIDID(id int) *ScholarshipCreate {
	sc.mutation.SetSemesterIDID(id)
	return sc
}

// SetNillableSemesterIDID sets the SemesterID edge to Semester by id if the given value is not nil.
func (sc *ScholarshipCreate) SetNillableSemesterIDID(id *int) *ScholarshipCreate {
	if id != nil {
		sc = sc.SetSemesterIDID(*id)
	}
	return sc
}

// SetSemesterID sets the SemesterID edge to Semester.
func (sc *ScholarshipCreate) SetSemesterID(s *Semester) *ScholarshipCreate {
	return sc.SetSemesterIDID(s.ID)
}

// SetUserIDID sets the UserID edge to User by id.
func (sc *ScholarshipCreate) SetUserIDID(id int) *ScholarshipCreate {
	sc.mutation.SetUserIDID(id)
	return sc
}

// SetNillableUserIDID sets the UserID edge to User by id if the given value is not nil.
func (sc *ScholarshipCreate) SetNillableUserIDID(id *int) *ScholarshipCreate {
	if id != nil {
		sc = sc.SetUserIDID(*id)
	}
	return sc
}

// SetUserID sets the UserID edge to User.
func (sc *ScholarshipCreate) SetUserID(u *User) *ScholarshipCreate {
	return sc.SetUserIDID(u.ID)
}

// Mutation returns the ScholarshipMutation object of the builder.
func (sc *ScholarshipCreate) Mutation() *ScholarshipMutation {
	return sc.mutation
}

// Save creates the Scholarship in the database.
func (sc *ScholarshipCreate) Save(ctx context.Context) (*Scholarship, error) {
	if _, ok := sc.mutation.Schlolarshipname(); !ok {
		return nil, &ValidationError{Name: "Schlolarshipname", err: errors.New("ent: missing required field \"Schlolarshipname\"")}
	}
	if v, ok := sc.mutation.Schlolarshipname(); ok {
		if err := scholarship.SchlolarshipnameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Schlolarshipname", err: fmt.Errorf("ent: validator failed for field \"Schlolarshipname\": %w", err)}
		}
	}
	if _, ok := sc.mutation.Organization(); !ok {
		return nil, &ValidationError{Name: "organization", err: errors.New("ent: missing required field \"organization\"")}
	}
	if v, ok := sc.mutation.Organization(); ok {
		if err := scholarship.OrganizationValidator(v); err != nil {
			return nil, &ValidationError{Name: "organization", err: fmt.Errorf("ent: validator failed for field \"organization\": %w", err)}
		}
	}
	var (
		err  error
		node *Scholarship
	)
	if len(sc.hooks) == 0 {
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScholarshipMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ScholarshipCreate) SaveX(ctx context.Context) *Scholarship {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *ScholarshipCreate) sqlSave(ctx context.Context) (*Scholarship, error) {
	s, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}

func (sc *ScholarshipCreate) createSpec() (*Scholarship, *sqlgraph.CreateSpec) {
	var (
		s     = &Scholarship{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: scholarship.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: scholarship.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Schlolarshipname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: scholarship.FieldSchlolarshipname,
		})
		s.Schlolarshipname = value
	}
	if value, ok := sc.mutation.Organization(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: scholarship.FieldOrganization,
		})
		s.Organization = value
	}
	if nodes := sc.mutation.ScholarshiptypeIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   scholarship.ScholarshiptypeIDTable,
			Columns: []string{scholarship.ScholarshiptypeIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scholarshiptype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.EducationlevelIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   scholarship.EducationlevelIDTable,
			Columns: []string{scholarship.EducationlevelIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: educationlevel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StudyplanIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   scholarship.StudyplanIDTable,
			Columns: []string{scholarship.StudyplanIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: studyplan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.SemesterIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   scholarship.SemesterIDTable,
			Columns: []string{scholarship.SemesterIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: semester.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.UserIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   scholarship.UserIDTable,
			Columns: []string{scholarship.UserIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return s, _spec
}
