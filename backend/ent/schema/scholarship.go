package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Scholarship holds the schema definition for the Scholarship entity.
type Scholarship struct {
	ent.Schema
}

// Fields of the Scholarship.
func (Scholarship) Fields() []ent.Field {
	return []ent.Field{
		field.String("Schlolarshipname").NotEmpty(),
		field.String("organization").NotEmpty(),
	}
}

// Edges of the Scholarship.
func (Scholarship) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ScholarshiptypeID", Scholarshiptype.Type).Ref("ScholarshiptypeID").Unique(),
		edge.From("EducationlevelID", Educationlevel.Type).Ref("EducationlevelID").Unique(),
		edge.From("StudyplanID", Studyplan.Type).Ref("StudyplanID").Unique(),
		edge.From("SemesterID", Semester.Type).Ref("SemesterID").Unique(),
		edge.From("UserID", User.Type).Ref("UserID").Unique(),
	}
}
