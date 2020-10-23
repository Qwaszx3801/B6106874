package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Semester holds the schema definition for the Semester entity.
type Semester struct {
	ent.Schema
}

// Fields of the Semester.
func (Semester) Fields() []ent.Field {
	return []ent.Field{
		field.String("Semestername").NotEmpty(),
	}
}

// Edges of the Semester.
func (Semester) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("SemesterID", Scholarship.Type),
	}
}
