package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Studyplan holds the schema definition for the Studyplan entity.
type Studyplan struct {
	ent.Schema
}

// Fields of the Studyplan.
func (Studyplan) Fields() []ent.Field {
	return []ent.Field{
		field.String("Studyplanname").NotEmpty(),
	}
}

// Edges of the Studyplan.
func (Studyplan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("StudyplanID", Scholarship.Type),
	}
}
