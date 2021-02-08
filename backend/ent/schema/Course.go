package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Course schema.
type Course struct {
	ent.Schema
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.String("COURSE"),
	}
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
    return []ent.Edge{
		edge.To("bookcourse", Bookcourse.Type),
		edge.To("promotion", Promotion.Type),
    }
}