package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Bookcourse schema.
type Bookcourse struct {
	ent.Schema
}

// Fields of the Bookcourse.
func (Bookcourse) Fields() []ent.Field {
	return []ent.Field{
		field.Time("BOOKTIME"),
	}
}

// Edges of the Bookcourse.
func (Bookcourse) Edges() []ent.Edge {
    return []ent.Edge{
		edge.From("course", Course.Type).
			Ref("bookcourse").
			Unique(),
		edge.From("employee", Employee.Type).
			Ref("bookcourse").
			Unique(),
		edge.From("member", Member.Type).
			Ref("bookcourse").
			Unique(),	
    }
}