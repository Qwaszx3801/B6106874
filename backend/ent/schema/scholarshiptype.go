package schema

import (
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent"
)

// Scholarshiptype holds the schema definition for the Scholarshiptype entity.
type Scholarshiptype struct {
	ent.Schema
}

// Fields of the Scholarshiptype.
func (Scholarshiptype) Fields() []ent.Field {
	return []ent.Field{
		field.String("Scholarshiptypename").NotEmpty(),
	}
}

// Edges of the Scholarshiptype.
func (Scholarshiptype) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ScholarshiptypeID",Scholarship.Type),
	}
}