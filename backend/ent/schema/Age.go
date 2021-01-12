package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Age schema.
type Age struct {
	ent.Schema
}

// Fields of the Age.
func (Age) Fields() []ent.Field {
	return []ent.Field{
		field.Int("AGE").
		Unique(),
	}
}

// Edges of the Age.
func (Age) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("employee", Employee.Type),
	}
}
