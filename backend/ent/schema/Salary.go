package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Salary schema.
type Salary struct {
	ent.Schema
}

// Fields of the Salary.
func (Salary) Fields() []ent.Field {
	return []ent.Field{
		field.Int("SALARY").Unique(),
	}
}

// Edges of the Salary.
func (Salary) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("employee", Employee.Type),
	}
}
